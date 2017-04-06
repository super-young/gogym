package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

var basedir string

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run server.go [basedir]")
		os.Exit(1)
	}

	basedir = os.Args[1]

	fmt.Printf("serving content: %s\n", basedir)

	// http.Handle("/", http.FileServer(http.Dir(basedir)))

	http.HandleFunc("/", serveContent)
	http.ListenAndServe(":7777", nil)
}

func serveContent(res http.ResponseWriter, req *http.Request) {
	reqpath := req.URL.Path

	if reqpath == "/" {
		reqpath = "index.html"
	}

	filepath := path.Join(basedir, reqpath)
	f, err := os.Open(filepath)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}

	_, err = io.Copy(res, f)
	if err != nil {
		log.Println("serve content:", err)
	}
}
