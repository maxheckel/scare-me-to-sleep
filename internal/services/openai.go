package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/maxheckel/scare-me-to-sleep/internal/config"
	"github.com/maxheckel/scare-me-to-sleep/internal/domain"
	"io/ioutil"
	"net/http"
	"strings"
)

//AddOpenIDAuthHeader Used for the http transport which adds the open AI client secret to every request
type AddOpenIDAuthHeader struct {
	T      http.RoundTripper
	Config *config.Config
}

func (adt AddOpenIDAuthHeader) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", adt.Config.OpenAIKey))
	return adt.T.RoundTrip(req)
}

type OpenAIClient struct {
	Config     *config.Config
	HttpClient *http.Client
}

type OpenAI interface {
	Generate(prompt *domain.Prompt) (*domain.Response, error)
}

func NewOpenAIClient(config *config.Config) OpenAI {
	httpClient := &http.Client{
		Transport: AddOpenIDAuthHeader{T: http.DefaultTransport, Config: config},
	}
	return OpenAIClient{
		Config:     config,
		HttpClient: httpClient,
	}
}

func (oac OpenAIClient) Generate(prompt *domain.Prompt) (*domain.Response, error) {
	bodyObj := domain.GenerateRequest{
		Model:            "curie:ft-google-2022-07-18-17-48-55",
		Prompt:           prompt.Text,
		Temperature:      0.7,
		MaxTokens:        1024,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		Stop:             "+++",
	}
	bodyData, err := json.Marshal(bodyObj)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(bodyData))
	req.Header.Add("Content-Type", "application/json")
	resp, err := oac.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	generated := domain.GenerateResponse{}
	err = json.Unmarshal(body, &generated)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	generated.Choices[0].Text = strings.ReplaceAll(generated.Choices[0].Text, "+", "")
	generated.Choices[0].Text = strings.ReplaceAll(generated.Choices[0].Text, "-> ", "")
	generated.Choices[0].Text = strings.TrimSpace(generated.Choices[0].Text)
	return &domain.Response{
		Text:     generated.Choices[0].Text,
		PromptID: prompt.ID,
		Votes:    0,
	}, nil
}
