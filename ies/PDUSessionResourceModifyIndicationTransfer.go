package ies

import (
	"bytes"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionResourceModifyIndicationTransfer struct {
	DLQosFlowPerTNLInformation           QosFlowPerTNLInformation       `madatory`
	AdditionalDLQosFlowPerTNLInformation []QosFlowPerTNLInformationItem `lb:1,ub:maxnoofMultiConnectivityMinusOne,optional`
	// IEExtensions *PDUSessionResourceModifyIndicationTransferExtIEs `optional`
}

func (ie *PDUSessionResourceModifyIndicationTransfer) Encode() (b []byte, err error) {
	var buf bytes.Buffer
	w := aper.NewWriter(&buf)
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.AdditionalDLQosFlowPerTNLInformation != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.DLQosFlowPerTNLInformation.Encode(w); err != nil {
		err = utils.WrapError("Encode DLQosFlowPerTNLInformation", err)
		return
	}
	if len(ie.AdditionalDLQosFlowPerTNLInformation) > 0 {
		tmp := Sequence[*QosFlowPerTNLInformationItem]{
			Value: []*QosFlowPerTNLInformationItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofMultiConnectivityMinusOne},
			ext:   false,
		}
		for _, i := range ie.AdditionalDLQosFlowPerTNLInformation {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode AdditionalDLQosFlowPerTNLInformation", err)
			return
		}
	}
	err = w.Close()
	b = buf.Bytes()
	return
}
func (ie *PDUSessionResourceModifyIndicationTransfer) Decode(wire []byte) (err error) {
	r := aper.NewReader(bytes.NewBuffer(wire))
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.DLQosFlowPerTNLInformation.Decode(r); err != nil {
		err = utils.WrapError("Read DLQosFlowPerTNLInformation", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_AdditionalDLQosFlowPerTNLInformation := Sequence[*QosFlowPerTNLInformationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMultiConnectivityMinusOne},
			ext: false,
		}
		fn := func() *QosFlowPerTNLInformationItem { return new(QosFlowPerTNLInformationItem) }
		if err = tmp_AdditionalDLQosFlowPerTNLInformation.Decode(r, fn); err != nil {
			err = utils.WrapError("Read AdditionalDLQosFlowPerTNLInformation", err)
			return
		}
		ie.AdditionalDLQosFlowPerTNLInformation = []QosFlowPerTNLInformationItem{}
		for _, i := range tmp_AdditionalDLQosFlowPerTNLInformation.Value {
			ie.AdditionalDLQosFlowPerTNLInformation = append(ie.AdditionalDLQosFlowPerTNLInformation, *i)
		}
	}
	return
}
