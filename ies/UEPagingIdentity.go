package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	UEPagingIdentityPresentNothing uint64 = iota
	UEPagingIdentityPresentFivegSTmsi
	UEPagingIdentityPresentChoiceExtensions
)

type UEPagingIdentity struct {
	Choice     uint64
	FiveGSTMSI *FiveGSTMSI
	// ChoiceExtensions *UEPagingIdentityExtIEs
}

func (ie *UEPagingIdentity) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case UEPagingIdentityPresentFivegSTmsi:
		err = ie.FiveGSTMSI.Encode(w)
	}
	return
}
func (ie *UEPagingIdentity) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return
	}
	switch ie.Choice {
	case UEPagingIdentityPresentFivegSTmsi:
		var tmp FiveGSTMSI
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read FiveGSTMSI", err)
			return
		}
		ie.FiveGSTMSI = &tmp
	}
	return
}
