build:
	go build -o ./bin/blog-rss

run: build
	./bin/blog-rss

test:
	go test -v ./...

