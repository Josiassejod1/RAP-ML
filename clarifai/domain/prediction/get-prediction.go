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
		ReqID       string `json:"req_id"`
	} `json:"status"`
	Outputs []struct {
		ID     string `json:"id"`
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		CreatedAt time.Time `json:"created_at"`
		Model     struct {
			ID         string    `json:"id"`
			Name       string    `json:"name"`
			CreatedAt  time.Time `json:"created_at"`
			AppID      string    `json:"app_id"`
			OutputInfo struct {
				OutputConfig struct {
					ConceptsMutuallyExclusive   bool   `json:"concepts_mutually_exclusive"`
					ClosedEnvironment           bool   `json:"closed_environment"`
					MaxConcepts                 int    `json:"max_concepts"`
					MinValue                    int    `json:"min_value"`
					TestSplitPercent            int    `json:"test_split_percent"`
					EmbedModelVersionID         string `json:"embed_model_version_id"`
					InvalidDataTolerancePercent int    `json:"invalid_data_tolerance_percent"`
				} `json:"output_config"`
				Message string `json:"message"`
				Type    string `json:"type"`
				TypeExt string `json:"type_ext"`
			} `json:"output_info"`
			ModelVersion struct {
				ID        string    `json:"id"`
				CreatedAt time.Time `json:"created_at"`
				Status    struct {
					Code        int    `json:"code"`
					Description string `json:"description"`
				} `json:"status"`
				TotalInputCount int       `json:"total_input_count"`
				CompletedAt     time.Time `json:"completed_at"`
			} `json:"model_version"`
		} `json:"model"`
		Input struct {
			ID   string `json:"id"`
			Data struct {
				Image struct {
					URL    string `json:"url"`
					Base64 string `json:"base64"`
				} `json:"image"`
			} `json:"data"`
		} `json:"input"`
		Data struct {
			Regions []struct {
				ID         string `json:"id"`
				RegionInfo struct {
					BoundingBox struct {
						TopRow    float64 `json:"top_row"`
						LeftCol   float64 `json:"left_col"`
						BottomRow float64 `json:"bottom_row"`
						RightCol  float64 `json:"right_col"`
					} `json:"bounding_box"`
				} `json:"region_info"`
				Data struct {
					Concepts []struct {
						ID    string  `json:"id"`
						Name  string  `json:"name"`
						Value float64 `json:"value"`
						AppID string  `json:"app_id"`
					} `json:"concepts"`
				} `json:"data"`
			} `json:"regions"`
		} `json:"data"`
	} `json:"outputs"`
}