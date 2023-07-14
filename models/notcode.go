package models

type Req struct {
	Language string `json:"language"`
	Version  string `json:"version"`
	Code     string `json:"code"`
	Cases    string `json:"cases"`
	Test     string `json:"test"`
}
