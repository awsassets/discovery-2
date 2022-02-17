package discovery

import (
	"context"
	"testing"
)

func TestDiscover(t *testing.T) {
	r := NewResolver("consul:8600")
	d, err := r.Discover(
		context.Background(),
		"",
		"",
		"vault.service.dc1.consul",
	)
	if err != nil {
		t.Fatal(err)
	}

	if d[0].Target != "ac110003.addr.dc1.consul" {
		t.Fatalf("want ac110003.addr.dc1.consul, got %v", d[0].Target)
	}

	if d[0].Address.String() != "172.17.0.3" {
		t.Fatalf("want 172.17.0.3, got %v", d[0].Address.String())
	}

	if d[0].Port != uint16(8200) {
		t.Fatalf("want 8200, got %v", d[0].Port)
	}
}
