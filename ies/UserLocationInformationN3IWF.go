package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UserLocationInformationN3IWF struct {
	IPAddress  aper.BitString `lb:1,ub:160,madatory,valExt`
	PortNumber []byte         `lb:2,ub:2,madatory`
	// IEExtensions *UserLocationInformationN3IWFExtIEs `optional`
}

func (ie *UserLocationInformationN3IWF) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_IPAddress := NewBITSTRING(ie.IPAddress, aper.Constraint{Lb: 1, Ub: 160}, true)
	if err = tmp_IPAddress.Encode(w); err != nil {
		err = utils.WrapError("Encode IPAddress", err)
		return
	}
	tmp_PortNumber := NewOCTETSTRING(ie.PortNumber, aper.Constraint{Lb: 2, Ub: 2}, false)
	if err = tmp_PortNumber.Encode(w); err != nil {
		err = utils.WrapError("Encode PortNumber", err)
		return
	}
	return
}
func (ie *UserLocationInformationN3IWF) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_IPAddress := BITSTRING{
		c:   aper.Constraint{Lb: 1, Ub: 160},
		ext: true,
	}
	if err = tmp_IPAddress.Decode(r); err != nil {
		err = utils.WrapError("Read IPAddress", err)
		return
	}
	ie.IPAddress = aper.BitString{Bytes: tmp_IPAddress.Value.Bytes, NumBits: tmp_IPAddress.Value.NumBits}
	tmp_PortNumber := OCTETSTRING{
		c:   aper.Constraint{Lb: 2, Ub: 2},
		ext: false,
	}
	if err = tmp_PortNumber.Decode(r); err != nil {
		err = utils.WrapError("Read PortNumber", err)
		return
	}
	ie.PortNumber = tmp_PortNumber.Value
	return
}
