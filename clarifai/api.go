package api

import (
  "fmt"
  "github.com/josiassejod1/train-ml/clarifai/domain/image"
  "github.com/josiassejod1/train-ml/clarifai/domain/model"
)

func SendImage(url string, id string, value bool) {
  body :=  Image.PostImage {
    Inputs: []Image.Input{
      {
        Data: Image.Data{
            Image: Image.Image { Url: url},
            Concepts: []Image.Concept{
              {
                Id: id,
                Value: value,
              },
            },
        },
      },
    },
  }

  fmt.Println(body)
}

func CreateModel(
  url string,
  model_id string,
  id string,
  value bool,
  concept bool,
  env bool) {
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

  fmt.Println(body)
}
