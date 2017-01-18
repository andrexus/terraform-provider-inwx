package inwx

import (
	"log"

	"github.com/andrexus/goinwx"
	"os"
)

type Config struct {
	Username string
	Password string
	Sandbox  bool
}

func (c *Config) Client() (*goinwx.Client, error) {
	client := goinwx.NewClient(c.Username, c.Password)
	err := client.Login()
	if err != nil {
		return nil, err
	}
	if c.Sandbox == true {
		os.Setenv("GOINWX_SANDBOX", "true")
	}

	log.Printf("[INFO] INWX client configured for URL: %s", client.BaseURL.String())

	return client, nil
}
