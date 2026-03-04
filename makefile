up:
	npx kill-port 8082
	goose -dir db/migrations postgres "postgres://dogs:dogs123@localhost:5450/dogs_db?sslmode=disable" up
	docker compose up -d
	go run cmd/dogs-service/main.go
kill:
	npx kill-port 8082

