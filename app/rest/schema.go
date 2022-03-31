package rest

import "time"

type Base struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type HTTPResponse struct {
	Msg string `json:"msg,omitempty" example:"any message"`
}

type IDResponse struct {
	ID string `json:"id"`
}

type CreateGuestRequest struct {
	Name string `json:"name"`
	Doc  string `json:"doc"`
}

type Guest struct {
	Base `json:",inline"`
	Name string `json:"name"`
	Doc  string `json:"doc"`
}

type SearchGuestRequest struct {
	Last     int `json:"last" query:"last"`
	PageSize int `json:"page_size" query:"page_size"`
}

type SearchGuestResponse struct {
	Guests []*Guest `json:"guests"`
	Last   int      `json:"last"`
}
