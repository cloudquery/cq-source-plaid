package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/plaid/plaid-go/v10/plaid"
	"github.com/rs/zerolog"
)

func TestServer(t *testing.T, data any) *httptest.Server {
	t.Helper()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data, _ := json.Marshal(data)
		_, _ = w.Write(data)
	}))
	return ts
}

func TestHelper(t *testing.T, table *schema.Table, ts *httptest.Server) {
	version := "vDev"
	table.IgnoreInTests = false
	t.Helper()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro}).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source, opts source.Options) (schema.ClientMeta, error) {
		configuration := plaid.NewConfiguration()
		configuration.UseEnvironment(plaid.Development)
		urlParts := strings.Split(ts.URL, "://")
		configuration.Scheme = urlParts[0]
		configuration.Host = urlParts[1]
		client := plaid.NewAPIClient(configuration)
		s := Spec{
			ClientId:    "test",
			Secret:      "test",
			AccessToken: "test",
		}
		s.SetDefaults()
		err := s.Validate()
		if err != nil {
			return nil, err
		}
		return &Client{
			Logger:   l,
			Services: client,
			ClientId: s.ClientId,
			Secret:   s.Secret,
		}, nil
	}
	p := source.NewPlugin(
		table.Name,
		version,
		[]*schema.Table{
			table,
		},
		newTestExecutionClient)
	p.SetLogger(l)
	source.TestPluginSync(t, p, specs.Source{
		Name:         "dev",
		Path:         "cloudquery/dev",
		Version:      version,
		Tables:       []string{table.Name},
		Destinations: []string{"mock-destination"},
	})
}
