// BSD 3-Clause License
//
// Copyright (c) 2022, InterlockLedger
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// This package contains the implementation of the commands used to test each
// IL2 API.
package core

import (
	"fmt"

	"github.com/interlockledger/go-interlockledger-rest-client/client"
)

// This is the application's core singleton.
var AppCore ApplicationCore

// The application's core type.
type ApplicationCore struct {
	Config Configuration
}

// Loads the configuration from the file.
func (c *ApplicationCore) LoadConfig(configFile string) error {
	if c.Config.Loaded {
		return nil
	}
	return c.Config.Load(configFile)
}

// Creates a new API client based on the loaded configuration.
func (c *ApplicationCore) NewClient() (*client.APIClient, error) {
	if !c.Config.Loaded {
		return nil, fmt.Errorf("The configuration was not loaded.")
	}
	// Creates a new configuration.
	configuration := client.NewConfiguration()
	// Set the name of the server here
	configuration.BasePath = c.Config.BasePath
	configuration.NoServerVerification = true
	configuration.CertFile = c.Config.CertFile
	configuration.KeyFile = c.Config.KeyFile
	configuration.PFXFile = c.Config.PFXFile
	configuration.PFXPassword = c.Config.PFXPassword

	if err := configuration.Init(); err != nil {
		return nil, fmt.Errorf("unable to load the client certificate: %w", err)
	}
	// Create the new client
	return client.NewAPIClient(configuration), nil
}
