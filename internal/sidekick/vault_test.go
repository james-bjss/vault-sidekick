package sidekick

import (
	"log"
	"net"
	"testing"

	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/http"
	"github.com/hashicorp/vault/vault"
)

func TestVault(t *testing.T) {
	ln, _ := createInMemoryVault(t)

	// TODO: Need to make vault serice testable by explicitly passing in the vault options
	vaultService, err := NewVaultService(ln.Addr().String())

	if err != nil {
		t.Logf("Failed to Create Vault Service: %v", err)
		t.Fail()
	}

	if vaultService == nil {
		t.Logf("Should return a valid pointer to vault service")
		t.Fail()
	}
}

func createInMemoryVault(t *testing.T) (net.Listener, *api.Client) {
	core, _, token := vault.TestCoreUnsealed(t)
	ln, addr := http.TestServer(t, core)

	conf := api.DefaultConfig()
	conf.Address = addr

	log.Print(conf.Address)

	client, err := api.NewClient(conf)
	if err != nil {
		t.Fatal(err)
	}
	client.SetToken(token)

	// Check we can write a secret
	_, err = client.Logical().Write("secret/foo", map[string]interface{}{
		"secret": "bar",
	})

	if err != nil {
		t.Fatal(err)
	}

	return ln, client
}
