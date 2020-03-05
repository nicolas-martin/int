package decider

import (
	"bufio"
	"log"
	"testing"

	"github.com/nicolas-martin/int/internal/parser"
	"github.com/nicolas-martin/int/internal/repository"
	"github.com/nicolas-martin/int/internal/types"
	"github.com/nicolas-martin/int/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestDecider_Acceptance(t *testing.T) {
	repo := repository.NewRepository()
	decider := NewDecider(repo)

	f := util.TryOpen("../input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	responses := make([]*types.DepositResponse, 0)
	for scanner.Scan() {
		deposit, err := parser.ParseString(scanner.Text())
		assert.Nil(t, err, "Fail to parse input string")

		resp, err := decider.ProcessDeposit(deposit)
		if err != nil {
			log.Fatal(err)
		}

		if resp.Accepted {
			err := repo.Create(deposit)
			assert.Nil(t, err, "Fail to create deposit")

		}
		responses = append(responses, resp)

	}

	// Read output results
	fo := util.TryOpen("../output.txt")
	defer fo.Close()
	scannerOut := bufio.NewScanner(fo)
	outputResponses := make(map[int]*types.DepositResponse)
	for scannerOut.Scan() {
		depositResponse, err := parser.ParseOutput(scannerOut.Text())
		assert.Nil(t, err, "Fail to parse output")
		outputResponses[depositResponse.ID] = depositResponse
	}

	// Using the transactionID, find the expected result in the map
	// and compare it against the generated response
	for _, v := range responses {
		assert.Equal(t, outputResponses[v.ID], v, "response mismatch")
	}
}
