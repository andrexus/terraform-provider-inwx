package inwx

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a schema.Provider for INWX.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INWX_USERNAME", nil),
				Description: "Username for API operations.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INWX_PASSWORD", nil),
				Description: "Password for API operations.",
			},
			"TAN": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("INWX_TAN", nil),
				Description: "TAN for account unlock.",
			},
			"sandbox": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("INWX_SANDBOX", false),
				Description: "Use sandbox environment (api.ote.domrobot.com).",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"inwx_record": resourceINWXRecord(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		TAN:      d.Get("TAN").(string),
		Sandbox:  d.Get("sandbox").(bool),
	}

	return config.Client()
}
