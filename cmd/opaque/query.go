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

package opaque

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/interlockledger/go-interlockledger-rest-client-tester/cmd/flags"
	"github.com/interlockledger/go-interlockledger-rest-client-tester/core"
)

var opaqueQueryCmdFlags = struct {
	AppID          int64
	PayloadTypeIds []int64
	HowMany        int64
	OutputFile     string
}{}

// Implements GET /peers
var opaqueQueryCmd = &cobra.Command{
	Use:   "query",
	Short: "Get a list of peers kwown in the network.",
	Long: `Get a list of peers kwown in the network.
	
Calls GET /peers`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := flags.Flags.RequireChainId(); err != nil {
			return err
		}
		if opaqueQueryCmdFlags.AppID <= 0 {
			return fmt.Errorf("appId is required")
		}
		if opaqueQueryCmdFlags.OutputFile == "" {
			return fmt.Errorf("payloadFile is required")
		}
		if len(opaqueQueryCmdFlags.PayloadTypeIds) <= 0 {
			return fmt.Errorf("at least one payloadTypeIds is required")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, err := core.AppCore.NewClient()
		if err != nil {
			return err
		}
		ret, lastChangedRecordSerial, _, err := apiClient.OpaqueApi.Query(
			context.TODO(), flags.Flags.ChainId, opaqueQueryCmdFlags.AppID,
			opaqueQueryCmdFlags.PayloadTypeIds, opaqueQueryCmdFlags.HowMany,
			flags.Flags.LastToFirst, int(flags.Flags.Page), int(flags.Flags.PageSize))
		if err != nil {
			return core.FormatRequestResponseCommandError(err)
		}
		fmt.Printf("lastChangedRecordSerial: %d\n", lastChangedRecordSerial)
		fmt.Printf("Output written to: %s\n", opaqueQueryCmdFlags.OutputFile)
		err = os.WriteFile(opaqueQueryCmdFlags.OutputFile, ret, 0644)
		if err != nil {
			return core.FormatRequestResponseCommandError(err)
		}
		return nil
	},
}

func init() {
	// Adding the parameters shared by all commands.
	opaqueQueryCmd.Flags().Int64VarP(&opaqueQueryCmdFlags.AppID, "appId", "a", 0, "The application ID.")
	opaqueQueryCmd.Flags().Int64SliceVarP(&opaqueQueryCmdFlags.PayloadTypeIds, "payloadTypeId", "p", nil, "Add a payload type ID to query.")
	opaqueQueryCmd.Flags().Int64VarP(&opaqueQueryCmdFlags.HowMany, "count", "n", 0, "Number of records.")
	opaqueQueryCmd.Flags().StringVarP(&opaqueQueryCmdFlags.OutputFile, "output", "o", "", "The output file.")
	flags.Flags.RegisterPagingParams(opaqueQueryCmd.Flags())
	flags.Flags.RegisterPagingReverseParams(opaqueQueryCmd.Flags())
}