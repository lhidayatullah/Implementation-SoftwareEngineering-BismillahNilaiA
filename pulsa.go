package main

import (
	"fmt"
)

var balance int = 50000

func main() {
	var choice int

	for {

		fmt.Println("=====================================")
		fmt.Println("\nWelcome to the *858# Service")
		fmt.Println("1. Transfer Pulsa")
		fmt.Println("2. Check Balance")
		fmt.Println("3. Exit")
		fmt.Println(" ")
		fmt.Println("=====================================")
		fmt.Println("\n")

		fmt.Print("Please choose an option: ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			transferPulsa()
		case 2:
			checkBalance()
		case 3:
			fmt.Println("Thank you for using our service!")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

func transferPulsa() {
	var phoneNumber string
	var amount int
	var done bool = false
	var done1 bool = false
	var done2 bool = false
	for done == false {
		for done1 == false {
			fmt.Print("Enter destination phone number (Example : 08xxx and 628xxx) : ")
			fmt.Scan(&phoneNumber)
			if len(phoneNumber) >= 3 && phoneNumber[1] == '8' {
				done1 = true
			} else {
				fmt.Println("Phone Number not valid!")
			}
		}

		for done2 == false {
			fmt.Print("Enter amount to transfer: ")
			fmt.Scan(&amount)
			if amount > balance {
				fmt.Println("Not enough Money!")
			} else {
				done2 = true
				balance = balance - amount
			}
		}

		fmt.Printf("Transferring %d pulsa to %s...\n", amount, phoneNumber)
		fmt.Println("Transfer successful!")
		fmt.Println("Remaining balance:", balance)
		done = true
	}

}

func checkBalance() {
	fmt.Println("Your current balance is", balance)
}
