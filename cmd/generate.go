/*
Copyright © 2021 JPI Technologies Ltd <oss@jpi.io>

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

	"github.com/karasz/gomesh/wireguard"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate configs",
	Long:  `Generate will create the configs file in the specified folder`,
	Run: func(cmd *cobra.Command, args []string) {
		out, _ := cmd.Flags().GetString("output")
		id, _ := cmd.Flags().GetInt("network_id")
		peername, _ := cmd.Flags().GetString("peer_name")
		usestdout, _ := cmd.Flags().GetBool("useStdOut")
		wireguard.SetOutput(usestdout)
		err := thePeers.GenerateConfigs(out, id, peername)
		if err != nil {
			fmt.Println("generate", err)
		}
	},
}

func init() {
	generateCmd.Flags().StringP("output", "o", "output", "Directory where to output configs.")
	generateCmd.Flags().IntP("network_id", "i", 0, "ID of the network to generate")
	generateCmd.Flags().StringP("peer_name", "p", "", "Generate config for this peer")
	generateCmd.Flags().BoolP("useStdOut", "s", false, "Use StdOut instead of files")
	rootCmd.AddCommand(generateCmd)
}
