build:
	GOOS=linux go build -o bin/create
	GOOS=linux go build -o bin/list
	GOOS=linux go build -o bin/get
	GOOS=linux go build -o bin/update
	GOOS=linux go build -o bin/delete