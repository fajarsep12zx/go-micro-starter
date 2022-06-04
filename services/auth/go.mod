module github.com/fajarsep12zx/go-micro-starter/services/auth

go 1.13

require (
	github.com/fajarsep12zx/go-micro-starter/core v0.0.0
	github.com/joho/godotenv v1.3.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	github.com/patrickmn/go-cache v2.1.0+incompatible
)

replace github.com/fajarsep12zx/go-micro-starter/core => ../../core

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
