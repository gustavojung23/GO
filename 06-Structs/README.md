# 06 - Structs

## 🏗️ Introdução às Structs em Go
Em Go, **structs** são tipos compostos que permitem agrupar diferentes tipos de dados sob um único nome. Elas são fundamentais para a modelagem de dados complexos e para a criação de estruturas mais organizadas e reutilizáveis no seu código. Com structs, você pode representar entidades do mundo real de maneira mais eficiente e clara.

## 🛠 Exemplos de Código
A seguir, um exemplo simples que demonstra como definir e utilizar structs em Go.

### Código de Exemplo

```go
package main

import "fmt"

// Definição da struct Endereco
type Endereco struct {
    Rua    string
    Numero uint8
}

// Definição da struct Usuario que inclui uma struct Endereco
type Usuario struct {
    Nome     string
    Idade    uint8
    Endereco Endereco
}

func main() {
    var u Usuario
    u.Nome = "John"
    u.Idade = 18
    u.Endereco = Endereco{
        Rua:    "Rua Flor",
        Numero: 12,
    }

    fmt.Println(u)
}
```

## Explicação do Código

### 1. Declaração das Structs
```go
type Endereco struct {
    Rua    string
    Numero uint8
}

type Usuario struct {
    Nome     string
    Idade    uint8
    Endereco Endereco
}
```

- **Endereco**: Representa um endereço com campos ``Rua`` (string) e ``Numero`` (uint8).
- **Usuario**: Representa um usuário com campos ``Nome`` (string), ``Idade`` (uint8) e ``Endereco`` (do tipo ``Endereco``).

### 2. Inicialização e Atribuição de Valores
```go
func main() {
    var u Usuario
    u.Nome = "John"
    u.Idade = 18
    u.Endereco = Endereco{
        Rua:    "Rua Flor",
        Numero: 12,
    }

    fmt.Println(u)
}
```

- **Declaração da Variável** ``u``: Cria uma variável ``u`` do tipo ``Usuario``.
- **Atribuição de Campos**:
    - ``u.Nome = "John"``: Define o nome do usuário.
    - ``u.Idade = 18``: Define a idade do usuário.
    - ``u.Endereco = Endereco{...}``: Inicializa o campo Endereco com os valores especificados.
- **Impressão da Struct**: ``fmt.Println(u)`` exibe todos os campos da struct ``Usuario``.

### 3. Saída Esperada
```
{John 18 {Rua Flor 12}}
```

## 🔍 Conceitos Importantes

### 1. Definição de Structs
Structs permitem agrupar dados relacionados em um único tipo composto, facilitando a manipulação e a organização das informações.

```go
type Pessoa struct {
    Nome  string
    Idade int
}
```

### 2. Inicialização de Structs
Existem várias maneiras de inicializar structs em Go:

- **Declaração e Atribuição Posterior**
```go
var p Pessoa
p.Nome = "Alice"
p.Idade = 30
```

- **Inicialização Composta**
```go
p := Pessoa{
    Nome:  "Bob",
    Idade: 25,
}
```

- **Inicialização Anônima**
```go
p := Pessoa{"Charlie", 40}
```

### 3. Structs Aninhadas
Structs podem conter outras structs, permitindo a criação de estruturas de dados complexas.
```go
type Endereco struct {
    Rua    string
    Numero int
}

type Empresa struct {
    Nome     string
    Endereco Endereco
}
```

### 4. Métodos em Structs
Go permite associar métodos a structs, tornando-os mais funcionais e orientados a objetos.
```go
func (p Pessoa) Saudacao() {
    fmt.Printf("Olá, meu nome é %s.\n", p.Nome)
}
```

### 5. Structs como Referências
Embora Go seja uma linguagem de passagem por valor, você pode usar ponteiros para structs para modificar os dados originais.
```go
func (p *Pessoa) Aniversario() {
    p.Idade += 1
}
```

## 🛡 Boas Práticas

- **Nomenclatura Clara**

Utilize nomes de structs que representem claramente a entidade ou conceito que estão modelando.

```go
Copiar código
type Produto struct {
    Nome  string
    Preco float64
}
```

- **Exportação de Campos Necessários**

Exporte apenas os campos que precisam ser acessados de outros pacotes, mantendo a encapsulação.

```go
type Usuario struct {
    Nome     string // Exportado
    idade    int    // Não exportado
}
```

- **Documentação dos Campos**

Comente os campos das structs para melhorar a legibilidade e facilitar a manutenção.

```go
type Endereco struct {
    Rua    string // Nome da rua
    Numero int    // Número da residência
}
```

- **Evite Structs Excessivamente Grandes**

Mantenha as structs focadas em uma única responsabilidade para facilitar a reutilização e a testabilidade.