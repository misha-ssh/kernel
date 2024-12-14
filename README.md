# Module kernel 

## Tests

Tests for this package put in dir ```/test```

Run all test

```bash
go test ./... -test.v -tags module
```

### Units

Run test
```bash
go test ./... -test.v 
```

Generate coverage file
```bash
go test -coverprofile=coverage.out ./... 
```

Transform coverage file to html
```bash
 go tool cover -html=coverage.out -o coverage.html
```

### Modules

Run test

```bash
go test ./tests/modules -test.v -tags module
```