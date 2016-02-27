package gopixite

import (
	"github.com/jpillora/velox"
	"image"
	"log"
	"net/http"
	"os"
	"time"
)

type Serveable struct {
	velox.State
	Id          int64
	Name        string
	Description string
	MainImage   image.Image
}

func main() {
	serv1 := &Serveable{Id: 1, Name: "Hello World", Description: "First test"}
	go func() {
		for { //ever and ever and ever and ....
			serv1.MainImage, _ = Watch("/home/fo/yadda.jpg") //SHOULD watch file system for yadda.jpg to change
			serv1.Push()
			time.Sleep(250 * time.Millisecond)
		}
	}()

	http.Handle("/velox.js", velox.JS)
	http.Handle("/sync", velox.SyncHandler(serv1))
	//index handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write(indexhtml)
	})
	//listen!
	port := os.Getenv("PORT")
	if port == "" {
		port = "7070"
	}
	log.Printf("Listening on :%s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

var indexhtml = []byte(`
<!-- documentation -->
<pre id="code">
&lt;pre id="example">&lt;/pre>
&lt;script src="/velox.js">&lt;/script>
&lt;script>
var foo = {};
var v = velox.sse("/sync", foo);
v.onupdate = function() {
	example.innerHTML = JSON.stringify(foo, null, 2);
};
&lt;/script>
</pre>
<a href="https://github.com/jpillora/velox"><img style="position: absolute; z-index: 2; top: 0; right: 0; border: 0;" src="https://s3.amazonaws.com/github/ribbons/forkme_right_darkblue_121621.png" alt="Fork me on GitHub"></a>
<hr>
<!-- example -->
<pre id="example"></pre>
<script src="/velox.js"></script>
<script>
var serv1 = {};
var v = velox.sse("/sync", serv1);
v.onupdate = function() {
	example.innerHTML = JSON.stringify(foo, null, 2);
};
</script>
`)
