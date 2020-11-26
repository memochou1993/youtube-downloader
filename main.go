package main

import (
	"github.com/memochou1993/youtube-downloader/app/controller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controller.Download)

	log.Fatal(http.ListenAndServe(":8083", nil))
}
