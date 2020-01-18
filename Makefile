APP=pohls
TESTINPUTDIR=./
TESTOUTPUTDIR=out
PHL_TMDB_TOKEN="tmdb_token"

go-build:
	go build -o ${APP}-${TARGET} main.go

run:
	go run main.go -i ${TESTINPUTDIR} -o ${TESTOUTPUTDIR} -d

dev:
	make run

go-test:
	go test  ./... -v
