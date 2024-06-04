all: dev

build-js:
	@npm --prefix ./app install
	@npm --prefix ./app run build

build: build-js
	@go run github.com/a-h/templ/cmd/templ@latest fmt .
	@go run github.com/a-h/templ/cmd/templ@latest generate
	@go build -o ./tmp/main ./cmd/nicklesseos.com/main.go
	@chmod +x ./tmp/main

dev:
	@go run github.com/air-verse/air@latest

run:
	@./tmp/main