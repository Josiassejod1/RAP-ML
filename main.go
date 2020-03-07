package main

import (
  api "github.com/josiassejod1/train-ml/clarifai"
  "github.com/joho/godotenv"
  "log"
)

`
Sample Spotify out put 
curl -X "GET" "https://api.spotify.com/v1/search?q=Muse&type=track%2Cartist&market=US&limit=10&offset=5" -H 
"Accept: application/json" -H 
"Content-Type: application/json" -H 
"Authorization: Bearer YOUR TOKEN"
`

func main() {
 // api.SendImage("http://png.com", "Pop Smoke", true)
 metadata := []string{"Roddy Rich", "Roddy Ricch"}
 api.SendImage("https://i1.sndcdn.com/avatars-000735067801-x1vpo6-t500x500.jpg", metadata)
  //api.CreateModel("test", "pika chu", "string", true, false, false)
}

func init(){
  if err := godotenv.Load();
  err != nil {
    log.Print("No .env file found")
  }
}
