package inwx

import (
	"log"
	"strings"

	"github.com/andrexus/goinwx"
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
	log.Printf("[INFO] Trying to login with provided credentials")
	err := client.Account.Login()
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] Login successful")
	if c.TAN != "" {
		log.Printf("[INFO] 2-factor-auth is configured. Trying to unlock account")
		if unlockErr := client.Account.Unlock(strings.Replace(c.TAN, " ", "", -1)); unlockErr != nil {
			return nil, unlockErr
		}
	}

	log.Printf("[INFO] INWX client configured for URL: %s", client.BaseURL.String())

	return client, nil
}
