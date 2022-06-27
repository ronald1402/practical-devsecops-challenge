package dto

type JDoodleRequestDto struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	ClientScript string `json:"script"`
	Language     string `json:"language"`
	VersionIndex string `json:"versionIndex"`
}
