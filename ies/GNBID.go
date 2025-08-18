package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	GNBIDPresentNothing uint64 = iota
	GNBIDPresentGnbId
	GNBIDPresentChoiceExtensions
)

type GNBID struct {
	Choice uint64
	GNBID  *aper.BitString `lb:22,ub:32`
	// ChoiceExtensions *GNBIDExtIEs
}

func (ie *GNBID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case GNBIDPresentGnbId:
		tmp := NewBITSTRING(*ie.GNBID, aper.Constraint{Lb: 22, Ub: 32}, false)
		err = tmp.Encode(w)
	}
	return
}
func (ie *GNBID) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return
	}
	switch ie.Choice {
	case GNBIDPresentGnbId:
		tmp := BITSTRING{c: aper.Constraint{Lb: 22, Ub: 32}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read GNBID", err)
			return
		}
		ie.GNBID = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	}
	return
}
