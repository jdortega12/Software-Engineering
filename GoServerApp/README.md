# Go Server App
Server application portion of SE 2022 Group Project

### Dependencies
1. [Go Programming Language](https://go.dev/dl/)
2. [VSCode Go Extension (only if using VSCode)](https://marketplace.visualstudio.com/items?itemName=golang.Go)
3. [Go Gin-Gonic HTTP Framework](https://github.com/gin-gonic/gin)
4. [Gin-Gonic Sessions Package](https://github.com/gin-contrib/sessions)
5. [GORM (Go Object-Relational Mapping)](https://gorm.io/index.html)
6. [Gobco (for testing branch coverage)](https://github.com/rillig/gobco)
7. [Windows GCC Port (GORM uses some cgo libraries)](https://sourceforge.net/projects/tdm-gcc/)

### Build
```go build```

### Run
```./GoServerApp.exe```

### Test
#### Statement coverage
```go test ./... -cover```
#### Branch coverage (one package at a time)
```gobco [package]```

####
`doc.go` files are just for documenting that package's functionality.