package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type QosFlowPerTNLInformation struct {
	UPTransportLayerInformation UPTransportLayerInformation `madatory`
	AssociatedQosFlowList       []AssociatedQosFlowItem     `lb:1,ub:maxnoofQosFlows,madatory`
	// IEExtensions *QosFlowPerTNLInformationExtIEs `optional`
}

func (ie *QosFlowPerTNLInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.UPTransportLayerInformation.Encode(w); err != nil {
		err = utils.WrapError("Encode UPTransportLayerInformation", err)
		return
	}
	if len(ie.AssociatedQosFlowList) > 0 {
		tmp := Sequence[*AssociatedQosFlowItem]{
			Value: []*AssociatedQosFlowItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext:   false,
		}
		for _, i := range ie.AssociatedQosFlowList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode AssociatedQosFlowList", err)
			return
		}
	} else {
		err = utils.WrapError("AssociatedQosFlowList is nil", err)
		return
	}
	return
}
func (ie *QosFlowPerTNLInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.UPTransportLayerInformation.Decode(r); err != nil {
		err = utils.WrapError("Read UPTransportLayerInformation", err)
		return
	}
	tmp_AssociatedQosFlowList := Sequence[*AssociatedQosFlowItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
		ext: false,
	}
	fn := func() *AssociatedQosFlowItem { return new(AssociatedQosFlowItem) }
	if err = tmp_AssociatedQosFlowList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read AssociatedQosFlowList", err)
		return
	}
	ie.AssociatedQosFlowList = []AssociatedQosFlowItem{}
	for _, i := range tmp_AssociatedQosFlowList.Value {
		ie.AssociatedQosFlowList = append(ie.AssociatedQosFlowList, *i)
	}
	return
}
