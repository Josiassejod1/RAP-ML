package Concept

type SendConcept struct {
	Inputs []Input `json:"inputs"`
	}
	
	type Input struct {
	  Data Data `json:"data"`
	}
	
	type Data struct {
	  Image Image `json:"image"`
	  Concepts Conceptss `json:"concepts"`
	}
	
	type Image struct {
	  Url string `json:"url"`
	}

	type Conceptss struct {
		ID    string `json:"id"`
		Value bool   `json:"value"`
	  }

	