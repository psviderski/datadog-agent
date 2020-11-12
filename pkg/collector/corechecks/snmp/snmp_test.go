// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

package snmp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicConfiguration(t *testing.T) {
	check := Check{}
	rawInstanceConfig := []byte(`
ip_address: 1.2.3.4
`)
	err := check.Configure(rawInstanceConfig, []byte(``), "test")

	assert.Nil(t, err)
	assert.Equal(t, "1.2.3.4", check.config.instance.IPAddress)
}
