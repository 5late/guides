package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rs/cors"
)

type videoData struct {
	Title string `json:"title"`
	Date  string `json:"date"`
}

//var templates = template.Must(template.ParseFiles("public/upload.html"))

func displayHTML(w http.ResponseWriter, page string, r *http.Request) {
	http.ServeFile(w, r, "./public/"+page+".html")
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 1 GB files
	r.ParseMultipartForm(1024 << 20)

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	//displayHTML(w, "uploadfinish", r)

	f, err := os.Create("./articles/" + handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	filename := handler.Filename
	currentTime := time.Now()

	CreateJSON(`./guides/`+filename, currentTime.String())

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(f, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func CreateJSON(vTitle string, vDate string) {
	filename := "guides.json"
	file, err := ioutil.ReadFile("guides.json")
	if err != nil {
		log.Println(err)
	}

	datas := []videoData{}

	json.Unmarshal(file, &datas)

	newStruct := &videoData{
		Title: vTitle,
		Date:  vDate,
	}

	datas = append(datas, *newStruct)

	dataBytes, err := json.MarshalIndent(datas, "", "   ")
	if err != nil {
		log.Println(err)
	}

	err = ioutil.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		log.Println(err)
	}
}

func setupRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		displayHTML(w, "upload", r)
	case "POST":
		uploadFile(w, r)
	}
}

func ServeJSON(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "guides.json")
}

func main() {
	mux := http.NewServeMux()

	fmt.Println("Hello World!")

	mux.HandleFunc("/json", ServeJSON)
	mux.HandleFunc("/upload", setupRoutes)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":1337", handler)
}
