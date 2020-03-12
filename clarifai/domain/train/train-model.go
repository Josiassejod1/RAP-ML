package train 

import("time")

type TrainModel struct {
	Status struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	} `json:"status"`
	Model struct {
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
}