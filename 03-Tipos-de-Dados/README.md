# 03 - Tipos de Dados

## 📋 Introdução aos Tipos de Dados em Go

Go é uma linguagem fortemente tipada, o que significa que cada valor tem um tipo específico. Compreender os tipos de dados disponíveis e como utilizá-los corretamente é fundamental para escrever código eficiente e sem erros.

Neste módulo, abordaremos os principais tipos de dados em Go: **booleanos**, **erro**, **numéricos**, **strings** e **valor zero**.

## 🛠 Exemplos de Código

A seguir, exemplos de cada tipo de dado:

### 1. Booleanos

Booleanos são variáveis que podem armazenar apenas dois valores: `true` ou `false`.

#### Código de Exemplo

```go
package main

import "fmt"

func main() {
    var booleanoTrue bool = true
    fmt.Println(booleanoTrue)
    
    var booleanoFalse bool = false
    fmt.Println(booleanoFalse)
}
```

#### Explicação
- `bool` é o tipo de dado que representa booleanos.
- `true` e `false` são os únicos valores possíveis para esse tipo.

### 2. Erro

Go tem um tipo de dados dedicado a tratar erros chamado `error`. É amplamente utilizado nas funções que podem falhar.

#### Código de Exemplo

```go
package main

import (
    "errors"
    "fmt"
)

func main() {
    var erro error
    fmt.Println(erro)

    var newErro error = errors.New("Erro interno")
    fmt.Println(newErro)
}
```

#### Explicação
- O tipo `error` é usado para representar e manipular erros.
- `errors.New()` e `false` cria uma nova instância de erro.

### 3. Tipos Numéricos

Go oferece suporte a diversos tipos numéricos, incluindo inteiros e números de ponto flutuante.

#### Código de Exemplo

```go
package main

import "fmt"

func main() {
    numero := 100000
    fmt.Println("Numero int: ", numero)

    var numeroInt8 int8 = 127
    fmt.Println("Numero int8: ", numeroInt8)

    var numeroInt16 int16 = 32767
    fmt.Println("Número int16: ", numeroInt16)

    var numeroInt32 int32 = 2147483647
    fmt.Println("Número int32: ", numeroInt32)

    var numeroInt64 int64 = 9223372036854775807
    fmt.Println("Número int64: ", numeroInt64)

    var numeroUint uint = 2147483647
    fmt.Println("Número uint: ", numeroUint)

    var numeroUint8 uint8 = 255
    fmt.Println("Número uint8: ", numeroUint8)

    var numeroUint16 uint16 = 65535
    fmt.Println("Número uint16: ", numeroUint16)

    var numeroUint32 uint32 = 4294967295
    fmt.Println("Número uint32: ", numeroUint32)

    var numeroUint64 uint64 = 18446744073709551615
    fmt.Println("Número uint64: ", numeroUint64)

    var numeroRune rune = 2147483647
    fmt.Println("Número rune: ", numeroRune)

    var numeroByte byte = 255
    fmt.Println("Número byte: ", numeroByte)

    var numeroFloat32 float32 = 1.12345678901234567890
    fmt.Println("Número float64: ", numeroFloat32)

    var numeroFloat64 float64 = 1.12345678901234567890
    fmt.Println("Número float64: ", numeroFloat64)
}
```

#### Explicação
- **Inteiros**: `int`, `int8`, `int16`, `int32`, `int64`, bem como suas versões sem sinal (`uint` etc.).
- **Ponto Flutuante**: `float32` e `float64` são usados para números com casas decimais.

### 4. Strings

Strings em Go são cadeias de caracteres imutáveis, ou seja, não podem ser alteradas após a sua criação.

#### Código de Exemplo

```go
package main

import "fmt"

func main() {
    var texto string = "Texto"
    fmt.Println(texto)
}
```

#### Explicação
- `string` é o tipo utilizado para representar sequências de caracteres.
- Strings são imutáveis, ou seja, não podem ser modificadas após serem criadas.

### 5. Valor Zero

Em Go, variáveis que não são inicializadas explicitamente recebem um valor padrão chamado valor zero. O valor zero depende do tipo da variável:

- **Booleanos**: `false`
- **Inteiros**: `0`
- **Float**: `0.0`
- **Strings**: `""` (string vazia)
- **Erro**: `nil` (sem erro)

#### Código de Exemplo

```go
package main

import "fmt"

func main() {
	var defaultBool bool
	var defaultInt int
	var defaultFloat float64
	var defaultString string

	fmt.Println(defaultBool)
	fmt.Println(defaultInt)
	fmt.Println(defaultFloat)
	fmt.Println(defaultString)
}
```

#### Explicação
- Variáveis não inicializadas recebem automaticamente o valor zero de seu tipo correspondente.

## 🔍 Conceitos Importantes

### 1. Inferência de Tipo
Em Go, podemos omitir o tipo ao declarar uma variável, e o compilador o inferirá automaticamente:

```go
name := "Jhon"   // inferido como string
age := 30        // inferido como int
```

### 2. Conversão de Tipos
Em Go, a conversão entre tipos deve ser explícita, o que evita erros de compilação.

```go
var i int = 42
var f float64 = float64(i)
```

### 3. Zero Values
- Ao trabalhar com variáveis que ainda não foram inicializadas, é importante entender os valores padrão.
- **Nil**: Tipos de dados como ponteiros, slices, maps, e interfaces têm o valor zero `nil`.

## ✅ Próximos Passos
Após entender os tipos básicos de dados em Go, explore tópicos mais avançados, como:

- **Tipos Compostos**: Structs, Arrays, Slices, Maps
- **Manipulação de Erros**: Como trabalhar com o tipo error
- **Conversão de Tipos**: Como converter entre diferentes tipos numéricos