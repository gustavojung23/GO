# 14 - Métodos 🛠️

## 📖 Visão Geral

Em Go, **métodos** são funções associadas a tipos específicos (geralmente structs), que permitem manipular e operar sobre os dados desses tipos de forma encapsulada. Métodos ajudam a manter o código mais organizado e orientado a objetos, permitindo que funcionalidades sejam diretamente associadas a tipos e seus dados.

Go permite definir dois tipos de métodos:
- **Métodos com recebedores por valor**: onde uma cópia do valor é passada para o método.
- **Métodos com recebedores por ponteiro**: onde o ponteiro do valor é passado, permitindo modificar os dados originais.

### 🔍 Objetivos deste Módulo

- **Compreender** como definir e utilizar métodos em Go.
- **Explorar** a diferença entre métodos com recebedores por valor e por ponteiro.
- **Ver na prática** como métodos podem encapsular comportamento e lógica diretamente em tipos.

## 🛠 Exemplo de Código

Aqui está um exemplo que demonstra como criar e usar métodos associados a uma struct em Go.

### Código de Exemplo

```go
package main

import (
    "fmt"
)

type User struct {
    name string
    age  uint8
}

// Método com recebedor por valor
func (u User) saveDB() {
    fmt.Printf("Salvando dados do Usuário %s com idade %d anos\n", u.name, u.age)
}

// Método com recebedor por ponteiro
func (u *User) alterDate() {
    u.age++
    fmt.Printf("Alterando a Idade do Usuário: %d\n", u.age)
}

func main() {
    usuario := User{"John", 20}
    usuario.saveDB()      // Chamando método com recebedor por valor
    usuario.alterDate()    // Chamando método com recebedor por ponteiro
}
```

## Explicação do Código
### 1. Definição da Struct `User`:
```go
type User struct {
    name string
    age  uint8
}
```
- ``User`` é uma struct que contém dois campos: ``name`` (nome do usuário) e ``age`` (idade do usuário).

### 2. Método ``saveDB`` com Recebedor por Valor:
```go
func (u User) saveDB() {
    fmt.Printf("Salvando dados do Usuário %s com idade %d anos\n", u.name, u.age)
}
```
- Este método recebe um valor de cópia da struct ``User`` e imprime as informações do usuário.
- Como ele recebe uma cópia, qualquer modificação no valor não afeta o valor original fora do método.

### 3. Método ``alterDate`` com Recebedor por Ponteiro:
```go
func (u *User) alterDate() {
    u.age++
    fmt.Printf("Alterando a Idade do Usuário: %d\n", u.age)
}
```
- Este método recebe um **ponteiro** para a struct ``User``, permitindo que ele altere diretamente o valor de ``age`` no objeto original.
- Com o uso de um ponteiro, o método pode modificar o estado original da struct.

### 4. Função `main`:
```go
func main() {
    usuario := User{"John", 20}
    usuario.saveDB()
    usuario.alterDate()
}
```
- Cria uma struct ``User`` chamado usuario com nome "John" e idade 20.
- Chama o método ``saveDB``, que imprime os dados do usuário.
- Chama o método ``alterDate``, que incrementa a idade do usuário e imprime o novo valor.

### Saída Esperada:
```makefile
Salvando dados do Usuário John com idade 20 anos
Alterando a Idade do Usuário: 21
```

## 🔍 Conceitos Importantes
### 1. Recebedor por Valor
- O método recebe uma cópia do valor original da struct.
- Modificações feitas na cópia **não afetam** o valor original.
#### Exemplo:
```go
func (u User) saveDB() {
    // Recebe uma cópia dos dados da struct User
}
```

### 2. Recebedor por Ponteiro
- O método recebe um ponteiro para o valor original da struct.
- Modificações feitas no valor do ponteiro **afetam diretamente** o valor original.
#### Exemplo:
```go
func (u *User) alterDate() {
    // Recebe um ponteiro para os dados da struct User, permitindo modificar o valor original
}
```

### Quando Usar Cada Um?
- **Recebedor por valor** é útil quando você não precisa modificar o valor original e quer garantir que o método não altere o estado do objeto.
- **Recebedor por ponteiro** é necessário quando você quer modificar o valor original da struct dentro do método.

## 🛡 Boas Práticas
- **Usar Recebedor por Ponteiro Quando Necessário**: Prefira recebedores por ponteiro se o método precisa modificar o estado do objeto ou quando o objeto é grande, evitando cópias desnecessárias.
```go
func (u *User) updateAge() {
    u.age++
}
```

- **Usar Recebedor por Valor para Simplicidade**: Use recebedores por valor quando você não precisa alterar o estado da struct e quer garantir que o método não afete os dados originais.
```go
func (u User) displayInfo() {
    fmt.Println(u.name, u.age)
}
```

- **Mantenha o Método Relacionado ao Tipo**: Métodos devem ser diretamente relacionados ao comportamento ou função do tipo em que estão associados. Evite métodos genéricos ou que não façam sentido no contexto da struct.