package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	sync "github.com/Be3751/sync-HLS-multipage"
)

func main() {
	root, err := sync.GetRootPath()
	if err != nil {
		log.Fatal("failed to get the root path in this project")
		return
	}

	fileServer := http.FileServer(http.Dir(fmt.Sprintf("%s/web/public/static", root)))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
	http.HandleFunc("/F1", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile(fmt.Sprintf("%s/web/public/static/index.html", root))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})
	http.HandleFunc("/F2", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile(fmt.Sprintf("%s/web/public/static/index.html", root))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	http.ListenAndServe(":8080", nil)
}
