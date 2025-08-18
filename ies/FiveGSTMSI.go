package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type FiveGSTMSI struct {
	AMFSetID   aper.BitString `lb:10,ub:10,madatory`
	AMFPointer aper.BitString `lb:6,ub:6,madatory`
	FiveGTMSI  []byte         `lb:4,ub:4,madatory`
	// IEExtensions *FiveGSTMSIExtIEs `optional`
}

func (ie *FiveGSTMSI) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_AMFSetID := NewBITSTRING(ie.AMFSetID, aper.Constraint{Lb: 10, Ub: 10}, false)
	if err = tmp_AMFSetID.Encode(w); err != nil {
		err = utils.WrapError("Encode AMFSetID", err)
		return
	}
	tmp_AMFPointer := NewBITSTRING(ie.AMFPointer, aper.Constraint{Lb: 6, Ub: 6}, false)
	if err = tmp_AMFPointer.Encode(w); err != nil {
		err = utils.WrapError("Encode AMFPointer", err)
		return
	}
	tmp_FiveGTMSI := NewOCTETSTRING(ie.FiveGTMSI, aper.Constraint{Lb: 4, Ub: 4}, false)
	if err = tmp_FiveGTMSI.Encode(w); err != nil {
		err = utils.WrapError("Encode FiveGTMSI", err)
		return
	}
	return
}
func (ie *FiveGSTMSI) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_AMFSetID := BITSTRING{
		c:   aper.Constraint{Lb: 10, Ub: 10},
		ext: false,
	}
	if err = tmp_AMFSetID.Decode(r); err != nil {
		err = utils.WrapError("Read AMFSetID", err)
		return
	}
	ie.AMFSetID = aper.BitString{Bytes: tmp_AMFSetID.Value.Bytes, NumBits: tmp_AMFSetID.Value.NumBits}
	tmp_AMFPointer := BITSTRING{
		c:   aper.Constraint{Lb: 6, Ub: 6},
		ext: false,
	}
	if err = tmp_AMFPointer.Decode(r); err != nil {
		err = utils.WrapError("Read AMFPointer", err)
		return
	}
	ie.AMFPointer = aper.BitString{Bytes: tmp_AMFPointer.Value.Bytes, NumBits: tmp_AMFPointer.Value.NumBits}
	tmp_FiveGTMSI := OCTETSTRING{
		c:   aper.Constraint{Lb: 4, Ub: 4},
		ext: false,
	}
	if err = tmp_FiveGTMSI.Decode(r); err != nil {
		err = utils.WrapError("Read FiveGTMSI", err)
		return
	}
	ie.FiveGTMSI = tmp_FiveGTMSI.Value
	return
}
