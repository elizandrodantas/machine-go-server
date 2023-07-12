package tool

import "github.com/elizandrodantas/machine-go-server/util"

type MachineProtocol struct {
	Valid   bool
	Key     []byte
	Iv      []byte
	Message []byte
}

func MachineProtocolEncrypted(enc []byte) *MachineProtocol {
	if len(enc) < 32 {
		return &MachineProtocol{
			Valid: false,
		}
	}

	iv := enc[:16]
	key := enc[len(enc)-16:]
	message := enc[16 : len(enc)-16]

	return &MachineProtocol{true, key, iv, message}
}

func (y *MachineProtocol) IsValidProtocol() bool {
	return y.Valid
}

func (y *MachineProtocol) GetKey() []byte {
	return y.Key
}

func (y *MachineProtocol) GetIv() []byte {
	return y.Iv
}

func (y *MachineProtocol) GetMessage() []byte {
	return y.Message
}

func (y *MachineProtocol) ToByteArray() []byte {
	var b []byte
	b = append(b, y.Iv...)
	b = append(b, y.Message...)
	b = append(b, y.Key...)

	return b
}

func (y *MachineProtocol) ToBase64() string {
	return util.Encode64(y.ToByteArray())
}
