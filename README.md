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

Build for linux
```bash
env GOOS=linux GOARCH=arm go build -v
```

RUN
```bash
./peon -r ../../test/data/project build module_a
```

EXEC
```bash
./peon -r ../../test/data/project exec module_d 'python setup.py test'
```

TEST
```bash
go test ./...
```

Get all dependencies
```bash
go get -d ./...
```
