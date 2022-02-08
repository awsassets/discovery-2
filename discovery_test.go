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

	t.Log(d)
}
