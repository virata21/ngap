package ies

import "github.com/lvdund/ngap/aper"

const (
	NewSecurityContextIndTrue aper.Enumerated = 0
)

type NewSecurityContextInd struct {
	Value aper.Enumerated
}

func (ie *NewSecurityContextInd) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, true)
	return
}
func (ie *NewSecurityContextInd) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true)
	ie.Value = aper.Enumerated(v)
	return
}
