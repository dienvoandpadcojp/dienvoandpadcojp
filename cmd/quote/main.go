package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
)

type Quote struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	// Step 1: Fetch a random quote from the API
	baseURL := os.Getenv("API_BASE_URL")
	if baseURL == "" {
		log.Fatalln("API_BASE_URL is not set")
	}

	resp, err := http.Get(baseURL)
	if err != nil {
		log.Fatalln("Error fetching quote:", err)
	}
	defer resp.Body.Close()

	// Step 2: Parse the JSON response
	var quotes []Quote
	if err := json.NewDecoder(resp.Body).Decode(&quotes); err != nil {
		log.Fatalln("Error decoding JSON:", err)
	}
	if len(quotes) == 0 {
		log.Fatalln("No quotes found")
	}
	quote := quotes[0]

	// Step 3: Read the contents of README.md
	readmePath := "README.md"
	readmeContent, err := os.ReadFile(readmePath)
	if err != nil {
		log.Fatalln("Error reading README.md:", err)
	}

	// Step 4: Append the fetched quote to the end of README.md
	quotePattern := regexp.MustCompile(`(?s)> .*\n> - .*\n?$`)
	newQuote := fmt.Sprintf("> %s\n> - %s\n", quote.Content, quote.Author)
	newContent := quotePattern.ReplaceAllString(string(readmeContent), newQuote)

	// If no existing quote was found, append the new quote
	if newContent == string(readmeContent) {
		newContent = fmt.Sprintf("%s\n\n%s", string(readmeContent), newQuote)
	}
	// Step 5: Save the updated contents back to README.md
	if err := os.WriteFile(readmePath, []byte(newContent), 0644); err != nil {
		log.Fatalln("Error writing to README.md: ", err)
		return
	}

	log.Println("Quote appended to README.md successfully.")
}
