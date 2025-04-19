# 🧠 Backend Residência Técnica

API RESTful desenvolvida em Go para gerenciamento de agentes, mensagens, ferramentas e tokens de uso. Utiliza GORM com PostgreSQL e está implantada na Railway.

---

## 🌐 URL Base da API

**Use esta URL:**

https://backend-residencia-production-ef55.up.railway.app/

---

## 🛠️ Tecnologias

- Go 1.22
- Gin (framework HTTP)
- GORM (ORM)
- PostgreSQL (via Neon)
- Railway (deploy)
- Docker (ambiente de desenvolvimento)

---

## 📦 Estrutura do Banco de Dados

O sistema possui as seguintes entidades principais:

- **Agents**
- **Messages**
- **UsageTokens**
- **Tools**
- **AgentTools**
- **OutputStructures**
- **AgentOutputStructures**

---

## 🔗 Rotas da API

### 🔹 Agentes

| Método | Rota              | Descrição                   |
|--------|-------------------|-----------------------------|
| GET    | `/agents`         | Listar todos os agentes     |
| GET    | `/agents/:id`     | Obter agente por ID         |
| POST   | `/agents`         | Criar novo agente           |
| PUT    | `/agents/:id`     | Atualizar agente existente  |
| DELETE | `/agents/:id`     | Deletar agente              |

### 🔹 Mensagens

| Método | Rota              | Descrição                     |
|--------|-------------------|-------------------------------|
| GET    | `/messages`       | Listar todas as mensagens     |
| GET    | `/messages/:id`   | Obter mensagem por ID         |
| POST   | `/messages`       | Criar nova mensagem           |
| PUT    | `/messages/:id`   | Atualizar mensagem existente  |
| DELETE | `/messages/:id`   | Deletar mensagem              |

### 🔹 Tokens de Uso

| Método | Rota                  | Descrição                          |
|--------|-----------------------|------------------------------------|
| GET    | `/usage-tokens`       | Listar todos os tokens             |
| GET    | `/usage-tokens/:id`   | Obter token por ID                 |
| POST   | `/usage-tokens`       | Criar token para mensagem assistant|
| PUT    | `/usage-tokens/:id`   | Atualizar token                    |
| DELETE | `/usage-tokens/:id`   | Deletar token                      |

### 🔹 Ferramentas e Estruturas

| Método | Rota                         | Descrição                         |
|--------|------------------------------|-----------------------------------|
| GET    | `/tools`                     | Listar ferramentas                |
| GET    | `/output-structures`         | Listar estruturas de output       |
| GET    | `/agent-tools`               | Listar vínculos agente-ferramenta |
| GET    | `/agent-output-structures`   | Listar vínculos agente-estrutura  |

---

## 🧪 Testes

Utilize ferramentas como:

- [Postman](https://postman.com)
- [Insomnia](https://insomnia.rest)
- `curl` no terminal:

```bash
curl https://backend-residencia-production-ef55.up.railway.app/agents
