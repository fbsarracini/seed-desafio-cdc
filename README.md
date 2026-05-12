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

Analisa a complexidade dos arquivos `.go` do projeto com base nas métricas derivadas do **CDD — Cognitive-Driven Development**, teoria criada pelo brasileiro [Alberto Souza](https://github.com/asouza). O CDD propõe limitar a carga cognitiva por unidade de código, atribuindo pontos de complexidade a construções que exigem mais esforço mental do desenvolvedor (como branches, lambdas, acoplamentos, etc.). Cada arquivo deve respeitar um teto máximo de pontos para manter o código sustentável.

```bash
make build-cdd      # compila o binário em bin/metrics
make start-analyser # compila e roda o analisador
```

## Referências

- [ICSME 2020 — Cognitive-Driven Development (artigo oficial)](https://github.com/asouza/pilares-design-codigo/blob/master/ICSME-2020-cognitive-driven-development.pdf)
