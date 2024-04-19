package models

type FileApiResponse struct {
	Success int    `json:"success"`
	Message string `json:"message"`
	Url     string `json:"url"`
}
