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
	ALLOWEDORIGIN  string
	GITHUBTOKEN    string
	GITHUBUSERNAME string
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
	if strings.EqualFold("0", os.Getenv("SETUP")) {
		logger.Info().Msg("Running portfolio setup in dev environment")
		return Configuration{
			SQLURI:         "root@tcp(mysql:3306)/db",
			MongoURI:       "mongodb://mongodb:27017/test",
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
			Todos:    os.Getenv("TODOS_PORT"),
			GitHub:   os.Getenv("GITHUB_PORT"),
			Schedule: os.Getenv("SCHEDULE_PORT"),
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
