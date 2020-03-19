package main

import (
  api "github.com/josiassejod1/train-ml/clarifai"
  "github.com/joho/godotenv"
  "log"
  "net/http"
  "encoding/json"
  "fmt"
)

func main() {
 // api.SendImage("http://png.com", "Pop Smoke", true)
 //metadata := []string{"Roddy Rich", "Roddy Ricch"}
 //api.SendImage("https://i1.sndcdn.com/avatars-000735067801-x1vpo6-t500x500.jpg", metadata)
  //api.CreateModel("test", "pika chu", "string", true, false, false)
  //api.SearchArtistImage("Jay-Z")
  //api.GetWikiImage("Jay-Z")
  http.HandleFunc("/predict", predict)
  http.HandleFunc("/artist", search)
  http.HandleFunc("/upload", upload)
  fmt.Printf("Starting server for testing HTTP")
  if err := http.ListenAndServe(":3002", nil); err != nil {
      log.Fatal(err)
  }
}

type ArtistData struct {
  MetaData string
  Image string
}

type Prediction struct {
  Image string
}

func upload(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/upload" {
    http.Error(w, "crapped out", http.StatusNotFound)
    return
  } 

  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Method", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
  w.Header().Set("Access-Control-Allow-Header", "Origin, Content-Type, X-Auth-Token")

  log.Println(r.Body)
  decoder := json.NewDecoder(r.Body)
  var data ArtistData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
    //artist := r.FormValue("artist")
   // fmt.Println(r)
   log.Println(r.Form)
   image := data.Image
   metadata := data.MetaData
    description := []string{metadata}
    url := api.SendImage(image, description)
    fmt.Fprintf(w, url)
  }


func search(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/artist" {
    http.Error(w, "crap", http.StatusNotFound)
    return
  } 

  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Method", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
  w.Header().Set("Access-Control-Allow-Header", "Origin, Content-Type, X-Auth-Token")
  keys, ok := r.URL.Query()["artist"]
  log.Println(r.URL.Query())
    
  if !ok || len(keys[0]) < 1 {
      log.Println("Url Param 'key' is missing")
      return
  } else {
    //artist := r.FormValue("artist")
   // fmt.Println(r)

    key := keys[0]
    
    url := api.SearchArtistImage(key)
    fmt.Fprintf(w, url)
  }

  
}

func predict(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/predict" {
    http.Error(w, "prediction not found", http.StatusNotFound)
    return
  } 

  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Method", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
  w.Header().Set("Access-Control-Allow-Header", "Origin, Content-Type, X-Auth-Token")

  keys, ok := r.URL.Query()["image"]

    
  if !ok || len(keys[0]) < 1 {
      log.Println("Url Param 'key' is missing")
      return
  } else {

    key := keys[0]
    
    data := api.GetPrediction(key)
    fmt.Fprintf(w, data)
  }
  
}

func init(){
  if err := godotenv.Load();
  err != nil {
    log.Print("No .env file found")
  }
}
