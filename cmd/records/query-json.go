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

package records

import (
	"github.com/spf13/cobra"

	"github.com/interlockledger/go-interlockledger-rest-client-tester/cmd/flags"
	"github.com/interlockledger/go-interlockledger-rest-client-tester/core"
	"github.com/interlockledger/go-interlockledger-rest-client/client"
)

// Implements GET /records@{chain}/asJson/query
var recordQueryJSONCmd = &cobra.Command{
	Use:   "query-json",
	Short: "Query using InterlockQL and returns the result as JSON.",
	Long: `Query using InterlockQL and returns the result as JSON..

Calls GET /records@{chain}/asJson/query`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := flags.Flags.RequireChainId(); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, err := core.AppCore.NewClient()
		if err != nil {
			return err
		}

		var options client.RecordApiRecordsQueryAsJsonOpts
		options.QueryAsInterlockQL = recordFlags.OptionalQuery()
		options.HowMany = recordFlags.OptionalHowMany()
		options.Page = flags.Flags.OptionalPage()
		options.PageSize = flags.Flags.OptionalPageSize()
		options.LastToFirst = flags.Flags.OptionalLastToFirst()
		ret, _, err := apiClient.RecordApi.RecordsQueryAsJson(nil, flags.Flags.ChainId, &options)
		if err != nil {
			return core.FormatRequestResponseCommandError(err)
		}
		core.PrintAsJSON(ret)
		return nil
	},
}

func init() {
	recordFlags.RegisterRecordListParams(recordQueryJSONCmd.Flags())
	recordFlags.RegisterQueryParams(recordQueryJSONCmd.Flags())
	recordFlags.RegisterRecordHowManyParams(recordQueryJSONCmd.Flags())
	flags.Flags.RegisterPagingParams(recordQueryJSONCmd.Flags())
	flags.Flags.RegisterPagingReverseParams(recordQueryJSONCmd.Flags())
}
