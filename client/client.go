package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/plaid/plaid-go/v10/plaid"
	"github.com/rs/zerolog"
)

type Client struct {
	Logger      zerolog.Logger
	Services    *plaid.APIClient
	ClientId    string
	Secret      string
	AccessToken string
}

func (c *Client) ID() string {
	return "plaid"
}

type httpLogger struct {
	zerolog.Logger
}

func (l httpLogger) Printf(format string, v ...interface{}) {
	if strings.Contains(format, "retrying") {
		l.Logger.Info().Msgf(format, v...)
	} else {
		l.Logger.Debug().Msgf(format, v...)
	}
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}

	pluginSpec.SetDefaults()
	if err := pluginSpec.Validate(); err != nil {
		return nil, err
	}

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10
	retryClient.Logger = httpLogger{logger}

	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", pluginSpec.ClientId)
	configuration.AddDefaultHeader("PLAID-SECRET", pluginSpec.Secret)
	configuration.UseEnvironment(Environments[pluginSpec.Environment])
	configuration.HTTPClient = retryClient.StandardClient()

	client := plaid.NewAPIClient(configuration)

	return &Client{
		Logger:      logger,
		Services:    client,
		ClientId:    pluginSpec.ClientId,
		Secret:      pluginSpec.Secret,
		AccessToken: pluginSpec.AccessToken,
	}, nil
}
