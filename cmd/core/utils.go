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

package core

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/interlockledger/go-interlockledger-rest-client/client"
)

// Converts the object o into a formatted JSON. It is used to output JSON
// responses when necessary.
func ToPrettyJSON(o any) string {
	bin, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		panic(fmt.Sprintf("Unable to convert the object into a JSON string: %v", err))
	}
	return string(bin)
}

// Prints the given object as a JSON into stdout.
func PrintAsJSON(o any) {
	fmt.Println(ToPrettyJSON(o))
}

// Loads a JSON file from a file.
func LoadJSONFile(file string, value any) error {
	bin, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(bin, value)
}

/*
Formats the errors returned by the API calls in a human readable form that is
them wrapped into an error that can be returned to Cobra framework.

This is a quick and dirty way to implement error handling using this framework
but may not be the best way to do so.
*/
func FormatRequestResponseCommandError(err error) error {
	switch err.(type) {
	case *client.GenericSwaggerError:
		e := err.(*client.GenericSwaggerError)
		return fmt.Errorf("Request failed with the error: %w\n%s", err,
			ToPrettyJSON(e.Model()))
	case client.GenericSwaggerError:
		e := err.(client.GenericSwaggerError)
		return fmt.Errorf("Request failed with the error: %w\n%s", err,
			ToPrettyJSON(e.Model()))
	default:
		return fmt.Errorf("Request failed with the error: %w", err)
	}
}
