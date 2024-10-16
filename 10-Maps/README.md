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