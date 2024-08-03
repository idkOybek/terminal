package models

type ExportRequest struct {
	Filename string                   `json:"filename"`
	Objects  []map[string]interface{} `json:"objects"`
}
