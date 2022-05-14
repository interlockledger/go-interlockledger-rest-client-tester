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

package json

import (
	stdjson "encoding/json"
	"fmt"
	"os"

	"github.com/interlockledger/go-interlockledger-rest-client-tester/cmd/commands/flags"
	"github.com/interlockledger/go-interlockledger-rest-client/crypto"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var JSONRootCmd = &cobra.Command{
	Use:   "json",
	Short: "Execute JSON document API calls.",
}

func init() {
	JSONRootCmd.AddCommand(jsonAddCmd)
	JSONRootCmd.AddCommand(jsonGetCmd)
	JSONRootCmd.AddCommand(jsonAddWithKeyCmd)

	JSONRootCmd.PersistentFlags().StringVar(&flags.Flags.JSONFile, "json", "", "The JSON file to add. Defaults to \"{\"dummy\": \"DUMMY\"}\"")
	JSONRootCmd.PersistentFlags().StringVar(&flags.Flags.CertFile, "cert", "", "The public key certificate.")
	JSONRootCmd.PersistentFlags().Int64VarP(&flags.Flags.Id, "id", "i", int64(-1), "The ID of the document. It may be required by some commands.")
	JSONRootCmd.PersistentFlags().StringVar(&flags.Flags.ReaderCertFile, "reader-cert", "", "The certificate file that contains the reader key.")
}

func loadJSON() (map[string]any, error) {

	if flags.Flags.JSONFile == "" {
		return map[string]any{"dummy": "DUMMY"}, nil
	}
	bytes, err := os.ReadFile(flags.Flags.JSONFile)
	if err != nil {
		return nil, err
	}
	ret := make(map[string]any)
	err = stdjson.Unmarshal(bytes, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func LoadReaderCertificate(file string) (crypto.ReaderKey, error) {
	// Load the reader certificate
	cert, err := crypto.LoadCertificate(file)
	if err != nil {
		return nil, fmt.Errorf("Unable to load the reader certificate: %w", err)
	}
	readerKey, err := crypto.NewReaderKey(cert[0].PublicKey, nil)
	if err != nil {
		return nil, fmt.Errorf("Unable to extract the public key: %w", err)
	}
	return readerKey, nil
}
