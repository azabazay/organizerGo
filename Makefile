build:
	go build -o bin/fact

run: build
	./bin/fact
