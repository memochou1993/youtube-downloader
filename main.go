package main

import (
	"context"
	"fmt"
	"github.com/memochou1993/youtube-downloader/app"
)

func main() {
	client := &app.Client{}

	id := ""

	video := client.GetVideo(context.Background(), id)

	// TODO
	fmt.Println(video)
}
