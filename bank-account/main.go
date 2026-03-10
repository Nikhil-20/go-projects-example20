package main

import (
	"bank-account-management/accountmanager"
	"bank-account-management/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	sa := accountmanager.InitializeStoreData()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(" Enter any of the Options ")
	fmt.Println(" Enter 1 to Add Account ")
	fmt.Println(" Enter 2 Deposit Amount")
	fmt.Println(" Enter 3 Get Account Details")
	fmt.Println(" Enter 4 Withdraw Amount")
	fmt.Println(" Enter 0 to Exit")

	for {
		fmt.Println("\n Enter Options:")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		options, _ := strconv.Atoi(input)

		switch options {

		case 0:
			fmt.Println("Exiting..")
			return
		case 1:

			fmt.Println("Enter Name:")
			scanner.Scan()
			name := scanner.Text()

			fmt.Println("Enter Age:")
			scanner.Scan()
			age, _ := strconv.Atoi(scanner.Text())

			fmt.Println("Enter Address:")
			scanner.Scan()
			address := scanner.Text()

			fmt.Println("Enter Nationality:")
			scanner.Scan()
			nationality := scanner.Text()

			fmt.Println("Enter Account Type (Savings/Current):")
			scanner.Scan()

			accountType, err := model.ValidateAccountType(scanner.Text())
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Enter Aadhaar Number:")
			scanner.Scan()
			aadharNumber := scanner.Text()

			accdetails := model.AccountManager{
				Name:           name,
				Age:            age,
				Address:        address,
				Nationality:    nationality,
				Accounttype:    accountType,
				Aadhaar_number: aadharNumber,
			}
			id, err := sa.CreateAccount(accdetails)
			if err != nil {
				fmt.Println("Error creating account:", err)
			} else {
				fmt.Printf("Account created successfully! Your Account ID: %d\n", id)
			}

		case 2:

			fmt.Println("Enter Account Number:")
			scanner.Scan()
			accNum, _ := strconv.Atoi(scanner.Text())

			fmt.Println("Enter Amount to Deposit:")
			scanner.Scan()

			amount, _ := strconv.ParseFloat(scanner.Text(), 64)

			amt, err := sa.DepositAmount(accNum, amount)

			if err != nil {
				fmt.Println("Error in depositing the amount:", err)

			} else {
				fmt.Printf("Deposited the amount %d successfully", amt)
			}

		case 3:
			fmt.Println("Enter Account Number:")
			scanner.Scan()
			accNum, _ := strconv.Atoi(scanner.Text())

			accmanager, err := sa.GetAccountDetails(accNum)
			if err != nil {
				fmt.Println("Error in getting account details:", err)

			} else {
				fmt.Println("Account details:", accmanager)
			}
		case 4:

			fmt.Println("Enter Account Number:")
			scanner.Scan()
			accNum, _ := strconv.Atoi(scanner.Text())

			fmt.Println("Enter Amount to Withdraw:")
			scanner.Scan()
			amount, _ := strconv.ParseFloat(scanner.Text(), 64)

			amt, err := sa.WithdrawAmount(accNum, amount)

			if err != nil {
				fmt.Println("Error in withdrawing the amount:", err)

			} else {
				fmt.Printf("Withdrawn the amount successfully, remaining balance %d", amt)
			}

		default:
			fmt.Println("Invalid options please tru again")

		}
	}

}
