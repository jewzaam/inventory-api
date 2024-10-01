package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOptions(t *testing.T) {
	t.Run("default options", func(t *testing.T) {
		opts := NewOptions()

		expectedOpts := &Options{
			Addr:    "0.0.0.0:8000",
			Timeout: 300,
			CertOpt: 3,
		}

		assert.Equal(t, expectedOpts, opts)
	})
}
