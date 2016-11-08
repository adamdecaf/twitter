.PHONY: build test vet

vet:
	./bin/cmd.sh 'vet' $(proj)

build: vet
	./bin/cmd.sh 'build' $(proj)

test: build
	./bin/cmd.sh 'test' $(proj)
