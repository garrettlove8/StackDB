run-db:
	go run ./cmd/db/.

test:
	go test -v -cover ./...

clean:
	rm -rf ./stackdb
	rm -rf ./sdb
	go clean

build: clean
	go build