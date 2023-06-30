package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/marcbudd/linkup-service/linkuperrors"
)

func CreatePostGPT() (string, *linkuperrors.LinkupError) {

	url := os.Getenv("OPENAI_API_URL")
	apiKey := os.Getenv("OPENAI_API_KEY")
	payload := `{
		"prompt": "Schreibe mir einen witzigen Tweet",
		"max_tokens": 30
	}`

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", linkuperrors.New("openai api error: "+err.Error(), http.StatusInternalServerError)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Body = ioutil.NopCloser(strings.NewReader(payload))

	resp, err := client.Do(req)
	if err != nil {
		return "", linkuperrors.New("openai api error: "+err.Error(), http.StatusInternalServerError)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", linkuperrors.New(err.Error(), http.StatusInternalServerError)
	}

	string := string(body)

	if strings.Contains(string, "error") {
		return "", linkuperrors.New("openai api error: "+string, http.StatusInternalServerError)
	}

	return string, nil

}
