package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionInfo(t *testing.T) {
	assert.Equal(t, "unknown", versionInfo())

	buildVersion = "v1.0.0"
	assert.Equal(t, "v1.0.0", versionInfo())
}
