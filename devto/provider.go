package devto

import (
	"context"

	"github.com/cuotos/devto-go/devto"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	provider := &schema.Provider{}

	provider.ResourcesMap = map[string]*schema.Resource{
		"devto_article": resourceArticle(),
	}

	provider.DataSourcesMap = map[string]*schema.Resource{}

	provider.Schema = map[string]*schema.Schema{
		"api_key": {
			Type:     schema.TypeString,
			Required: true,
		},
		"api_url": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}

	provider.ConfigureContextFunc = func(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
		apiKey := data.Get("api_key").(string)
		apiURL := data.Get("api_url").(string)

		opts := []devto.Option{}

		if apiURL != "" {
			opts = append(opts, devto.WithAPIURL(apiURL))
		}

		client, err := devto.New(apiKey, opts...)

		if err != nil {
			return nil, diag.FromErr(err)
		}

		return client, nil
	}

	return provider
}
