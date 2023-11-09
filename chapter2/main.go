package main

import (
	"fmt"
	"log"
	"nomad-go/chapter2/accounts"
)

func main() {
	account := accounts.NewAccount("hoo")
	account.Deposit(30)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())

	// GO 는 보안을 위해 복사본을 만든다. 그걸 원치 않을 때 포인터를 사용하자.
}
