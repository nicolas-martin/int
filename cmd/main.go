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

		printResp(resp)
		_ = resp

	}

}

func printResp(resp *types.DepositResponse) {

	respStr, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(respStr))

}
