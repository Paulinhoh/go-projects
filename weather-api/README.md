# 🌦️ Weather API - Aplicação Go

Uma aplicação em Go que consome dados de clima de uma API de terceiros e implementa cache com Redis para otimizar as requisições.

![img](https://assets.roadmap.sh/guest/weather-api-f8i1q.png)

## 📋 Índice

- [Visão Geral](#visão-geral)
- [Funcionalidades](#funcionalidades)
- [Arquitetura](#arquitetura)
- [Tecnologias](#tecnologias)
- [Pré-requisitos](#pré-requisitos)
- [Instalação](#instalação)
- [Configuração](#configuração)
- [Como Usar](#como-usar)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Fluxo de Funcionamento](#fluxo-de-funcionamento)
- [Exemplos de Uso](#exemplos-de-uso)
- [Melhorias Futuras](#melhorias-futuras)

---

## 🎯 Visão Geral

Weather API é uma aplicação CLI desenvolvida em Go que busca informações meteorológicas de uma cidade usando a API de terceiros **Visual Crossing**. A aplicação implementa um sistema de cache com **Redis** para armazenar resultados por 2 minutos, evitando requisições desnecessárias à API externa.

---

## ✨ Funcionalidades

- 🔍 **Busca de Clima**: Consulta dados meteorológicos de qualquer cidade
- 💾 **Cache Inteligente**: Armazena resultados em Redis por 2 minutos
- ⚡ **Performance**: Retorna dados em cache antes de fazer requisições desnecessárias
- 🌍 **Timezone Suportado**: Retorna informações de fuso horário da cidade
- 📝 **Descrição do Clima**: Fornece descrição detalhada do clima atual
- 🕐 **Timestamp**: Registra o momento da busca

---

## 🏗️ Arquitetura

```
                    ┌─────────────────┐
                    │   Usuário/CLI   │
                    └────────┬────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │  Weather API    │
                    │  (Go App)       │
                    └────────┬────────┘
                    ┌────────┴────────┐
                    │                 │
                    ▼                 ▼
            ┌──────────────┐  ┌──────────────┐
            │ Redis Cache  │  │ Visual Cross │
            │   (2 min)    │  │  Weather API │
            └──────────────┘  └──────────────┘
```

**Fluxo de Requisição:**

1. ✅ Verifica se dados estão em cache no Redis
2. 📦 Se em cache: retorna dados armazenados
3. 🌐 Se não em cache: faz requisição à API Visual Crossing
4. 💾 Armazena resultado no Redis com TTL de 2 minutos
5. 📤 Retorna os dados para o usuário

---

## 💻 Tecnologias

| Tecnologia              | Versão  | Descrição                              |
| ----------------------- | ------- | -------------------------------------- |
| **Go**                  | 1.26.1  | Linguagem de programação               |
| **Redis**               | v9.20.0 | Cache em memória                       |
| **godotenv**            | v1.5.1  | Gerenciamento de variáveis de ambiente |
| **Visual Crossing API** | -       | API de dados meteorológicos            |

---

## 📋 Pré-requisitos

Antes de iniciar, certifique-se de ter instalado:

- ✅ **Go 1.26.1+** - [Download](https://golang.org/dl/)
- ✅ **Redis** - [Download](https://redis.io/download)
- ✅ **Git** (opcional)
- ✅ **Chave API** - [Obter em Visual Crossing](https://www.visualcrossing.com/)

### Verificar Instalação

```bash
# Verificar Go
go version

# Verificar Redis (servidor rodando)
redis-cli ping
# Deve retornar: PONG
```

---

## 🚀 Instalação

### 1. Clonar o Repositório

```bash
git clone https://github.com/seu-usuario/weather-api.git
cd weather-api
```

### 2. Instalar Dependências

```bash
go mod download
go mod tidy
```

### 3. Iniciar o Redis

```bash
# No Windows (se instalado via WSL ou Docker)
redis-server

# Ou via Docker
docker run -d -p 6379:6379 redis
```

---

## ⚙️ Configuração

### 1. Criar Arquivo `.env`

Copie o arquivo `.env.example` para `.env`:

```bash
cp .env.example .env
```

### 2. Preencher Variáveis de Ambiente

Edite o arquivo `.env` com suas credenciais:

```env
API_KEY="sua_api_key_aqui"
REDIS="localhost:6379"
```

**Como obter a API Key:**

1. Acesse [Visual Crossing Weather API](https://www.visualcrossing.com/)
2. Registre uma conta (versão gratuita disponível)
3. Copie sua API Key
4. Cole no arquivo `.env`

**Configurar Redis:**

- **Localmente**: `localhost:6379` (padrão)
- **Remoto**: `seu-redis-host:porta`
- **Docker**: `redis-container:6379` or `redis://redis:6379`

---

## 📖 Como Usar

### Via Go Run

```bash
go run cmd/api/main.go -s "nome_da_cidade"
```

### Via Build (Executável)

```bash
# Build do projeto
make build
# Ou manualmente
go build ./cmd/api/.

# Executar
./api -s "nome_da_cidade"
```

### Exemplos de Comando

```bash
# Buscar clima no Rio de Janeiro
go run cmd/api/main.go -s "rio de janeiro"

# Buscar clima em São Paulo
go run cmd/api/main.go -s "são paulo"

# Buscar clima em cidades com espaços
go run cmd/api/main.go -s "new york"

# Buscar clima em cidades com acentos
go run cmd/api/main.go -s "são paulo"
```

### Saída Esperada

```
[29/05/2026-14:30:45] - Address: Rio de Janeiro, Brazil | TimeZone: America/Sao_Paulo | Description: Rainy
```

---

## 📁 Estrutura do Projeto

```
weather-api/
├── cmd/
│   └── api/
│       └── main.go              # Ponto de entrada da aplicação
├── internal/
│   ├── cache/
│   │   └── redis.go             # Lógica de cache com Redis
│   ├── handlers/
│   │   └── weather.go           # Handler principal (busca e cache)
│   └── models/
│       └── weather.go           # Structs de dados
├── .env                         # Variáveis de ambiente (não versionar)
├── .env.example                 # Template de variáveis
├── .gitignore                   # Arquivos ignorados no Git
├── go.mod                       # Dependências Go
├── go.sum                       # Hash das dependências
├── makefile                     # Comandos de build e execução
└── README.md                    # Este arquivo
```

### Descrição dos Arquivos

| Arquivo                 | Descrição                                                |
| ----------------------- | -------------------------------------------------------- |
| `main.go`               | Entrada da aplicação, processa flags de linha de comando |
| `weather.go` (handlers) | Função principal que orquestra cache e requisições       |
| `redis.go`              | Funções `Get()` e `Set()` para operações de cache        |
| `weather.go` (models)   | Definição da struct `Weather` com campos JSON            |

---

## 🔄 Fluxo de Funcionamento

### Sequência de Execução

```
1. Usuário executa: go run cmd/api/main.go -s "rio de janeiro"
                    │
2. main.go processa flag "-s"
                    │
3. handlers.GetWeather("rio de janeiro") é chamado
                    │
4. Carrega variáveis de .env
                    │
5. Verifica cache Redis
    ├─ Encontrado? ✓ Retorna dados em cache
    └─ Não encontrado? Continue...
                    │
6. Faz requisição HTTP à Visual Crossing API
                    │
7. Verifica StatusCode 200
    ├─ Sucesso? ✓ Decodifica JSON
    └─ Erro? ✗ Log error e para
                    │
8. Armazena resultado no Redis (TTL: 2 min)
                    │
9. Retorna dados para main.go
                    │
10. Formata e exibe resultado
```

### Tratamento de Cache

```
Cache Hit (Acerto):
  Redis: GET weather:rio+de+janeiro
         └─> Encontrado ✓
             └─> Retorna após 0ms

Cache Miss (Erro):
  Redis: GET weather:rio+de+janeiro
         └─> Não encontrado
             └─> HTTP Request à API
                 └─> Redis: SET weather:rio+de+janeiro (2 min)
                     └─> Retorna após ~500ms
```

---

## 💡 Exemplos de Uso

### Exemplo 1: Primeira Busca (Sem Cache)

```bash
$ go run cmd/api/main.go -s "tokyo"
[29/05/2026-14:35:20] - Address: Tokyo, Japan | TimeZone: Asia/Tokyo | Description: Partly cloudy
# Tempo: ~600ms (requisição à API)
```

### Exemplo 2: Segunda Busca (Com Cache)

```bash
$ go run cmd/api/main.go -s "tokyo"
[29/05/2026-14:35:25] - Address: Tokyo, Japan | TimeZone: Asia/Tokyo | Description: Partly cloudy
# Tempo: ~50ms (retornou do cache)
```

### Exemplo 3: Busca Diferente

```bash
$ go run cmd/api/main.go -s "paris"
[29/05/2026-14:36:00] - Address: Paris, France | TimeZone: Europe/Paris | Description: Sunny
# Tempo: ~600ms (nova requisição à API)
```

---

## 🛠️ Comandos Úteis

### Development

```bash
# Executar a aplicação
make run

# Build do projeto
make build

# Executar com argumentos
go run cmd/api/main.go -s "sua_cidade"

# Testar Redis localmente
redis-cli
# > SET weather:test "dados"
# > GET weather:test
```

### Debugging

```bash
# Verificar conexão Redis
redis-cli ping

# Ver todas as chaves armazenadas
redis-cli KEYS "*"

# Ver tempo de vida de uma chave
redis-cli TTL weather:rio+de+janeiro

# Limpar cache
redis-cli FLUSHDB
```

---

## 📊 Resposta da API

A aplicação retorna um objeto `Weather` com:

```go
type Weather struct {
    Address     string    `json:"resolvedAddress"`  // Endereço resolvido
    Timezone    string    `json:"timezone"`         // Fuso horário
    Description string    `json:"description"`      // Descrição do clima
    Timestamp   time.Time                           // Hora da busca
}
```

**Exemplo de JSON em Cache:**

```json
{
    "resolvedAddress": "Rio de Janeiro, Brazil",
    "timezone": "America/Sao_Paulo",
    "description": "Rainy",
    "Timestamp": "2026-05-29T14:30:45Z"
}
```

---

## 🚨 Tratamento de Erros

A aplicação trata os seguintes cenários:

| Erro                              | Causa                   | Solução                              |
| --------------------------------- | ----------------------- | ------------------------------------ |
| "o parametro não pode ser vazio!" | Sem flag `-s`           | `go run cmd/api/main.go -s "cidade"` |
| "Erro ao carregar .env"           | Arquivo `.env` faltando | Criar `.env` com variáveis           |
| "Failed to connect redis"         | Redis não está rodando  | `redis-server` ou Docker             |
| "API retornou status code: 401"   | API Key inválida        | Verificar credenciais em `.env`      |
| "API retornou status code: 404"   | Cidade não encontrada   | Verificar nome da cidade             |

---

## 🔐 Segurança

### Boas Práticas Implementadas

✅ Variáveis sensíveis em `.env`  
✅ `.env` não versionado no Git  
✅ Validação de parâmetros  
✅ Tratamento de erros

### Recomendações Adicionais

- 🔐 Não compartilhe sua `API_KEY`
- 🔒 Use ambiente seguro para Redis em produção
- 📝 Implemente autenticação Redis com senha
- 🔍 Valide dados de entrada antes de usar

---

## 🐛 Troubleshooting

### Problema: "Erro ao carregar .env"

**Solução:**

```bash
# Certifique-se que .env existe na raiz do projeto
ls -la | grep .env

# Se não existir, copie do .env.example
cp .env.example .env
```

### Problema: "Failed to connect in the redis instance"

**Solução:**

```bash
# Verifique se Redis está rodando
redis-cli ping
# Deve retornar PONG

# Se não estiver rodando:
redis-server  # Linux/Mac
# Ou use Docker:
docker run -d -p 6379:6379 redis
```

### Problema: "API retornou status code: 401"

**Solução:**

```bash
# API Key inválida
# 1. Verifique se a chave está correta em .env
# 2. Obtenha uma nova chave em: https://www.visualcrossing.com/
# 3. Atualize o arquivo .env
```

### Problema: Variáveis de ambiente não carregam

**Solução:**

```bash
# Certifique-se que está no diretório correto
pwd
# Deve estar em: weather-api/

# Verifique o formato do .env
cat .env
# Deve ter: API_KEY="seu_valor"
# Deve ter: REDIS="localhost:6379"
```

---

## 📈 Melhorias Futuras

Sugestões de melhorias para o projeto:

- [ ] **API REST**: Converter para HTTP server com endpoints
- [ ] **Múltiplas Cidades**: Buscar clima de várias cidades em paralelo
- [ ] **Banco de Dados**: Persistir histórico de buscas
- [ ] **Testes Unitários**: Adicionar testes com `testing` package
- [ ] **Docker**: Containerizar a aplicação e Redis
- [ ] **Logging**: Sistema de logs estruturado (zap, logrus)
- [ ] **Configuração**: Suporte a múltiplos ambientes (dev, test, prod)
- [ ] **Métricas**: Prometheus metrics para monitoramento
- [ ] **Tratamento Avançado**: Retry logic e circuit breaker
- [ ] **Documentação**: Swagger/OpenAPI se virar REST API
- [ ] **CI/CD**: GitHub Actions para testes e build

---

## 📚 Recursos Úteis

### Documentações

- [Go Documentation](https://golang.org/doc/)
- [Redis Documentation](https://redis.io/docs/)
- [go-redis Package](https://github.com/redis/go-redis)
- [godotenv Package](https://github.com/joho/godotenv)
- [Visual Crossing API](https://www.visualcrossing.com/resources/documentation/weather-api/timeline-weather-api/)

### Tutoriais

- [Go Getting Started](https://golang.org/doc/tutorial/getting-started)
- [Redis with Go](https://redis.io/docs/clients/go/)

---

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.

---

## 👨‍💻 Contribuição

Contribuições são bem-vindas! Para contribuir:

1. Faça um Fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

---

## 📞 Suporte

Se encontrar problemas ou tiver dúvidas:

- 📧 Abra uma issue no repositório
- 💬 Consulte a seção [Troubleshooting](#troubleshooting)
- 🔍 Verifique os [Exemplos de Uso](#exemplos-de-uso)

---

## ✨ Créditos

- **API de Clima**: [Visual Crossing](https://www.visualcrossing.com/)
- **Cache**: [Redis](https://redis.io/)
- **Linguagem**: [Go](https://golang.org/)

---

**Última atualização**: 30 de Maio de 2026  
**Versão**: 1.0.0
