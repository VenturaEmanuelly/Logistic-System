
### Sistema de Gerenciamento de Produtos

Sistema de gerenciamento de produtos desenvolvido em Go, utilizando o framework Fiber para expor uma API RESTful e PostgreSQL como banco de dados. O sistema oferece funcionalidades para manipular informações de produtos e suas métricas associadas.

#### Funcionalidades

- **Gerenciamento de Produtos:** Operações CRUD (Criar, Ler, Atualizar, Deletar) para produtos.
- **Métricas de Produtos:** Registro e cálculo das quantidades totais e movimentações de produtos.

#### Arquitetura

- **Entidades:** Definições das estruturas de dados (`Product`, `TotalProduct`).
- **Repositórios:** Implementações de acesso ao banco de dados para operações sobre produtos e métricas.
- **Casos de Uso:** Lógica de negócios para gerenciar produtos e métricas.
- **Manipuladores:** Endpoints HTTP que interagem com a lógica de negócios e retornam respostas apropriadas.
- **Configuração e Inicialização:** Configuração do servidor e injeção de dependências.

#### Tecnologias

- **Go:** Linguagem de programação.
- **Fiber:** Framework para desenvolvimento de APIs.
- **PostgreSQL:** Banco de dados relacional.
- **github.com/lib/pq:** Driver PostgreSQL para Go.

#### Como Usar

1. **Instale Dependências:**

   ```bash
   go mod tidy
   ```

2. **Configure o Banco de Dados:** Atualize a string de conexão no arquivo `main.go`.

3. **Execute o Servidor:**

   ```bash
   go run main.go
   ```

4. **Endpoints API:**

   - **POST /**: Cria um novo produto.
   - **GET /**: Recupera um produto.
   - **PUT /**: Atualiza um produto.
   - **DELETE /**: Remove um produto.


