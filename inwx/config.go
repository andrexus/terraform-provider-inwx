package inwx

import (
	"log"

	"github.com/andrexus/goinwx"
)

type Config struct {
	Username string
	Password string
	Sandbox  bool
}

func (c *Config) Client() (*goinwx.Client, error) {
	clientOpts := &goinwx.ClientOptions{Sandbox:c.Sandbox}
	client := goinwx.NewClient(c.Username, c.Password, clientOpts)
	err := client.Login()
	if err != nil {
		return nil, err
	}

	log.Printf("[INFO] INWX client configured for URL: %s", client.BaseURL.String())

	return client, nil
}
