package main

// importing fmt and encoding/json pacakges

//AccountDetails struct
type AccountDetails struct {
	id          string
	accountType string
}

//Account struct
type AccountP struct {
	details      *AccountDetails
	CustomerName string
}

// Account class method setDetails
func (account *AccountP) setDetails(id string, accountType string) {

	account.details = &AccountDetails{id, accountType}
}

//Account class method getId
func (account *AccountP) getId() string {

	return account.details.id
}

//Account class method getAccountType
func (account *AccountP) getAccountType() string {

	return account.details.accountType
}

// func main() {

// 	var account *AccountP = &AccountP{CustomerName: "John Smith"}
// 	account.setDetails("4532", "current")

// 	jsonAccount, _ := json.Marshal(account)
// 	fmt.Println("Private Class hidden", string(jsonAccount))
// 	fmt.Println(*account)

// 	fmt.Println("Account Id", account.getId())

// 	fmt.Println("Account Type", account.getAccountType())

// }
