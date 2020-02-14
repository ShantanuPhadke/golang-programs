package bitcoin

import (
	"errors"
	"fmt"
)

// CHANGE LOG:
// (1) Started out with a struct and functions that just took in Wallet objects.
//	   Realized that wallet's balanced being referenced inside of the Deposit
//	   function and the Balance function were different.
// (2) Instead started to pass in pointers to Wallet objects and work on those.
//     These pointers help us to work on the objects we desire directly instead of
//     copies of those objects.
// (3) Created a custom Bitcoin type (float64) to use in our Wallet functions. Changed
//     all of the previous float64 references to reference the new Bitcoin type.
// (4) Defined an interface for Stringer with a single method called String()
// (5) Implemented the Stringer interface for the type Bitcoin
//     - Used the new stringer implementation to print Bitcoins in wallet_test.go
// (6) Wrote a Withdraw function that represents withdrawing a certain amount of
// 	   Bitcoin from a given wallet.
// (7) Added in support for withdrawing when a wallet has insufficient funds by
//     forcing the function to return an error.
// (8) Refactored error message outside of function into a variable.
//	   - Used the same error variable in our bitcoin wallet test program as well.

/*Stringer ... interface
String() = returns a string representing the type
*/
type Stringer interface {
	String() string
}

/*Bitcoin ... float64 */
type Bitcoin float64

/*Stringer's String implementation for the Bitcoin type */
func (b Bitcoin) String() string {
	return fmt.Sprintf("%g BTC", b)
}

/*Wallet ...
balance = the current balance in the bitcoin wallet
*/
type Wallet struct {
	balance Bitcoin
}

/*Deposit ... (for Wallets) */
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

/*ErrorInsufficientFunds error for Withdraw method */
var ErrorInsufficientFunds = errors.New("not enough funds to withdraw")

/*Withdraw ... (for Wallets) */
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrorInsufficientFunds
	}
	w.balance -= amount
	return nil
}

/*Balance ... (for Wallets) */
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
