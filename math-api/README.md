# Math API

Uma API REST simples e eficiente desenvolvida em **Go** que fornece operações matemáticas básicas através de endpoints HTTP.

## 📋 Descrição

Math API é um servidor HTTP que expõe quatro operações matemáticas fundamentais:

- **Adição (Sum)** - Soma dois números
- **Subtração (Sub)** - Subtrai dois números
- **Multiplicação (Mult)** - Multiplica dois números
- **Divisão (Div)** - Divide dois números

A API foi desenvolvida seguindo boas práticas de estrutura de projeto Go, com separação clara de responsabilidades entre handlers e modelos de dados.

## 🚀 Quickstart

### Requisitos

- **Go 1.26.1** ou superior
- **Make** (para usar os comandos do Makefile)

### Instalação

1. Clone o repositório:

```bash
git clone <seu-repositorio>
cd math-api
```

2. Instale as dependências (não há dependências externas, apenas a biblioteca padrão do Go):

```bash
go mod download
```

3. Execute o projeto:

```bash
make run
```

Ou execute diretamente:

```bash
go run cmd/api/main.go
```

O servidor iniciará na porta **8080** e estará pronto para receber requisições.

## 🔌 Endpoints da API

### 1. Adição (Sum)

**Endpoint:** `POST /api/sum`

Realiza a soma de dois números.

**Request Body:**

```json
{
    "numA": 10,
    "numB": 5
}
```

**Response (201 Created):**

```json
{
    "result": 15
}
```

### 2. Subtração (Sub)

**Endpoint:** `POST /api/sub`

Realiza a subtração de dois números.

**Request Body:**

```json
{
    "numA": 10,
    "numB": 5
}
```

**Response (201 Created):**

```json
{
    "result": 5
}
```

### 3. Multiplicação (Mult)

**Endpoint:** `POST /api/mult`

Realiza a multiplicação de dois números.

**Request Body:**

```json
{
    "numA": 10,
    "numB": 5
}
```

**Response (201 Created):**

```json
{
    "result": 50
}
```

### 4. Divisão (Div)

**Endpoint:** `POST /api/div`

Realiza a divisão de dois números. Valida se o divisor é zero.

**Request Body:**

```json
{
    "numA": 10,
    "numB": 5
}
```

**Response (201 Created):**

```json
{
    "result": 2
}
```

**Response com Divisão por Zero (400 Bad Request):**

```json
{
    "error": "divisão por zero não é permitida"
}
```

## 📝 Exemplos de Uso

### Com cURL

```bash
# Adição
curl -X POST http://localhost:8080/api/sum \
  -H "Content-Type: application/json" \
  -d '{"numA": 10, "numB": 5}'

# Subtração
curl -X POST http://localhost:8080/api/sub \
  -H "Content-Type: application/json" \
  -d '{"numA": 10, "numB": 5}'

# Multiplicação
curl -X POST http://localhost:8080/api/mult \
  -H "Content-Type: application/json" \
  -d '{"numA": 10, "numB": 5}'

# Divisão
curl -X POST http://localhost:8080/api/div \
  -H "Content-Type: application/json" \
  -d '{"numA": 10, "numB": 5}'
```

### Com JavaScript (Fetch API)

```javascript
// Adição
fetch("http://localhost:8080/api/sum", {
    method: "POST",
    headers: {
        "Content-Type": "application/json",
    },
    body: JSON.stringify({
        numA: 10,
        numB: 5,
    }),
})
    .then((response) => response.json())
    .then((data) => console.log(data))
    .catch((error) => console.error("Erro:", error));
```

### Com Python

```python
import requests

url = 'http://localhost:8080/api/sum'
payload = {
    'numA': 10,
    'numB': 5
}

response = requests.post(url, json=payload)
print(response.json())  # {'result': 15}
```

## 🗂️ Estrutura do Projeto

```
math-api/
├── cmd/
│   └── api/
│       └── main.go              # Ponto de entrada da aplicação
├── internal/
│   ├── handlers/
│   │   └── math.go              # Implementação dos handlers de operações
│   ├── services/
│   │   └── math.go              # Lógica de negócio das operações matemáticas
│   └── models/
│       └── math.go              # Modelo de dados (Operacao)
├── go.mod                        # Definição do módulo Go
├── makefile                      # Comandos de automação
└── README.md                     # Este arquivo
```

### Detalhes dos Diretórios

