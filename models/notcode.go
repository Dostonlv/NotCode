package models

type Req struct {
	Language string `json:"language"`
	Version  string `json:"version"`
	Code     string `json:"code"`
}

type Test struct {
	Inputs  [3]string
	Outputs [3]string
}
