package ies

import "github.com/lvdund/ngap/aper"

const (
	SourceOfUEActivityBehaviourInformationSubscriptioninformation aper.Enumerated = 0
	SourceOfUEActivityBehaviourInformationStatistics              aper.Enumerated = 1
)

type SourceOfUEActivityBehaviourInformation struct {
	Value aper.Enumerated
}

func (ie *SourceOfUEActivityBehaviourInformation) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}
func (ie *SourceOfUEActivityBehaviourInformation) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
