package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type CompletedCellsInTAIEUTRAItem struct {
	EUTRACGI EUTRACGI `madatory`
	// IEExtensions *CompletedCellsInTAIEUTRAItemExtIEs `optional`
}

func (ie *CompletedCellsInTAIEUTRAItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.EUTRACGI.Encode(w); err != nil {
		err = utils.WrapError("Encode EUTRACGI", err)
		return
	}
	return
}
func (ie *CompletedCellsInTAIEUTRAItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.EUTRACGI.Decode(r); err != nil {
		err = utils.WrapError("Read EUTRACGI", err)
		return
	}
	return
}
