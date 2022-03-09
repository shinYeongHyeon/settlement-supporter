# Settlement

## API

### Run
```shell
docker-compose -f docker-compose-db.yml up -d
docker-compose up -d 
```

### Documentation update
```shell
swag fmt # (Option) 포맷팅
swag init
```

### Formatting
```shell
gofmt -w src/*
```

### Test
```shell
go test ./...
```

## TODO
-[ ] Replace ID/PW to OAuth