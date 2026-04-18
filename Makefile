# Змінні
BINARY_NAME=app.exe
LINTER_PATH=C:\Users\Андрей\go\bin\golangci-lint

.PHONY: all fmt lint test build

# Головна ціль: виконує все по черзі
all: fmt lint test build

# Форматування коду згідно зі стандартами Go
fmt:
	go fmt ./...

# Запуск статичного аналізу
lint:
	$(LINTER_PATH) run

# Запуск тестів з детектором стану гонитви (Race Detector)
test:
	go test -v -race ./...

# Компіляція бінарного файлу
build:
	if not exist bin mkdir bin
	go build -o bin/$(BINARY_NAME) ./cmd/app/main.go