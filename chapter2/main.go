package main

import (
	"fmt"
	"nomad-go/chapter2/banking"
)

func main() {
	// 이 상태는 모든 클라이언트에게 열려있는 상태 -> 누구나 값을 변경할 수 있다.
	//account := banking.Account{Owner: "hoo", Balance: 1000000}
	//account.Balance = 100
	//fmt.Println(account)

	account := banking.NewAccount("hoo")
	fmt.Println(account)

}

// Golang 은 생성자가 없기 때문에 function 으로 생성하거나 struct 를 만들어야한다.
