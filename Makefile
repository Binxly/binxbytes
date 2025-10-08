# ---------------------------------------------
# binxbytes Makefile
# ---------------------------------------------

FUNCTION_NAME = binxbytes
CLOUDFRONT_DIST_ID = E1NOQF6M7H13UN

.PHONY: help dev build build-lambda package deploy clean

help:
	@echo "Available commands:"
	@echo "  make dev            - Run local dev server"
	@echo "  make build          - Build local binary"
	@echo "  make build-lambda   - Build Linux binary for Lambda (ARM64)"
	@echo "  make package        - Zip binary + static folders"
	@echo "  make deploy         - Deploy ZIP to Lambda and invalidate CloudFront"
	@echo "  make clean          - Remove built files"

dev:
	go run . -dev

build:
	go build -o binxbytes .

build-lambda:
	GOOS=linux GOARCH=arm64 go build -o bootstrap .

package: build-lambda
	zip -r function.zip bootstrap static/ blog/ templates/

deploy: package
	aws lambda update-function-code \
		--function-name $(FUNCTION_NAME) \
		--zip-file fileb://function.zip

	aws cloudfront create-invalidation \
		--distribution-id $(CLOUDFRONT_DIST_ID) \
		--paths "/*"

clean:
	rm -f bootstrap function.zip binxbytes