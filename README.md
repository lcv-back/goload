# goload

## Structure

```
goload/
├── Makefile
├── buf.yaml
├── go.mod
├── go.sum
├── .golangci.yaml
├── buf.gen.yaml
├── buf.yaml
├── VERSION
├── api/
│   ├── go_load.proto
│   ├── validate.proto
│   ├── go_load.swagger.json
│   ├── api/
│   │   └── go_load.swagger.json
│   │
├── cmd/
│   └── main.go
│   │
├── configs/
│   └── configs.go
│   └── local.yaml
│   │
├── deployments/
│   ├── docker-compose.dev.yml
│   │
├── internal/
│   ├── configs/
│   │   └── auth.go
│   │   └── cache.go
│   │   └── config.go
│   │   └── db.go
│   │   └── grpc.go
│   │   └── http.go
│   │   └── log.go
│   │   └── wireset.go
│   │
│   ├── dataaccess/
│   │   ├── cache/
│   │   ├── database/
│   │   │   ├── account_password.go
│   │   │   ├── account.go
│   │   │   ├── db.go
│   │   │   ├── download_task.go
│   │   │   ├── token_public_key.go
│   │   │   ├── wireset.go
│   │   │   └── migrations/
│   │   │   │   └── mysql/
│   │   │   │       ├── 0001.initalize.sql
│   │   ├── wireset.go
│   ├── generated/
│   │   └── grpc/
│   │       └── go_load/
│   │   │   │   └── go_load_grpc.pb.go
│   │   │   │   ├── go_load.pb.go
│   │   │   │   ├── go_load.pb.gw.go
│   │   │   │   ├── go_load.pb.validate.go
│   ├── handler/
│   │   ├── consumers/
│   │   │   └── .gitkeep
│   │   ├── grpc/
│   │   │   └── handler.go
│   │   │   └── server.go
│   │   │   └── wireset.go
│   │   └── http/
│   │   │   └── server.go
│   │   │   └── wireset.go
│   │   └── jobs/
│   │   │   └── .gitkeep
│   │   ├── wireset.go
│   │   │
│   ├── logic/
│   │   ├── account.go
│   │   └── download_task.go
│   │   └── hash.go
│   │   └── token.go
│   │   └── wireset.go
│   ├── utils/
│   │   ├── jwt.go
│   │   └── log.go
│   │   └── wireset.go
│   ├── wiring/
│   │   ├── wire.go
│   │   └── wire_gen.go
├── vendor/
│   ├── github.com/
│   └── go.uber.org/
│   ├── golang.org/
│   ├── google.golang.org/
│   └── gopkg.in/
└── build/
    └── (compiled binaries)
```
