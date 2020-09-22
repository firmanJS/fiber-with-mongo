package models

type ResultsData struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Results interface{} `json:"results"`
	Meta    interface{} `json:"_meta"`
}