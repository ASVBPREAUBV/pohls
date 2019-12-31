APP=pohls

go-build:
	go build -o ${APP} main.go

run:
	go run main.go

dev:
	make run

go-test:
	go test  ./internal/resolver/ -v
