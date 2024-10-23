# 15 - Interfaces

## 📖 Visão Geral

Em Go, **interfaces** são um recurso poderoso que define um conjunto de métodos que um tipo deve implementar. Elas permitem que o código seja mais flexível e genérico, pois você pode definir comportamentos sem especificar como exatamente eles devem ser implementados. Qualquer tipo que implemente os métodos exigidos por uma interface é considerado uma implementação daquela interface, sem a necessidade de declaração explícita.

Neste módulo, veremos como criar e utilizar interfaces em Go, implementando métodos específicos para diferentes tipos e aplicando a polimorfia.

### 🔍 Objetivos deste Módulo

- **Compreender** como criar interfaces em Go.
- **Implementar** interfaces para diferentes tipos.
- **Utilizar** métodos com interfaces como parâmetros, permitindo maior flexibilidade no código.

## 🛠 Exemplo de Código

Aqui está um exemplo que demonstra a criação e o uso de interfaces em Go.

### Código de Exemplo

```go
package main

import (
    "fmt"
    "math"
)

// Interface que define o método area
type Area interface {
    area() float64
}

// Struct Retangulo com altura e largura
type Retangulo struct {
    Altura  float64
    Largura float64
}

// Struct Circulo com raio
type Circulo struct {
    Raio float64
}

// Função que recebe uma interface como argumento e imprime a área
func escreverArea(a Area) {
    fmt.Printf("A forma da área é: %0.2f\n", a.area())
}

// Implementação do método area para Retangulo
func (r Retangulo) area() float64 {
    return r.Altura * r.Largura
}

// Implementação do método area para Circulo
func (c Circulo) area() float64 {
    return math.Pi * math.Pow(c.Raio, 2)
}

func main() {
// Criação de um Retangulo e chamada da função escreverArea
    r := Retangulo{10, 15}
    escreverArea(r)

// Criação de um Circulo e chamada da função escreverArea
    c := Circulo{5}
    escreverArea(c)
}
```

## Explicação do Código

### 1. Definição da Interface ``Area``:
```go
type Area interface {
    area() float64
}
```
- A interface ``Area`` exige que qualquer tipo que a implemente forneça uma função ``area()`` que retorna um ``float64``.

### 2. Struct ``Retangulo``:
```go
type Retangulo struct {
    Altura  float64
    Largura float64
}
```
- ``Retangulo`` é uma struct que armazena a altura e a largura de um retângulo.

### 3. Struct ``Circulo``:
```go
type Circulo struct {
    Raio float64
}
```
- ``Circulo`` é uma struct que armazena o raio de um círculo.

### 4. Implementação do Método ``area`` para ``Retangulo``:
```go
func (r Retangulo) area() float64 {
    return r.Altura * r.Largura
}
```
- Este método calcula a área de um retângulo multiplicando a altura pela largura.

### 5. Implementação do Método ``area`` para ``Circulo``:
```go
func (c Circulo) area() float64 {
    return math.Pi * math.Pow(c.Raio, 2)
}
```
- Este método calcula a área de um círculo usando a fórmula matemática ``π * r²``.

### 6. Função ``escreverArea``:
```go
func escreverArea(a Area) {
    fmt.Printf("A forma da área é: %0.2f\n", a.area())
}
```
- Esta função recebe um argumento do tipo ``Area`` e chama o método ``area()`` para calcular e imprimir a área da forma geométrica.

### 7. Função ``main``:
- No ``main``, criamos um retângulo com altura 10 e largura 15, e um círculo com raio 5. Em seguida, chamamos ``escreverArea`` para ambos, calculando e exibindo as áreas.

### Saída Esperada
```makefile
A forma da área é: 150.00
A forma da área é: 78.54
```

## 🔍 Conceitos Importantes
### 1. Interface
- Em Go, uma **interface** define um conjunto de métodos que um tipo deve implementar.
- Qualquer tipo que implemente os métodos de uma interface é implicitamente considerado como um tipo daquela interface.
#### Exemplo:
```go
type Area interface {
    area() float64
}
```

### 2. Polimorfismo
- Interfaces permitem **polimorfismo** em Go, já que diferentes tipos podem implementar a mesma interface e serem tratados da mesma forma.
#### Exemplo:
```go
func escreverArea(a Area) {
    fmt.Printf("A forma da área é: %0.2f\n", a.area())
}
```
- Aqui, tanto ``Retangulo`` quanto ``Circulo`` podem ser passados para a função ``escreverArea`` pois ambos implementam o método ``area()`` exigido pela interface ``Area``.

### 3. Implementação Implícita
- Em Go, não é necessário declarar explicitamente que um tipo implementa uma interface. Se o tipo possui todos os métodos da interface, ele já é considerado uma implementação da mesma.
#### Exemplo:
```go
func (r Retangulo) area() float64 {
    return r.Altura * r.Largura
}
```

## 🛡 Boas Práticas:
- **Defina Interfaces Pequenas**: É uma boa prática criar interfaces pequenas, que definam apenas o comportamento necessário para um contexto específico.
```go
type Area interface {
    area() float64
}
```

- **Utilize Interfaces como Abstração**: Utilize interfaces para abstrair comportamento e evitar acoplamento direto a tipos concretos, facilitando a substituição ou modificação futura.
```go
func escreverArea(a Area) {
    fmt.Printf("A forma da área é: %0.2f\n", a.area())
}
```

- **Implemente Interfaces Implicitamente**: Em Go, a implementação implícita de interfaces simplifica a reutilização de código e evita dependências desnecessárias entre os tipos.