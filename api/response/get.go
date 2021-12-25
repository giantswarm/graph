package response

type Get struct {
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"entity,omitempty"`
}
