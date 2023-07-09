package remotewrite

import (
	"testing"
)

func Test_hash(t *testing.T) {
	var name string = "allsite"
	id := hashTenant(name)
	if "3234220490:2606740347" != id {
		t.Fatalf("failed to hash tenant: %s", name)
	}
}
