package main

import (
	"awesomeProject/invertedIndex"
	"io/ioutil"
	"log"
	"os"
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
	tokens := map[string]invertedIndex.Index{}

	tokens, err = invertedIndex.GetInvertedIndex(pathDir, files)
	if err != nil {
		log.Fatal(err)
	}

	invertedIndex.SearchBestStringMatch(tokens)

	_, err = indexFile.WriteString(invertedIndex.FormOutputString(tokens))
	if err != nil {
		log.Fatal(err)
	}
}
