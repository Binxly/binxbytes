# BinxBytes

BinxBytes is a simple static site and blog built with Go, deployed to AWS Lambda thanks to [algnhsa](https://github.com/akrylysov/algnhsa).

## Development

To run on a local server:

```bash
make dev
```

Build a local binary with:

```bash
make build
```

## Deployment

To deploy, build for Lambda and upload using:

```bash
make deploy
```

Then clean any temporary files from the build in the root directory using:

```bash
make clean
```

## Notes

- `algnhsa` is used to run Go's standard standard `net/http` handler on Lambda.

- Static files and blog posts are bundled with the binary, but need to change this so that posts and static files are hosted in a S3 bucket.

- Deployment uses the AWS CLI, will need changes for SAM or CloudFormation.