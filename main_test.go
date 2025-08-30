package kernel

import (
	"github.com/zalando/go-keyring"
	"testing"
)

// TestMain - a special function that runs all tests
// mocks keyring to successfully pass tests
func TestMain(m *testing.M) {
	keyring.MockInit()

	m.Run()
}
