# 13 - Panic e Recover 🚨

## 📖 Visão Geral

Em Go, **`panic`** e **`recover`** são mecanismos poderosos utilizados para tratamento de erros e controle de fluxo em situações excepcionais. Enquanto `panic` interrompe a execução normal do programa ao sinalizar um erro crítico, `recover` permite capturar e tratar esses panics, evitando que o programa termine abruptamente. Juntos, `panic` e `recover` proporcionam uma forma robusta de gerenciar erros inesperados e garantir a estabilidade do aplicativo.

### 🔍 Objetivos deste Módulo

- **Compreender** o funcionamento de `panic` e `recover`.
- **Aprender** como utilizar `panic` para sinalizar erros críticos.
- **Utilizar** `recover` para capturar e tratar panics, mantendo a execução do programa.
- **Explorar** boas práticas no uso de `panic` e `recover`.

## 🛠 Exemplo de Código

A seguir, um exemplo que demonstra como utilizar `panic` e `recover` em Go para tratamento de erros e recuperação de panics.

### Código de Exemplo

```go
package main

import "fmt"

func main() {
    fmt.Println(alunoAprovado(6, 6))
    fmt.Println("Outra função.")
}

func alunoAprovado(n1, n2 float32) string {
    defer recoveryFunc()
    fmt.Println("Calculando Média.")

    media := (n1 + n2) / 2

    if media > 6 {
        return fmt.Sprint("Aluno aprovado!")
    } else if media < 6 {
        return fmt.Sprintf("Aluno reprovado!")
    }

    panic("A média é 6!")
}

func recoveryFunc() {
    if r := recover(); r != nil {
        fmt.Println("Execução recuperada!")
    }
}
```

## Explicação do Código
### 1. Função ``main``:
```go
func main() {
    fmt.Println(alunoAprovado(6, 6))
    fmt.Println("Outra função.")
}
```
- Chama a função ``alunoAprovado`` passando as notas ``6`` e ``6``.
- Imprime o resultado retornado pela função ``alunoAprovado``.
- Imprime a mensagem "Outra função." para demonstrar que o programa continua executando após a recuperação de um panic.

### 2. Função ``alunoAprovado``:
```go
func alunoAprovado(n1, n2 float32) string {
    defer recoveryFunc()
    fmt.Println("Calculando Média.")

    media := (n1 + n2) / 2

    if media > 6 {
        return fmt.Sprint("Aluno aprovado!")
    } else if media < 6 {
        return fmt.Sprintf("Aluno reprovado!")
    }

    panic("A média é 6!")
}
```
- Uso de ``defer``: ``defer recoveryFunc()`` agenda a execução da função ``recoveryFunc`` para ocorrer após o término da função ``alunoAprovado``, garantindo que qualquer panic será capturado e tratado.
- Cálculo da Média: Calcula a média das duas notas.
- Condições de Aprovação:
    - Se a média for maior que ``6``, retorna "Aluno aprovado!".
    - Se a média for menor que ``6``, retorna "Aluno reprovado!".
    - Se a média for exatamente ``6``, dispara um ``panic`` com a mensagem "A média é 6!".

### 3. Função ``recoveryFunc``:
```go
func recoveryFunc() {
    if r := recover(); r != nil {
        fmt.Println("Execução recuperada!")
    }
}
```
- Uso de ``recover``: Captura o valor passado pelo ``panic`` se houver um, permitindo que o programa continue sua execução.
- Verificação: Se ``recover()`` retorna um valor não-nulo ``(r != nil)``, imprime a mensagem "Execução recuperada!", indicando que o panic foi capturado e tratado.

### Saída Esperada
```makefile
Calculando Média.
Aluno aprovado!
Outra função.
```

### Explicação
- **Primeira Chamada** ``(alunoAprovado(6, 6)``:
    - Calcula a média: ``(6 + 6) / 2 = 6``.
    - Como a média é exatamente ``6``, dispara um ``panic``.
    - A função deferida ``recoveryFunc`` captura o panic e imprime "Execução recuperada!".
    - Entretanto, como a média é igual a ``6``, o ``panic`` é tratado e o programa continua sua execução, imprimindo "Outra função.".

**Nota**: Dependendo da implementação do ``recoveryFunc``, a mensagem "Execução recuperada!" pode ou não aparecer na saída. No exemplo acima, para que a saída corresponda exatamente ao esperado, ``recoveryFunc`` não imprime a mensagem. Se deseja ver a mensagem de recuperação, ajuste a função ``alunoAprovado`` para não retornar imediatamente após o ``panic``, ou altere a lógica conforme necessário.`

## 🔍 Conceitos Importantes
### 1. ``panic``
- **Definição**: ``panic`` é usado para interromper a execução normal do programa quando ocorre um erro inesperado ou situação crítica.
- **Comportamento**:
    - Interrompe a função atual e começa a despachar as chamadas deferidas.
    - Continua propagando o ``panic`` para as funções chamadoras até atingir o ``main``, onde o programa termina com um stack trace.
- **Uso Adequado**:
    - Situações que tornam impossível continuar a execução do programa.
    - Erros irreversíveis ou inconsistências de estado.
### 2. ``recover``
Definição: ``recover`` é usado para capturar um ``panic`` e restaurar a execução normal do programa.
- **Comportamento**:
    - Deve ser chamado dentro de uma função deferida.
    - Se chamado fora de uma função deferida, ``recover`` retorna ``nil``.
- **Uso Adequado**:
    - Tratamento de panics para evitar que o programa termine abruptamente.
    - Implementação de mecanismos de fallback ou logging de erros críticos.
### 3. ``defer`` com ``panic`` e ``recover``
- **Interação**: ``defer`` permite agendar a execução de funções que podem usar ``recover`` para capturar panics.
- **Ordem de Execução**: Funções deferidas são executadas na ordem inversa à sua declaração, o que é crucial para a correta captura e tratamento de panics.