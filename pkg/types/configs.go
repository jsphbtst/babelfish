package types

type Defaults struct {
	TargetLanguage string `json:"targetLanguage"`
	Stream         bool   `json:"stream"`
}

type Configs struct {
	Defaults Defaults `json:"defaults"`
}
