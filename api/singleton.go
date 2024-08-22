package api

import "fmt"

var client *KoboldClient

func Generate(prompt string) (string, error) {
	if client == nil {
		var err error
		client, err = NewKoboldClient("http://localhost:5001/api")
		if err != nil {
			fmt.Println("failed to create koboldcpp client")
			panic(err)
		}
	}

	params := &KoboldParams{
		MaxContextLength: 16384,
		MaxLength:        512,
		Temperature:      1.0,
		DynaTempRange:    0,
		TopP:             1,
		MinP:             0.05,
		TopK:             0,
		TopA:             0,
		Typical:          1.0,
		Tfs:              1.0,
		RepPen:           1.0,
		RepPenRange:      1024,
		RepPenSlope:      0,
		SamplerOrder:     []int{6, 0, 1, 3, 4, 2, 5},
		SamplerSeed:      -1,
		StopSequence:     []string{"</output>"},
		BanTokens:        false,
		TrimStop:         true,
	}

	instructionSeq := "<instructions>"
	responseSeq := "</instructions><output>"

	params.Prompt = instructionSeq + string(prompt) + responseSeq
	response, err := client.Generate(params)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	if response.Status != "ok" {
		fmt.Println("received " + response.Status + " status")
		return "", fmt.Errorf("received " + response.Status + " status")
	}

	return response.Text, nil
}
