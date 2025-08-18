package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SONInformationReply struct {
	XnTNLConfigurationInfo *XnTNLConfigurationInfo `optional`
	// IEExtensions *SONInformationReplyExtIEs `optional`
}

func (ie *SONInformationReply) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.XnTNLConfigurationInfo != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.XnTNLConfigurationInfo != nil {
		if err = ie.XnTNLConfigurationInfo.Encode(w); err != nil {
			err = utils.WrapError("Encode XnTNLConfigurationInfo", err)
			return
		}
	}
	return
}
func (ie *SONInformationReply) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(XnTNLConfigurationInfo)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read XnTNLConfigurationInfo", err)
			return
		}
		ie.XnTNLConfigurationInfo = tmp
	}
	return
}
