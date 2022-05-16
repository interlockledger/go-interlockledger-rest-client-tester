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
	"github.com/antihax/optional"
	"github.com/interlockledger/go-interlockledger-rest-client-tester/cmd/commands/flags"
	"github.com/spf13/pflag"
)

type RecordFlags struct {
	First   int64
	Last    int64
	HowMany int64
	Query   string
}

var recordFlags RecordFlags

func (f *RecordFlags) RegisterRecordListParams(flagSet *pflag.FlagSet) {
	flagSet.Int64Var(&f.First, "first", -1, "Serial of the first block.")
	flagSet.Int64Var(&f.Last, "last", -1, "Serial of the last block.")
}

func (f *RecordFlags) RegisterRecordHowManyParams(flagSet *pflag.FlagSet) {
	flagSet.Int64Var(&f.HowMany, "how-many", -1, "Maximum number of records to return.")
}

func (f *RecordFlags) RegisterQueryParams(flagSet *pflag.FlagSet) {
	flagSet.StringVarP(&f.Query, "query", "q", "", "InterlockQL query.")
}

func (f *RecordFlags) OptionalFirst() optional.Int64 {
	return flags.OptionalInt64(f.First)
}

func (f *RecordFlags) OptionalLast() optional.Int64 {
	return flags.OptionalInt64(f.Last)
}

func (f *RecordFlags) OptionalHowMany() optional.Int64 {
	return flags.OptionalInt64(f.Last)
}

func (f *RecordFlags) OptionalQuery() optional.String {
	return flags.OptionalString(f.Query)
}
