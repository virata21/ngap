package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type COUNTValueForPDCPSN12 struct {
	PDCPSN12    int64 `lb:0,ub:4095,madatory`
	HFNPDCPSN12 int64 `lb:0,ub:1048575,madatory`
	// IEExtensions *COUNTValueForPDCPSN12ExtIEs `optional`
}

func (ie *COUNTValueForPDCPSN12) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PDCPSN12 := NewINTEGER(ie.PDCPSN12, aper.Constraint{Lb: 0, Ub: 4095}, false)
	if err = tmp_PDCPSN12.Encode(w); err != nil {
		err = utils.WrapError("Encode PDCPSN12", err)
		return
	}
	tmp_HFNPDCPSN12 := NewINTEGER(ie.HFNPDCPSN12, aper.Constraint{Lb: 0, Ub: 1048575}, false)
	if err = tmp_HFNPDCPSN12.Encode(w); err != nil {
		err = utils.WrapError("Encode HFNPDCPSN12", err)
		return
	}
	return
}
func (ie *COUNTValueForPDCPSN12) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_PDCPSN12 := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4095},
		ext: false,
	}
	if err = tmp_PDCPSN12.Decode(r); err != nil {
		err = utils.WrapError("Read PDCPSN12", err)
		return
	}
	ie.PDCPSN12 = int64(tmp_PDCPSN12.Value)
	tmp_HFNPDCPSN12 := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 1048575},
		ext: false,
	}
	if err = tmp_HFNPDCPSN12.Decode(r); err != nil {
		err = utils.WrapError("Read HFNPDCPSN12", err)
		return
	}
	ie.HFNPDCPSN12 = int64(tmp_HFNPDCPSN12.Value)
	return
}
