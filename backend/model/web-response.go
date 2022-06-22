package model

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Token  string      `json:"token ,omitempty"`
	Data   interface{} `json:"data"`
}
