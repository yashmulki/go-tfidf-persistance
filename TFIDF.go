package TFIDF_Persistant

// NOTE: Closing is the responsibility of the caller. This does not implement any direct file interaction

import (
	"encoding/json"
	"errors"
	"github.com/bbalet/stopwords"
	"log"
	"os"
	"regexp"
	"strings"
)

func NewInstance(name string) TFIDFInstance {
	instance := TFIDFInstance{
		Name:           name,
		Terms: 			map[string]int{},
		TotalDocuments: 0,
	}
	return instance
}

func LoadInstance(file os.File) (*TFIDFInstance, error) {
	var data []byte
	_, err := file.Read(data)
	if err != nil {
		return nil, err
	}
	var instance TFIDFInstance
	err = json.Unmarshal(data, instance)
	if err != nil {
		return nil, err
	}
	return &instance, nil
}

func CreateNewFile(path string) (*os.File, error) {
	f, err := os.Create(path+".json")
	if err != nil {
		return nil, err
	}
	return f, nil
}


type TFIDFInstance struct {
	Name string `json:"Name"`
	Terms map[string]int `json:"Terms"`
	TotalDocuments int `json:"TotalDocuments"`
}

func (i TFIDFInstance) SaveToFile(file os.File) error {
	data, err := json.MarshalIndent(i, "", " ")
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

func (i TFIDFInstance) TFIDFScores(doc string, addToCorpus bool) (map[string]float64, error) {
	words, err := i.process(doc, true)
	if err != nil {
		return nil, err
	}

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
	if addToCorpus {
		err = i.AddToCorpus(doc, true)
		if err != nil {
			return tfidfResults, err
		}
	}
	return tfidfResults, nil
}

func (i TFIDFInstance) AddToCorpus(doc string, processed bool) error {

	var words []string
	var err error

	if !processed {
		words, err = i.process(doc, true)
		if err != nil {
			return err
		}
	}

	docLen := len(words)
	if docLen == 0 {
		return errors.New("empty document")
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
	// REMINDER: It is the responsibility of callers to save
	return nil
}

func (i TFIDFInstance) process(doc string, removeStopwords bool) ([]string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return nil, err
	}
	processedString := strings.ToLower(doc)
	processedString = reg.ReplaceAllString(processedString, "")

	if removeStopwords {
		processedString = stopwords.CleanString(processedString, "en", false)
	}

	words := strings.Fields(processedString)
	return words, nil
}
