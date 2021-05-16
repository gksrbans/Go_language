package main

import (
	"fmt"

	"github.com/gksrbans/learngo/mydict"

	"github.com/gksrbans/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("kyumoon")
	account.Deposit(10)

	// Go언어는 err 처리를 강제하기 때문에 개꿀이다.
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println("------------------------------------")
	// String Method
	fmt.Println(account)
	fmt.Println("------------------------------------")

	// Dictionary Part #2
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	err2 := dictionary.Add("two", "testing...")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(dictionary)
	}

	// Update part # 2-1
	err3 := dictionary.Update("two", "Updeated!")
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(dictionary)
	}

	// Delete part # 2-2
	dictionary.Delete("two")
	fmt.Println(dictionary)

}
