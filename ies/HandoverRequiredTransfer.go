package ies

import (
	"bytes"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type HandoverRequiredTransfer struct {
	DirectForwardingPathAvailability *DirectForwardingPathAvailability `optional`
	// IEExtensions *HandoverRequiredTransferExtIEs `optional`
}

func (ie *HandoverRequiredTransfer) Encode() (b []byte, err error) {
	var buf bytes.Buffer
	w := aper.NewWriter(&buf)
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.DirectForwardingPathAvailability != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.DirectForwardingPathAvailability != nil {
		if err = ie.DirectForwardingPathAvailability.Encode(w); err != nil {
			err = utils.WrapError("Encode DirectForwardingPathAvailability", err)
			return
		}
	}
	err = w.Close()
	b = buf.Bytes()
	return
}
func (ie *HandoverRequiredTransfer) Decode(wire []byte) (err error) {
	r := aper.NewReader(bytes.NewBuffer(wire))
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(DirectForwardingPathAvailability)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read DirectForwardingPathAvailability", err)
			return
		}
		ie.DirectForwardingPathAvailability = tmp
	}
	return
}
