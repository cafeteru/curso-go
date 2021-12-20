# curso-go

To init project -> `go mod init main`

To execute test use this command

```shell
go test *.go
go test ./...
```

To format code -> `gofmt`

To inspect code -> `go vet` or `golint`

To generate Documentation

```shell
go install golang.org/x/tools/cmd/godoc
godoc -http=:8080
```
