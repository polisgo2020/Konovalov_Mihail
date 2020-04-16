package invertedIndex

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"sync"
)

//Func search best string match looking for the file with the most tokens and print them in decreasing order
func SearchBestStringMatch(fileWithIndexPath string, userStrings []string) []string {
	tokens := readFileWithInvertedIndex(fileWithIndexPath)
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
	arrayWithFilesPriority := []string{}
	for _, filePriority := range fileWithPriority {
		arrayWithFilesPriority = append(arrayWithFilesPriority, filePriority.fileName+"{"+
			strconv.Itoa(filePriority.numberWords)+"}")
		fmt.Println(filePriority.fileName + "{" + strconv.Itoa(filePriority.numberWords) + "}")
	}
	return arrayWithFilesPriority
}

func readFileWithInvertedIndex(filePath string) map[string]Index {

	text, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(text), "\n")
	m := make(chan map[string]Index)
	wg := &sync.WaitGroup{}
	for _, line := range lines {
		wg.Add(1)
		go writeString(line, m, wg)
	}
	go func(wg *sync.WaitGroup, readChan chan map[string]Index) {
		wg.Wait()
		close(readChan)
	}(wg, m)
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
	return tokensFilePosition
}

func writeString(line string, m chan<- map[string]Index, wg *sync.WaitGroup) {
	defer wg.Done()
	token := parseString(line)
	m <- token
	return
}

func parseString(line string) map[string]Index {
	endToken := map[string]Index{}
	var fileName string
	var position int
	var positions []int
	stringsInLine := strings.Fields(line)
	i := 0
	if len(stringsInLine) == 0 {
		return nil
	}
	token := trueStringForInvertedIndex(stringsInLine[0])
	stringsInLine = stringsInLine[1:]
	for _, files := range stringsInLine {
		if files == "|" {
			fileName = ""
			continue
		}
		i = 0
		for ; string(files[i]) != "{"; i++ {
			fileName += string(files[i])
		}

		endIndex := len(files) - 2
		j := 1
		for ; endIndex > i; endIndex-- {
			if string(files[endIndex]) == "," {
				positions = append(positions, position)
				if endToken[token] == nil {
					endToken[token] = Index{}
				}
				endToken[token][fileName] = positions
				j = 1
				position = 0
				continue
			}
			position += int(rune(files[endIndex])-48) * j
			j *= 10
		}
		positions = append(positions, position)
		if endToken[token] == nil {
			endToken[token] = Index{}
		}
		endToken[token][fileName] = positions
		positions = []int{}
		position = 0
	}

	return endToken
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
