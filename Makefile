# ---------------------------------------------
# binxbytes Makefile
# ---------------------------------------------

FUNCTION_NAME = binxbytes

.PHONY: help dev build build-lambda package deploy clean

help:
	@echo "Available commands:"
	@echo "  make dev            - Run local dev server"
	@echo "  make build          - Build local binary"
	@echo "  make build-lambda   - Build Linux binary for Lambda (ARM64)"
	@echo "  make package        - Zip binary + static folders"
	@echo "  make deploy         - Deploy ZIP to Lambda"
	@echo "  make clean          - Remove built files"

dev:
	go run main.go -dev

build:
	go build -o binxbytes main.go

build-lambda:
	GOOS=linux GOARCH=arm64 go build -o bootstrap main.go

package: build-lambda
	zip -r function.zip bootstrap static/ blog/ templates/

deploy: package
	aws lambda update-function-code \
		--function-name $(FUNCTION_NAME) \
		--zip-file fileb://function.zip

clean:
	rm -f bootstrap function.zip binxbytes