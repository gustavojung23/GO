# 09 - Arrays e Slices 📚

## 📖 Visão Geral

Em Go, **arrays** e **slices** são estruturas de dados fundamentais para armazenar coleções de elementos. Embora ambos permitam armazenar múltiplos valores do mesmo tipo, eles possuem diferenças significativas em termos de flexibilidade e funcionalidade. Compreender como declarar, inicializar e manipular arrays e slices é essencial para escrever programas eficientes e organizados na linguagem Go.

### 🔍 Objetivos deste Módulo

- Entender a diferença entre arrays e slices.
- Aprender a declarar e inicializar arrays e slices.
- Manipular arrays e slices utilizando funções incorporadas.
- Explorar as capacidades e limitações de cada estrutura.

## 🛠 Exemplos de Código

A seguir, um exemplo completo que demonstra como declarar e utilizar arrays e slices em Go.

### Código de Exemplo

```go
package main

import "fmt"

func main() {
    // Declaração e Inicialização de Arrays

    // Array de inteiros com 5 elementos
    var array1 [5]int
    array1[0] = 1
    array1[1] = 2
    array1[2] = 3
    array1[3] = 4
    array1[4] = 5
    fmt.Println("Array1:", array1)

    // Array de strings com 2 elementos
    array2 := [2]string{"Posição 1", "Posição 2"}
    fmt.Println("Array2:", array2)

    // Array com tamanho inferido
    array3 := [...]int{1, 2, 3, 4}
    fmt.Println("Array3:", array3)

    // Declaração e Manipulação de Slices

    // Slice de inteiros com 4 elementos
    slice1 := []int{10, 11, 12, 13}
    fmt.Println("Slice1:", slice1)

    // Adicionando um elemento ao slice
    slice1 = append(slice1, 14)
    fmt.Println("Slice1 após append:", slice1)

    // Criando um slice a partir de um array ou outro slice
    slice2 := slice1[0:3]
    fmt.Println("Slice2 (slice1[0:3]):", slice2)

    // Criando um slice com a função make
    slice3 := make([]int, 3)
    fmt.Println("Slice3 - Tamanho:", len(slice3))
    fmt.Println("Slice3 - Capacidade:", cap(slice3))
}
```

## Explicação do Código

### 1. Declaração e Inicialização de Arrays
- **Array1**: Declara um array de inteiros com 5 elementos e atribui valores a cada posição.
```go
var array1 [5]int
array1[0] = 1
array1[1] = 2
array1[2] = 3
array1[3] = 4
array1[4] = 5
fmt.Println("Array1:", array1)
```

**Saída:**
```makefile
Array1: [1 2 3 4 5]
```

- **Array2**: Declara e inicializa um array de strings com 2 elementos.
```go
array2 := [2]string{"Posição 1", "Posição 2"}
fmt.Println("Array2:", array2)
```

**Saída:**
```makefile
Array2: [Posição 1 Posição 2]
```

- **Array3**: Declara um array com tamanho inferido com base nos elementos fornecidos.
```go
array3 := [...]int{1, 2, 3, 4}
fmt.Println("Array3:", array3)
```

**Saída:**
```makefile
Array3: [1 2 3 4]
```

### 2. Declaração e Manipulação de Slices
- **Slice1**: Declara e inicializa um slice de inteiros com 4 elementos.
```go
slice1 := []int{10, 11, 12, 13}
fmt.Println("Slice1:", slice1)
```

**Saída:**
```makefile
Slice1: [10 11 12 13]
```

- **Append**: Adiciona um elemento ao slice utilizando a função append.
```go
slice1 = append(slice1, 14)
fmt.Println("Slice1 após append:", slice1)
```

**Saída:**
```makefile
Slice1 após append: [10 11 12 13 14]
```

- **Slice2**: Cria um novo slice a partir de slice1, incluindo elementos do índice 0 ao 2.
```go
slice2 := slice1[0:3]
fmt.Println("Slice2 (slice1[0:3]):", slice2)
```

**Saída:**
```makefile
Slice2 (slice1[0:3]): [10 11 12]
```

- **Slice3**: Cria um slice utilizando a função make, especificando o tamanho.
```go
slice3 := make([]int, 3)
fmt.Println("Slice3 - Tamanho:", len(slice3))
fmt.Println("Slice3 - Capacidade:", cap(slice3))
```

**Saída:**
```makefile
Slice3 - Tamanho: 3
Slice3 - Capacidade: 3
```

## 🔍 Conceitos Importantes
### 1. Arrays
- **Definição**: Estruturas de dados com tamanho fixo que armazenam elementos do mesmo tipo.
- **Sintaxe**: ``[tamanho]tipo``, onde tamanho é o número de elementos e tipo é o tipo dos elementos.
- **Características**:
    - Tamanho fixo após a declaração.
    - Acesso direto aos elementos através de índices.
    - Valor de tipo diferente de slices, mesmo que contenham os mesmos tipos e valores.

### 2. Slices
- **Definição**: Estruturas de dados dinâmicas que representam uma sequência de elementos de um array subjacente.
- **Sintaxe**: ``[]tipo``, onde ``tipo`` é o tipo dos elementos.
- **Características**:
    - Tamanho flexível que pode crescer ou diminuir conforme necessário.
    - Slices referenciam arrays subjacentes, permitindo compartilhamento de dados.
    - Utilizam funções como ``append`` para adicionar elementos.

### 3. Diferenças entre Arrays e Slices
| Característica   | Arrays                 | Slices                           |
| :--------------- | :--------------------- | :------------------------------- |
| Tamanho          | Fixado na declaração   | Dinâmico                         |
| Flexibilidade    | Menos flexíveis        | Mais flexíveis                   |
| Memória          | Alocação contínua      | Referenciam arrays               |
| Passagem         | Passagem por valor     | Passagem por referência          |

### 4. Funções Comuns com Slices
- ``append``: Adiciona elementos a um slice.
```go
slice := []int{1, 2, 3}
slice = append(slice, 4)
```

- ``make``: Cria um slice com tamanho e capacidade especificados.
```go
slice := make([]int, 5, 10)
```

- ``copy``: Copia elementos de um slice para outro.
```go
slice1 := []int{1, 2, 3}
slice2 := make([]int, len(slice1))
copy(slice2, slice1)
```

## 🛡 Boas Práticas
- **Use Slices ao Invés de Arrays Quando Possível**

Slices oferecem maior flexibilidade e são mais idiomáticos em Go para a maioria das situações.
```go
// Preferir slice
slice := []int{1, 2, 3}

// Evitar arrays fixos
var array [3]int
```

- **Gerencie o Tamanho e a Capacidade dos Slices**

Entenda como o tamanho (``len``) e a capacidade (``cap``) dos slices afetam a performance e a memória.
```go
slice := make([]int, 5, 10)
```

- **Evite Modificar Slices Enquanto Itera**

Modificar slices durante iterações pode levar a comportamentos inesperados. Prefira criar novas slices se precisar modificar os dados.

- **Use ``copy`` para Duplicar Slices**

Ao precisar de uma cópia independente de um slice, utilize a função ``copy``.
```go
slice1 := []int{1, 2, 3}
slice2 := make([]int, len(slice1))
copy(slice2, slice1)
```
- **Documente a Intenção do Uso de Slices**

Comentários claros ajudam a entender por que um slice está sendo usado de determinada forma, especialmente em contextos complexos.
```go
// slice contém os IDs ativos dos usuários
slice := []int{101, 102, 103}
```