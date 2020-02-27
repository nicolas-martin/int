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

	f := tryOpen("input.txt")
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

	verify(responses)
}

func verify(responses []*types.DepositResponse) {
	f := tryOpen("output.txt")
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
	for _, v := range responses {
		if v.Accepted != outputResponses[v.ID].Accepted {
			// fmt.Printf("My output %+v, expected %+v \r\n", v, outputResponses[v.ID])
			count++
		}

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

func tryOpen(path string) *os.File {
	var f *os.File
	f, err := os.Open(fmt.Sprintf("../%s", path))
	if err != nil {
		f, err = os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		return f

	}
	return f

}
