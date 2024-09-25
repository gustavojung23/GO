# Go Projeto de Estudo

Este repositório contém estudos e exemplos de código em Go, organizados em pastas:

## Estrutura do Projeto

### 1. Pacotes
Nesta pasta estão exemplos de código que utilizam pacotes nativos do Go, explorando conceitos básicos.

### 2. Pacotes Externos
Aqui estão exemplos de como integrar pacotes externos no Go, utilizando bibliotecas populares para estender as funcionalidades da linguagem.

### 3. Variáveis
Esta pasta contém exemplos sobre o uso de variáveis no Go.

### 4. Tipos de Dados
Esta pasta contém exemplos de diferentes tipos de dados numéricos em Go, com foco em inteiros (com ou sem sinal) e números de ponto flutuante.

#### Numéricos

Go oferece uma ampla variedade de tipos de dados numéricos para atender diferentes necessidades de precisão e uso da memória. Aqui estão os principais tipos numéricos e seus intervalos:

##### Inteiros
- `int`: Inteiro com tamanho dependente da arquitetura do computador (32 ou 64 bits).
- `int8`: Inteiro de 8 bits, varia de -128 a 127.
- `int16`: Inteiro de 16 bits, varia de -32.768 a 32.767.
- `int32`: Inteiro de 32 bits, varia de -2.147.483.648 a 2.147.483.647.
- `int64`: Inteiro de 64 bits, varia de -9.223.372.036.854.775.808 a 9.223.372.036.854.775.807.

##### Inteiros sem sinal
- `uint`: Inteiro sem sinal, com tamanho dependente da arquitetura (32 ou 64 bits).
- `uint8`: Inteiro de 8 bits sem sinal, varia de 0 a 255.
- `uint16`: Inteiro de 16 bits sem sinal, varia de 0 a 65.535.
- `uint32`: Inteiro de 32 bits sem sinal, varia de 0 a 4.294.967.295.
- `uint64`: Inteiro de 64 bits sem sinal, varia de 0 a 18.446.744.073.709.551.615.

##### Tipos especiais
- `rune`: Alias para `int32`, usado para representar caracteres Unicode.
- `byte`: Alias para `uint8`, normalmente utilizado para representar bytes.

##### Pontos Flutuantes
- `float32`: Números com ponto flutuante de 32 bits, com precisão de aproximadamente 6 a 9 dígitos decimais.
- `float64`: Números com ponto flutuante de 64 bits, com precisão de aproximadamente 15 a 17 dígitos decimais.

#### Strings e Caracteres

Go oferece suporte a strings, que são sequências imutáveis de caracteres, e a `rune`, que representa caracteres Unicode de 32 bits. Aqui estão os principais pontos sobre strings e caracteres:

##### Strings
- `string` é uma sequência imutável de bytes. Cada byte representa um caractere codificado em UTF-8.
- Strings em Go são delimitadas por aspas duplas (`"`), e uma vez criadas, não podem ser modificadas.
  
##### Caracteres
- `char` em Go é representado pelo tipo `rune`, que é um alias para `int32` e é usado para representar caracteres Unicode.
- Um `rune` pode armazenar qualquer caractere Unicode e é delimitado por aspas simples (`'`).


---

Este projeto está em desenvolvimento e é voltado para o aprendizado contínuo em Go.
