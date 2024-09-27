# Go Projeto de Estudo

Este repositório contém estudos e exemplos de código em Go, organizados em pastas:

## 1. Pacotes
Nesta pasta estão exemplos de código que utilizam pacotes nativos do Go, explorando conceitos básicos.

## 2. Pacotes Externos
Aqui estão exemplos de como integrar pacotes externos no Go, utilizando bibliotecas populares para estender as funcionalidades da linguagem.

## 3. Variáveis
Esta pasta contém exemplos sobre o uso de variáveis no Go.

## 4. Tipos de Dados
Esta pasta contém exemplos de diferentes tipos de dados numéricos em Go, com foco em inteiros (com ou sem sinal) e números de ponto flutuante.

### 4.1 Numéricos

Go oferece uma ampla variedade de tipos de dados numéricos para atender diferentes necessidades de precisão e uso da memória. Aqui estão os principais tipos numéricos e seus intervalos:

#### 4.1.1 Inteiros
- `int`: Inteiro com tamanho dependente da arquitetura do computador (32 ou 64 bits).
- `int8`: Inteiro de 8 bits, varia de -128 a 127.
- `int16`: Inteiro de 16 bits, varia de -32.768 a 32.767.
- `int32`: Inteiro de 32 bits, varia de -2.147.483.648 a 2.147.483.647.
- `int64`: Inteiro de 64 bits, varia de -9.223.372.036.854.775.808 a 9.223.372.036.854.775.807.

#### 4.1.2 Inteiros sem sinal
- `uint`: Inteiro sem sinal, com tamanho dependente da arquitetura (32 ou 64 bits).
- `uint8`: Inteiro de 8 bits sem sinal, varia de 0 a 255.
- `uint16`: Inteiro de 16 bits sem sinal, varia de 0 a 65.535.
- `uint32`: Inteiro de 32 bits sem sinal, varia de 0 a 4.294.967.295.
- `uint64`: Inteiro de 64 bits sem sinal, varia de 0 a 18.446.744.073.709.551.615.

#### 4.1.3 Tipos especiais
- `rune`: Alias para `int32`, usado para representar caracteres Unicode.
- `byte`: Alias para `uint8`, normalmente utilizado para representar bytes.

#### 4.1.4 Pontos Flutuantes
- `float32`: Números com ponto flutuante de 32 bits, com precisão de aproximadamente 6 a 9 dígitos decimais.
- `float64`: Números com ponto flutuante de 64 bits, com precisão de aproximadamente 15 a 17 dígitos decimais.

### 4.2 Strings e Caracteres

Go oferece suporte a strings, que são sequências imutáveis de caracteres, e a `rune`, que representa caracteres Unicode de 32 bits. Aqui estão os principais pontos sobre strings e caracteres:

#### 4.2.1 Strings
- `string` é uma sequência imutável de bytes. Cada byte representa um caractere codificado em UTF-8.
- Strings em Go são delimitadas por aspas duplas (`"`), e uma vez criadas, não podem ser modificadas.
  
#### 4.2.2 Caracteres
- `char` em Go é representado pelo tipo `rune`, que é um alias para `int32` e é usado para representar caracteres Unicode.
- Um `rune` pode armazenar qualquer caractere Unicode e é delimitado por aspas simples (`'`).

### 4.3 Booleanos

O tipo `bool` em Go representa um valor booleano, que pode ser **verdadeiro** (`true`) ou **falso** (`false`). Ele é utilizado para controlar a lógica do programa, sendo comumente aplicado em expressões condicionais e estruturas de controle como `if` e `for`.

### 4.4 Erros

Em Go, o tipo `error` é uma interface embutida que representa a condição de erro. É usada para retornar informações sobre falhas em operações ou funções. A interface `error` tem um método, `Error()`, que retorna uma mensagem descritiva sobre o erro.

#### 4.4.1 Criando Erros
Você pode criar erros personalizados usando a função `errors.New` da biblioteca padrão.

### 4.5 Valores Zero

Em Go, todas as variáveis são inicializadas com um **valor zero** padrão se não forem explicitamente inicializadas. O valor zero depende do tipo da variável e garante que elas nunca contenham "lixo" de memória, como em outras linguagens.

#### 4.5.1 Valores Zero por Tipo:
- **Tipos numéricos** (`int`, `float`, etc.): O valor zero é `0`.
- **Strings**: O valor zero é uma string vazia `""`.
- **Booleanos**: O valor zero é `false`.
- **Ponteiros, slices, maps, funções, interfaces e canais**: O valor zero é `nil`.

## 5. Funções
Contém exemplos básicos de como utilizar funções em Go, incluindo:
- Funções que retornam valores
- Funções que retornam múltiplos valores
- Declaração de funções como variáveis

## 6. Operadores
Contém exemplos do uso de operadores em Go.

### 6.1 - Operadores Aritméticos
Os operadores aritméticos em Go são usados para operações matemáticas básicas.

- `+`: Soma
- `-`: Subtração
- `*`: Multiplicação
- `/`: Divisão
- `%`: Módulo (resto da divisão)

### 6.2 - Operadores de Atribuição
Operadores de atribuição são usados para armazenar valores em variáveis.

- `=`: Atribuição simples
- `+=`: Atribuição com soma
- `-=`: Atribuição com subtração
- `=*`: Atribuição com multiplicação
- `=/`: Atribuição com divisão
- `=%`: Atribuição com módulo

### 6.3 - Operadores Relacionais
Operadores relacionais comparam dois valores e retornam um valor booleano (`true` ou `false`).

- `>`: Maior
- `<`: Menor
- `>=`: Maior ou igual
- `<=`: Menor ou igual
- `==`: Igual
- `!=`: Diferente

### 6.4 - Operadores Lógicos
Os operadores lógicos são usados para combinar expressões booleanas.

- `&&`: **E** lógico
- `||`: **OU** lógico
- `!`: Negação

### 6.5 - Operadores Unários
Os operadores unários modificam o valor de uma variável em uma unidade.

- `++` ou `+=`: Incremento
- `--` ou `-=`: Decremento
- `*=`: Multiplicação
- `/=`: Divisão
- `%=`: Operação de módulo

## 7. Structs
Structs são tipos de dados compostos que permitem agrupar variáveis de diferentes tipos sob um único nome.

### 7.1 Definindo uma Struct
No Go, as structs são definidas com a palavra-chave type, seguida pelo nome da struct e sua composição.

### 7.2 Usando uma Struct
Após definir uma struct, você pode criar instâncias dela, atribuir valores aos seus campos e utilizá-los no seu programa.

---

Este projeto está em desenvolvimento e é voltado para o aprendizado contínuo em Go.
