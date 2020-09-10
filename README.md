# peon

Work work

Init module
```bash
go mod init
```

BUILD
```bash
cd cmd/peon
go build
```

RUN
```bash
./peon -r ../../test/data/project build -m module_a
```

TEST
```bash
go test ./...
```

Get all dependencies
```bash
go get -d ./...
```
