package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
)

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

func GenerateRandomString(length int) string {
	if length == 0 {
		length = 6
	}
	var randomStr string = ""
	var alphabets string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	AlphaList := strings.Split(alphabets, "")

	for i := 0; i < length; i++ {
		randomStr += AlphaList[rand.Intn(len(AlphaList)-1)]
	}

	return randomStr
}

func GenerateShortUrl(LongUrl string) string {
	return fmt.Sprintf("https://go.trim/%s", GenerateRandomString(4))
}
