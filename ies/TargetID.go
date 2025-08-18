package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	TargetIDPresentNothing uint64 = iota
	TargetIDPresentTargetrannodeid
	TargetIDPresentTargetenbId
	TargetIDPresentChoiceExtensions
)

type TargetID struct {
	Choice          uint64
	TargetRANNodeID *TargetRANNodeID
	TargeteNBID     *TargeteNBID
	// ChoiceExtensions *TargetIDExtIEs
}

func (ie *TargetID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case TargetIDPresentTargetrannodeid:
		err = ie.TargetRANNodeID.Encode(w)
	case TargetIDPresentTargetenbId:
		err = ie.TargeteNBID.Encode(w)
	}
	return
}
func (ie *TargetID) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case TargetIDPresentTargetrannodeid:
		var tmp TargetRANNodeID
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read TargetRANNodeID", err)
			return
		}
		ie.TargetRANNodeID = &tmp
	case TargetIDPresentTargetenbId:
		var tmp TargeteNBID
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read TargeteNBID", err)
			return
		}
		ie.TargeteNBID = &tmp
	}
	return
}
