package models

type HttpResponse struct {
	Code    string `json:"code,omitempty"`
	Massage string `json:"massage,omitempty"`
	Data    any    `json:"data,omitempty"`
}
