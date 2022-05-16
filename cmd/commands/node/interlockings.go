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

package node

import (
	"github.com/antihax/optional"
	"github.com/spf13/cobra"

	"github.com/interlockledger/go-interlockledger-rest-client-tester/cmd/commands/flags"
	"github.com/interlockledger/go-interlockledger-rest-client-tester/cmd/core"
	"github.com/interlockledger/go-interlockledger-rest-client/client"
)

var (
	nodeInterlockingsFlags = struct {
		LastKnownBlock int64
		LastToFirst    bool
		Page           int32
		PageSize       int32
	}{}
)

// testCmd represents the test command
var nodeInterlockingsCmd = &cobra.Command{
	Use:   "interlockings",
	Short: "Get the version of the API.",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := flags.Flags.RequireChainId(); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, err := core.AppCore.NewClient()

		var optionalParams client.NodeApiInterlockingsListOpts
		if nodeInterlockingsFlags.LastKnownBlock >= 0 {
			optionalParams.LastKnownBlock = optional.NewInt64(nodeInterlockingsFlags.LastKnownBlock)
		}
		optionalParams.LastToFirst = optional.NewBool(nodeInterlockingsFlags.LastToFirst)
		if nodeInterlockingsFlags.Page >= 0 {
			optionalParams.Page = optional.NewInt32(nodeInterlockingsFlags.Page)
		}
		if nodeInterlockingsFlags.PageSize >= 0 {
			optionalParams.PageSize = optional.NewInt32(nodeInterlockingsFlags.PageSize)
		}
		ret, _, err := apiClient.NodeApi.InterlockingsList(nil, flags.Flags.ChainId, &optionalParams)
		if err != nil {
			return core.FormatRequestResponseCommandError(err)
		}
		core.PrintAsJSON(ret)
		return nil
	},
}

func init() {
	nodeInterlockingsCmd.Flags().Int64Var(&nodeInterlockingsFlags.LastKnownBlock, "last-known-block", -1, "Last known block.")
	nodeInterlockingsCmd.Flags().BoolVar(&nodeInterlockingsFlags.LastToFirst, "last-to-first", false, "Last to first order.")
	nodeInterlockingsCmd.Flags().Int32Var(&nodeInterlockingsFlags.Page, "page", -1, "Page.")
	nodeInterlockingsCmd.Flags().Int32Var(&nodeInterlockingsFlags.PageSize, "page-size", -1, "Page size.")
	flags.Flags.RegisterChainIdParameter(nodeInterlockingsCmd.Flags())
}
