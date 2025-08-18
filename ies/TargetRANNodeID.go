package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TargetRANNodeID struct {
	GlobalRANNodeID GlobalRANNodeID `madatory`
	SelectedTAI     TAI             `madatory`
	// IEExtensions *TargetRANNodeIDExtIEs `optional`
}

func (ie *TargetRANNodeID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.GlobalRANNodeID.Encode(w); err != nil {
		err = utils.WrapError("Encode GlobalRANNodeID", err)
		return
	}
	if err = ie.SelectedTAI.Encode(w); err != nil {
		err = utils.WrapError("Encode SelectedTAI", err)
		return
	}
	return
}
func (ie *TargetRANNodeID) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.GlobalRANNodeID.Decode(r); err != nil {
		err = utils.WrapError("Read GlobalRANNodeID", err)
		return
	}
	if err = ie.SelectedTAI.Decode(r); err != nil {
		err = utils.WrapError("Read SelectedTAI", err)
		return
	}
	return
}
