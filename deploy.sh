#!/usr/bin/env bash
# ---------------------------------------------
# Lambda Deployment Script
# ---------------------------------------------
set -e

FUNCTION_NAME="${LAMBDA_FUNCTION_NAME:-binxbytes}"
AWS_REGION="${AWS_REGION:-us-east-1}"

echo "🚀 Building Lambda binary..."
GOOS=linux GOARCH=arm64 go build -o bootstrap main.go

echo "📦 Creating ZIP..."
zip -r function.zip bootstrap static/ blog/ templates/

echo "🔄 Deploying to Lambda..."
aws lambda update-function-code \
  --function-name $FUNCTION_NAME \
  --zip-file fileb://function.zip \
  --region $AWS_REGION

echo "✅ Deployment complete!"
echo "🧹 Cleaning up..."
rm -f bootstrap function.zip

echo "🎉 Done!"
