.PHONY: build clean deploy

build:
	go build
	env GOOS=linux go build -ldflags="-s -w" -o bin/sendMail ./main.go

clean:
	rm -rf ./bin 

deploy: clean build
	sls deploy --verbose
