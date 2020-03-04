package main

import (
  api "github.com/josiassejod1/train-ml/clarifai"
)

func main() {
  api.SendImage("http://png.com", "Pop Smoke", true)
  api.CreateModel("test", "pika chu", "string", true, false, false)
}
