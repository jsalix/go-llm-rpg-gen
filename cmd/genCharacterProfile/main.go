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
	promptArg := "Theldor the Unlucky"

	fmt.Printf("Generating character profile given prompt: %s\n\n", promptArg)

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

	characterProfilePrompt, err := gen.CharacterProfilePrompt(string(worldOverview), promptArg)
	if err != nil {
		fmt.Println(err)
		return
	}

	characterProfile, err := api.Generate(characterProfilePrompt)
	if err != nil {
		fmt.Println(err)
		return
	}

	path, err := filepath.Abs("../../output/character/" + strings.ReplaceAll(promptArg, " ", "_") + time.Now().Format("20060102150405") + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile(path, []byte(characterProfile), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(characterProfile)
}
