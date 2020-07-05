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
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/vault/api"
)

// token authentication plugin
type authTokenPlugin struct {
	// the vault client
	client *api.Client
}

// NewUserTokenPlugin creates a new User Token plugin
func NewUserTokenPlugin(client *api.Client) AuthInterface {
	return &authTokenPlugin{
		client: client,
	}
}

// Create retrieves the token from an environment variable or file
func (r authTokenPlugin) Create(cfg *vaultAuthOptions) (string, error) {
	if cfg.FileName != "" {
		content, err := readConfigFile(cfg.FileName, cfg.FileFormat)
		if err != nil {
			return "", err
		}
		// check: ensure we have a token in the file
		token := content.Token
		if token == "" {
			return "", fmt.Errorf("the auth file: %s does not contain a token", cfg.FileName)
		}

		return token, nil
	}

	// step: check the VAULT_TOKEN
	if val := os.Getenv("VAULT_TOKEN"); val != "" {
		return val, nil
	}

	// step: check the VAULT_TOKEN_FILE
	if filepath := os.Getenv("VAULT_TOKEN_FILE"); filepath != "" {
		content, err := ioutil.ReadFile(filepath)
		if err != nil {
			return "", err
		}
		return string(content), nil
	}

	return "", fmt.Errorf("no token provided")
}
