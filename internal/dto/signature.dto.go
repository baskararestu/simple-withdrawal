package dto

type SignatureRequestDto struct {
	Method       string                 `json:"method" example:"POST"`
	RelativePath string                 `json:"relativePath" example:"/api/withdraw/create"`
	Param         map[string]interface{} `json:"body" swaggertype:"object"`
	Timestamp    string                 `json:"timestamp" example:"2024-01-01T12:00:00Z"`
	SecretKey    string                 `json:"secretKey" example:"dev-secret"`
}

type SignatureResponseDto struct {
	Signature string `json:"signature" example:"f7321c4c3c99..."`
}

type SignatureRequestExample struct {
	Method       string `json:"method" example:"POST"`
	RelativePath string `json:"relativePath" example:"/api/withdraw/create"`
	Body         struct {
		UserID int `json:"userId" example:"1"`
		Amount int `json:"amount" example:"10000"`
	} `json:"body"`
	Timestamp string `json:"timestamp" example:"2024-01-01T12:00:00Z"`
	SecretKey string `json:"secretKey" example:"dev-secret"`
}
