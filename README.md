# some server
start:
```
go run main.go
```

in docker:
```
docker build --tag some-server .
docker run -p 8090:8090 some-server
```