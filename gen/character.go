package gen

import (
	"fmt"
	"os"
	"path/filepath"
)

func CharacterProfilePrompt(worldOverview string, promptArg string) (string, error) {
	path, err := filepath.Abs("../../gen/prompts/npc-profile.txt")
	if err != nil {
		fmt.Println("unable to read prompt file")
		return "", err
	}

	prompt, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("unable to read prompt file")
		return "", err
	}
	return fmt.Sprintf(string(prompt), worldOverview, promptArg), nil
}

func CharacterMemoryPrompt(worldOverview, characterProfile, promptArg string) (string, error) {
	path, err := filepath.Abs("../../gen/prompts/npc-memory.txt")
	if err != nil {
		fmt.Println("unable to read prompt file")
		return "", err
	}

	prompt, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("unable to read prompt file")
		return "", err
	}
	return fmt.Sprintf(string(prompt), worldOverview, characterProfile, promptArg), nil
}
