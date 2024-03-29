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

// This package implements the root command of this application.
package cmd

import (
	"github.com/interlockledger/go-interlockledger-rest-client-tester/cmd/chain"
	"github.com/interlockledger/go-interlockledger-rest-client-tester/cmd/flags"
	"github.com/interlockledger/go-interlockledger-rest-client-tester/cmd/json"
	"github.com/interlockledger/go-interlockledger-rest-client-tester/cmd/opaque"
	"github.com/interlockledger/go-interlockledger-rest-client-tester/core"

	"github.com/interlockledger/go-interlockledger-rest-client-tester/cmd/docs"
	"github.com/interlockledger/go-interlockledger-rest-client-tester/cmd/node"
	"github.com/interlockledger/go-interlockledger-rest-client-tester/cmd/records"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	rootCmd = &cobra.Command{
		Use:     "IL2 Go REST Client Tester",
		Short:   "Test program for the IL2 Go REST Client",
		Long:    "This is a very simple test program to test and show how to use the IL2 Go REST Client.",
		Version: "0.1.0",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return core.AppCore.LoadConfig(flags.Flags.ConfigFile)
		},
	}
)

func init() {
	// Register the API commands
	rootCmd.AddCommand(chain.ChainRootCmd)
	rootCmd.AddCommand(json.JSONRootCmd)
	rootCmd.AddCommand(node.NodeRootCmd)
	rootCmd.AddCommand(records.RecordRootCmd)
	rootCmd.AddCommand(docs.DocsRootCmd)
	rootCmd.AddCommand(opaque.OpaqueRootCmd)

	// Adding the parameters shared by all commands.
	rootCmd.PersistentFlags().StringVar(&flags.Flags.ConfigFile, "config", "config.json", "The configuration file.")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
