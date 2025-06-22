package configuration

import "encoding/json"

type Configuration struct {
	UseEnv bool `json:"use_console"`

	UseConfigFile  bool   `json:"use_config_file"`
	ConfigFilePath string `json:"config_file_path"`

	UseConfigString bool   `json:"use_config_string"`
	ConfigString    string `json:"config_string"`

	UseApi bool   `json:"use_api"`
	ApiUrl string `json:"api_url"`
}

func NewConfiguration() *Configuration {
	return &Configuration{}
}

func FromJsonString(jsonString string) (*Configuration, error) {
	var config Configuration
	err := json.Unmarshal([]byte(jsonString), &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *Configuration) Init() *IConfiguration {
	var config *IConfiguration = nil

	if c.UseApi {
		if c.ApiUrl != "" {
			api := NewAPI(c.ApiUrl)
			iconfig := IConfiguration(api)
			config = &iconfig
		}
	}

	return config
}
