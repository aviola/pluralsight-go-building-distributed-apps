.PHONY: clean
clean: 
	rm -rf ./bin

.PHONY: build-registryservice
build-registryservice: 
	go build -o bin/registryservice ./cmd/registryservice

.PHONY: build-logservice
build-logservice: 
	go build -o bin/logservice ./cmd/logservice

.PHONY: build
build: build-registryservice build-logservice


