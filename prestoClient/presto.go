package prestoClient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//Config to create a client instance
type Config struct {
	URL     string
	User    string
	Catalog string
	Schema  string
}

//CreateClient creates a Client instance
func CreateClient(config Config) (Client, error) {
	return &client{&config}, nil
}

type Result struct {
}

//Client is the instance user uses to interact with the library
type Client interface {
	Config() *Config
	ExecuteSync(query string) (*Result, error)
}

type client struct {
	config *Config
}

func (client *client) Config() *Config {
	return client.config
}

func (client *client) ExecuteSync(query string) (*Result, error) {

	req, err := http.NewRequest("POST", prestoURL(client.Config().URL), strings.NewReader(query))
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Presto-User", client.Config().User)
	req.Header.Add("X-Presto-Catalog", client.Config().Catalog)
	req.Header.Add("X-Presto-Schema", client.Config().Schema)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	fmt.Printf("Response code %d \n", res.StatusCode)

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Error in Executing Query %v", res.Status)
	}
	var result queryResults
	json.NewDecoder(res.Body).Decode(&result)
	fmt.Printf("Result %v", result.NextURI)
	return &Result{}, nil
}

func prestoURL(url string) string {
	return fmt.Sprintf("http://%s/v1/statement", url)
}
