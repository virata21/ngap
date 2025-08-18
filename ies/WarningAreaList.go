package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	WarningAreaListPresentNothing uint64 = iota
	WarningAreaListPresentEutraCgilistforwarning
	WarningAreaListPresentNrCgilistforwarning
	WarningAreaListPresentTailistforwarning
	WarningAreaListPresentEmergencyareaidlist
	WarningAreaListPresentChoiceExtensions
)

type WarningAreaList struct {
	Choice                 uint64
	EUTRACGIListForWarning []EUTRACGI
	NRCGIListForWarning    []NRCGI
	TAIListForWarning      []TAI
	EmergencyAreaIDList    []EmergencyAreaID
	// ChoiceExtensions *WarningAreaListExtIEs
}

func (ie *WarningAreaList) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return
	}
	switch ie.Choice {
	case WarningAreaListPresentEutraCgilistforwarning:
		tmp := Sequence[*EUTRACGI]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellIDforWarning},
			ext: false,
		}
		for _, i := range ie.EUTRACGIListForWarning {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case WarningAreaListPresentNrCgilistforwarning:
		tmp := Sequence[*NRCGI]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellIDforWarning},
			ext: false,
		}
		for _, i := range ie.NRCGIListForWarning {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case WarningAreaListPresentTailistforwarning:
		tmp := Sequence[*TAI]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforWarning},
			ext: false,
		}
		for _, i := range ie.TAIListForWarning {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case WarningAreaListPresentEmergencyareaidlist:
		tmp := Sequence[*EmergencyAreaID]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofEmergencyAreaID},
			ext: false,
		}
		for _, i := range ie.EmergencyAreaIDList {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	}
	return
}
func (ie *WarningAreaList) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return
	}
	switch ie.Choice {
	case WarningAreaListPresentEutraCgilistforwarning:
		tmp := NewSequence[*EUTRACGI](nil, aper.Constraint{Lb: 1, Ub: maxnoofCellIDforWarning}, false)
		fn := func() *EUTRACGI {
			return new(EUTRACGI)
		}
		if err = tmp.Decode(r, fn); err != nil {
			err = utils.WrapError("Read EUTRACGIListForWarning", err)
			return
		}
		for _, i := range tmp.Value {
			ie.EUTRACGIListForWarning = append(ie.EUTRACGIListForWarning, *i)
		}
	case WarningAreaListPresentNrCgilistforwarning:
		tmp := NewSequence[*NRCGI](nil, aper.Constraint{Lb: 1, Ub: maxnoofCellIDforWarning}, false)
		fn := func() *NRCGI {
			return new(NRCGI)
		}
		if err = tmp.Decode(r, fn); err != nil {
			err = utils.WrapError("Read NRCGIListForWarning", err)
			return
		}
		for _, i := range tmp.Value {
			ie.NRCGIListForWarning = append(ie.NRCGIListForWarning, *i)
		}
	case WarningAreaListPresentTailistforwarning:
		tmp := NewSequence[*TAI](nil, aper.Constraint{Lb: 1, Ub: maxnoofTAIforWarning}, false)
		fn := func() *TAI {
			return new(TAI)
		}
		if err = tmp.Decode(r, fn); err != nil {
			err = utils.WrapError("Read TAIListForWarning", err)
			return
		}
		for _, i := range tmp.Value {
			ie.TAIListForWarning = append(ie.TAIListForWarning, *i)
		}
	case WarningAreaListPresentEmergencyareaidlist:
		tmp := NewSequence[*EmergencyAreaID](nil, aper.Constraint{Lb: 1, Ub: maxnoofEmergencyAreaID}, false)
		fn := func() *EmergencyAreaID {
			return new(EmergencyAreaID)
		}
		if err = tmp.Decode(r, fn); err != nil {
			err = utils.WrapError("Read EmergencyAreaIDList", err)
			return
		}
		for _, i := range tmp.Value {
			ie.EmergencyAreaIDList = append(ie.EmergencyAreaIDList, *i)
		}
	}
	return
}
