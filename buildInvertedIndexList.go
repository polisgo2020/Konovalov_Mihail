package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("There are no arguments!!!")
	}
	pathDir := os.Args[1]
	files, err := ioutil.ReadDir(pathDir)
	if err != nil {
		log.Fatal(err)
	}

	indexFile, err := os.Create("InvertedIndexList.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer indexFile.Close()

	tokens, err := getInvertedIndex(pathDir, files)
	if err != nil {
		log.Fatal(err)
	}
	for key, value := range tokens {
		_, err = indexFile.WriteString(key + ": {" + value + "}\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getInvertedIndex(pathDir string, files []os.FileInfo) (map[string]string, error) {
	tokens := map[string]string{}
	for i, file := range files {
		fileWithStrings, err := ioutil.ReadFile(pathDir + "/" + file.Name())
		if err != nil {
			return nil, err
		}
		stringsInFile := strings.Fields(string(fileWithStrings))
		for j := range stringsInFile {
			if stringsInFile[j] == "" {
				continue
			}
			value, ok := tokens[stringsInFile[j]]
			if ok && strconv.Itoa(int(value[len(value)-1])-48) != strconv.Itoa(i+1) {
				tokens[stringsInFile[j]] = value + "," + strconv.Itoa(i+1)
			} else if !ok {
				tokens[stringsInFile[j]] = strconv.Itoa(i + 1)
			}
		}
	}
	return tokens, nil
}
