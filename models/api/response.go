package api

type Response struct {
	Error        bool        `json:"error"`
	ErrorMessage string      `json:"errorMessage"`
	Payload      interface{} `json:"payload"`
}
