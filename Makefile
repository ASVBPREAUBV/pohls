APP=pohls

build:
	go build -o ${APP} main.go

run:
	go run main.go

dev:
	make run

test:
	go test -v
