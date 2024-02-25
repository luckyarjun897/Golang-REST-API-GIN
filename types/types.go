package types

type Movies struct {
	Name       string  `json:"name"`
	Language   string  `json:"language"`
	Budget     float32 `json:"budget"`
	Collection float32 `json:"collection"`
}