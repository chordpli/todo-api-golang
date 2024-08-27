# 종속성 추가

```shell
# framework
go get -u github.com/gin-gonic/gin

## cors
go get github.com/go-chi/cors

## middleware
go get -u github.com/go-chi/chi/v5/middleware

# log
go get -u github.com/go-chi/httplog/v2

# mysql
go get -u github.com/go-sql-driver/mysql

# entd
go install entgo.io/ent/cmd/ent@latest
go get -d entgo.io/ent/cmd/ent

# zerolog
go get github.com/rs/zerolog
go get github.com/go-chi/httplog

# swagger
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/http-swagger
swag init

# test
go install github.com/golang/mock/mockgen@latest
go get github.com/golang/mock/gomock
go get github.com/stretchr/testify/assert

# viper
go get -u github.com/spf13/viper

# validator
go get -u github.com/go-playground/validator/v10

## mock 생성
mockgen -source=internal/api/user/repository.go -destination=test/{domain}/mock/mock_{domain}_repository.go -package=user

```

# 실행
```shell
make api
```

# ENTD
`go run -mod=mod entgo.io/ent/cmd/ent new --target ./edge/ent/schema <entity>`  
`go generate ./edge/ent/`

go run -mod=mod entgo.io/ent/cmd/ent new --target ./edge/ent/schema Clothes

# DOCS
```shell
swag init -g ./cmd/app/main.go
```

http://localhost:8000/swagger/index.html
