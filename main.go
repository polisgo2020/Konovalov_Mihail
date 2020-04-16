package main

import (
	"fmt"
	"github.com/polisgo2020/Konovalov_Mihail/invertedIndex"
	"github.com/polisgo2020/Konovalov_Mihail/web"
	"io/ioutil"
	"log"
	"os"
)

//Choose mode  create file with inverted index, search files with the best match. If you want create file
//enter "create", than dir with text and name file for inverted index, if you have file with inverted index
//and you want search, enter "search", than directory were is the file and string for search. If you dont
//have file with index and you want search enter "createAndSearch", than enter directory with files, than
//name for file with index and string what you want. If you want create web server and searching enter web
//than address for server, if don't enter address for server program chose :8080.
//examples: 1)create: create filesDir InvertedIndex
//2)search: search /InvertedIndex.txt I love go
//3)searchAndBuild: searchAndBuild filesDir InvertedIndex I love go
//4)web: web :8081
func main() {
	if len(os.Args) == 1 {
		log.Fatal("There are no arguments!!!")
	}
	if os.Args[1] == "create" {
		if len(os.Args[2]) == 0 || len(os.Args[3]) == 0 {
			log.Fatal("Empty path or file name!!!")
		}
		createProgram(os.Args[2], os.Args[3])
		return
	}
	if os.Args[1] == "search" {
		if len(os.Args[3:]) == 0 {
			return
		}
		searchProgram(os.Args[2], os.Args[3:])
		return
	}

	if os.Args[1] == "searchAndBuild" {
		if len(os.Args[2]) == 0 || len(os.Args[3]) == 0 || len(os.Args[4:]) == 0 {
			log.Fatal("Empty path, file name or searching string!!!")
		}
		createProgram(os.Args[2], os.Args[3])
		searchProgram(os.Args[3]+".txt", os.Args[4:])
		return
	}

	if os.Args[1] == "web" {
		if len(os.Args[2:]) == 0 {
			fmt.Println("Choosing standard address :8080 ")
			web.ServerSearch(":8080")
		}
		web.ServerSearch(os.Args[2])
	}
}

func createProgram(pathDirectory string, indexFileName string) {
	files, err := ioutil.ReadDir(pathDirectory)
	if err != nil {
		log.Fatal(err)
	}

	indexFile, err := os.Create(indexFileName + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer indexFile.Close()

	tokens := invertedIndex.GetInvertedIndex(pathDirectory, files)

	_, err = indexFile.WriteString(invertedIndex.FormOutputString(tokens))
	if err != nil {
		log.Fatal(err)
	}
}

func searchProgram(fileWithIndexPath string, userStrings []string) {
	invertedIndex.SearchBestStringMatch(fileWithIndexPath, userStrings)
}
