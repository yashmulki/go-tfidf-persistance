package TFIDF_Persistant

import (
	"errors"
	"log"
	"os"
	"regexp"
	"strings"
)

func NewInstance(file os.File) error {
	return nil
}

func LoadInstance(file os.File, name string) error {
	return nil
}

func CreateNewFile() (os.File, error) {
	return os.File{}, nil
}

type TFIDFInstance struct {
	Name string
	Terms map[string]int
	TotalDocuments int
}

func (i TFIDFInstance) TFIDFScores(doc string, addToCorpus bool) (map[string]float64, error) {
	// Strip out punctuation
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := strings.ToLower(doc)
	processedString = reg.ReplaceAllString(processedString, "")
	words := strings.Fields(processedString)

	docLen := len(words)
	if docLen == 0 {
		return nil, errors.New("no words in document")
	}
	var countMap map[string]int

	for _, word := range words {
		if _, exists := countMap[word]; exists {
			countMap[word] += 1
		} else {
			countMap[word] = 1
		}
	}
	var tfidfResults map[string]float64
	for k, v := range countMap {
		tfScore := float64(v) / float64(docLen)
		var idfScore float64
		if val, exists := i.Terms[k]; exists {
			idfScore = ln(float64(i.TotalDocuments) / float64(val))
		} else {
			idfScore = -1
		}
		 // TODO: Calculate IDF score from existing corpus
		tfidf := tfScore * idfScore
		tfidfResults[k] = tfidf
	}
	return tfidfResults, nil
}

func (i TFIDFInstance) AddToCorpus(doc string) {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := strings.ToLower(doc)
	processedString = reg.ReplaceAllString(processedString, "")
	words := strings.Fields(processedString)

	docLen := len(words)
	if docLen == 0 {
		return
	}

	var countMap map[string]int
	for _, word := range words {
		countMap[word] = 1
	}

	for k := range countMap {
		if _, exists := i.Terms[k]; exists {
			i.Terms[k] += 1
		} else {
			i.Terms[k] = 1
		}
	}
	i.TotalDocuments += 1
	// Make sure to save this file
}
