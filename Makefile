all: dev

build:
	@go run github.com/a-h/templ/cmd/templ@latest generate
	@npm --prefix app install
	@npm --prefix app run build
	@go build -o ./tmp/main cmd/nicklesseos.com/main.go

dev:
	@go run github.com/cosmtrek/air@latest

run: build
	@./tmp/main