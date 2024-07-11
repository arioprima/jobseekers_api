package schemas

type SchemaDatabaseError struct {
	Code int    `json:"code"`
	Type string `json:"type"`
}

type SchemaErrorResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Error  interface{} `json:"error"`
}

type SchemaUnauthorizedResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ErrorMetaConfig struct {
	Tag     string `json:"tag"`
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   string `json:"value,omitempty"`
}

type ErrorConfig struct {
	Options []ErrorMetaConfig `json:"options"`
}
