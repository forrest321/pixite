package gopixite

import (
	"github.com/jpillora/velox"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/forrest321/pixite/constants"
	"github.com/forrest321/pixite/models"
	"github.com/gorilla/mux"
	"github.com/forrest321/pixite/handlers"
	"github.com/gorilla/context"
)

func main() {
	serv1 := &models.Serveable{Id: 1, Name: "Hello World!", Description: "First test"}
	go func() {
		for { //ever and ever and ever and ....
			serv1.MainImage, _ = Watch(constants.FileToWatch) //SHOULD watch file system for yadda.jpg to change
			serv1.Push()
			time.Sleep(250 * time.Millisecond)
		}
	}()

	r := mux.NewRouter()
	context.Set(r, constants.ServeableOneKey, serv1)

	r.Handle("/velox.js", velox.JS)
	r.Handle("/sync", velox.SyncHandler(serv1))
	r.HandleFunc("/", handlers.Index)
	http.Handle("/", r)

	//listen!
	port := os.Getenv("PORT")
	if port == "" {
		port = constants.Port
	}
	log.Printf("Listening on :%s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

var indexHtml = []byte(``)


