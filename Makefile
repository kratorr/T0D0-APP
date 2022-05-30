
build:
	go build ./cmd/main.go
docs:
	swag init -g ./pkg/handler/handler.go
test:
	echo "test"