package ies

import "github.com/lvdund/ngap/aper"

const (
	EmergencyServiceTargetCNFivegc aper.Enumerated = 0
	EmergencyServiceTargetCNEpc    aper.Enumerated = 1
)

type EmergencyServiceTargetCN struct {
	Value aper.Enumerated
}

func (ie *EmergencyServiceTargetCN) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}
func (ie *EmergencyServiceTargetCN) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
