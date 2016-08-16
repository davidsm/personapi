package person

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Set a constant seed for all tests for reproducible results
	Seed(12345)
	os.Exit(m.Run())
}
