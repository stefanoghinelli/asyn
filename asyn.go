package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

const datamuseAPI = "https://api.datamuse.com/words"

type Synonym struct {
	Word string "json:'word'"
}

func main() {
	var rootCmd = &cobra.Command{Use: "asyn"}

	var word string
	var results int

	var listCmd = &cobra.Command{
		Use:   "list",
		Run: func(cmd *cobra.Command, args []string) {
			synonyms, err := getSynonyms(word, results)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Synonyms for %s:\n", word)
			for _, synonym := range synonyms {
				fmt.Println(" -", synonym.Word)
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

func getSynonyms(word string, results int) ([]Synonym, error) {
	client := resty.New()

	resp, err := client.R().
		SetQueryParam("rel_syn", word).
		SetQueryParam("max", strconv.Itoa(results)).
		Get(datamuseAPI)

	if err != nil {
		return nil, err
	}

	var synonyms []Synonym
	err = json.Unmarshal(resp.Body(), &synonyms)
	if err != nil {
		return nil, err
	}

	if len(synonyms) == 0 {
		return nil, fmt.Errorf("no synonyms found for the word %s", word)
	}

	return synonyms, nil
}
