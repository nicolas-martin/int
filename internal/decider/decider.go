package decider

import (
	"github.com/nicolas-martin/int/internal/repository"
	"github.com/nicolas-martin/int/internal/types"
)

// Decider evaluates the message and compares the rules
type Decider struct {
	Rules      *types.Rules
	repository repository.IRepository
}

// NewDecider creates a new instance of decider
func NewDecider(repo repository.IRepository) *Decider {
	r := &types.Rules{
		DayLimit:              5_000.0,
		WeekLimit:             20_000.0,
		TransationLimitPerDay: 3,
	}

	return &Decider{Rules: r, repository: repo}

}

// IsTransactionAllowed checks if the deposit is allowed given a set of rules.
func (d *Decider) IsTransactionAllowed(deposit *types.Deposit) (*types.DepositResponse, error) {
	deposits, err := d.repository.RetrieveAll(deposit.ID)
	response := &types.DepositResponse{
		ID:         deposit.ID,
		CustomerID: deposit.CustomerID,
		Accepted:   false,
	}

	if err != nil {
		return nil, err
	}

	lastDay := make([]*types.Deposit, 0)

	// Check for lastday first
	for _, v := range deposits {
		delta := deposit.Time.Sub(v.Time).Hours()
		if delta < 24 {
			lastDay = append(lastDay, v)
		}
	}

	totalDaily := 0.0
	for _, v := range lastDay {
		totalDaily += v.LoadAmount
	}

	lastWeek := make([]*types.Deposit, 0)
	for _, v := range deposits {
		delta := deposit.Time.Sub(v.Time).Hours() / 24
		if delta < 7 {
			lastWeek = append(lastWeek, v)
		}
	}

	totalWeek := 0.0
	for _, v := range lastWeek {
		totalWeek += v.LoadAmount
	}

	if totalDaily+deposit.LoadAmount > d.Rules.DayLimit || len(lastDay) > d.Rules.TransationLimitPerDay || totalWeek+deposit.LoadAmount > d.Rules.WeekLimit {
		return response, nil
	}
	response.Accepted = true
	return response, nil

}
