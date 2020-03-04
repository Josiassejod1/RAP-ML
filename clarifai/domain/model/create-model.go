package Models

type CreateModel struct {
  Model Model `json: "model"`
}

type Model struct {
  Id string `json: "id"`
  OutputInfo OutputInfo `json: "output_info"`
}

type OutputInfo struct {
  Data []Concept `json: "data"`
  OutputConfig OutputConfig `json: "output_config"`
}

type Concept struct {
  Id string `json: "id"`
}

type OutputConfig struct {
  ConceptMutal bool `json: "concepts_mutually_exclusive"`
  Env bool `json: "closed_environment"`
}
