# 11 - Estruturas de Controle 🛠️

## 📖 Visão Geral

Em Go, as estruturas de controle são usadas para controlar o fluxo de execução do programa com base em condições e iterações. Isso inclui declarações como `if`, `else`, `for` e `switch`, que desempenham um papel crucial em tomar decisões e iterar sobre coleções.

### 🔍 Objetivos deste Módulo

- Entender como utilizar as estruturas de controle `if/else`, `for` e `switch`.
- Aprender a tomar decisões com base em condições.
- Realizar loops em arrays e slices.
- Controlar o fluxo de execução com base em diferentes casos de um `switch`.

## 🛠 Exemplo de Código

A seguir, exemplos de uso das principais estruturas de controle em Go.

### 1. If / Else

A estrutura `if/else` é usada para tomar decisões com base em condições booleanas. Dependendo do resultado da condição, o programa executa blocos de código diferentes.

#### Código de Exemplo

```go
package main

import "fmt"

func main() {
    numero1 := 10

    // Condicional simples
    if numero1 >= 15 {
    fmt.Println("Maior")
    } else {
        fmt.Println("Menor")
    }

    // Condicional com inicialização
    if numero2 := numero1; numero2 > 9 {
        fmt.Println("Maior")
    } else {
        fmt.Println("Menor")
    }
}
```

#### Explicação
- A primeira condição verifica se ``numero1`` é maior ou igual a 15, imprimindo "Maior" ou "Menor" com base nisso.
- A segunda condição inicializa uma nova variável ``numero2`` e faz uma verificação sobre ela.

### 2. Loops

O loop ``for`` é a única estrutura de loop em Go e pode ser usada de diversas formas, como um loop clássico ou como um loop para iterar sobre coleções.

#### Código de Exemplo

```go
package main

import "fmt"

func main() {
    // Loop tradicional
    for i := 0; i < 10; i++ {
        fmt.Println("Laço I: ", i)
    }

    // Iterando sobre um array com range
    cidades := [3]string{"Nova York", "Brasília", "Paris"}
    for i, cidade := range cidades {
        fmt.Println(i, cidade)
    }
}
```

#### Explicação
- O primeiro ``for`` é um loop tradicional que itera de 0 a 9 e imprime o valor de ``i``.
- O segundo exemplo utiliza ``range`` para iterar sobre um array de cidades e imprime o índice e o valor de cada elemento.

### 2. Switch

A estrutura ``switch`` permite executar diferentes blocos de código com base no valor de uma expressão. É uma alternativa para uma série de ``if/else`` quando há múltiplos casos a serem verificados.

#### Código de Exemplo

```go
package main

import "fmt"

func main() {
    dia := diaDaSemana(1)
    fmt.Println(dia)
}

func diaDaSemana(numero int) string {
    switch numero {
        case 1:
            return "Domingo"
        case 2:
            return "Segunda-Feira"
        case 3:
            return "Terça-Feira"
        case 4:
            return "Quarta-Feira"
        case 5:
            return "Quinta-Feira"
        case 6:
            return "Sexta-Feira"
        case 7:
            return "Sábado"
        default:
            return "Número Inválido"
    }
}
```

#### Explicação
- A função ``diaDaSemana`` retorna o nome do dia com base no número passado.
- A estrutura ``switch`` verifica o valor do número e retorna o nome correspondente do dia da semana.
- Um ``default`` é fornecido para lidar com entradas inválidas.

## 🔍 Conceitos Importantes
### 1. If / Else
- A estrutura ``if/else`` executa diferentes blocos de código com base no valor booleano de uma expressão.
- É possível inicializar variáveis dentro da condição ``if`` e usá-las tanto no ``if`` quanto no ``else``.

### 2. For
O loop ``for`` é a principal ferramenta de iteração em Go e é suficiente para a maioria dos casos de uso. No entanto, é importante entender que ele não é a única forma de implementar a iteração em Go. ``Goroutines``, ``canais`` e a recursividade oferecem alternativas poderosas, especialmente em cenários mais complexos ou quando se deseja um alto nível de concorrência.

### 3. Switch
- O ``switch`` em Go avalia uma expressão e executa o primeiro caso correspondente.
- Não é necessário utilizar ``break``, pois Go já "quebra" automaticamente após a execução de um caso.
- Pode ser usado como uma alternativa mais elegante ao ``if/else`` para condições múltiplas.

## 🛡️ Boas Práticas
- **Use Switch ao Invés de Múltiplos If/Else**

O switch pode ser mais legível e eficiente quando há múltiplos casos a serem avaliados.

- **Evite Loops Infinitos Desnecessários**

Um loop sem uma condição de término clara pode causar problemas de performance ou travamentos. Verifique sempre as condições de saída dos loops.