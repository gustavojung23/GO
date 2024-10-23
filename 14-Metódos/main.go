package main

import (
	"fmt"
)

type User struct {
	name string
	age  uint8
}

func (u User) saveDB() {
	fmt.Printf("Salvando dados do Usuário %s com idade %d anos\n", u.name, u.age)
}

func (u *User) alterDate() {
	u.age++
	fmt.Printf("Alterando a Idade do Usuário: %d", u.age)
}

func main() {
	usuario := User{"John", 20}
	usuario.saveDB()
	usuario.alterDate()
}
