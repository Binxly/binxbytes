# BinxBytes

BinxBytes is a simple static site and blog built with Go, deployed on AWS.

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

Then clean root directory using:

```bash
make clean
```

## Notes

- `algnhsa` is used to run Go's standard `net/http` handler on Lambda.

- Static files and blog posts are bundled with the binary, but might change this so static files are hosted on S3.

- Deployment uses the AWS CLI, removed the use of SAM or CloudFormation.
