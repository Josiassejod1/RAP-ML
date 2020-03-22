package api

import (
  "encoding/json"
  "net/http"
  "net/url"
  "fmt"
  "bytes"
  "os"
  "log"
  "context"
	"golang.org/x/oauth2/clientcredentials"
	"github.com/zmb3/spotify"
  "github.com/josiassejod1/train-ml/clarifai/domain/image"
  "github.com/josiassejod1/train-ml/clarifai/domain/model"
  "github.com/josiassejod1/train-ml/clarifai/domain/spotify"
  "github.com/josiassejod1/train-ml/clarifai/domain/prediction"
  "github.com/patrickmn/go-cache"
  "time"
)


var c = cache.New(5*time.Minute, 10*time.Minute)


/*
  ToDO
  - Add API That Trains Model
  - Add API That Deletes Old Models (Refresh)
*/

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

func validateModelKey() (string, string) {
  key, err := os.LookupEnv("MODEL_VERSION_NAME")
  key2, err2 := os.LookupEnv("MODEL_ID")

  if !err2 {
    fmt.Println("Model ID Not Found")
    os.Exit(1)
  } else {
    fmt.Println("Model ID Found")
  }


  if !err {
    fmt.Println("Model Version Name Not Found")
    os.Exit(1)
  } else {
    fmt.Println("Model Version Found")
  }
  return key, key2
}


func GetPrediction(search string) string {
  key := validateKey()
  _, model_id := validateModelKey()
  var response = ""
  
  body := PredictionInput.Format{
     InputType: []PredictionInput.InputType{
        {
          Data: PredictionInput.Data{
            Image: PredictionInput.Image{
              Url: search,
            },
          },
        },
      }, 
      }

      byte, _  := json.Marshal(&body)
      urlStr := "https://api.clarifai.com/v2/models/" + model_id + "/outputs"
      client := &http.Client {}
      
      
      req, _ := http.NewRequest("POST", urlStr, bytes.NewBuffer(byte))
      req.Header.Add("Authorization", "Key " + key)
      req.Header.Add("Content-Type", "application/json")
      resp,
      _ := client.Do(req)
    
      if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
        var data PredictionInput.PredictionResponse
        decoder := json.NewDecoder(resp.Body)
        err := decoder.Decode(&data)
        if (err == nil) {
          fmt.Println("Image Succesfully Uploaded")
          json, err := json.Marshal(data)
          if err != nil {
              response = "JSON Not Formatted Correctly"
          } 
          response = string(json)
        }
      } else {
        fmt.Println(resp.Status);
        response = resp.Status 
      }

      return response
    }



// func GetWikiImage(search string) {
//   urlStr := url.Values{}
//   urlStr.Set("action", "query")
//   urlStr.Set("prop", "pageimages")
//   urlStr.Set("format", "json")
//   urlStr.Set("piprop", "original")
//   urlStr.Set("titles", search)

//   client := &http.Client {}
 

//   wikiUrl := "https://en.wikipedia.org/w/api.php?" + urlStr.Encode()
//   fmt.Println(wikiUrl)
//   req, _ := http.NewRequest("GET", wikiUrl, nil)

//   resp,
//   _ := client.Do(req)

//   if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
//     var data map[string] interface {}
//     decoder := json.NewDecoder(resp.Body)
//     err := decoder.Decode( & data)
//     if (err == nil) {
//       fmt.Println("WIKI API Accessed")
//       fmt.Println(data)
//     }
//   } else {
//     fmt.Println(resp.Status);
//   }
// }

func SendImage(urlstring string, metadata []string) string {
  var response = ""
  key := validateKey()
  

  //Adds on urls to prevent over use of the model being trained
  fmt.Println(urlstring)

    _ , found := c.Get(urlstring)
    fmt.Println(found)
    if found {
      response = "Artist Image Already Used, Upload One Using Url Option"
      return response
    } else {
      c.Set(urlstring, "", cache.NoExpiration)
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
          response = "Image Successfully Uploaded"
        }
      } else {
        fmt.Println(resp.Status);
        fmt.Println(resp);
        response = resp.Status 
      }
    }

 
  return response
}

func SearchArtistImage(search string) string {
	config := &clientcredentials.Config{
		ClientID: os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL: spotify.TokenURL,
  }
  
  urlStr := url.Values{}
  urlStr.Set("q", search)
  urlStr.Set("type", "artist")
  urlStr.Set("limit", "1")

  client := &http.Client {}
 
  log.Println(urlStr)

  token, err := config.Token(context.Background())
  if err != nil {
		log.Fatalf("could't get token: %v", err)
  }
  
  log.Println(token)

  spotifyUrl := "https://api.spotify.com/v1/search?" + urlStr.Encode()

  req, _ := http.NewRequest("GET", spotifyUrl, nil)
  req.Header.Add("Authorization", "Bearer " + token.AccessToken)

  resp,
  _ := client.Do(req)

  var url = "unknown.jpg"

  if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
    var data artist.GetArtist
    decoder := json.NewDecoder(resp.Body)
    err := decoder.Decode( & data)
    if (err == nil) {
      fmt.Println("Spotify API Accessed")
      fmt.Println(data)
      if (len(data.Artists.Items) > 0) {
        url = data.Artists.Items[0].Images[0].URL
      }
    }
  } else {
    fmt.Println(resp.Status);
  }
  return url
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
    }
  } else {
    fmt.Println(resp.Status);
  }
}
