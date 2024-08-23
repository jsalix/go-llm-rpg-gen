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
	worldOverviewArg := "the_lands_of_Skyrim,_near_Windhelm_and_chaos-ridden_Winterhold20240822174006.txt"
	characterArg := "Theldor_the_Unlucky20240822174822.txt"
	promptArg := "Theldor talking about the time he got drunk and lost his temper, bringing shame upon himself and his family"

	fmt.Printf("Generating character memory given prompt: %s\n\n", promptArg)

	worldOverviewPath, err := filepath.Abs("../../output/worldoverview/" + worldOverviewArg)
	if err != nil {
		fmt.Println(err)
		return
	}

	worldOverview, err := os.ReadFile(worldOverviewPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	characterProfilePath, err := filepath.Abs("../../output/character/" + characterArg)
	if err != nil {
		fmt.Println(err)
		return
	}

	characterProfile, err := os.ReadFile(characterProfilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	characterMemoryPrompt, err := gen.CharacterMemoryPrompt(string(worldOverview), string(characterProfile), promptArg)
	if err != nil {
		fmt.Println(err)
		return
	}

	memory, err := api.Generate(characterMemoryPrompt)
	if err != nil {
		fmt.Println(err)
		return
	}

	path, err := filepath.Abs("../../output/memory/" + strings.ReplaceAll(promptArg, " ", "_") + time.Now().Format("20060102150405") + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile(path, []byte(memory), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(memory)
}
