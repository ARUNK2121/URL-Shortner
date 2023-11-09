package models

type ShortResponse struct {
	URL   string `json:"shorted_url"`
	Error string `json:"error"`
}
