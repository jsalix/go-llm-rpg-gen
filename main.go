package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jsalix/go-llm-rpg-gen/api"
	"github.com/jsalix/go-llm-rpg-gen/gen"
)

func main() {
	seedConcept := strings.Join(os.Args[1:], " ")

	worldOverviewPrompt, err := gen.WorldOverviewPrompt(seedConcept)
	if err != nil {
		fmt.Println("unable to read prompt file")
		return
	}

	worldOverviewResponse, err := api.Generate(fmt.Sprintf(string(worldOverviewPrompt), seedConcept))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Response: " + worldOverviewResponse)
}
