package accounts

// Account struct
type Account struct {
	Owner   string
	Balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{
		Owner:   owner,
		Balance: 0,
	}
	return &account
}
