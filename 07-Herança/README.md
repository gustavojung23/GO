# 07 - Herança 🦴

## 📖 Visão Geral

Go é uma linguagem que promove a **composição** em vez da herança tradicional encontrada em linguagens orientadas a objetos como Java ou C++. Embora Go não suporte herança por meio de classes, ele permite que você reutilize e estenda funcionalidades usando **composição de structs** e **interfaces**. Este módulo explora como simular comportamentos de herança em Go através da composição.

## 🛠 Exemplos de Código

A seguir, um exemplo que demonstra como implementar herança-like em Go utilizando composição de structs e métodos.

### Código de Exemplo

```go
package main

import "fmt"

// Definição da struct Animal
type Animal struct {
    Nome string
}

// Definição da struct Cachorro que compõe a struct Animal
type Cachorro struct {
    Animal
    Raça string
}

// Método Falar para a struct Cachorro
func (c Cachorro) Falar() {
    fmt.Println(c.Nome, "diz: Au au!")
}

// Método Falar para a struct Animal
func (a Animal) Falar() {
    fmt.Println(a.Nome, "emite um som.")
}

func main() {
    c := Cachorro{
        Animal: Animal{Nome: "Pingo"},
        Raça:   "Vira Lata",
    }

    c.Falar()           // Chama o método Falar do Cachorro
    c.Animal.Falar()    // Chama o método Falar do Animal
}
```

## Explicação do Código

### 1. Declaração das Structs

```go
type Animal struct {
    Nome string
}

type Cachorro struct {
    Animal
    Raça string
}
```

- **Animal**: Representa uma entidade básica com um campo ``Nome``.
- **Cachorro**: Compoe a struct ``Animal`` e adiciona um campo específico ``Raça``. Isso permite que Cachorro herde os campos e métodos de ``Animal``.

### 2. Definição dos Métodos
```go
func (c Cachorro) Falar() {
    fmt.Println(c.Nome, "diz: Au au!")
}

func (a Animal) Falar() {
    fmt.Println(a.Nome, "emite um som.")
}
```

- **Falar para Cachorro**: Método específico que imprime uma mensagem típica de cachorro.
- **Falar para Animal**: Método genérico que imprime uma mensagem genérica para qualquer animal.

### 3. Função Principal
```go
func main() {
    c := Cachorro{
        Animal: Animal{Nome: "Pingo"},
        Raça:   "Vira Lata",
    }

    c.Falar()           // Chama o método Falar do Cachorro
    c.Animal.Falar()    // Chama o método Falar do Animal
}
```
- **Inicialização do Cachorro**: Cria uma instância de ``Cachorro`` com nome "Pingo" e raça "Vira Lata".
- **Chamada dos Métodos**:
    - ``c.Falar()``: Executa o método ``Falar`` definido para Cachorro, exibindo "Pingo diz: Au au!".
    - ``c.Animal.Falar()``: Executa o método ``Falar`` definido para Animal, exibindo "Pingo emite um som."

### 4. Saída Esperada
```go
Pingo diz: Au au!
Pingo emite um som.
```

## 🔍 Conceitos Importantes

### 1. Composição de Structs
Em Go, a composição é usada para reutilizar funcionalidades de outras structs. Ao incluir uma struct dentro de outra, você herda seus campos e métodos, permitindo a extensão das funcionalidades sem usar herança tradicional.
```go
type Pessoa struct {
    Nome string
}

type Estudante struct {
    Pessoa
    Matricula string
}
```

### 2. Métodos Associados a Structs
Go permite associar métodos a qualquer tipo, incluindo structs. Isso é essencial para implementar comportamentos específicos e simular herança.
```go
func (p Pessoa) Saudacao() {
    fmt.Printf("Olá, meu nome é %s.\n", p.Nome)
}
```


### 3. Interfaces
Interfaces em Go definem comportamentos que tipos podem implementar. Elas são uma ferramenta poderosa para criar código flexível e orientado a contratos.
```go
type Falante interface {
    Falar()
}

func cumprimentar(f Falante) {
    f.Falar()
}
```


### 4. Sobrescrita de Métodos
Ao compor structs, você pode definir métodos com o mesmo nome que os métodos da struct incorporada. Isso permite sobrescrever comportamentos específicos.
```go
func (c Cachorro) Falar() {
    fmt.Println(c.Nome, "diz: Au au!")
}
```


### 5. Diferenças entre Herança e Composição
- **Herança**: Relação de "é um", onde uma classe herda atributos e métodos de outra.
- **Composição**: Relação de "tem um", onde uma struct inclui outra para reutilizar suas funcionalidades.

Go incentiva a composição sobre a herança para promover maior flexibilidade e menor acoplamento entre componentes.

## 🛡 Boas Práticas

- **Nomenclatura Clara**

Utilize nomes de structs que representem claramente a entidade ou conceito que estão modelando.
```go
type Produto struct {
    Nome  string
    Preco float64
}
```

- **Exportação de Campos Necessários**

Exporte apenas os campos que precisam ser acessados de outros pacotes, mantendo a encapsulação.
```go
type Usuario struct {
    Nome  string // Exportado
    idade int    // Não exportado
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