/*
Copyright © 2023-present, Meta Platforms, Inc. and affiliates
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	listTTPsCmd = &cobra.Command{
		Use:              "ttps",
		Short:            "list TTPForge repos (in which TTPs live) that you have installed",
		TraverseChildren: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ttpRefs, err := Conf.repoCollection.ListTTPs()
			if err != nil {
				return err
			}
			for _, ttpRef := range ttpRefs {
				fmt.Println(ttpRef)
			}
			return nil
		},
	}
)

func init() {
	listCmd.AddCommand(listTTPsCmd)
}
