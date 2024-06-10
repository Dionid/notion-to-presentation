ifneq ("$(wildcard $(.env))","")
    include .env
	export $(shell sed 's/=.*//' .env)
endif

PROJECT_NAME=ntp
BINARY_NAME=${PROJECT_NAME}

# Run

parse:
	cd cmd/cli && go run .

serve-saas:
	npx tailwindcss build -i tailwind.css -o cmd/saas/public/style.css
	cd cmd/saas \
	&& go generate ./... \
	&& go run . serve

live-saas:
	templ generate --watch --proxy="http://localhost:8090" --cmd="make serve-saas"

# Generate

templ:
	templ generate

# Build

build-cli:
	make build-cli-mac && make build-cli-linux

build-cli-mac:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin ./cmd/cli

clean-mac:
	go clean
	rm ${BINARY_NAME}-darwin

build-cli-linux:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux ./cmd/cli

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux

# Setup

setup:
	npm i
	go install github.com/a-h/templ/cmd/templ@latest
	go get
	go mod tidy