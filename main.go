package main

import (
	"fmt"
	"simple/payment-wallet/ledger"
)

func main() {
	l := ledger.NewLedger()

	aliAcc := l.CreateAccount("Ali")
	mahmoodAcc := l.CreateAccount("Mahmood")

	fmt.Printf("create two account in the ledger: 1. %v, 2. %v \n", aliAcc.Name, mahmoodAcc.Name)

	fmt.Printf("toping up %v a 1000 halala \n", aliAcc.Name)
	l.TopUpAccount(aliAcc, 1000)
	fmt.Printf("Account '%v' current balance %v \n", aliAcc.Name, aliAcc.Balance)

	fmt.Printf("Transfer 5000 halala from %v to %v \n", aliAcc.Name, mahmoodAcc.Name)
	_, err := l.FundTransafer(aliAcc, mahmoodAcc, 5000)
	if err != nil {
		fmt.Println("Error Transfering fund: ", err.Error())
	}
	fmt.Printf("%v & %v balance are %v, %v \n", aliAcc.Name, mahmoodAcc.Name, aliAcc.Balance, mahmoodAcc.Balance)

	fmt.Printf("Transfer 7000 halala from %v to %v \n", aliAcc.Name, mahmoodAcc.Name)
	_, err = l.FundTransafer(aliAcc, mahmoodAcc, 7000)
	if err != nil {
		fmt.Println("Error Transfering fund: ", err.Error())
	}
	fmt.Printf("%v & %v balance are %v, %v \n", aliAcc.Name, mahmoodAcc.Name, aliAcc.Balance, mahmoodAcc.Balance)
}
