run:
	go run -race main.go
	
watch:
	nodemon --exec go run main.go --signal SIGTERM

clean:
	go clean

build:
	@GOOS=linux GOARCH=amd64
	@go build -o gin-mongodb ./
	@echo ">> Finished"

prod:
	@./gin-mongodb