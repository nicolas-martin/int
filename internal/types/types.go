package types

import "time"

// Deposit is the input message of a Deposit action
type Deposit struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	LoadAmount float64   `json:"load_amount"`
	Time       time.Time `json:"time"`
}

// DepositResponse is the response of a deposit
type DepositResponse struct {
	ID         int  `json:"id"`
	CustomerID int  `json:"customer_id"`
	Accepted   bool `json:"accepted"`
}

// Rules represents specific transaction rules
// Could have chosen to do map[time.Duration]float64
// if we wanted to have more control over the "look back
// period". For now I'll use set intervals
type Rules struct {
	DayLimit              float64
	WeekLimit             float64
	TransationLimitPerDay int
}
