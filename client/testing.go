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

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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

func remove(s schema.ColumnList, i string) schema.ColumnList {
	for j, c := range s {
		if c.Name == i {
			return append(s[:j], s[j+1:]...)
		}
	}

	return s
}

func TestHelper(t *testing.T, table *schema.Table, ts *httptest.Server) {
	table.IgnoreInTests = false
	t.Helper()

	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	sched := scheduler.NewScheduler(scheduler.WithLogger(l))

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
		t.Fatal(err)
	}

	c := &Client{
		Logger:      l,
		Services:    client,
		ClientId:    s.ClientId,
		Secret:      s.Secret,
		AccessToken: s.AccessToken,
	}

	tables := schema.Tables{table}
	if err := transformers.TransformTables(tables); err != nil {
		t.Fatal(err)
	}

	// We need to remove additional_properties column from the table as faker cannot generate data with interface{} type
	err = transformers.Apply(tables, func(table *schema.Table) error {
		table.Columns = remove(table.Columns, "additional_properties")
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	messages, err := sched.SyncAll(context.Background(), c, tables)
	if err != nil {
		t.Fatalf("failed to sync: %v", err)
	}
	plugin.ValidateNoEmptyColumns(t, tables, messages)
}
