up:
	docker compose up -d
	go run cmd/dogs-service/main.go
kill:
	npx kill-port 8082