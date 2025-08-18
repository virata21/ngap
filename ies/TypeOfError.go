package ies

import "github.com/lvdund/ngap/aper"

const (
	TypeOfErrorNotunderstood aper.Enumerated = 0
	TypeOfErrorMissing       aper.Enumerated = 1
)

type TypeOfError struct {
	Value aper.Enumerated
}

func (ie *TypeOfError) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}
func (ie *TypeOfError) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
