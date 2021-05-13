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
package keys

import (
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// GenerateKey will return a base64 encoded string
// suitable for using as a Wireguyard PrivateKey
func GenerateKey() (string, error) {
	privatekey, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		return "", err
	}
	privkeystring := privatekey.String()
	return privkeystring, nil

}

//PublicKey will return a base64 encoded
// public key from a base64 encoded PrivateKey
func PublicKey(a string) (string, error) {
	k, err := wgtypes.ParseKey(a)
	if err != nil {
		return "", err
	}
	return k.PublicKey().String(), nil
}
