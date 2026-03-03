package http

import (
	"context"
	"dogs-service/http/handlers"
	"dogs-service/http/response"
	"dogs-service/logger"
	"dogs-service/service/health"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	dogHandler *handlers.DogHandler
	healthSvc  *health.HealthCheckerService
}

func NewServer(dogHandler *handlers.DogHandler, healthSvc *health.HealthCheckerService) *Server {
	return &Server{
		dogHandler: dogHandler,
		healthSvc:  healthSvc,
	}
}

func (s *Server) Listen(ctx context.Context, addr string) error {

	// Router
	r := chi.NewRouter()

	// Middlewares (same as production)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Routes
	r.Route("/api/v1", func(r chi.Router) {

		r.Get("/health", s.HealthCheckHandler)

		r.Route("/dogs", func(r chi.Router) {

			// Wrap handlers that return (any, int, error)
			r.Post("/", WrapHandler(s.dogHandler.CreateDog))
			r.Get("/", WrapHandler(s.dogHandler.GetDogs))

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", WrapHandler(s.dogHandler.GetByID))
				r.Put("/", WrapHandler(s.dogHandler.Update))
				r.Delete("/", WrapHandler(s.dogHandler.Delete))
			})
		})
	})

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	errCh := make(chan error, 1)

	go func() {
		logger.Info("Server started on " + addr)
		errCh <- server.ListenAndServe()
	}()

	select {

	case err := <-errCh:
		return err

	case <-ctx.Done():
		logger.Info("Shutting down server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return server.Shutdown(shutdownCtx)
	}
}

// WrapHandler converts a handler that returns (any, int, error) into http.HandlerFunc
func WrapHandler(handler func(r *http.Request) (any, int, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				logger.Error("panic recovered", rec)
				response.RespondMessage(w, http.StatusInternalServerError, "internal server error")
			}
		}()

		res, status, err := handler(r)
		if err != nil {
			logger.Error("request error", err)
			response.RespondMessage(w, status, err.Error())
			return
		}

		if res != nil {
			response.RespondJSON(w, status, res)
		}
	}
}

func (s *Server) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if ok := s.healthSvc.Health(r.Context()); !ok {
		response.RespondMessage(w, http.StatusServiceUnavailable, "health check failed")
		return
	}
	response.RespondMessage(w, http.StatusOK, "!!! We are RunninGoo !!!")
}