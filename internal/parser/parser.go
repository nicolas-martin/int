package parser

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/nicolas-martin/int/internal/types"
)

type tmpDeposit struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	LoadAmount string `json:"load_amount"`
	Time       string `json:"time"`
}

// ParseString parses []byte to deposit
func ParseString(s string) (*types.Deposit, error) {

	tmpD := &tmpDeposit{}

	err := json.Unmarshal([]byte(s), tmpD)
	if err != nil {
		return nil, errors.Wrap(err, "Could not parse deposit")
	}
	// Remove the $
	tmpD.LoadAmount = strings.Replace(tmpD.LoadAmount, "$", "", 1)

	idInt, err := strconv.Atoi(tmpD.ID)
	if err != nil {
		return nil, errors.Wrap(err, "Could not convert id")
	}

	customerIDInt, err := strconv.Atoi(tmpD.CustomerID)
	if err != nil {
		return nil, errors.Wrap(err, "Could not convert customerID")
	}

	loadAmountFloat, err := strconv.ParseFloat(tmpD.LoadAmount, 64)
	if err != nil {
		return nil, errors.Wrap(err, "Could not conver load amount")
	}

	//Mon Jan 2 15:04:05 -0700 MST 2006
	// RFC3339     = "2006-01-02T15:04:05Z07:00"

	parsedTime, err := time.Parse(time.RFC3339, tmpD.Time)
	if err != nil {
		return nil, errors.Wrap(err, "Could not parse time")
	}

	d := &types.Deposit{
		ID:         idInt,
		CustomerID: customerIDInt,
		LoadAmount: loadAmountFloat,
		Time:       parsedTime,
	}

	return d, nil

}

type tmpDepositResponse struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Accepted   bool   `json:"accepted"`
}

// ParseOutput for verification
func ParseOutput(s string) (*types.DepositResponse, error) {

	tmpDR := &tmpDepositResponse{}

	err := json.Unmarshal([]byte(s), tmpDR)
	if err != nil {
		return nil, errors.Wrap(err, "Could not parse deposit")
	}

	idInt, err := strconv.Atoi(tmpDR.ID)
	if err != nil {
		return nil, errors.Wrap(err, "Could not convert id")
	}

	customerIDInt, err := strconv.Atoi(tmpDR.CustomerID)
	if err != nil {
		return nil, errors.Wrap(err, "Could not convert customerID")
	}

	dr := &types.DepositResponse{
		ID:         idInt,
		CustomerID: customerIDInt,
		Accepted:   tmpDR.Accepted,
	}

	return dr, nil
}
