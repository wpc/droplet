package main

import (
	"bufio"
	"net/http"
	"os"
)

func dropletHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	filename := r.URL.Path[len("/"):]
	file, error := os.Create(filename)
	if error != nil {
		http.Error(w, "Fail to new file\n", http.StatusInternalServerError)
		return
	}

	// all, _ := ioutil.ReadAll(r.Body)
	// println(string(all))

	writer := bufio.NewWriter(file)
	writer.ReadFrom(r.Body)
	writer.Flush()

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("File created\n"))
}

func main() {
	http.HandleFunc("/", dropletHandler)
	http.ListenAndServe(":3721", nil)
}
