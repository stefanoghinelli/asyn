package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func fetchData(word string, results int) ([]string, error) {
	var thesaurusURL = "https://api.datamuse.com/words"
	client := resty.New()
	resp, err := client.R().SetQueryParam("rel_syn", word).SetQueryParam("max", strconv.Itoa(results)).Get(thesaurusURL)

	if err != nil {
		return nil, err
	}

	synonyms := getSynonyms(resp.Body())

	if len(synonyms) == 0 {
		return nil, fmt.Errorf("no synonyms found for the word %s", word)
	}

	return synonyms, nil
}

func getSynonyms(response []byte) []string {
	var synonyms []string
	var data []map[string]interface{}
	err := json.Unmarshal(response, &data)
	if err != nil {
		return nil
	}

	for _, item := range data {
		if synonym, ok := item["word"].(string); ok {
			synonyms = append(synonyms, synonym)
		}
	}

	return synonyms
}

func main() {
	var rootCmd = &cobra.Command{Use: "asyn"}
	var word string
	var results int
	var listCmd = &cobra.Command{
		Use:   "list",
		Run: func(cmd *cobra.Command, args []string) {
			synonyms, err := fetchData(word, results)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Synonyms for %s:\n", word)
			for _, synonym := range synonyms {
				fmt.Println(" -", synonym)
			}
		},
	}

	listCmd.Flags().StringVarP(&word, "word", "w", "", "The word to find synonyms for")
	listCmd.MarkFlagRequired("word")
	listCmd.Flags().IntVarP(&results, "results", "r", 10, "Number of results in output")
	rootCmd.AddCommand(listCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

