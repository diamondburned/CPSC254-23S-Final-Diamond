package bank

// User represents a user of the bank.
type User struct {
	Username string         `json:"username"`
	Password string         `json:"-"` // bcrypt-hashed
	Activity []UserActivity `json:"activity"`
	Balance  int            `json:"balance"`
}

// UserActivity represents an activity of a user.
type UserActivity struct {
	Type   string `json:"type"` // "deposit", "withdraw", "transfer"
	Amount int    `json:"amount"`
}

type UserService interface {
	// Login returns a user if the username and password are correct.
	Login(username, password string) (User, error)
}

type BankService interface {
	// Deposit deposits the amount to the user's balance.
	Deposit(user User, amount int) error
	// Withdraw withdraws the amount from the user's balance.
	Withdraw(user User, amount int) error
	// Transfer transfers the amount from the user to the recipient.
	// Transfer(user User, recipient string, amount int) error
}
