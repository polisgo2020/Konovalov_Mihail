package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func searchBestStringMatch(tokens map[string]Index) {
	userStrings := os.Args[2:]
	if len(userStrings) == 0 {
		return
	}
	userStringsForInvertedIndex := makeStringsForInvertedIndex(userStrings)

	filesWithTokens := map[string]Index{}
	for _, invertedString := range userStringsForInvertedIndex {
		files, ok := tokens[invertedString]
		if ok {
			for fileName, position := range files {
				tokensInFile := filesWithTokens[fileName]
				if tokensInFile == nil {
					tokensInFile = Index{}
				}
				tokensInFile[invertedString] = position
				filesWithTokens[fileName] = tokensInFile
			}
		}
	}

	fileWithPriority := []FilePriority{}
	i := 0
	for file, tokensWithPosition := range filesWithTokens {
		for _, token := range userStringsForInvertedIndex {
			position, ok := tokensWithPosition[token]
			if ok && len(position) != 0 {
				if len(fileWithPriority) == i {
					fileWithPriority = append(fileWithPriority, *newFilePriority(file, 1))
					continue
				}
				fileWithPriority[i].numberWords += 1
			}
		}
		if len(fileWithPriority) != 0 {
			i++
		}
	}
	sort.Slice(fileWithPriority, func(i, j int) bool {
		return fileWithPriority[i].numberWords > fileWithPriority[j].numberWords
	})

	for _, filePriority := range fileWithPriority {
		fmt.Println(filePriority.fileName + "{" + strconv.Itoa(filePriority.numberWords) + "}")
	}
}

func makeStringsForInvertedIndex(strings []string) []string {
	for i := range strings {
		strings[i] = trueStringForInvertedIndex(strings[i])
	}
	return strings
}

type FilePriority struct {
	fileName    string
	numberWords int
}

func newFilePriority(fileName string, numberWords int) *FilePriority {
	return &FilePriority{fileName: fileName, numberWords: numberWords}
}
