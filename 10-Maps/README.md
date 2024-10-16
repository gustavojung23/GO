# 10 - Maps 📚

## 📖 Visão Geral

**Maps** em Go são estruturas de dados que armazenam pares de chave-valor, oferecendo uma maneira eficiente de mapear chaves a valores correspondentes. Eles são semelhantes a dicionários em outras linguagens de programação, e permitem uma busca rápida de valores com base nas suas chaves.

### 🔍 Objetivos deste Módulo

- Compreender o conceito de **maps** em Go.
- Aprender a declarar e inicializar **maps**.
- Manipular maps e acessar seus valores.
- Trabalhar com maps aninhados (maps dentro de maps).

## 🛠 Exemplo de Código

A seguir, um exemplo que demonstra como declarar, inicializar e manipular **maps** em Go.

### Código de Exemplo

```go
package main

import "fmt"

func main() {
    // Declaração e Inicialização de um Map Simples
    usuario := map[string]string{
        "nome":      "Jhon",
        "sobrenome": "Doe",
    }
    fmt.Println("Map de Usuário:", usuario)

    // Declaração e Inicialização de um Map Aninhado
    curso := map[string]map[string]string{
        "Faculdade": {
            "Campus": "UNI",
            "Curso":  "ADS",
        },
    }
    fmt.Println("Map de Curso:", curso)
}
```

## Explicação do Código

### 1. Map Simples
- O map ``usuario`` associa strings a strings. A chave representa o tipo de informação (por exemplo, ``nome``, ``sobrenome``) e o valor representa o dado correspondente.
```go
usuario := map[string]string{
    "nome":      "Jhon",
    "sobrenome": "Doe",
}
fmt.Println("Map de Usuário:", usuario)
```
#### Saída:
```makefile
Map de Usuário: map[nome:Jhon sobrenome:Doe]
```

### 2. Map Aninhado
- O map ``curso`` é um exemplo de um map cujos valores também são maps. Isso permite armazenar informações hierarquicamente.
```go
curso := map[string]map[string]string{
    "Faculdade": {
        "Campus": "UNI",
        "Curso":  "ADS",
    },
}
fmt.Println("Map de Curso:", curso)
```
#### Saída:
```makefile
Map de Curso: map[Faculdade:map[Campus:UNI Curso:ADS]]
```

## 🔍 Conceitos Importantes
### 1. Maps
- **Definição**: Uma coleção de pares chave-valor, onde cada chave única mapeia para um valor.
- **Sintaxe**: ``map[tipo_chave]tipo_valor``, onde ``tipo_chave`` é o tipo das chaves e ``tipo_valor`` é o tipo dos valores.
- **Características**:
    - As chaves são únicas e cada uma delas pode ter apenas um valor correspondente.
    - O valor de um map pode ser qualquer tipo, inclusive outro map (aninhamento).
    - Maps em Go são implementados como referências internas, portanto, são passados por referência.

### 2. Declaração de Maps
- **Declaração e Inicialização**: Um map pode ser declarado e inicializado ao mesmo tempo usando a sintaxe literal, como mostrado no exemplo.
```go
usuario := map[string]string{
    "nome":      "Jhon",
    "sobrenome": "Doe",
}
```

- **Declaração Vazia**: Um map também pode ser declarado vazio e posteriormente preenchido com valores.
```go
var m map[string]int
m = make(map[string]int)
```

- **Acesso a Elementos**: Valores em um map são acessados usando a chave correspondente.
```go
fmt.Println(usuario["nome"])  // "Jhon"
```

### 3. Funções Comuns
- **Adicionar ou Atualizar Elementos**: Use a chave para adicionar ou atualizar um valor.
```go
usuario["nome"] = "Jhon"
```

- **Deletar Elementos**: A função delete remove uma chave e seu valor correspondente do map.
```go
delete(usuario, "sobrenome")
```

- **Verificar a Existência de uma Chave**: Use uma atribuição dupla para verificar se uma chave está presente no map.
```go
valor, ok := usuario["idade"]
if ok {
    fmt.Println("Idade:", valor)
} else {
    fmt.Println("Idade não encontrada")
}
```

## 🛡 Boas Práticas
- **Evite Modificar Maps Durante a Iteração**

Modificar um map enquanto está iterando sobre ele pode levar a comportamentos inesperados.
```go
for k, v := range usuario {
    fmt.Println(k, v)
}
```

- **Inicialize Maps com ``make``**

Sempre use a função ``make`` para inicializar um map vazio antes de começar a adicionar elementos a ele.
```go
m := make(map[string]int)
```