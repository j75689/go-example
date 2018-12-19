## Build
```bash
cd plugin
go tool compile -I $GOPATH/pkg/darwin_amd64 ./*.go
```

## Run
```bash
go run main.go
```