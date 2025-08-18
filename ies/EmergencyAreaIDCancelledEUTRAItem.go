package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type EmergencyAreaIDCancelledEUTRAItem struct {
	EmergencyAreaID          []byte                         `lb:3,ub:3,madatory`
	CancelledCellsInEAIEUTRA []CancelledCellsInEAIEUTRAItem `lb:1,ub:maxnoofCellinEAI,madatory`
	// IEExtensions *EmergencyAreaIDCancelledEUTRAItemExtIEs `optional`
}

func (ie *EmergencyAreaIDCancelledEUTRAItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_EmergencyAreaID := NewOCTETSTRING(ie.EmergencyAreaID, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_EmergencyAreaID.Encode(w); err != nil {
		err = utils.WrapError("Encode EmergencyAreaID", err)
		return
	}
	if len(ie.CancelledCellsInEAIEUTRA) > 0 {
		tmp := Sequence[*CancelledCellsInEAIEUTRAItem]{
			Value: []*CancelledCellsInEAIEUTRAItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofCellinEAI},
			ext:   false,
		}
		for _, i := range ie.CancelledCellsInEAIEUTRA {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode CancelledCellsInEAIEUTRA", err)
			return
		}
	} else {
		err = utils.WrapError("CancelledCellsInEAIEUTRA is nil", err)
		return
	}
	return
}
func (ie *EmergencyAreaIDCancelledEUTRAItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_EmergencyAreaID := OCTETSTRING{
		c:   aper.Constraint{Lb: 3, Ub: 3},
		ext: false,
	}
	if err = tmp_EmergencyAreaID.Decode(r); err != nil {
		err = utils.WrapError("Read EmergencyAreaID", err)
		return
	}
	ie.EmergencyAreaID = tmp_EmergencyAreaID.Value
	tmp_CancelledCellsInEAIEUTRA := Sequence[*CancelledCellsInEAIEUTRAItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofCellinEAI},
		ext: false,
	}
	fn := func() *CancelledCellsInEAIEUTRAItem { return new(CancelledCellsInEAIEUTRAItem) }
	if err = tmp_CancelledCellsInEAIEUTRA.Decode(r, fn); err != nil {
		err = utils.WrapError("Read CancelledCellsInEAIEUTRA", err)
		return
	}
	ie.CancelledCellsInEAIEUTRA = []CancelledCellsInEAIEUTRAItem{}
	for _, i := range tmp_CancelledCellsInEAIEUTRA.Value {
		ie.CancelledCellsInEAIEUTRA = append(ie.CancelledCellsInEAIEUTRA, *i)
	}
	return
}
