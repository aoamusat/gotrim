package utils

import "encoding/json"

type UrlResponse struct {
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

func (r *UrlResponse) ToJSON() ([]byte, error) {
	JsonData, err := json.MarshalIndent(r, "", "   ")
	if err != nil {
		return JsonData, err
	}

	return JsonData, nil
}
