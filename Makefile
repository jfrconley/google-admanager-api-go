.PHONY: generate clean tidy

generate:
	./scripts/generate.sh

clean:
	rm -rf services/

tidy:
	go mod tidy
