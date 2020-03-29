package invertedIndex

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Index map[string][]int

func GetInvertedIndex(pathDir string, files []os.FileInfo) (map[string]Index, error) {
	tokens := map[string]Index{}
	for _, file := range files {
		fileWithStrings, err := ioutil.ReadFile(pathDir + "/" + file.Name())
		if err != nil {
			return nil, err
		}
		stringsInFile := strings.Fields(string(fileWithStrings))
		for j, stringInFile := range stringsInFile {
			stringInFile = trueStringForInvertedIndex(stringInFile)
			if stringInFile == "" {
				continue
			}
			tokensInFiles, ok := tokens[stringInFile]
			if ok {
				positionInFile := tokensInFiles[file.Name()]
				positionInFile = append(positionInFile, j+1)
				tokensInFiles[file.Name()] = positionInFile
			} else if !ok {
				tokensInFiles = Index{}
				tokensInFiles[file.Name()] = []int{j + 1}
				tokens[stringInFile] = tokensInFiles
			}
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
