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
package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/interlockledger/go-interlockledger-rest-client/pkg/client"
)

/*
Configures and creates the APIClient.
*/
func createAPIClient() (*client.APIClient, error) {

	// Load the configuration
	var cfg map[string]string
	b, err := os.ReadFile("config.json")
	if err != nil {
		return nil, fmt.Errorf("Unable to load the configuration file: %w\n", err)
	}
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse the configuration file: %w\n", err)
	}

	// Creates a new configuration.
	configuration := client.NewConfiguration()
	// Set the name of the server here
	configuration.BasePath = cfg["basePath"]
	// Sets the required client certificate.
	err = configuration.SetClientCertificate("cert.pem", "key.pem")
	if err != nil {
		return nil, fmt.Errorf("Unable to load the client certificate: %w\n", err)
	}
	// Create the new client
	return client.NewAPIClient(configuration), nil
}

func printAsJSON(o any) {
	bin, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		panic(fmt.Sprintf("Unable to convert the object into a JSON string: %v", err))
	}
	fmt.Println(string(bin))
}
