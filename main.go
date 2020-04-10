package main

import (
	"github.com/polisgo2020/Konovalov_Mihail/invertedIndex"
	"io/ioutil"
	"log"
	"os"
)

//Choose mode or create file with inverted index or search files with the best match. If you want create file
//enter "create", if you want search enter "search". Than you mast enter the directory path with files and if
//you want search enter the string statement witch you want.
func main() {
	if len(os.Args) == 1 {
		log.Fatal("There are no arguments!!!")
	}

	pathDir := os.Args[2]
	files, err := ioutil.ReadDir(pathDir)
	if err != nil {
		log.Fatal(err)
	}

	tokens := map[string]invertedIndex.Index{}

	tokens, err = invertedIndex.GetInvertedIndex(pathDir, files)
	if err != nil {
		log.Fatal(err)
	}

	if os.Args[1] == "search" {
		searchProgram(&tokens)
	} else if os.Args[1] == "create" {
		createProgram(&tokens)
	} else {
		log.Fatal("Not correct program mod!!!")
	}

}

func createProgram(tokens *map[string]invertedIndex.Index) {
	indexFile, err := os.Create("InvertedIndexList.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer indexFile.Close()

	_, err = indexFile.WriteString(invertedIndex.FormOutputString(*tokens))
	if err != nil {
		log.Fatal(err)
	}
}

func searchProgram(tokens *map[string]invertedIndex.Index) {
	invertedIndex.SearchBestStringMatch(*tokens)
}
