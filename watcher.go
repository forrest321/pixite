package gopixite

import (
	"image"
	"log"
	"github.com/fsnotify/fsnotify"
)

func Watch(location string) (image.Image, error) {
	var imageToReturn image.Image
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					imageToReturn = image.Decode(event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(location)
	if err != nil {
		log.Fatal(err)
	}
	<-done



	return imageToReturn, err
}
