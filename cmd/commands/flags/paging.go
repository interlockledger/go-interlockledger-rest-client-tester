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

package flags

import (
	"github.com/antihax/optional"
	"github.com/spf13/pflag"
)

// This struct holds the global flags used by all or most commands.
type PagingFlags struct {
	// Paging
	Page        int32
	PageSize    int32
	LastToFirst bool
}

func (f *PagingFlags) RegisterPagingParams(flagSet *pflag.FlagSet) {
	flagSet.Int32Var(&f.Page, "page", -1, "The page.")
	flagSet.Int32Var(&f.PageSize, "page-size", -1, "The page size.")
}

func (f *PagingFlags) RegisterPagingReverseParams(flagSet *pflag.FlagSet) {
	flagSet.BoolVar(&f.LastToFirst, "last-to-first", false, "Invert the list order.")
}

func (f *PagingFlags) OptionalPage() optional.Int32 {
	if f.Page == -1 {
		return optional.EmptyInt32()
	} else {
		return optional.NewInt32(f.Page)
	}
}

func (f *PagingFlags) OptionalPageSize() optional.Int32 {
	if f.PageSize == -1 {
		return optional.EmptyInt32()
	} else {
		return optional.NewInt32(f.PageSize)
	}
}

func (f *PagingFlags) OptionalLastToFirst() optional.Bool {
	return optional.NewBool(f.LastToFirst)
}
