package accounting

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"runtime"
)

type (
	Configuration struct {
		// Listener http listener binding options.
		Listener Listener `json:"listener"`

		// DSN database connection string.
		DSN string `json:"dsn"`

		// Database is name of primary database name for entire app.
		// default: mml_be
		Database string `json:"database"`

		// JWT is all config for make token, ..
		JWT Jwt `json:"jwt"`
	}

	// Listener contains Server https listener options.
	Listener struct {
		// Host is network address for bind Server http listener to it.
		// default: 127.0.0.1
		Host string `json:"host" mapstructure:"host"`

		// Port is network port for bind Server http listener to it.
		// default: 8080
		Port int `json:"port" mapstructure:"port"`

		// Cert is path to TLS certificate file.
		// if Cert is not specified, Server listener runs without TLS.
		Cert string `json:"cert" mapstructure:"cert"`

		// Key is path to TLS certificate PrivateKey file.
		// it ignored if Cert is not specified.
		Key string `json:"key" mapstructure:"key"`

		// AllowedHosts is allowed host for CORS configuration.
		// It applied in production mode
		AllowedHosts []string `json:"allowed_hosts" mapstructure:"allowed_hosts"`

		// SSLHost is ssl host for gin secure configuration.
		// It applied in production mode
		SSLHost string `json:"ssl_host" mapstructure:"ssl_host"`
	}

	// Jwt contains JWT configuration options.
	Jwt struct {
		Secret        string `json:"secret"`
		TokenExpire   int64  `json:"token_expire"`
		RefreshExpire int64  `json:"refresh_expire"`
		Issuer        string `json:"issuer"`
		Audience      string `json:"audience"`
		SubjectKey    string `json:"subject_key"`
		IdentityKey   string `json:"identity_key"`
		RoleKey       string `json:"role_key"`
	}
)

// Initialize configuration
func NewConfiguration() (*Configuration, error) {
	config := &Configuration{}

	err := config.loadConfig()
	if err != nil {
		return nil, err
	}

	return config, nil
}

// load config file
func (c *Configuration) loadConfig() error {
	path := ""

	if runtime.GOOS == "windows" {
		path = ".\\config\\config.json"
	} else {
		path = "./config/config.json"
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer closeFile(file)

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, c)
	if err != nil {
		return err
	}

	return nil
}

// close file system
func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Panicln(err.Error())
	}
}
