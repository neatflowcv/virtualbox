//go:build virtualbox

package virtualbox_test

import (
	"testing"

	"github.com/neatflowcv/virtualbox/pkg/virtualbox"
)

func TestListVMs(t *testing.T) {
	t.Parallel()

	vms, err := virtualbox.ListVMs(t.Context())
	if err != nil {
		t.Fatalf("failed to list vms: %v", err)
	}

	for _, vm := range vms {
		t.Logf("vm: (%s) (%s)", vm.Name(), vm.UUID())
	}
}
