package configuration

import "encoding/json"

type Configuration struct {
	UseEnv bool `json:"use_console"`

	UseConfigFile  bool   `json:"use_config_file"`
	ConfigFilePath string `json:"config_file_path"`

	UseConfigString bool   `json:"use_config_string"`
	ConfigString    string `json:"config_string"`

	UseApi       bool   `json:"use_api"`
	ApiUrl       string `json:"api_url"`
	ApiNamespace string `json:"api_namespace"`
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

func (c *Configuration) Init() IConfiguration {
	if c.UseApi {
		if c.ApiUrl != "" {
			return NewAPI(c.ApiUrl, c.ApiNamespace)
		}
	}

	return nil
}
