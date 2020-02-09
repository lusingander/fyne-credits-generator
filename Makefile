BINARY_NAME=fyne-credits-generator
SAMPLE_BINARY_NAME=sample

.PHONY: build
build:
	go build -o $(BINARY_NAME) ./cmd/fyne-credits-generator

.PHONY: generate
generate: build
	./$(BINARY_NAME) > ./cmd/sample/credits.go

.PHONY: sample
sample: generate
	go build -o $(SAMPLE_BINARY_NAME) ./cmd/sample

.PHONY: run
run: sample
	./$(SAMPLE_BINARY_NAME)

clean:
	rm $(BINARY_NAME) $(SAMPLE_BINARY_NAME)