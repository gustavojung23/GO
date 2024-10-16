# 12 - Defer ⏳

## 📖 Visão Geral

Em Go, a palavra-chave `defer` é usada para agendar a execução de uma função após o término da função que a chamou. Independentemente de como a função atual termina—seja retornando normalmente, ocorrendo um `panic`, ou por qualquer outro motivo—as funções deferidas serão executadas na ordem inversa em que foram declaradas. Isso é particularmente útil para garantir a liberação de recursos, como fechar arquivos ou liberar locks, mesmo quando erros ocorrem.

## 🔍 Objetivos deste Módulo

- Compreender o funcionamento da palavra-chave `defer` em Go.
- Aprender a agendar a execução de funções deferidas.
- Entender a ordem de execução das funções deferidas.
- Explorar exemplos práticos do uso de `defer` para gerenciamento de recursos e controle de fluxo.

## 🛠 Exemplo de Código

A seguir, um exemplo que demonstra como utilizar `defer` em Go para garantir a execução de funções de limpeza após o término de uma função principal.

### Código de Exemplo

```go
package main

import "fmt"

func main() {
    fmt.Println(alunoAprovado(7, 6))
}

func alunoAprovado(n1, n2 float32) string {
    defer fmt.Println("Média calculada. Resultado será retornado!")
    fmt.Println("Calculando Média.")

    media := (n1 + n2) / 2

    if media >= 6 {
        return fmt.Sprint("Aluno aprovado!")
    }
    return fmt.Sprintf("Aluno reprovado!")
}
```

## Explicação do Código

### 1. Função ``main``:

```go
func main() {
    fmt.Println(alunoAprovado(7, 6))
}
```
- Chama a função ``alunoAprovado`` passando as notas ``7`` e ``6``.
- Imprime o resultado retornado pela função ``alunoAprovado``.

### 2. Função ``alunoAprovado``:
```go
func alunoAprovado(n1, n2 float32) string {
    defer fmt.Println("Média calculada. Resultado será retornado!")
    fmt.Println("Calculando Média.")

    media := (n1 + n2) / 2

    if media >= 6 {
        return fmt.Sprint("Aluno aprovado!")
    }
    return fmt.Sprintf("Aluno reprovado!")
}
```
- Uso de ``defer``: ``defer fmt.Println("Média calculada. Resultado será retornado!")`` agenda a impressão dessa mensagem para ocorrer após o retorno da função, garantindo que a mensagem sempre seja exibida, independentemente do caminho de retorno.
- Cálculo da Média: Calcula a média das duas notas.
- Condição de Aprovação:
    - Se a média for maior ou igual a ``6``, retorna "Aluno aprovado!".
    - Caso contrário, retorna "Aluno reprovado!".

### Saída Esperada
```makefile
Calculando Média.
Aluno aprovado!
Média calculada. Resultado será retornado!
```

#### Ordem da Saída:
1. "Calculando Média." é impresso primeiro.
2. "Aluno aprovado!" é retornado e impresso pela função main.
3. "Média calculada. Resultado será retornado!" é impresso após o retorno da função alunoAprovado devido ao defer.

## 🔍 Conceitos Importantes

### 1. Ordem de Execução das Funções Deferidas
- **LIFO (Last In, First Out)**: Múltiplas chamadas de defer são executadas na ordem inversa à sua declaração.
```go
func main() {
    defer fmt.Println("Primeiro defer")
    defer fmt.Println("Segundo defer")
    defer fmt.Println("Terceiro defer")
}
```
### Saída
```makefile
Terceiro defer
Segundo defer
Primeiro defer
```

### 2. Avaliação de Argumentos
- **Momento da Avaliação**: Os argumentos das funções deferidas são avaliados no momento da declaração do defer, não no momento da execução.
```go
func main() {
    x := 10
    defer fmt.Println("Valor de x:", x)
    x = 20
}
```
### Saída
```makefile
Valor de x: 10
```

### 3. Uso de ``defer`` para Gerenciamento de Recursos
- **Fechando Arquivos**: Garantir que um arquivo seja fechado após sua utilização.
```go
file, err := os.Open("example.txt")
if err != nil {
    // tratar erro
}
defer file.Close()
// trabalhar com o arquivo
```
- **Liberando Locks**: Garantir que um lock seja liberado após uma operação crítica.
```go
var mu sync.Mutex
mu.Lock()
defer mu.Unlock()
// operação crítica
```

### 4. ``defer`` com ``panic`` e ``recover``
- **Recuperando de Panics**: Utilizar defer para capturar e recuperar de panics, evitando que o programa termine abruptamente.
```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recuperado de:", r)
        }
    }()
    panic("Um erro ocorreu!")
}
```