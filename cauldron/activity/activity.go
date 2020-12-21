package activity

import "time"

// Expense dao
type Expense struct {
	Date     time.Time `db:"time"`
	Amount   int32     `db:"amount"`
	Type     string    `db:"type"`
	ItemName string    `db:"name"`
}