- **cmd/api/**: Contém o código de entrada principal da aplicação
    - Configura o multiplexador HTTP
    - Define as rotas para cada operação
    - Inicia o servidor na porta 8080

- **internal/handlers/**: Contém os handlers (controladores) das requisições HTTP
    - Responsável por receber e validar os dados de entrada
    - Chama a camada de serviços para executar a operação
    - Retorna o resultado em formato JSON

- **internal/services/**: Contém a lógica de negócio das operações matemáticas
    - Implementa as funções `Sum()`, `Sub()`, `Mult()` e `Div()`
    - Separa a lógica de negócio dos handlers
    - Facilita testes unitários das operações

````

Executa a API em modo de desenvolvimento usando `go run`.

### Compilar a aplicação

```bash
make build
````

Compila o projeto gerando um executável chamado `math-api`.

## 🔄 Fluxo de Funcionamento

1. **Cliente** envia uma requisição POST para um dos endpoints (`/api/sum`, `/api/sub`, `/api/mult`, `/api/div`)
2. **Handler** recebe a requisição e decodifica o JSON do corpo
3. **Validação** do formato do JSON (se inválido, retorna erro 400)
4. **Services** executa a operação matemática com os valores fornecidos
5. **Resposta** com o resultado em formato JSON (status 201 Created)

### Arquitetura em Camadas

O projeto segue o padrão de arquitetura em 3 camadas:

```
Cliente HTTP
     ↓
  Handler (Validação & Roteamento)
     ↓
  Services (Lógica de Negócio)
     ↓
  Models (Estruturas de Dados)
```

Essa separação oferece:

- ✅ **Testabilidade** - Services podem ser testados independentemente
- ✅ **Manutenibilidade** - Lógica isolada em módulos específicos
- ✅ **Reutilização** - Services podem ser usados em diferentes contextos
- ✅ **Escalabilidade** - Fácil adicionar novas operações

## 🔍 Tratamento de Erros

Caso o corpo da requisição seja inválido ou mal formatado, a API retorna:

**Status:** 400 Bad Request

**Response:**

```json
{
    "error": "body invalido"
}
```

### Cenários de Erro

- Enviar JSON malformado
- Omitir os campos `numA` ou `numB`
- Usar tipos de dados não suportados (não números)
- **Divisão por zero** - A API retorna erro 400 com a mensagem: "divisão por zero não é permitida"

## 💡 Melhorias Futuras

- [x] Separar lógica de negócio em camada de services
- [x] Adicionar validação para divisão por zero
- [ ] Implementar logs estruturados (com `slog` ou `logrus`)
- [ ] Adicionar middleware de CORS
- [ ] Implementar autenticação/autorização
- [ ] Criar testes unitários para services e handlers
- [ ] Adicionar documentação OpenAPI/Swagger
- [ ] Implementar rate limiting
- [ ] Adicionar suporte para operações mais complexas (potência, raiz, etc)
- [ ] Containerizar a aplicação (Docker)
- [ ] Implementar health checks
- [ ] Adicionar tratamento de metricas (Prometheus)

## 📦 Dependências

A aplicação utiliza **apenas a biblioteca padrão do Go**, sem dependências externas:

- `net/http` - Para criar o servidor HTTP
- `encoding/json` - Para codificar/decodificar JSON
- `log` - Para logging básico

## 🧪 Testabilidade

Graças à separação em camada de services, é fácil testar a lógica de negócio de forma isolada:

```go
package services

import "testing"

func TestSum(t *testing.T) {
    result := Sum(10, 5)
    if result != 15 {
        t.Errorf("Expected 15, got %v", result)
    }
}

func TestDiv(t *testing.T) {
    result, err := Div(10, 5)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if result != 2 {
        t.Errorf("Expected 2, got %v", result)
    }
}

func TestDivByZero(t *testing.T) {
    _, err := Div(10, 0)
    if err == nil {
        t.Error("Expected error for division by zero")
    }
}
```

Os testes podem ser executados sem iniciar o servidor HTTP, tornando o desenvolvimento mais rápido e seguro!

### Para adicionar novas operações

1. **Adicione a função em `internal/services/math.go`:**

```go
func Power(base, exp float64) float64 {
    // Implementar a lógica da operação
    return math.Pow(base, exp)
}
```

2. **Crie o handler em `internal/handlers/math.go`:**

```go
func Power(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var operacao models.Operacao
    if err := json.NewDecoder(r.Body).Decode(&operacao); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{
            "error": "body invalido",
        })
        return
    }

    result := services.Power(operacao.NumA, operacao.NumB)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]float64{
        "result": result,
    })
}
```

3. **Registre a rota em `cmd/api/main.go`:**

```go
mux.HandleFunc("POST /api/power", handlers.Power)
```

**Benefício:** A lógica fica isolada em services, facilitando testes e manutenção!

## 📄 Licença

Este projeto está disponível sob a licença MIT.

## 👥 Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests.

## 📞 Suporte

Para dúvidas ou problemas, por favor abra uma issue no repositório do projeto.

---

**Versão:** 1.2.0  
**Última atualização:** Maio de 2026

### Changelog

#### v1.2.0

- 🔒 Implementação de validação para divisão por zero na camada de services
- 🛡️ Tratamento de erros robusto no handler de divisão
- 📝 Retorno de mensagens de erro dinâmicas

#### v1.1.0

- ✨ Refatoração da arquitetura com separação em camada de services
- 📦 Melhor organização do código seguindo padrão de 3 camadas
- 🧪 Facilita testes unitários da lógica de negócio

#### v1.0.0

- 🚀 Release inicial da Math API
- 4 operações matemáticas básicas (Sum, Sub, Mult, Div)
