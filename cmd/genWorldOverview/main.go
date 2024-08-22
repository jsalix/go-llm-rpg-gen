package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jsalix/go-llm-rpg-gen/api"
	"github.com/jsalix/go-llm-rpg-gen/gen"
)

func main() {
	seedConcept := strings.Join(os.Args[1:], " ")

	fmt.Printf("Generating world overview given prompt: %s\n\n", seedConcept)

	worldOverviewPrompt, err := gen.WorldOverviewPrompt(seedConcept)
	if err != nil {
		fmt.Println("unable to read prompt file")
		return
	}

	worldOverview, err := api.Generate(worldOverviewPrompt)
	if err != nil {
		fmt.Println(err)
		return
	}

	path, err := filepath.Abs("../../output/worldoverview/" + strings.ReplaceAll(seedConcept, " ", "_") + time.Now().Format("20060102150405") + ".txt")
	if err != nil {
		fmt.Println("unable to create path")
		return
	}

	err = os.WriteFile(path, []byte(worldOverview), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(worldOverview)
}
