package main

import (
	"fmt"
	"log"
	"nomad-go/chapter2/banking"
)

func main() {
	// 이 상태는 모든 클라이언트에게 열려있는 상태 -> 누구나 값을 변경할 수 있다.
	//account := banking.Account{Owner: "hoo", Balance: 1000000}
	//account.Balance = 100
	//fmt.Println(account)

	account := banking.NewAccount("hoo")
	fmt.Println("Account : ", account)
	account.Deposit(2000000)
	fmt.Println(account.Balance())
	err := account.Withdraw(1000000)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account) // java 의 .toString() 같은 것 -> Struct 의 String 메서드로 선언

}

// Golang 은 생성자가 없기 때문에 function 으로 생성하거나 struct 를 만들어야한다.
// balance 를 increase 하려면 또 다른 function 을 만들어야한다. 근데 이제 method 라고 부른다.
