package main

type StoryNode struct {
	Text    string            `json:"text"`
	Choices map[string]string `json:"choices"`
}

var story = map[string]StoryNode{
	"start": {
		Text: "You find yourself in a dark forest. Do you want to go left or right?",
		Choices: map[string]string{
			"1": "You walk left and discover a hidden treasure. Do you want to take it or leave it?",
			"2": "You go right and encounter a fierce dragon! Do you want to fight or flee?",
		},
	},
	// More nodes will be added by LLM
}
