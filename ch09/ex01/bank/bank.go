package bank

var deposits = make(chan int)  // send amount to deposit
var balances = make(chan int)  // receive balance
var withdraws = make(chan int) // receive balance
var withdrawOk = make(chan bool)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	withdraws <- amount
	return <-withdrawOk
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case amount := <-withdraws:
			if balance >= amount {
				balance -= amount
				withdrawOk <- true
			} else {
				withdrawOk <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
