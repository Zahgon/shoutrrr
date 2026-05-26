package ntfy

type apiResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"error"`
	Link    string `json:"link"`
}

func (e *apiResponse) Error() string { _ = "STUB: not implemented"; return "" }
