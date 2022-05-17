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

package docs

import (
	"fmt"

	"github.com/spf13/pflag"
)

var docsFlags DocsFlags

type DocsFlags struct {
	TransactionId string
	DocumentFile  string
	ContentType   string
	Path          string
	Comment       string
	Locator       string
	Index         int32
	OutputFile    string
}

func (f *DocsFlags) RequireTransactionId() error {
	if f.TransactionId == "" {
		return fmt.Errorf("The transaction id is missing.")
	} else {
		return nil
	}
}

func (f *DocsFlags) RequireDocument() error {
	if f.DocumentFile == "" {
		return fmt.Errorf("The document file is missing.")
	}
	if f.ContentType == "" {
		return fmt.Errorf("The content type is missing.")
	}
	return nil
}

func (f *DocsFlags) RequireLocator() error {
	if f.Locator == "" {
		return fmt.Errorf("The document locator is missing.")
	}
	return nil
}

func (f *DocsFlags) RequireOutput() error {
	if f.OutputFile == "" {
		return fmt.Errorf("The output file is missing.")
	}
	return nil
}

func (f *DocsFlags) RequireLocatorAndIndex() error {
	if f.Index == -1 {
		return fmt.Errorf("The document index is missing.")
	}
	return nil
}

func (f *DocsFlags) RegisterTransactionIDParameter(flagSet *pflag.FlagSet) {
	flagSet.StringVarP(&docsFlags.TransactionId, "transaction-id", "t", "", "Transaction ID.")
}

func (f *DocsFlags) RegisterDocumentParameter(flagSet *pflag.FlagSet) {
	flagSet.StringVarP(&docsFlags.DocumentFile, "document", "d", "", "The document file.")
	flagSet.StringVarP(&docsFlags.ContentType, "content-type", "m", "", "Content-type of the document.")
	flagSet.StringVar(&docsFlags.Path, "path", "/", "Path of the document.")
	flagSet.StringVar(&docsFlags.Comment, "comment", "-", "Comment of the document.")
}

func (f *DocsFlags) RegisterLocatorParameter(flagSet *pflag.FlagSet) {
	flagSet.StringVarP(&docsFlags.Locator, "locator", "l", "", "The document locator.")
}

func (f *DocsFlags) RegisterIndexParameter(flagSet *pflag.FlagSet) {
	flagSet.Int32VarP(&docsFlags.Index, "index", "i", -1, "Index of the document.")
}

func (f *DocsFlags) RegisterZipFileParameter(flagSet *pflag.FlagSet) {
	flagSet.StringVarP(&docsFlags.OutputFile, "output", "o", "", "The output file.")
}
