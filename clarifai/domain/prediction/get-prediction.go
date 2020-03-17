package PredictionInput

import ("time")
type Format struct {
	InputType []InputType `json:"inputs"`
}

type InputType struct {
	Data Data `json:"data"`
}

type Data struct {
	Image Image `json:"image"`
}

type Image struct {
	Url string `json:"base64"`
}


type PredictionResponse struct {
	Status struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	} `json:"status"`
	Outputs[] struct {
		ID     string `json:"id"`
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		CreatedAt time.Time `json:"created_at"`
		Model     struct {
			Name       string    `json:"name"`
			ID         string    `json:"id"`
			CreatedAt  time.Time `json:"created_at"`
			AppID      string    `json:"app_id"`
			OutputInfo struct {
				Message      string `json:"message"`
				Type         string `json:"type"`
				OutputConfig struct {
					ConceptsMutuallyExclusive bool `json:"concepts_mutually_exclusive"`
					ClosedEnvironment         bool `json:"closed_environment"`
				} `json:"output_config"`
			} `json:"output_info"`
			ModelVersion struct {
				ID        string    `json:"id"`
				CreatedAt time.Time `json:"created_at"`
				Status    struct {
					Code        int    `json:"code"`
					Description string `json:"description"`
				} `json:"status"`
			} `json:"model_version"`
		} `json:"model"`
		Input struct {
			ID   string `json:"id"`
			Data struct {
				Image struct {
					URL string `json:"url"`
				} `json:"image"`
			} `json:"data"`
		} `json:"input"`
		Data struct {
			Concepts []struct {
				ID    string  `json:"id"`
				Name  string  `json:"name"`
				AppID string  `json:"app_id"`
				Value float64 `json:"value"`
			} `json:"concepts"`
		} `json:"data"`
	} `json:"outputs"`
}