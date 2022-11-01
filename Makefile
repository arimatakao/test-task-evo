all:
	go run ./cmd/main.go

build:
	go build ./cmd/main.go

parse:
	go run ./cmd/main.go example.csv
