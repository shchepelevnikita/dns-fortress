# Project Structure

myproject/
├── cmd/
│   └── myproject/
│       └── main.go
├── pkg/
│   └── myproject/
│       ├── dns/
│       │   ├── handler.go
│       │   └── resolver.go
│       └── util/
│           └── utils.go
├── internal/
│   └── myproject/
│       ├── config/
│       │   └── config.go
│       ├── service/
│       │   └── service.go
│       └── api/
│           └── api.go
├── api/
│   └── swagger.yaml
├── deployments/
│   ├── ansible/
│   │   └── playbook.yml
│   └── kubernetes/
│       ├── deployment.yaml
│       └── service.yaml
├── go.mod
├── go.sum
└── README.md