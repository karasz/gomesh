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
package wireguard

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/tabwriter"
)

// Peers is a map containing all
// registered peers
type Peers []Peer

var dbFile string
var useStdOut bool

//Peer is a Wireguard Peer
type Peer struct {
	Name       string
	PrivateKey string
	Address    []string
	ListenPort int
	Endpoint   string
	AllowedIPs []string
	FwMark     int
	DNS        string
	MTU        int
	Table      string
	PreUp      string
	PostUp     string
	PreDown    string
	PostDown   string
	SaveConfig bool
}

// LoadPeers will load the register with Peers
// from the specified JSON file
func LoadPeers(peersPath string) (Peers, error) {

	var p Peers
	var err error
	var peersFile *os.File
	peersFile, err = os.OpenFile(peersPath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer peersFile.Close()

	jsonParser := json.NewDecoder(peersFile)
	err = jsonParser.Decode(&p)
	if err != nil && err.Error() != "EOF" {
		return nil, err
	}
	dbFile = peersPath

	return p, nil
}

// SetOutput will instruct to use standard out if called with true
// argument or files if called with false
func SetOutput(out bool) {
	useStdOut = out
}

func (p Peers) peerExists(pr Peer) bool {
	result := false

	for _, k := range p {
		if strings.EqualFold(pr.Name, k.Name) {
			result = true
		}
	}

	return result

}

//AddPeer will add a Peer to the register
func (p Peers) AddPeer(pr Peer) error {
	if pr.PrivateKey == "" {
		k, err := GenerateKey()
		if err != nil {
			return err
		}
		pr.PrivateKey = k
	}

	if pr.ListenPort == 0 {
		pr.ListenPort = 51820
	}
	p = append(p, pr)
	err := p.DumpPeers(true)
	return err
}

//DeletePeer will delete the named Peer from the register
func (p Peers) DeletePeer(pr string) {
	if len(p) == 0 {
		return
	}
	index := 0
	for i := range p {
		if p[i].Name == pr {
			index = i
		}
	}

	p = append(p[:index], p[index+1:]...)
	err := p.DumpPeers(true)
	if err != nil {
		fmt.Println(err)
	}
}

//GenerateConfigs will generate the Wireguard mesh
//configs in the specified folder
func (p Peers) GenerateConfigs(folder string, id int, peername string) error {
	var err error
	if err = os.MkdirAll(folder, 0775); err != nil {
		return err
	}
	if peername == "" {
		for i := range p {
			err = p.dumpConfig(p[i], folder, id)
		}
	} else {
		for j := range p {
			if p[j].Name == peername {
				err = p.dumpConfig(p[j], folder, id)
			}
		}
	}
	return err
}

func (p Peers) dumpConfig(pr Peer, folder string, id int) error {
	var err error
	var strToWrite string

	// write the interface section
	strToWrite = "[Interface]\n"
	strToWrite = strToWrite + fmt.Sprintf("# Name: %s\n", strings.ToLower(pr.Name))
	strToWrite = strToWrite + fmt.Sprintf("Address = %s\n", strings.Join(pr.Address, ","))
	strToWrite = strToWrite + fmt.Sprintf("PrivateKey = %s\n", pr.PrivateKey)
	strToWrite = strToWrite + fmt.Sprintf("Endpoint = %s:%d\n", pr.Endpoint, pr.ListenPort)

	// write the peers section
	for j := range p {
		if p[j].Name != pr.Name {
			strToWrite = strToWrite + "\n[Peer]\n"
			strToWrite = strToWrite + fmt.Sprintf("# Name: %s\n", strings.ToLower(p[j].Name))
			pub, err := PublicKey(p[j].PrivateKey)
			if err != nil {
				return err
			}
			strToWrite = strToWrite + fmt.Sprintf("PublicKeyKey = %s\n", pub)
			strToWrite = strToWrite + fmt.Sprintf("Endpoint = %s:%d\n", p[j].Endpoint, p[j].ListenPort)
			allIPs := append(p[j].Address, p[j].AllowedIPs...)
			strToWrite = strToWrite + fmt.Sprintf("AllowedIPs = %s\n", strings.Join(allIPs, ","))

		}

	}
	if !useStdOut {
		configFile := folder + "/" + pr.Name + "_wg" + fmt.Sprintf("%d", id) + ".conf"
		f, err := os.OpenFile(configFile, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		_, err = f.WriteString(strToWrite)
		return err
	} else {
		fmt.Println(strToWrite)
	}
	return err
}

// DumpPeers will generate a JSON file at the provided location with the
// registered Peers
func (p Peers) DumpPeers(overwrite bool) error {
	// we chose to represent that a file exists with value 2 (linux read ACL)
	// and that we want to overwrite with value 4 (linux write ACL) so that
	// we can check all the possibilities with one if statement
	var err error
	var exists int = 2
	var overwritebits = 0

	if overwrite {
		overwritebits = 4
	}

	peersJson, _ := json.MarshalIndent(p, "", "    ")

	if _, err = os.Stat(dbFile); os.IsNotExist(err) {
		exists = 0
		err = os.MkdirAll(filepath.Dir(dbFile), 0755)
		if err != nil {
			return err
		}
	}

	if overwritebits+exists != 2 {
		err = os.WriteFile(dbFile, peersJson, 0644)
		if err != nil {
			return err
		}
	} else {
		err = fmt.Errorf("Peers database exists and I am not allowed to overwrite")
	}
	return err
}

// PrettyPrint will print a table with the peers
func (p Peers) PrettyPrint(brief bool) {
	if len(p) == 0 {
		return
	}
	const padding = 3
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.AlignRight|tabwriter.Debug)

	if !brief {
		fmt.Fprintln(tw, "NAME\t", "PRIVATEKEY\t", "ADDRESS\t", "LISTENPORT\t", "ENDPOINT\t", "ALLOWEDIPS\t", "FWMARK\t", "DNS\t", "MTU\t", "TABLE\t", "PREUP\t", "POSTUP\t", "PREDOWN\t", "POSTDOWN\t", "SAVE\t")
		for _, v := range p {
			fmt.Fprintln(tw, v.Name+"\t", v.PrivateKey+"\t", strings.Join(v.Address, ",")+"\t", strconv.Itoa(v.ListenPort)+"\t", v.Endpoint+"\t", strings.Join(v.AllowedIPs, ",")+"\t", strconv.Itoa(v.FwMark)+"\t", v.DNS+"\t", strconv.Itoa(v.MTU)+"\t", v.Table+"\t", v.PreUp+"\t", v.PostUp+"\t", v.PreDown+"\t", v.PostDown+"\t", strconv.FormatBool(v.SaveConfig)+"\t")

		}
	} else {
		fmt.Fprintln(tw, "NAME\t", "PRIVATEKEY\t", "ADDRESS\t", "LISTENPORT\t", "ENDPOINT\t")
		for _, v := range p {
			fmt.Fprintln(tw, v.Name+"\t", v.PrivateKey+"\t", "["+strings.Join(v.Address, ",")+"]\t", strconv.Itoa(v.ListenPort)+"\t", v.Endpoint+"\t")
		}
	}

	tw.Flush()

}
