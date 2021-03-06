//code generated by CXDatum. DO NOT EDIT.

package model

import (
	"github.com/SkycoinProject/skycoin/src/cipher"
)

var donereq chan bool
var doneres chan bool
var req chan []byte
var res chan []byte

func init() {
	donereq = make(chan bool)
	doneres = make(chan bool)
	req = make(chan []byte, 1)
	res = make(chan []byte, 1)
}

func waitForPrgmReturn() []byte {
	<-doneres
	return <-res
}

func passToPrgm(byts []byte) {
	req <- byts
	donereq <- true
}

const CXDATUM_EGRESS_CODE_getTweet = 0

func getTweet(txid cipher.SHA256) (txt Tweet) {
	_cxdatum_tmp_136 := getTweet_arg_EgressExpect{}
	_cxdatum_tmp_136.txid = txid
	_cxdatum_tmp_137 := []byte{CXDATUM_EGRESS_CODE_getTweet}
	_cxdatum_tmp_137 = append(_cxdatum_tmp_137, serializegetTweet_arg_EgressExpect(_cxdatum_tmp_136)...)
	passToPrgm(_cxdatum_tmp_137)
	_cxdatum_tmp_137 = waitForPrgmReturn()
	_cxdatum_tmp_138 := deserializegetTweet_arg_EgressReturn(_cxdatum_tmp_137)
	txt = _cxdatum_tmp_138.txt
	return
}
