package config

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"

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
	os.Setenv("LOCAL", "0")
	os.Setenv("GH", "")
	os.Setenv("SECRET_KEY", "123")
	if strings.EqualFold("0", os.Getenv("LOCAL")) {
		return Configuration{
			SQLURI:         "mysql://localhost:3306/db",
			MongoURI:       "mongodb://localhost:27017/database",
			DBNAME:         "test",
			SERVERORIGIN:   "*",
			GITHUBTOKEN:    os.Getenv("GH"),
			GITHUBUSERNAME: "shashank-priyadarshi",
			ALLOWEDORIGIN:  "*",
			SECRETKEY:      fetchSecretKey(),
			Ports: Ports{
				Server:   "8085",
				Todos:    "8086",
				GitHub:   "8087",
				Schedule: "8088",
			},
			Collections: Collections{
				BIODATA:    "b",
				GITHUBDATA: "g",
				TODOS:      "t",
				GRAPHDATA:  "gr",
				SCHEDULE:   "s",
			},
			NewRelic: NewRelic{
				Application: "",
				License:     "",
				LogForward:  false,
			},
		}
	}
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
