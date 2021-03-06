package sources

import (
	"testing"

	"github.com/OWASP/Amass/config"
	"github.com/OWASP/Amass/requests"
	"github.com/OWASP/Amass/resolvers"
)

func TestSpyse(t *testing.T) {
	if *networkTest == false || *configPath == "" {
		return
	}

	cfg := setupConfig(domainTest)

	API := new(config.APIKey)
	API = cfg.GetAPIKey("spyse")

	if API == nil || API.Key == "" || API.Secret == "" {
		t.Errorf("API key data was not provided")
		return
	}

	bus, out := setupEventBus(requests.NewNameTopic)
	defer bus.Stop()

	pool := resolvers.NewResolverPool(nil)
	defer pool.Stop()

	srv := NewSpyse(cfg, bus, pool)

	result := testService(srv, out)
	if result < expectedTest {
		t.Errorf("Found %d names, expected at least %d instead", result, expectedTest)
	}
}
