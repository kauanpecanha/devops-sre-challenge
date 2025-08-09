# Documentação do Desafio Técnico de SRE/DevOps

## 🚀 Tecnologias

Este projeto foi desenvolvido com as seguintes tecnologias:

- go
- docker
- mongodb
- opentelemetry

## 📦 Instalação

1. Clone o repositório:

```bash
git clone https://github.com/kauanpecanha/devops-sre-challenge.git
```

## 📖 Comandos utilizados

criação do pacote
```bash
go mod init kauanpecanha/devops-challenge
```

adição de dependências do opentelemetry
```bash
go get "go.opentelemetry.io/otel" \
  "go.opentelemetry.io/otel/exporters/stdout/stdoutmetric" \
  "go.opentelemetry.io/otel/exporters/stdout/stdouttrace" \
  "go.opentelemetry.io/otel/exporters/stdout/stdoutlog" \
  "go.opentelemetry.io/otel/sdk/log" \
  "go.opentelemetry.io/otel/log/global" \
  "go.opentelemetry.io/otel/propagation" \
  "go.opentelemetry.io/otel/sdk/metric" \
  "go.opentelemetry.io/otel/sdk/resource" \
  "go.opentelemetry.io/otel/sdk/trace" \
  "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"\
  "go.opentelemetry.io/contrib/bridges/otelslog"
```

rodando o projeto
```bash
export OTEL_RESOURCE_ATTRIBUTES="service.name=dice,service.version=0.1.0"
go run .
```