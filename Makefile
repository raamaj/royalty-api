generate-doc:
	swag init -g cmd/main.go

start:
	go run cmd/main.go