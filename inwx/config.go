package inwx

import (
	"log"

	"github.com/andrexus/goinwx"
	"strings"
)

type Config struct {
	Username string
	Password string
	TAN      string
	Sandbox  bool
}

func (c *Config) Client() (*goinwx.Client, error) {
	clientOpts := &goinwx.ClientOptions{Sandbox: c.Sandbox}
	client := goinwx.NewClient(c.Username, c.Password, clientOpts)
	err := client.Account.Login()
	if err != nil {
		return nil, err
	}
	if c.TAN != "" {
		unlockErr := client.Account.Unlock(strings.Replace(c.TAN, " ", "", -1))
		if err != nil {
			return nil, unlockErr
		}
	}

	log.Printf("[INFO] INWX client configured for URL: %s", client.BaseURL.String())

	return client, nil
}
