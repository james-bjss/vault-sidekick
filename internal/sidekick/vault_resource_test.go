// Package sidekick the main vault-sidekick package

/*
Copyright 2015 Home Office All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sidekick

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceFilename(t *testing.T) {
	rn := VaultResource{
		path:     "test_secret",
		resource: "secret",
		options:  map[string]string{},
	}
	assert.Equal(t, "test_secret.secret", rn.GetFilename())
}

func TestIsValid(t *testing.T) {
	resource := defaultVaultResource()
	resource.path = "/test/name"
	resource.resource = "secret"

	assert.Nil(t, resource.IsValid())
	resource.resource = "nothing"
	assert.NotNil(t, resource.IsValid())
	resource.resource = "pki"
	assert.NotNil(t, resource.IsValid())
	resource.resource = "ssh"
	assert.NotNil(t, resource.IsValid())
}
