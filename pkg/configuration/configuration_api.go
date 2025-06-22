package configuration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type API struct {
	Url       string
	Headers   map[string]string
	Namespace string
}

func NewAPI(url string, namespace string) *API {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting hostname: %v\n", err)
		hostname = "unknown"
	}

	file, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err == nil {
		namespace = string(file)
	}

	headers := make(map[string]string)
	headers["hostname"] = hostname

	return &API{
		Url:       url,
		Headers:   headers,
		Namespace: namespace,
	}
}

func (c *API) Get(key string) (string, error) {
	fullUrl := c.Url
	if c.Namespace != "" {
		fullUrl = fullUrl + "." + c.Namespace
	}

	req, err := http.NewRequest("GET", fullUrl+"?key="+key, nil)
	if err != nil {
		return "", err
	}

	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		responder := resp.Header.Get("responder")
		responderLog := resp.Header.Get("responder_log")
		return "", fmt.Errorf("API request to %s failed with status: %d from repsponder: %s with log: %s", fullUrl, resp.StatusCode, responder, responderLog)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (c *API) GetFromJson(key string, structure any) error {
	strValue, err := c.Get(key)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(strValue), structure)
	if err != nil {
		return err
	}

	return nil
}
