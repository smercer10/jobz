package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	if err := LoadConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Println("Search term:", Cfg.SearchTerm)
	fmt.Println("Job location:", Cfg.JobLocation)

	dummyDescription := "The candidate will need to be experienced with Docker and Kubernetes."

	words := strings.Fields(dummyDescription)
	var reduced []string

	for _, word := range words {
		cleanWord := strings.ToLower(strings.Trim(word, ".,!?"))
		if _, isFiltered := Cfg.DescriptionBlacklistSet[cleanWord]; !isFiltered {
			reduced = append(reduced, word)
		}
	}

	fmt.Println("Reduced description:", strings.Join(reduced, " "))
}
