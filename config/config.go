package config

import (
	"errors"
	"os"
)

// LoadApikey is similar to LoadApikeyFromConfig. return MACKEREL_APIKEY environment value if defined MACKEREL_APIKEY
func LoadApikey() (apiKey string, err error) {
	apiKey = ""
	if apiKey = os.Getenv("MACKEREL_APIKEY"); apiKey != "" {
		return apiKey, nil
	}
	// ToDo: Support getting apiKey from mackerel-agent.conf file
	// key := LoadApikeyFromConfig(conf)
	return "", errors.New("failed to get mackerel apikey")
}
