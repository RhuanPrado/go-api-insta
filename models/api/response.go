package api

type Response struct {
	Error        bool   `json:"error"`
	ErrorMessage string `json:"errorMessage"`
	Status       string `json:"status"`
}
