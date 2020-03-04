package Image

type PostImage struct {
Inputs []Input `json:"inputs"`
}

type Input struct {
  Data Data `json:"data"`
}

type Concept struct {
  Id string `json:"id"`
  Value bool `json:"value"`
}

type Data struct {
  Image Image `json:"image"`
  Concepts []Concept `json:"concepts"`
}

type Image struct {
  Url string `json:"url"`
}
