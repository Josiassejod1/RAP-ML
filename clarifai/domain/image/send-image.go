package Image

type PostImage struct {
Inputs []Input `json:"inputs"`
}

type Input struct {
  Data Data `json:"data"`
}

type Data struct {
  Image Image `json:"image"`
  MetaData MetaData `json:"metadata"`
}

type Image struct {
  Url string `json:"url"`
}

type MetaData struct {
    List []string  `json:"list,omitempty"`
}
