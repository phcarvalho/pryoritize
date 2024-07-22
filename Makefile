build: 
	@go build -o ./bin/pry .
run:
	@go build -o ./bin/pry . && ./bin/pry
