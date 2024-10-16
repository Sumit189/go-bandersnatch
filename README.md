# Go Bandersnatch

## Overview
Go Bandersnatch is an interactive storytelling application built in Go, inspired by the "choose your own adventure" format. Users navigate through a story by making choices that lead to different outcomes, with the help of a language model to generate dynamic responses.

## Features
- Interactive storytelling with multiple choices at each node.
- Dynamic response generation using OpenAI's GPT-4o model.
- Memory management to keep track of the conversation history.

## Getting Started

### Prerequisites
- Go 1.23.2 or later
- An OpenAI API key

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/go-bandersnatch.git
   cd go-bandersnatch
   ```

2. Set your OpenAI API key in the `storyBuilder.go` file:
   ```go
   const apiKey = "YOUR_OPENAI_API_KEY"
   ```

3. Run the application:
   ```bash
   go run go run main.go memory.go story.go storyBuilder.go
   ```

### Usage
- Upon running the application, you will be presented with a story and a set of choices.
- Enter the number corresponding to your choice to navigate through the story.
- The application will generate new story nodes based on your choices.

## Code Structure
- `main.go`: Contains the main logic for the application, including the story flow and user interaction.
- `storyBuilder.go`: Handles communication with the OpenAI API to generate dynamic story responses.
- `memory.go`: Manages the conversation history.
- `story.go`: Defines the structure of story nodes and initializes the starting story.

## Contributing
Contributions are welcome! Please feel free to submit a pull request or open an issue for any suggestions or improvements.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

