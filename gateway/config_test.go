// Copyright (c) Alex Ellis 2017. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package main

import (
	"testing"
	"time"
)

type EnvBucket struct {
	Items map[string]string
}

func NewEnvBucket() EnvBucket {
	return EnvBucket{
		Items: make(map[string]string),
	}
}

func (e EnvBucket) Getenv(key string) string {
	return e.Items[key]
}

func (e EnvBucket) Setenv(key string, value string) {
	e.Items[key] = value
}

func TestRead_EmptyTimeoutConfig(t *testing.T) {
	defaults := NewEnvBucket()
	readConfig := ReadConfig{}

	config := readConfig.Read(defaults)

	if (config.readTimeout) != time.Duration(8)*time.Second {
		t.Log("readTimeout incorrect")
		t.Fail()
	}
	if (config.writeTimeout) != time.Duration(8)*time.Second {
		t.Log("writeTimeout incorrect")
		t.Fail()
	}
}

func TestRead_ReadAndWriteTimeoutConfig(t *testing.T) {
	defaults := NewEnvBucket()
	defaults.Setenv("read_timeout", "10")
	defaults.Setenv("write_timeout", "60")

	readConfig := ReadConfig{}
	config := readConfig.Read(defaults)

	if (config.readTimeout) != time.Duration(10)*time.Second {
		t.Logf("readTimeout incorrect, got: %d\n", config.readTimeout)
		t.Fail()
	}
	if (config.writeTimeout) != time.Duration(60)*time.Second {
		t.Logf("writeTimeout incorrect, got: %d\n", config.writeTimeout)
		t.Fail()
	}
}