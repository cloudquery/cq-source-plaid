package client

import (
	"fmt"
	"sort"
	"strings"

	"github.com/plaid/plaid-go/v10/plaid"
	"golang.org/x/exp/maps"
)

var Environments = map[string]plaid.Environment{
	"sandbox":     plaid.Sandbox,
	"development": plaid.Development,
	"production":  plaid.Production,
}

type Spec struct {
	ClientId    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
	Environment string `json:"environment"`
}

func (s *Spec) SetDefaults() {
	if s.Environment == "" {
		s.Environment = "sandbox"
	}
}

func (s *Spec) Validate() error {
	errors := make([]string, 0)
	if s.ClientId == "" {
		errors = append(errors, `"client_id" is required`)
	}

	if s.Secret == "" {
		errors = append(errors, `"secret" is required`)
	}

	if s.AccessToken == "" {
		errors = append(errors, `"access_token" is required`)
	}

	if _, ok := Environments[s.Environment]; !ok {
		validValues := maps.Keys(Environments)
		sort.Strings(validValues)
		errors = append(errors, fmt.Sprintf(`invalid "environment". Expected one of %q, got %q`, validValues, s.Environment))
	}

	if len(errors) > 0 {
		return fmt.Errorf("invalid plugin spec: %s", strings.Join(errors, ". "))
	}

	return nil
}
