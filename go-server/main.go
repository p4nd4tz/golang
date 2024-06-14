package main
import(
	"log"
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return;
	}

	fmt.Fprintf(w, "hello\n");
}


func main() {
	fileServer := http.FileServer(http.Dir("./static"));

	http.Handle("/", fileServer);
	http.HandleFunc("/hello", hello);

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
