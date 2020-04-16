package web

import (
	"fmt"
	"github.com/polisgo2020/Konovalov_Mihail/invertedIndex"
	"log"
	"net/http"
	"strings"
	"time"
)

var html = []byte(`
<html>
	<body>
		<form action="/" method="post">
			Enter the path to file with inverted index: <input type="file" name="InvertedIndexList">
			Enter the string for searching: <input type="text" name="userString">
			<value type="submit" name="invertedIndex">
		</form>
	</body>
</html>
`)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write(html)
		return
	}
	invertedIndexList := r.FormValue("InvertedIndexList")
	userString := strings.Fields(r.FormValue("userString"))
	if len(invertedIndexList) != 0 && len(userString) != 0 {
		files := invertedIndex.SearchBestStringMatch(invertedIndexList, userString)
		fmt.Fprintln(w, files)
	} else {
		fmt.Fprintln(w, "Empty file path or user string!!!")
	}
}

func ServerSearch(address string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	server := http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Starting server at:", address)
	log.Fatal(server.ListenAndServe())
}
