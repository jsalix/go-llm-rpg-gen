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
	worldOverviewArg := "the_world_of_dr._suess's_the_lorax_but_if_they_invented_guns20240822160628.txt"
	characterArg := "Mr_Meeks_the_wandering_Turk.txt"
	promptArg := "his last encounter with a human"

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
