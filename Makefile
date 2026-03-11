.PHONY: generate clean tidy

generate:
	go generate ./...

clean:
	rm -rf services/

tidy:
	go mod tidy
