package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nicolas-martin/int/internal/decider"
	"github.com/nicolas-martin/int/internal/parser"
	"github.com/nicolas-martin/int/internal/repository"
	"github.com/nicolas-martin/int/internal/types"
	"github.com/nicolas-martin/int/internal/util"
)

func main() {
	repo := repository.NewRepository()
	decider := decider.NewDecider(repo)

	f := util.TryOpen("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	responses := make([]*types.DepositResponse, 0)
	for scanner.Scan() {
		deposit, err := parser.ParseString(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		resp, err := decider.ProcessDeposit(deposit)
		if err != nil {
			log.Fatal(err)
		}

		if resp.Accepted {
			err := repo.Create(deposit)
			if err != nil {
				log.Print(err)
			}

		}

		responses = append(responses, resp)
		printResp(resp)

	}
}

func printResp(resp *types.DepositResponse) {
	respStr, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(respStr))
}
