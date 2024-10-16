package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	currentNode := "start"
	memory := &Memory{}
	fmt.Println("Loading, please wait...")

	for {
		node := story[currentNode]
		orderedChoices := []string{"1", "2"}

		responseChan := make(chan struct {
			choice    string
			storyText string
			choices   map[string]string
			err       error
		}, len(orderedChoices)) // buffered channel to avoid blocking

		var wg sync.WaitGroup

		// Generate responses concurrently for the current story node
		for _, choice := range orderedChoices {
			if text, exists := node.Choices[choice]; exists {
				wg.Add(1)
				go func(choice string, text string) {
					defer wg.Done()
					storyText, choices, choice, err := generateDynamicResponse(memory, text)

					// Send the result to the response channel
					responseChan <- struct {
						choice    string
						storyText string
						choices   map[string]string
						err       error
					}{choice, storyText, choices, err}
				}(choice, text)
			}
		}

		// Close the response channel once all responses are processed
		go func() {
			wg.Wait()
			close(responseChan)
		}()

		// Prompt the user for input while responses are being generated
		fmt.Println("______________________")
		fmt.Println(node.Text)
		memory.AddToMemory("assistant", node.Text)
		fmt.Println("+--------------------+")
		fmt.Println("Choices:")

		for _, choice := range orderedChoices {
			if text, exists := node.Choices[choice]; exists {
				fmt.Printf("%s: %s\n", choice, text)
			}
		}
		fmt.Println("+---------------------+")
		fmt.Print("Enter your choice: ")

		reader := bufio.NewReader(os.Stdin)
		userChoice, _ := reader.ReadString('\n')
		userChoice = userChoice[:len(userChoice)-1]

		// Check if the user's choice exists
		if nextNode, exists := node.Choices[userChoice]; exists {
			currentNode = nextNode
		} else {
			fmt.Println("Invalid choice. Please try again.")
			continue
		}

		// Process responses from the goroutines
		for response := range responseChan {
			if response.err != nil {
				fmt.Println("Error generating story:", response.err)
				return
			}

			// add new story nodes to the story
			story[response.choice] = StoryNode{
				Text:    response.storyText,
				Choices: response.choices,
			}
		}
	}
}
