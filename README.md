# Casa do Código — Desafio Jornada Dev Eficiente

Implementação do desafio da livraria **Casa do Código** proposto pelo Alberto na Jornada Dev Eficiente. O objetivo é construir uma API REST completa de forma incremental

---

## Pré-requisitos

- Go 1.22+
- Copiar `.env.sample` para `.env` e ajustar as variáveis se necessário

## Aplicação web

```bash
make build   # compila o binário em bin/api
make start   # compila e sobe a API
```

A API sobe por padrão em `http://localhost:8080`.

## Ferramenta de métricas (dev)

Analisa a complexidade dos arquivos `.go` do projeto e gera um relatório.

```bash
make build-cdd      # compila o binário em bin/metrics
make start-analyser # compila e roda o analisador
```
