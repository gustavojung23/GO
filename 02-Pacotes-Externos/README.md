# 02 - Pacotes Externos

## 🌐 Introdução aos Pacotes Externos em Go

Em Go, além dos pacotes padrão fornecidos pela linguagem, você pode utilizar **pacotes externos** para adicionar funcionalidades adicionais ao seu projeto. Pacotes externos são desenvolvidos por terceiros e podem ser facilmente integrados ao seu código, ampliando as capacidades da linguagem sem a necessidade de reinventar a roda.

Este documento explica como utilizar pacotes externos em Go, utilizando o pacote [`github.com/badoux/checkmail`](https://github.com/badoux/checkmail) como exemplo. Este pacote é útil para validar formatos de e-mails.

## 🛠 Exemplo de Código

A seguir, um exemplo simples que demonstra como importar e utilizar um pacote externo em Go.

### Código de Exemplo

```go
package main

import (
    "fmt"
    "github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello World")

	erro := checkmail.ValidateFormat("teste@teste.com")
	fmt.Println(erro)

	erro2 := checkmail.ValidateFormat("123")
	fmt.Print(erro2)
}
```

### Explicação do Código
1. **Declaração do Pacote Principal**
```go
package main
```
- Indica que este arquivo pertence ao pacote main.
- O pacote main é o ponto de entrada de um programa Go executável.

2. **Importação de Pacotes**
```go
import (
	"fmt"

	"github.com/badoux/checkmail"
)
```
- `fmt`: Pacote padrão do Go para formatação de entrada e saída.
- `github.com/badoux/checkmail`: Pacote externo para validação de formatos de e-mail.

3. **Função Principal**
```go
func main() {
	fmt.Println("Hello World")

	erro := checkmail.ValidateFormat("teste@teste.com")
	fmt.Println(erro)

	erro2 := checkmail.ValidateFormat("123")
	fmt.Print(erro2)
}
```
- `fmt.Println("Hello World")`: Imprime "Hello World" no console.
- `checkmail.ValidateFormat("teste@teste.com")`: Valida o formato do e-mail fornecido.
- `checkmail.ValidateFormat("123")`: Tenta validar uma string que não é um e-mail válido, retornando um erro.

## 📁 Estrutura de Diretórios
```go
02-Pacotes/
├── main.go
└── go.mod
```

## 🚀 Como Executar o Exemplo
1. **Clonar o Repositório**
Se ainda não clonou o repositório principal, faça isso primeiro:
```bash
git clone git clone https://github.com/gustavojung23/GO.git
cd GO/02-Pacotes-Externos
```

2. **Inicializar um Módulo Go**
Para gerenciar pacotes externos, é recomendado utilizar Go Modules. Inicialize um módulo Go na subpasta:
```bash
go mod init https://github.com/gustavojung23/GO/02-Pacotes-Externos
```

3. **Adicionar o Pacote Externo**
Use o comando `go get` para adicionar o pacote externo ao seu projeto:
```bash
go get github.com/badoux/checkmail
```
Isso adicionará o pacote ao seu go.mod e baixará as dependências necessárias.

4. **Executar o Programa**
Execute o programa utilizando o comando::
```go
go run main.go
```
**Saída Esperada:**
Execute o programa utilizando o comando:
```go
Hello World
<nil>
invalid email format
```

- `<nil>` indica que o primeiro e-mail está no formato correto.
- `invalid email format` indica que a segunda string não é um e-mail válido.

## 🔍 Conceitos Importantes
- **Go Modules**
    - **Go Modules** facilitam o gerenciamento de dependências em projetos Go.
    - Permitem especificar versões exatas de pacotes externos, garantindo consistência no ambiente de desenvolvimento.

- **Importação de Pacotes Externos**
    - Utilize o caminho completo do repositório (por exemplo, `github.com/badoux/checkmail`) para importar pacotes externos.

- **Gerenciamento de Dependências**
    - O arquivo go.mod gerencia as dependências do seu projeto.
    - O comando go get adiciona pacotes externos ao seu módulo.

## ✅ Próximos Passos
Após entender como utilizar pacotes externos, você pode explorar tópicos avançados, como:

- **Versionamento de Dependências**
    - Aprenda a especificar versões exatas de pacotes para garantir estabilidade no seu projeto.

- **Criação de Seus Próprios Pacotes Externos**
    - Desenvolva e publique seus próprios pacotes para reutilização em múltiplos projetos ou para a comunidade.

- **Testes com Pacotes Externos**
    - Implemente testes unitários que envolvem pacotes externos, garantindo que as integrações funcionem conforme esperado.

## 🛡 Boas Práticas
- **Verifique as Dependências**
    - Sempre verifique a reputação e a manutenção dos pacotes externos antes de adicioná-los ao seu projeto.

- **Mantenha as Dependências Atualizadas**
    - Regularmente atualize seus pacotes externos para obter melhorias e correções de segurança.

- **Isolamento de Dependências**
    - Utilize ferramentas como go mod tidy para remover dependências não utilizadas e manter o go.mod limpo.