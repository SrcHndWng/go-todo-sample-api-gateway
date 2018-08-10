build:
	GOOS=linux go build -o bin/post
	GOOS=linux go build -o bin/get
	GOOS=linux go build -o bin/put
	GOOS=linux go build -o bin/delete