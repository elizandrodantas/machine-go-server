package tool

import (
	"bytes"
	"testing"
)

func TestProtocol(t *testing.T) {
	iv := RandomByte(16)
	key := RandomByte(16)
	message := "hello test"

	concated := iv
	concated = append(concated, []byte(message)...)
	concated = append(concated, key...)

	protocol := MachineProtocolEncrypted(concated)

	if !bytes.Equal(key, protocol.GetKey()) {
		t.Error("invalid protocol, invalid key")
	}

	if !bytes.Equal(iv, protocol.GetIv()) {
		t.Error("invalid protocol, invalid iv")
	}

	if message != string(protocol.GetMessage()) {
		t.Errorf("invalid message, expected '%s' and received '%s'", message, string(protocol.GetMessage()))
	}
}
