package common

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	logger "github.com/rs/zerolog/log"
)

type Configuration struct {
	SQLURI         string
	MongoURI       string
	DBNAME         string
	SERVERORIGIN   string
	GITHUBTOKEN    string
	GITHUBUSERNAME string
	ALLOWEDORIGIN  string
	SECRETKEY      []byte
	Ports
	NewRelic
	Collections
}
type Ports struct {
	Server   string
	Todos    string
	GitHub   string
	Schedule string
}
type NewRelic struct {
	Application string
	License     string
	LogForward  bool
}
type Collections struct {
	BIODATA    string
	GITHUBDATA string
	TODOS      string
	GRAPHDATA  string
	SCHEDULE   string
}

func FetchConfig() Configuration {
	return Configuration{
		DBNAME:         os.Getenv("DB_NAME"),
		SQLURI:         os.Getenv("SQL_URI"),
		MongoURI:       os.Getenv("MONGO_URI"),
		GITHUBTOKEN:    os.Getenv("GITHUB_TOKEN"),
		ALLOWEDORIGIN:  os.Getenv("ALLOWED_ORIGIN"),
		GITHUBUSERNAME: os.Getenv("GITHUB_USERNAME"),
		SERVERORIGIN:   fmt.Sprintf("http://localhost:%v", os.Getenv("SERVER_PORT")),
		SECRETKEY:      fetchSecretKey(),
		Ports: Ports{
			Server:   os.Getenv("SERVER_PORT"),
			Todos:    os.Getenv("TODOS"),
			GitHub:   os.Getenv("GITHUB"),
			Schedule: os.Getenv("SCHEDULE"),
		},
		Collections: Collections{
			BIODATA:    os.Getenv("BIO"),
			GITHUBDATA: os.Getenv("GITHUB"),
			TODOS:      os.Getenv("TODOS"),
			GRAPHDATA:  os.Getenv("GRAPH"),
			SCHEDULE:   os.Getenv("SCHEDULE"),
		},
		NewRelic: NewRelic{
			Application: os.Getenv("NEWRELIC_APP"),
			License:     os.Getenv("NEWRELIC_LICENSE"),
			LogForward:  os.Getenv("NEWRELIC_LOG_FORWARD") == "true",
		},
	}
}

func fetchSecretKey() (key []byte) {
	// Read the secret key from the environment variable
	encodedKey := os.Getenv("SECRET_KEY")
	key, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil {
		logger.Info().Err(err).Msg("failed to decode secret key: ")
		return []byte("")
	}
	return
}

func BearerAuthAPICall(reqURL, authToken string, queryParams ...string) ([]byte, int) {
	client := http.Client{}
	request, err := http.NewRequest("GET", reqURL, bytes.NewBuffer([]byte("")))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", authToken))
	if err != nil {
		logger.Info().Err(err).Msg(fmt.Sprintf("err in making request to %v: %v", reqURL, err))
		return []byte{}, 503
	}

	if len(queryParams) > 0 {
		request.URL.RawQuery = addQueryParameters(request.URL.Query(), queryParams...)
	}

	resp, err := client.Do(request)

	if resp.StatusCode != http.StatusOK {
		logger.Info().Err(fmt.Errorf("error '%v' while making request to %v: %v", err, reqURL, resp.StatusCode))
		return []byte{}, resp.StatusCode
	}

	if err != nil {
		logger.Info().Err(err).Msg("err in reading req response: ")
		return []byte{}, resp.StatusCode
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Info().Err(err).Msg("err in reading req response: ")
		return []byte{}, 503
	}

	if strings.EqualFold(string(respBody), "[]") {
		return []byte{}, 503
	}
	return respBody, resp.StatusCode
}

func NoAuthAPICall(reqURL, origin string, reqBody []byte, queryParams ...string) ([]byte, int) {
	timeOut := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeOut,
	}

	request, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(reqBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Origin", origin)
	if err != nil {
		logger.Info().Err(err).Msg("err in creating new request: ")
		return []byte{}, 503
	}

	if len(queryParams) > 0 {
		request.URL.RawQuery = addQueryParameters(request.URL.Query(), queryParams...)
	}

	resp, err := client.Do(request)

	if resp.StatusCode != http.StatusOK {
		logger.Info().Err(fmt.Errorf("error '%v' while making request to %v: %v", err, reqURL, resp.StatusCode))
		return []byte{}, resp.StatusCode
	}

	if err != nil {
		logger.Info().Err(err).Msg("err in reading req response: ")
		return []byte{}, resp.StatusCode
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Info().Err(err).Msg("err in reading req response: ")
		return []byte{}, 503
	}

	return respBody, resp.StatusCode
}

func addQueryParameters(q url.Values, queryParams ...string) string {
	for _, param := range queryParams {
		keyVal := strings.SplitN(param, " ", 2)
		q.Add(keyVal[0], keyVal[1])
	}
	return q.Encode()
}
