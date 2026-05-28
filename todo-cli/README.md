# Todo CLI - Gerenciador de Tarefas em Linha de Comando

Uma aplicação simples e eficiente de gerenciamento de tarefas (TODO list) desenvolvida em **Go**, que funciona através de uma interface de linha de comando (CLI). O projeto armazena as tarefas em um arquivo JSON local, permitindo persistência de dados entre execuções.

---

## 📋 Características Principais

- ✅ **Criar tarefas**: Adicione novas tarefas à sua lista de afazeres
- 📝 **Listar tarefas**: Visualize todas as tarefas com seus detalhes
- ✓ **Marcar como completo**: Marque tarefas como concluídas
- 🗑️ **Deletar tarefas**: Remova tarefas que não são mais necessárias
- 💾 **Persistência de dados**: Todas as tarefas são salvas em arquivo JSON
- ⚡ **Leve e rápido**: Desenvolvido em Go para máxima performance
- 🔄 **ID automático**: Cada tarefa recebe um ID único e sequencial

---

## 🏗️ Estrutura do Projeto

```
todo-cli/
├── cmd/
│   └── main.go              # Ponto de entrada da aplicação
├── internal/
│   ├── models/
│   │   └── todo.go          # Definição da estrutura de dados Todo
│   ├── repository/
│   │   └── todo.go          # Gerenciamento de persistência (salvar/carregar)
│   └── services/
│       └── todo.go          # Lógica de negócio das operações
├── data/
│   └── todos.json           # Arquivo de armazenamento de dados
├── go.mod                   # Arquivo de módulo Go
├── makefile                 # Instruções de build
└── README.md                # Este arquivo

```

### Componentes Principais

#### 1. **Models** (`internal/models/todo.go`)

Define a estrutura de dados da tarefa:

```go
type Todo struct {
    ID          int       // Identificador único da tarefa
    Description string    // Descrição da tarefa
    Complete    bool      // Status de conclusão (true = concluída)
    CreatedAt   time.Time // Data e hora de criação
}
```

#### 2. **Repository** (`internal/repository/todo.go`)

Responsável pela persistência de dados:

- `LoadTodos()`: Carrega todas as tarefas do arquivo `dados/todos.json`
- `SaveTodo()`: Salva as tarefas no arquivo JSON
- Gerencia o slice em memória `Todos` que armazena as tarefas durante a execução

#### 3. **Services** (`internal/services/todo.go`)

Implementa a lógica de negócio:

- `CreateTodo(msg string)`: Cria uma nova tarefa
- `ListTodos()`: Exibe todas as tarefas formatadas
- `MarkCompleteTodo(id int)`: Marca uma tarefa como concluída
- `DeleteTodo(id int)`: Remove uma tarefa da lista

#### 4. **Main** (`cmd/main.go`)

Ponto de entrada que:

- Carrega as tarefas existentes
- Processa os argumentos da linha de comando
- Chama os serviços apropriados

---

## 🚀 Como Usar

### Pré-requisitos

- Go 1.26.1 ou superior instalado
- Acesso ao terminal/prompt de comando

### Instalação e Execução

**1. Clone ou baixe o projeto:**

```bash
cd todo-cli
```

**2. Compile o projeto (opcional):**

```bash
go build -o todo cmd/main.go
```

**3. Execute usando Go diretamente:**

```bash
go run cmd/main.go [flags]
```

---

## 🎯 Comandos Disponíveis

### Criar uma nova tarefa

```bash
go run cmd/main.go -c "Descrição da tarefa"
```

**Exemplo:**

```bash
go run cmd/main.go -c "Estudar Go"
go run cmd/main.go -c "Fazer compras no mercado"
```

**Resultado:** Uma nova tarefa é adicionada com ID sequencial e data/hora de criação automaticamente registradas.

---

### Listar todas as tarefas

```bash
go run cmd/main.go -l
```

**Exemplo de saída:**

```
ID: 1 | Descrição: ir para a academia | Completo: false | Criado em [28/05/2026-14:23:52]
ID: 2 | Descrição: ufs 19h - aula de gerencia de projeto | Completo: true | Criado em [28/05/2026-14:24:17]
ID: 3 | Descrição: prova: 05/07/2026 -> ESII | Completo: false | Criado em [28/05/2026-14:25:25]
ID: 4 | Descrição: prova: 25/07/2026 -> SO | Completo: false | Criado em [28/05/2026-14:25:40]
```

