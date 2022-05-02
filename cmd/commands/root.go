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

package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type rootCmdFlagsType struct {
	Chain string // Chain ID
	Id    int64  // Id
}

func (f *rootCmdFlagsType) RequireChainId() error {
	if f.Chain == "" {
		return fmt.Errorf("chain id is missing.")
	} else {
		return nil
	}
}

func (f *rootCmdFlagsType) RequireId() error {
	if f.Id == -1 {
		return fmt.Errorf("id is required.")
	} else {
		return nil
	}
}

var rootCmdFlags rootCmdFlagsType

func init() {
	rootCmd.AddCommand(chainRootCmd)
	rootCmd.AddCommand(jsonRootCmd)
	rootCmd.AddCommand(nodeRootCmd)

	jsonAddCmd.Flags().StringVarP(&rootCmdFlags.Chain, "chain", "c", "", "The ID of the chain. It may be required by some commands.")
	jsonAddCmd.Flags().Int64VarP(&rootCmdFlags.Id, "id", "i", int64(-1), "The ID of the document. It may be required by some commands.")

}

var (
	// Used for flags.
	rootCmd = &cobra.Command{
		Use:   "IL2 Go REST Client Tester",
		Short: "Test program for the IL2 Go REST Client",
		Long:  "This is a very simple test program for the IL2 Go REST Client.",
		//RunE: runOTP,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(chainListCmd)
	rootCmd.AddCommand(jsonAddCmd)
	rootCmd.AddCommand(jsonGetCmd)
	rootCmd.AddCommand(nodeVersionCmd)
}

func runOTP(cmd *cobra.Command, args []string) error {

	client, err := createAPIClient()
	version, _, err := client.NodeApi.ApiVersion(nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to query the server's version: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("The server's version is %s\n", version)

	return nil
}
