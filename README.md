# Korp Teste - Sistema de Emissão de Notas Fiscais

Projeto técnico desenvolvido para o processo seletivo de estágio da Korp ERP.

**Desenvolvedor:** Carlos Wylliam de Oliveira Soares

## Sumário

- [Visão Geral](#visão-geral)
- [Arquitetura](#arquitetura)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Como Rodar o Projeto](#como-rodar-o-projeto)
- [Funcionalidades](#funcionalidades)
- [Detalhamento Técnico](#detalhamento-técnico)

## Visão Geral

Sistema web para cadastro de produtos, emissão e gerenciamento de notas fiscais, desenvolvido com front-end em Angular e back-end em Go, seguindo arquitetura de microsserviços.

## Arquitetura

O sistema é dividido em três partes independentes:
```
KorpTeste/
├── frontend/
│ └── korpteste-frontend/ → Aplicação Angular
└── backend/
├── estoque/ → Microsserviço de Estoque (porta 8000)
└── faturamento/ → Microsserviço de Faturamento (porta 8001)
```


- **Serviço de Estoque**: responsável pelo cadastro, consulta, atualização e exclusão de produtos, incluindo controle de saldo.
- **Serviço de Faturamento**: responsável pelo cadastro, listagem, impressão e exclusão de notas fiscais. Se comunica com o Serviço de Estoque via HTTP para validar saldo e abater estoque.

Cada microsserviço possui seu próprio banco de dados SQLite, garantindo independência entre eles.

## Tecnologias Utilizadas

### Front-end
- **Angular** (Standalone Components)
- **RxJS**: utilizado internamente pelo `HttpClient` para lidar com as chamadas assíncronas à API, através de `Observables` tratados com `.subscribe()`.
- **jsPDF**: geração do documento fictício de impressão da nota fiscal.
- **Material Symbols (Google Fonts)**: ícones da interface.
- **Fontes**: Space Grotesk (títulos) e Geist (corpo de texto), via Google Fonts.
- CSS puro, sem bibliotecas de componentes visuais (Bootstrap, Angular Material, etc).

### Back-end
- **Go (Golang)**
- **Gin**: framework web utilizado para roteamento HTTP e middlewares.
- **gin-contrib/cors**: middleware de CORS para permitir requisições do front-end.
- **modernc.org/sqlite**: driver SQLite puro em Go (não requer CGO).
- **Gerenciamento de dependências**: `go.mod` / `go.sum`, gerenciados automaticamente via `go get` e `go mod init`.

### Banco de Dados
- **SQLite**, com um arquivo de banco separado para cada microsserviço (`estoque.db` e `faturamento.db`), criado e migrado automaticamente na inicialização de cada serviço.

## Como Rodar o Projeto

O projeto precisa de **3 processos rodando simultaneamente**, cada um em um terminal:

### 1. Serviço de Estoque
```bash
cd backend/estoque
go run main.go
```
Sobe em `http://localhost:8000`

### 2. Serviço de Faturamento
```bash
cd backend/faturamento
go run main.go
```
Sobe em `http://localhost:8001`

### 3. Front-end Angular
```bash
cd frontend/korpteste-frontend
npm install
ng serve
```
Sobe em `http://localhost:4200`

Acesse `http://localhost:4200` no navegador.

**Pré-requisitos:** Go instalado, Node.js e Angular CLI instalados (`npm install -g @angular/cli`).

## Funcionalidades

### Cadastro de Produtos
- Campos: código, descrição, saldo
- Validação de código duplicado
- Validação de saldo negativo e descrição mínima
- Exclusão de produtos

### Cadastro de Notas Fiscais
- Numeração sequencial automática (gerada pelo banco)
- Status inicial "Aberta" (definido automaticamente)
- Inclusão de múltiplos produtos com suas respectivas quantidades
- Validação de existência do produto e saldo disponível (considerando reservas de outras notas "Abertas") no momento da criação, via comunicação com o Serviço de Estoque
- Exclusão de notas (somente quando status "Aberta")

### Impressão de Notas Fiscais
- Botão de impressão visível na listagem
- Indicador visual de processamento ("Imprimindo...")
- Bloqueio de impressão para notas que não estão "Aberta"
- Ao imprimir: abate o saldo dos produtos utilizados (via comunicação com o Serviço de Estoque) e atualiza o status da nota para "Fechada"
- Geração de um documento PDF fictício da nota fiscal (via jsPDF)

## Detalhamento Técnico

### Ciclos de vida do Angular utilizados
- **`ngOnInit`**: utilizado nos componentes de listagem (`ProdutosList`, `NotasList`) para carregar os dados da API assim que o componente é inicializado.

### Uso da biblioteca RxJS
Sim. O `HttpClient` do Angular retorna `Observables` para todas as chamadas HTTP (GET, POST, PUT, DELETE). Essas chamadas são tratadas com `.subscribe()`, recebendo os blocos `next` (sucesso) e `error` (falha), utilizados para atualizar a interface e exibir mensagens de sucesso/erro ao usuário.

### Outras bibliotecas utilizadas
- **jsPDF**: geração do PDF fictício da nota fiscal no momento da impressão.

### Bibliotecas de componentes visuais
Nenhuma biblioteca de componentes visuais foi utilizada (sem Angular Material, Bootstrap, etc). Toda a interface foi estilizada com CSS puro, seguindo um design system próprio (cores, tipografia e componentes definidos no Figma). Os ícones utilizados são da biblioteca **Material Symbols** do Google, importados via CDN.

### Gerenciamento de dependências no Golang
Realizado através do sistema nativo de módulos do Go (`go mod`). Cada microsserviço possui seu próprio `go.mod` e `go.sum`, com dependências instaladas via `go get` (ex: `gin-gonic/gin`, `gin-contrib/cors`, `modernc.org/sqlite`).

### Frameworks utilizados no Golang
**Gin**, utilizado para roteamento HTTP, middlewares (CORS) e serialização/desserialização de JSON (`ShouldBindJSON`).

### Tratamento de erros e exceções no backend
- Erros de validação (campos obrigatórios, produto não encontrado, saldo insuficiente, código duplicado) retornam status HTTP apropriados (`400 Bad Request`, `404 Not Found`, `409 Conflict`) com mensagens descritivas em JSON.
- Falha de comunicação entre microsserviços (ex: Serviço de Estoque fora do ar durante criação/impressão de nota) é tratada com o status `503 Service Unavailable`, informando ao usuário que o serviço está indisponível, sem travar a aplicação.
- Erros inesperados de banco de dados retornam `500 Internal Server Error` com detalhes do erro para fins de depuração.
- No front-end, todos os erros retornados pela API são capturados no bloco `error` do `.subscribe()` e exibidos ao usuário através de mensagens visuais (toasts).

### Uso de C# e LINQ
Não se aplica. A implementação do back-end foi feita inteiramente em **Golang**, não em C#.