---

### Marcar tarefa como concluída

```bash
go run cmd/main.go -u <ID>
```

**Exemplo:**

```bash
go run cmd/main.go -u 1
```

**Resultado:**

```
ToDo atualizado com sucesso.
```

A tarefa com ID 1 agora terá `"complete": true` no arquivo JSON.

---

### Deletar uma tarefa

```bash
go run cmd/main.go -d <ID>
```

**Exemplo:**

```bash
go run cmd/main.go -d 2
```

**Resultado:**

```
toDo deletado com sucesso.
```

A tarefa com ID 2 é removida da lista permanentemente.

---

## 💾 Formato de Dados

Os dados são persistidos no arquivo `data/todos.json` em formato JSON:

```json
[
    {
        "id": 1,
        "description": "ir para a academia",
        "complete": false,
        "createdAt": "2026-05-28T14:23:52.5350897-03:00"
    },
    {
        "id": 2,
        "description": "ufs 19h - aula de gerencia de projeto",
        "complete": true,
        "createdAt": "2026-05-28T14:24:17.9907629-03:00"
    }
]
```

**Campos:**

- **id**: Número único sequencial que identifica a tarefa
- **description**: Texto descritivo da tarefa
- **complete**: Booleano indicando se a tarefa foi concluída
- **createdAt**: Timestamp ISO 8601 da criação da tarefa

---

## 🔧 Desenvolvimento

### Compilar e Executar

**Usando o Makefile:**

```bash
make run -c "Nova tarefa"
make run -l
```

**Compilar para um executável:**

```bash
go build -o todo cmd/main.go
./todo -c "Descrição"
./todo -l
```

### Estrutura de Pacotes

O projeto segue a arquitetura de camadas:

```
Models (dados)
    ↓
Repository (persistência)
    ↓
Services (lógica de negócio)
    ↓
Main (interface CLI)
```

---

## 📊 Fluxo de Operação

### Ao Iniciar a Aplicação

1. `LoadTodos()` carrega o arquivo JSON da pasta `dados/`
2. As tarefas são desserializadas para a estrutura `[]models.Todo`
3. A aplicação fica pronta para processar comandos

### Ao Modificar (criar, atualizar ou deletar)

1. A operação é realizada na memória no slice `Todos`
2. A função `SaveTodo()` é chamada para persistir as mudanças
3. O arquivo JSON é reescrito com os dados atualizados

### Ao Listar

1. Itera sobre todas as tarefas em memória
2. Formata e exibe cada tarefa no console

---

## ⚠️ Observações Importantes

- **Arquivo de Dados**: O arquivo `dados/todos.json` deve existir no diretório de execução
- **IDs Sequenciais**: Os IDs são baseados no comprimento do slice, então deletar e recriar pode resultar em reatribuição de IDs
- **Sincronização**: Não há sistema de lock, portanto não é recomendado executar múltiplas instâncias simultaneamente
- **Tratamento de Erros**: A aplicação imprime mensagens de erro, mas continua a execução
- **Caso Sensível**: Os flags da CLI são case-sensitive (`-c`, `-l`, `-u`, `-d`)

---

## 🎓 Conceitos Go Utilizados

- **Packages e Imports**: Organização modular do código
- **Structs e JSON Tags**: Serialização/desserialização de dados
- **File I/O**: Leitura e escrita de arquivos
- **CLI Flags**: Processamento de argumentos de linha de comando
- **Slices**: Coleções dinâmicas de dados
- **Defer**: Garantia de fechamento de recursos

---

## 📝 Exemplo de Fluxo Completo

```bash
# 1. Criar tarefas
go run cmd/main.go -c "Aprender Go"
go run cmd/main.go -c "Fazer projeto"
go run cmd/main.go -c "Revisar código"

# 2. Listar tarefas
go run cmd/main.go -l

# 3. Marcar como concluída
go run cmd/main.go -u 1

# 4. Deletar tarefa
go run cmd/main.go -d 3

# 5. Listar novamente para confirmar
go run cmd/main.go -l
```

---

## 📄 Licença

Este é um projeto educacional desenvolvido para fins de aprendizado com Go.

---
