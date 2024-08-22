package gen

import (
	"fmt"
	"os"
	"path/filepath"
)

func WorldOverviewPrompt(seedConcept string) (string, error) {
	path, err := filepath.Abs("../../gen/prompts/world-overview.txt")
	if err != nil {
		fmt.Println("unable to read prompt file")
		return "", err
	}

	worldOverviewPrompt, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("unable to read prompt file")
		return "", err
	}
	return fmt.Sprintf(string(worldOverviewPrompt), seedConcept), nil
}

// func ExtractWorldTitlePrompt(worldOverview string) (string, error) {
// 	path, err := filepath.Abs("../../gen/prompts/extract-world-name.txt")
// 	if err != nil {
// 		fmt.Println("unable to read prompt file")
// 		return "", err
// 	}

// 	worldTitlePrompt, err := os.ReadFile(path)
// 	if err != nil {
// 		fmt.Println("unable to read prompt file")
// 		return "", err
// 	}
// 	return fmt.Sprintf(string(worldTitlePrompt), worldOverview), nil
// }
