# rphpctl

CLI (Command Line Interface) escrita em Go com [Cobra](https://github.com/spf13/cobra) para interagir com microserviços pessoais via terminal. Funciona como um cliente de linha de comando para consumir APIs REST de três serviços:

- **IDP** — Serviço de autenticação (geração de tokens JWT)
- **Finance Manager** — Gestão financeira (contas, transferências, métodos de pagamento)
- **Password Vault** — Cofre de senhas (credenciais e códigos)

## Setup

### Pré-requisitos

- [Go](https://golang.org/) 1.24.2+

### Instalação

```bash
git clone https://github.com/ropehapi/rphpctl.git
cd rphpctl
go mod download
```

### Configuração

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

```env
KAIZEN_AUTH_SERVICE_HOST=http://localhost
KAIZEN_AUTH_SERVICE_PORT=8080

FINANCE_MANAGER_HOST=http://localhost
FINANCE_MANAGER_PORT=8081

PASSWORD_VAULT_HOST=http://localhost
PASSWORD_VAULT_PORT=8082

BEARER_TOKEN=seu_token_jwt_aqui
```

Opcionalmente, você pode criar um arquivo de configuração em `$HOME/.rphpctl.yaml`.

### Build

```bash
go build -o rphpctl .
```

### Execução

```bash
./rphpctl [comando] [subcomando] [flags]
```

## Comandos disponíveis

### Flags globais

| Flag | Shorthand | Descrição |
|---|---|---|
| `--config` | | Arquivo de configuração (padrão: `$HOME/.rphpctl.yaml`) |
| `--toggle` | `-t` | Help message for toggle |

---

### `idp` — Funções do IDP

| Subcomando | Descrição | Flags |
|---|---|---|
| `login` | Gera um token JWT no IDP | `-u, --username` (usuário), `-p, --password` (senha) |

**Exemplo:**

```bash
rphpctl idp login -u meu_usuario -p minha_senha
```

---

### `finance-manager` — Funções do Finance Manager

#### Contas

| Subcomando | Descrição | Flags |
|---|---|---|
| `create-account` | Cria uma conta | `-c, --currency` (moeda), `-b, --balance` (saldo), `-n, --name` (nome) |
| `get-accounts` | Lista contas | `-n, --name` (nome), `-c, --currency` (moeda), `-i, --id` (id) |
| `update-account` | Atualiza uma conta | `-i, --id` (id), `-c, --currency` (moeda), `-b, --balance` (saldo), `-n, --name` (nome) |
| `delete-account` | Deleta uma conta | `-i, --id` (id) |

**Exemplos:**

```bash
rphpctl finance-manager create-account -c BRL -b 1000 -n "Conta Corrente"
rphpctl finance-manager get-accounts
rphpctl finance-manager get-accounts -i abc123
rphpctl finance-manager update-account -i abc123 -n "Conta Poupança" -b 2000
rphpctl finance-manager delete-account -i abc123
```

#### Transferências

| Subcomando | Descrição | Flags |
|---|---|---|
| `create-cashin` | Cria uma entrada (cashin) | `--currency` (moeda), `-a, --amount` (valor), `-d, --description` (descrição), `-c, --category` (categoria), `--date` (data), `-o, --observations` (observações), `--accountId` (conta destino) |
| `create-cashout` | Cria uma saída (cashout) | `--currency` (moeda), `-a, --amount` (valor), `-d, --description` (descrição), `-c, --category` (categoria), `--date` (data), `-o, --observations` (observações), `-p, --paymentMethodId` (método de pagamento) |
| `get-transfers` | Lista transferências | `-t, --type` (tipo: cashin, cashout ou debt_payment), `-c, --category` (categoria), `-i, --id` (id) |
| `delete-transfers` | Deleta uma transferência | `-i, --id` (id) |

**Exemplos:**

```bash
rphpctl finance-manager create-cashin -a 500 -d "Salário" -c "renda" --date "2025-01-15"
rphpctl finance-manager create-cashout -a 50 -d "Almoço" -c "alimentação" -p method123
rphpctl finance-manager get-transfers
rphpctl finance-manager get-transfers -t cashin -c "renda"
rphpctl finance-manager delete-transfers -i abc123
```

#### Métodos de Pagamento

| Subcomando | Descrição | Flags |
|---|---|---|
| `create-payment-method` | Cria um método de pagamento | `-a, --accountId` (id da conta), `-t, --type` (tipo), `-n, --name` (nome) |
| `get-payment-methods` | Lista métodos de pagamento | `-n, --name` (nome), `-c, --currency` (moeda), `-i, --id` (id) |
| `update-payment-method` | Atualiza um método de pagamento | `-i, --id` (id), `-n, --name` (nome) |
| `delete-payment-method` | Deleta um método de pagamento | `-i, --id` (id) |

**Exemplos:**

```bash
rphpctl finance-manager create-payment-method -a acc123 -t "credit_card" -n "Nubank"
rphpctl finance-manager get-payment-methods
rphpctl finance-manager update-payment-method -i pm123 -n "Nubank Platinum"
rphpctl finance-manager delete-payment-method -i pm123
```

---

### `password-vault` — Funções do Password Vault

#### Contas (login/senha)

| Subcomando | Descrição | Flags |
|---|---|---|
| `create` | Cria um par login/senha | `-n, --name` (nome da conta), `-l, --login` (login), `-p, --password` (senha) |
| `get` | Lista pares login/senha | `-n, --name` (nome da conta) |
| `update` | Atualiza um par login/senha | `-i, --id` (id), `-n, --name` (nome), `-l, --login` (login), `-p, --password` (senha) |
| `delete` | Deleta um par login/senha | `-i, --id` (id) |

**Exemplos:**

```bash
rphpctl password-vault create -n "GitHub" -l meu_user -p minha_senha
rphpctl password-vault get
rphpctl password-vault get -n "GitHub"
rphpctl password-vault update -i abc123 -n "GitHub" -l novo_user -p nova_senha
rphpctl password-vault delete -i abc123
```

#### Códigos de conta

| Subcomando | Descrição | Flags |
|---|---|---|
| `create-codes` | Cria códigos para uma conta | `-n, --name` (nome da conta), `-c, --codes` (códigos) |
| `get-codes` | Lista códigos de uma conta | `-n, --name` (nome da conta) |
| `update-codes` | Atualiza códigos de uma conta | `-i, --id` (id), `-n, --name` (nome), `-c, --codes` (códigos) |
| `delete-codes` | Deleta códigos de uma conta | `-i, --id` (id) |

**Exemplos:**

```bash
rphpctl password-vault create-codes -n "GitHub" -c "abc123,def456"
rphpctl password-vault get-codes -n "GitHub"
rphpctl password-vault update-codes -i abc123 -n "GitHub" -c "novo123,novo456"
rphpctl password-vault delete-codes -i abc123
```

---

## Variáveis de ambiente

| Variável | Descrição |
|---|---|
| `KAIZEN_AUTH_SERVICE_HOST` | Host do serviço de autenticação (IDP) |
| `KAIZEN_AUTH_SERVICE_PORT` | Porta do serviço de autenticação (IDP) |
| `FINANCE_MANAGER_HOST` | Host do Finance Manager |
| `FINANCE_MANAGER_PORT` | Porta do Finance Manager |
| `PASSWORD_VAULT_HOST` | Host do Password Vault |
| `PASSWORD_VAULT_PORT` | Porta do Password Vault |
| `BEARER_TOKEN` | Token JWT para autenticação nas APIs |
