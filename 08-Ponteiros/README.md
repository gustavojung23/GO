# 08 - Ponteiros 🪡

## 📖 Visão Geral

Em Go, **ponteiros** são variáveis que armazenam o endereço de memória de outra variável. Eles são uma ferramenta poderosa que permite a manipulação direta da memória, otimizando o desempenho e permitindo a criação de estruturas de dados mais complexas. Compreender como utilizar ponteiros é essencial para escrever código eficiente e evitar erros comuns relacionados à memória.

Neste módulo, abordaremos os conceitos básicos de ponteiros, incluindo declaração, inicialização, desreferenciação e manipulação de ponteiros em Go.

## 🛠 Exemplos de Código

A seguir, um exemplo simples que demonstra como declarar e utilizar ponteiros em Go.

### Código de Exemplo

```go
package main

import "fmt"

func main() {
    var number1 int
    var numberPointer *int

    number1 = 100
    numberPointer = &number1

    fmt.Println("Valor de number1:", number1)
    fmt.Println("Endereço de number1:", numberPointer)
    fmt.Println("Valor através do ponteiro:", *numberPointer)
}
```

## Explicação do Código

### 1. Declaração das Variáveis
```go
var number1 int
var numberPointer *int
```
- **number1**: Uma variável do tipo ``int`` que armazenará um valor numérico.
- **numberPointer**: Uma variável do tipo ponteiro para ``int`` (``*int``) que armazenará o endereço de memória de ``number1``.

### 2. Atribuição de Valores
```go
number1 = 100
numberPointer = &number1
```
- ``number1 = 100``: Atribui o valor ``100`` à variável ``number1``.
- ``numberPointer = &number1``: Atribui o endereço de memória de ``number1 ``à variável ``numberPointer`` usando o operador ``&``.

### 3. Impressão dos Valores
```go
fmt.Println("Valor de number1:", number1)
fmt.Println("Endereço de number1:", numberPointer)
fmt.Println("Valor através do ponteiro:", *numberPointer)
```
- ``fmt.Println("Valor de number1:", number1)``: Imprime o valor armazenado em ``number1``.
- ``fmt.Println("Endereço de number1:", numberPointer)``: Imprime o endereço de memória armazenado em ``numberPointer``.
- ``fmt.Println("Valor através do ponteiro:", *numberPointer)``: Imprime o valor armazenado no endereço apontado por ``numberPointer`` usando o operador de desreferenciação ``*``.

### 4. Saída Esperada
```go
Valor de number1: 100
Endereço de number1: 0xc000014078
Valor através do ponteiro: 100
```
- O endereço exibido ``(0xc000014078)`` será diferente a cada execução, pois é o endereço de memória alocado para ``number1`` durante a execução do programa.

## 🔍 Conceitos Importantes

### 1.Declaração de Ponteiros
Ponteiros são declarados usando o asterisco ``(*)`` antes do tipo de dado. Por exemplo, ``*int`` é um ponteiro para um ``int``.
```go
var p *int
```

### 2. Inicialização de Ponteiros
Um ponteiro pode ser inicializado atribuindo-lhe o endereço de uma variável usando o operador ``&``.
```go
var a int = 10
var p *int = &a
```

### 3. Desreferenciação
A desreferenciação é feita usando o operador ``*`` para acessar o valor armazenado no endereço de memória apontado pelo ponteiro.
```go
fmt.Println(*p) // Imprime o valor de 'a', que é 10
```

### 4. Ponteiros e Funções
Passar ponteiros para funções permite que você modifique o valor original da variável.
```go
func incrementar(n *int) {
    *n += 1
}

func main() {
    var a int = 5
    incrementar(&a)
    fmt.Println(a) // Imprime 6
}
```

### 5. Ponteiros para Structs
Ponteiros também podem ser usados com structs para otimizar o uso de memória e facilitar a manipulação de dados complexos.
```go
type Pessoa struct {
    Nome string
    Idade int
}

func main() {
    p := Pessoa{Nome: "Jhon", Idade: 30}
    ptr := &p
    ptr.Idade = 31
    fmt.Println(p.Idade) // Imprime 31
}
```

## 🛡 Boas Práticas
- **Use Ponteiros com Cautela**

Ponteiros podem aumentar a flexibilidade do seu código, mas também podem introduzir complexidade e riscos de segurança. Use-os apenas quando necessário.

- **Evite Ponteiros Nulos**

Sempre inicialize ponteiros antes de usá-los para evitar erros de execução.
```go
var p *int
// p = new(int) // Inicialização segura
```

- **Documente o Uso de Ponteiros**

Explique por que um ponteiro está sendo usado em uma função ou método para facilitar a compreensão do código por outros desenvolvedores.

- **Prefira Passagem por Valor Quando Apropriado**

Nem sempre é necessário usar ponteiros. Para tipos pequenos e imutáveis, passar por valor pode ser mais simples e eficiente.

- **Combine Ponteiros com Structs**

Utilize ponteiros para structs para evitar cópias desnecessárias e permitir a modificação dos dados originais.
```go
type Carro struct {
    Marca string
    Modelo string
}

func atualizarModelo(c *Carro, novoModelo string) {
    c.Modelo = novoModelo
}
```