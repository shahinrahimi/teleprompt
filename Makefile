build: 
	@go build -o ./bin/teleprompt
run: build
	@./bin/teleprompt