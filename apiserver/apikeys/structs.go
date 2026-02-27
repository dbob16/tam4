package apikeys

type ApiKey struct {
	ApiKey string `json:"api_key"`
	Description string `json:"description"`
}

type ApiReq struct {
	ApiPW string `json:"api_pw"`
	Description string `json:"description"`
}
