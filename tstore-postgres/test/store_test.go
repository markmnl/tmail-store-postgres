package testpgstore

import (
	"testing"
	"github.com/markmnl/tmail-store/tstore/pkg"
	"github.com/markmnl/tmail-store-postgres/tstore-postgres/pkg"
)

// TestStoreMsg tests store method
func TestStoreMsg(t *testing.T) {
	msg := new(tstore.Msg)
	err := pgstore.Store(msg)
	if err != nil {
		t.Error(err)
	}
}