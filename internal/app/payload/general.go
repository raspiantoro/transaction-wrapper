package payload

type ErrorResponses struct {
	StatusCode uint64 `json:"status_code"`
	Message    string `json:"message"`
}
