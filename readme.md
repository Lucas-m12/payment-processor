# 💸 Processador de Pagamentos Bancários Concorrente

Simulador em Go para o processamento paralelo de boletos bancários, com foco em concorrência, resiliência e arquitetura limpa.

## 🧠 Objetivo

Demonstrar o uso prático de goroutines, channels, contextos com cancelamento e retry com backoff em cenários reais de back-end bancário.

## 🚀 Funcionalidades

- Processamento paralelo de múltiplos pagamentos
- Agrupamento por banco de destino
- Cancelamento de todo o lote em caso de falha
- Retry automático com backoff exponencial
- Relatório final com sucesso e falhas
- Estrutura modular e idiomática em Go

## 🛠️ Como rodar o projeto

```bash
# Clone o repositório
git clone https://github.com/seu-usuario/processador-pagamentos-go.git
cd processador-pagamentos-go

# Rode o projeto
go run ./cmd/processor
```

## 🧪 Rodando os testes

```bash
go test ./...
```

## 📂 Estrutura do projeto

```
cmd/               → Ponto de entrada (main.go)
internal/
  domain/          → Entidades e regras de negócio
  service/         → Lógica de orquestração
  infra/           → Simulação de serviços externos
  util/            → Funções auxiliares como retry
tests/             → Testes de integração
```

## 📋 Licença

MIT

---

# 💸 Concurrent Bank Payment Processor

Go-based simulation of parallel boleto processing, focusing on concurrency, resiliency, and clean architecture.

## 🧠 Purpose

Showcase practical usage of goroutines, channels, cancellation contexts, and retry with exponential backoff in real-world backend scenarios.

## 🚀 Features

- Parallel processing of multiple payments
- Grouping by target bank
- Cancel the entire batch on any failure
- Automatic retry with backoff
- Final report showing successes and failures
- Clean and idiomatic Go structure

## 🛠️ How to run

```bash
# Clone the repository
git clone https://github.com/your-username/processador-pagamentos-go.git
cd processador-pagamentos-go

# Run the project
go run ./cmd/processor
```

## 🧪 Running tests

```bash
go test ./...
```

## 📂 Project structure

```
cmd/               → Entry point (main.go)
internal/
  domain/          → Entities and business rules
  service/         → Orchestration logic
  infra/           → External service simulation
  util/            → Helper functions like retry
tests/             → Integration tests
```

## 📋 License

MIT
