APP=pohls

build: clean
	go build -o ${APP} main.go

run:
	go run main.go

dev:
	make run
