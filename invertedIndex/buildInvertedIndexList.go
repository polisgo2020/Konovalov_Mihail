package invertedIndex

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

type Index map[string][]int

func GetInvertedIndex(pathDir string, files []os.FileInfo) (map[string]Index, error) {
	m := make(chan map[string]Index)
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go readFile(m, &wg, pathDir, file.Name())
	}
	go func(wg *sync.WaitGroup, readChan chan map[string]Index) {
		wg.Wait()
		close(readChan)
	}(&wg, m)
	tokensFilePosition := map[string]Index{}
	for tokensF := range m {
		for token, fileWithPosition := range tokensF {
			tokensInFile, ok := tokensFilePosition[token]
			if !ok {
				tokensInFile = Index{}
			}
			for fileName, positions := range fileWithPosition {
				tokensInFile[fileName] = positions
				tokensFilePosition[token] = tokensInFile
			}
		}

	}
	return tokensFilePosition, nil
}

func readFile(outputChan chan<- map[string]Index, wg *sync.WaitGroup,
	pathDir string, fileName string) {
	defer wg.Done()
	fileWithStrings, err := ioutil.ReadFile(pathDir + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	tokens, err := makeMapForIndex(strings.Fields(string(fileWithStrings)), fileName)

	if err != nil {
		log.Fatal(err)
	} else {
		outputChan <- tokens
	}

	return
}

func makeMapForIndex(fileWithString []string, fileName string) (map[string]Index, error) {
	tokens := map[string]Index{}
	for i, stringInFile := range fileWithString {
		stringInFile = trueStringForInvertedIndex(stringInFile)
		if stringInFile == "" {
			continue
		}
		tokensInFiles, ok := tokens[stringInFile]
		if ok {
			positionInFile := tokensInFiles[fileName]
			positionInFile = append(positionInFile, i+1)
			tokensInFiles[fileName] = positionInFile
		} else if !ok {
			tokensInFiles = Index{}
			tokensInFiles[fileName] = []int{i + 1}
			tokens[stringInFile] = tokensInFiles
		}
	}
	return tokens, nil
}

func trueStringForInvertedIndex(stringInFile string) string {
	stringInFile = strings.TrimFunc(stringInFile, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	return strings.ToLower(stringInFile)
}

func FormOutputString(tokens map[string]Index) string {
	var outputString string
	for token, value := range tokens {
		outputString += token + ": "
		var i = 0

		for fileName, positions := range value {
			outputString += fileName + "{"

			for i, position := range positions {
				if i == len(positions)-1 {
					outputString += strconv.Itoa(position) + "}"
				} else {
					outputString += strconv.Itoa(position) + ", "
				}
			}

			if i == len(value)-1 {
				outputString += "\n"
			} else {
				outputString += " | "
			}
			i++
		}
	}
	return outputString
}
