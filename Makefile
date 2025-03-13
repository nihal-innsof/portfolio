build:
	npx tailwindcss -i views/css/styles.css -o public/css/styles.css
	@templ generate view
	@go build -o tmp/main .cmd/main.go

test:
	@go test -v ./...

go:
	@air

templ:
	@templ generate --watch --proxy="http://localhost:3000" --open-browser=true

tailwind:
	@npx tailwindcss -i internal/views/css/input.css -o static/css/output.css --watch

install-deps:
	@go mod tidy
	@npm install tailwindcss @tailwindcss/cli
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get github.com/a-h/templ
