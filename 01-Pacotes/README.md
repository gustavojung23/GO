# Pacotes

## Introdução aos Pacotes em Go

Em Go, **pacotes** (packages) são a principal forma de organizar e modularizar o código. Eles permitem que você divida seu programa em partes menores e reutilizáveis, facilitando a manutenção e a escalabilidade do projeto.

Cada arquivo Go começa com uma declaração de pacote que indica a qual pacote o arquivo pertence. O pacote `main` é especial e indica que o arquivo pode ser compilado e executado como um programa independente.

## Exemplo do Código
```go
package main

import (
	"fmt"
	"modules/auxiliar"
)

func main() {
	fmt.Println("Hello World!")
	auxiliar.Escrever()
}
```

## Explicação do Código

1. **Declaração do Pacote Principal**
```go
package main
```
- Indica que este arquivo pertence ao pacote `main`.
- O pacote `main` é o ponto de entrada de um programa Go executável.

2. **Declaração do Pacote Principal**
```go
import (
	"fmt"
	"modules/auxiliar"
)
```
- `fmt`: Pacote padrão do Go para formatação de entrada e saída.
- `modules/auxiliar`: Um pacote personalizado localizado no diretório `/auxiliar`.

3. **Função Principal**
```go
func main() {
	fmt.Println("Hello World!")
	auxiliar.Escrever()
}
```
- `fmt.Println("Hello World!")`: Imprime "Hello World!" no console.
- `auxiliar.Escrever()`: Chama a função `Escrever` do pacote `auxiliar`.

## Criando um Pacote Personalizado
Vamos criar o pacote `auxiliar` utilizado no exemplo acima.

1. **Estrutura de Diretórios**:
```go
1-Pacotes/
├── main.go
└── auxiliar/
	└── auxiliar.go
```

2. **Código do Pacote** `auxiliar`
```go
package auxiliar

import "fmt"

func Escrever() {
	fmt.Println("Pacote auxiliar")
}
```

- Declaração do Pacote
	* Define que este arquivo pertence ao pacote auxiliar.
```go
package auxiliar
```


- Importação do Pacote `fmt`
	* Utiliza o pacote padrão fmt para formatação de saída.
```go
import "fmt"
```
### Função Exportada `Escrever`
- A função Escrever está exportada porque começa com letra maiúscula.
- Funções exportadas podem ser acessadas de outros pacotes.
```go
func Escrever() {
	fmt.Println("Esta é uma função do pacote auxiliar.")
}
```

## Executando o Código
Para executar o exemplo completo:
1. **Navegue até o Diretório do Projeto**

```bash
cd 1-Pacotes
```

2. **Execute o Programa**
```bash
go run main.go
```

3. **Saída Esperada**:
```bash
Hello World!
Esta é uma função do pacote auxiliar.
```

## 🔍 Conceitos Importantes
- **Pacotes Padrão vs. Pacotes Personalizados**
	* Pacotes Padrão: Incluídos na biblioteca padrão do Go (como `fmt`, `math`, `net`, etc.).
	* Pacotes Personalizados: Criados pelo desenvolvedor para organizar código específico do projeto.

- **Exportação de Identificadores**
	* Identificadores (funções, variáveis, structs, etc.) que começam com letra maiúscula são exportados e acessíveis por outros pacotes.
	* Identificadores que começam com letra minúscula são não exportados e privados ao pacote.

- **Gerenciamento de Dependências**
	* Utilize ferramentas como `go mod` para gerenciar dependências de pacotes externos e versões.

## 🧰 Ferramentas Úteis
- **Go Modules**
	* Facilita o gerenciamento de dependências e versões de pacotes.
	* Inicialize um módulo com:

```bash
go mod init
```

- **GoDoc**
	* Gera documentação a partir dos comentários do código.
	* Adicione comentários claros e concisos para facilitar a geração de documentação.


## ✅ Próximos Passos
Após entender os conceitos básicos de pacotes, você pode explorar tópicos avançados, como:

- Criação de Pacotes Reutilizáveis
- Gerenciamento de Dependências com go mod
- Testes em Pacotes Go
- Documentação com GoDoc


## 📌 Dicas Adicionais
- **Organização de Código**: Mantenha uma estrutura de diretórios lógica e consistente para facilitar a navegação e manutenção.
- **Nomes Significativos**: Utilize nomes claros e descritivos para pacotes e funções.
- **Modularização**: Evite pacotes excessivamente grandes; divida funcionalidades em pacotes menores e coesos.
- **Documentação**: Comente seu código para melhorar a legibilidade e facilitar a colaboração.
