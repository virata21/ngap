package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	UPTransportLayerInformationPresentNothing uint64 = iota
	UPTransportLayerInformationPresentGtptunnel
	UPTransportLayerInformationPresentChoiceExtensions
)

type UPTransportLayerInformation struct {
	Choice    uint64
	GTPTunnel *GTPTunnel
	// ChoiceExtensions *UPTransportLayerInformationExtIEs
}

func (ie *UPTransportLayerInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case UPTransportLayerInformationPresentGtptunnel:
		err = ie.GTPTunnel.Encode(w)
	}
	return
}
func (ie *UPTransportLayerInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return
	}
	switch ie.Choice {
	case UPTransportLayerInformationPresentGtptunnel:
		var tmp GTPTunnel
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read GTPTunnel", err)
			return
		}
		ie.GTPTunnel = &tmp
	}
	return
}
