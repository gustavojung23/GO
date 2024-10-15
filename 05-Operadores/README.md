# 05 - Operadores

## 📖 Visão Geral

Os operadores são elementos fundamentais em qualquer linguagem de programação, e em Go não é diferente. Eles nos permitem realizar operações matemáticas, relacionais, lógicas, e manipular valores de variáveis. Neste exemplo, veremos os principais tipos de operadores:

- **Aritméticos**
- **Atribuição**
- **Relacionais**
- **Lógicos**
- **Unários**

## 🚀 Exemplos de Uso

Abaixo está um exemplo prático que demonstra como utilizar diferentes operadores na linguagem Go.

```go
package main

import "fmt"

func main() {
    // Operadores Aritméticos
    soma := 2 + 2
    subtracao := 2 - 1
    multiplicacao := 2 * 5
    divisao := 10 / 5
    restoDaDivisao := 10 % 2
    fmt.Println(soma, subtracao, multiplicacao, divisao, restoDaDivisao)

    // Operadores de Atribuição
    var variavel1 string = "String"
    variavel2 := "String 2"
    fmt.Println(variavel1, variavel2)

    // Operadores Relacionais
    fmt.Println(1 > 2)
    fmt.Println(1 >= 2)
    fmt.Println(1 < 2)
    fmt.Println(1 <= 2)
    fmt.Println(1 == 2)
    fmt.Println(1 != 2)

    // Operadores Lógicos
    verdadeiro, falso := true, false
    fmt.Println(verdadeiro && falso)
    fmt.Println(verdadeiro || falso)
    fmt.Println(!verdadeiro)
    fmt.Println(!falso)

    // Operadores Unários
    numero1 := 10
    numero1++
    numero1 += 15
    fmt.Println(numero1)

    numero2 := 10
    numero2--
    numero2 -= 5
    fmt.Println(numero2)

    numero3 := 10
    numero3 *= 2
    fmt.Println(numero3)

    numero4 := 14
    numero4 /= 2
    fmt.Println(numero4)

    numero5 := 10
    numero5 %= 2
    fmt.Println(numero5)
}
```

## ✨ Tipos de Operadores

### 1. Operadores Aritméticos

Realizam operações matemáticas básicas. Em Go, temos os seguintes operadores:
- `+`: Adição
- `-`: Subtração
- `*`: Multiplicação
- `/`: Divisão
- `%`: Resto da divisão

#### Exemplo:
```go
soma := 2 + 2
subtracao := 2 - 1
multiplicacao := 2 * 5
divisao := 10 / 5
restoDaDivisao := 10 % 2
fmt.Println(soma, subtracao, multiplicacao, divisao, restoDaDivisao)
```

### 2. Operadores de Atribuição

Usados para atribuir valores às variáveis. Alguns exemplos são:
- `=`: Atribui um valor à variável
- `+=`: Adiciona o valor à variável e atribui o resultado
- `-=`: Subtrai o valor e atribui o resultado

#### Exemplo:
```go
var variavel1 string = "String"
variavel2 := "String 2"
fmt.Println(variavel1, variavel2)
```
### 3. Operadores Relacionais

Comparam dois valores e retornam ``true`` ou ``false`` com base na relação entre eles.
- `>`: Maior que
- `<`: Menor que
- `>=`: Maior ou igual a
- `<=`: Menor ou igual a
- `==`: Igual a
- `!=`: Diferente de

#### Exemplo:
```go
fmt.Println(1 > 2)
fmt.Println(1 >= 2)
fmt.Println(1 < 2)
fmt.Println(1 <= 2)
fmt.Println(1 == 2)
fmt.Println(1 != 2)
```

### 4. Operadores Lógicos

Utilizados para combinar expressões booleanas.
- `&&`: E lógico (AND)
- `||`: Ou lógico (OR)
- `!`: Não lógico (NOT)

#### Exemplo:
```go
verdadeiro, falso := true, false
fmt.Println(verdadeiro && falso)
fmt.Println(verdadeiro || falso)
fmt.Println(!verdadeiro)
fmt.Println(!falso)
```

### 5. Operadores Unários

Esses operadores são usados para realizar operações com um único operando, como incremento ou decremento.
- `++`: Incrementa em 1
- `--`: Decrementa em 1
- `+=`: Aumenta o valor da variável em um número específico
- `-=`: Diminui o valor da variável em um número específico

#### Exemplo:
```go
numero1 := 10
numero1++
numero1 += 15
fmt.Println(numero1)

numero2 := 10
numero2--
numero2 -= 5
fmt.Println(numero2)

numero3 := 10
numero3 *= 2
fmt.Println(numero3)

numero4 := 14
numero4 /= 2
fmt.Println(numero4)

numero5 := 10
numero5 %= 2
fmt.Println(numero5)
```