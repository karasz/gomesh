/*
Copyright © 2021 Nagy Károly Gábriel <k@jpi.io>

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
	"github.com/karasz/gomesh/peers"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add or update a peer to/in the registry",
	Long:  `Wil add a peer to the registry or update`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		privatekey, _ := cmd.Flags().GetString("privatekey")
		address, _ := cmd.Flags().GetStringSlice("address")
		listenport, _ := cmd.Flags().GetInt("listenport")
		endpoint, _ := cmd.Flags().GetString("endpoint")
		allowedips, _ := cmd.Flags().GetStringSlice("allowedips")
		fwmark, _ := cmd.Flags().GetInt("fwmark")
		dns, _ := cmd.Flags().GetString("dns")
		mtu, _ := cmd.Flags().GetInt("mtu")
		table, _ := cmd.Flags().GetString("table")
		preup, _ := cmd.Flags().GetString("preup")
		predown, _ := cmd.Flags().GetString("predown")
		postup, _ := cmd.Flags().GetString("postup")
		postdown, _ := cmd.Flags().GetString("postdown")
		saveconfig, _ := cmd.Flags().GetBool("saveconfig")

		p := peers.Peer{
			Name:       name,
			PrivateKey: privatekey,
			Address:    address,
			ListenPort: listenport,
			Endpoint:   endpoint,
			AllowedIPs: allowedips,
			FwMark:     fwmark,
			DNS:        dns,
			MTU:        mtu,
			Table:      table,
			PreUp:      preup,
			PostUp:     postup,
			PreDown:    predown,
			PostDown:   postdown,
			SaveConfig: saveconfig,
		}
		thePeers.AddPeer(p)
	},
}

func init() {
	addCmd.Flags().BoolP("update", "u", false, "Update Peer if existing.")
	addCmd.Flags().StringP("name", "n", "", "endpoint. (Required)")
	addCmd.Flags().StringSliceP("address", "a", []string{}, "Address of the server. (Required)")
	addCmd.Flags().StringP("endpoint", "e", "", "The peer's endpoint")
	addCmd.Flags().StringSliceP("allowedips", "", []string{}, "additional allowed IP addresses")
	addCmd.Flags().StringP("privatekey", "p", "", "private key of server interface (if none given one will be generated")
	addCmd.Flags().IntP("listenport", "l", 51820, "Port to listen on, default 51820")
	addCmd.Flags().IntP("fwmark", "f", 0, "Mark the outgoing packets with")
	addCmd.Flags().StringP("dns", "d", "", "DNS server")
	addCmd.Flags().IntP("mtu", "m", 0, "Server interface MTU")
	addCmd.Flags().StringP("routing_table", "r", "", "Server routing table")
	addCmd.Flags().StringP("preUP", "", "", "Command to run before bringing the interface UP")
	addCmd.Flags().StringP("postUP", "", "", "Command to run after bringing the interface UP")
	addCmd.Flags().StringP("preDown", "", "", "Command to run before bringing the interface DOWN")
	addCmd.Flags().StringP("postDown", "", "", "Command to run after bringing the interface DOWN")
	addCmd.Flags().BoolP("saveconfig", "s", false, "Save config between reboots")
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("address")
	addCmd.MarkFlagRequired("endpoint")

	rootCmd.AddCommand(addCmd)
}
