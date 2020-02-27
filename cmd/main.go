package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nicolas-martin/int/internal/decider"
	"github.com/nicolas-martin/int/internal/parser"
	"github.com/nicolas-martin/int/internal/repository"
	"github.com/nicolas-martin/int/internal/types"
)

func main() {
	repo := repository.NewRepository()
	decider := decider.NewDecider(repo)

	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	responses := make([]*types.DepositResponse, 0)
	for scanner.Scan() {
		deposit, err := parser.ParseString(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		// rr, _ := json.Marshal(deposit)
		// fmt.Printf("\r\n %v \r\n", string(rr))

		resp, err := decider.IsTransactionAllowed(deposit)
		if err != nil {
			log.Fatal(err)
		}

		if resp.Accepted {
			repo.Create(deposit)
		}

		responses = append(responses, resp)
		// printResp(resp)

	}

	verify(responses)
}

func verify(responses []*types.DepositResponse) {
	f, err := os.Open("../output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	outputResponses := make(map[int]*types.DepositResponse)
	for scanner.Scan() {
		depositResponse, err := parser.ParseOutput(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		outputResponses[depositResponse.ID] = depositResponse
	}
	count := 0
	right := 0
	for _, v := range responses {
		if v.Accepted != outputResponses[v.ID].Accepted {
			// fmt.Printf("got = %+v, exptected = %+v\n", v, outputResponses[v.ID])
			count++
			continue
		}
		right++

	}

	fmt.Printf("Got %v wrong \r\n", count)
}

func printResp(resp *types.DepositResponse) {

	respStr, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(respStr))

}
