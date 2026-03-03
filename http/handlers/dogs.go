package handlers

import (
	"context"
	"dogs-service/models"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type DogService interface {
	CreateDog(ctx context.Context, dog models.Dog) error
	GetDogs(ctx context.Context) ([]models.Dog, error)
	GetDog(ctx context.Context, id int) (models.Dog, error)
	Update(ctx context.Context, dog models.Dog) error
	Delete(ctx context.Context, id int) error
}

type DogHandler struct {
	service DogService
}

func NewDogHandler(s DogService) *DogHandler {
	return &DogHandler{service: s}
}

// CreateDog now returns (any, int, error)
func (h *DogHandler) CreateDog(r *http.Request) (any, int, error) {
	var dog models.Dog
	if err := json.NewDecoder(r.Body).Decode(&dog); err != nil {
		return nil, http.StatusBadRequest, err
	}

	if err := h.service.CreateDog(r.Context(), dog); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return map[string]string{"message": "dog created successfully"}, http.StatusCreated, nil
}

func (h *DogHandler) GetDogs(r *http.Request) (any, int, error) {
	dogs, err := h.service.GetDogs(r.Context())
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return dogs, http.StatusOK, nil
}

func (h *DogHandler) GetByID(r *http.Request) (any, int, error) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		return nil, http.StatusBadRequest, errors.New("dog id is required")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("invalid dog id")
	}

	dog, err := h.service.GetDog(r.Context(), id)
	if err != nil {
		if err.Error() == "not found" {
			return map[string]string{"error": "dog not found"}, http.StatusNotFound, nil
		}
		return nil, http.StatusInternalServerError, err
	}

	return dog, http.StatusOK, nil
}
func (h *DogHandler) Update(r *http.Request) (any, int, error) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		return nil, http.StatusBadRequest, errors.New("dog id is required")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("invalid dog id")
	}

	var dog models.Dog
	if err := json.NewDecoder(r.Body).Decode(&dog); err != nil {
		return nil, http.StatusBadRequest, err
	}

	dog.ID = id
	if err := h.service.Update(r.Context(), dog); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return map[string]string{"message": "dog updated successfully"}, http.StatusOK, nil
}

func (h *DogHandler) Delete(r *http.Request) (any, int, error) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		return nil, http.StatusBadRequest, errors.New("dog id is required")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("invalid dog id")
	}

	err = h.service.Delete(r.Context(), id)
	if err != nil {
		if err.Error() == "not found" {
			return map[string]string{"error": "dog not found"}, http.StatusNotFound, nil
		}
		return nil, http.StatusInternalServerError, err
	}

	return map[string]string{"message": "dog deletd successfully"}, http.StatusOK, nil
}
