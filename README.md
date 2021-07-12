# vault

CLI tool that stores secrets in an encrypted file.

## usage

The CLI usage looks like the following:
```bash
go build -o ./vault cmd/cli.go 
./vault set some_key "some-value"
./vault get some_key
```