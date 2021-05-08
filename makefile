run-db:
	go run ./cmd/db/.

test:
	go test -v -cover ./...

clean:
	rm -rf ./stackdb