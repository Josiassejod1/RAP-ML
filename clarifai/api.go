package api

import (
  "encoding/json"
  "net/http"
  "fmt"
  "bytes"
  "os"
  "github.com/josiassejod1/train-ml/clarifai/domain/image"
  "github.com/josiassejod1/train-ml/clarifai/domain/model"
)

func validateKey() (string) {
  key, err := os.LookupEnv("KEY")
  if !err {
    fmt.Println("API Key Not Found")
    os.Exit(1)
  } else {
    fmt.Println("API KEY SET")
  }
  return key
}

func SendImage(urlstring string, metadata []string) {
  key := validateKey()
  body :=  Image.PostImage {
    Inputs: []Image.Input{
      {
        Data: Image.Data{
            Image: Image.Image { Url: urlstring},
            MetaData: Image.MetaData{
              List: metadata,
            },
        },
      },
    },
  }

  byte, _  := json.Marshal(&body)
  urlStr := "https://api.clarifai.com/v2/inputs"
  client := &http.Client {}
  
  req, _ := http.NewRequest("POST", urlStr, bytes.NewBuffer(byte))
  req.Header.Add("Authorization", "Key " + key)
  req.Header.Add("Content-Type", "application/json")
  fmt.Println(req)
  resp,
  _ := client.Do(req)

  if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
    var data map[string] interface {}
    decoder := json.NewDecoder(resp.Body)
    err := decoder.Decode( & data)
    if (err == nil) {
      fmt.Println("Image Succesfully Uploaded")
      fmt.Println(data["sid"])
    }
  } else {
    fmt.Println(resp.Status);
  }
}

func CreateModel(
  url string,
  model_id string,
  id string,
  value bool,
  concept bool,
  env bool) {
  
  key := validateKey()
  body := Models.CreateModel{
    Model: Models.Model{
        Id: model_id,
        OutputInfo: Models.OutputInfo{
          Data: []Models.Concept{
            {
              Id: id,
            },
          },
          OutputConfig: Models.OutputConfig {
              ConceptMutal: concept,
              Env: env,
          },
        },
    },
  }

  byte, _  := json.Marshal(&body)
  urlStr := "https://api.clarifai.com/v2/models"
  client := &http.Client {}
  
  req, _ := http.NewRequest("POST", urlStr, bytes.NewBuffer(byte))
  req.Header.Add("Authorization", "Key " + key)
  req.Header.Add("Content-Type", "application/json")
  fmt.Println(req)
  resp,
  _ := client.Do(req)

  if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
    var data map[string] interface {}
    decoder := json.NewDecoder(resp.Body)
    err := decoder.Decode( & data)
    if (err == nil) {
      fmt.Println("Model Successfully Created")
      fmt.Println(data["sid"])
    }
  } else {
    fmt.Println(resp.Status);
  }
}
