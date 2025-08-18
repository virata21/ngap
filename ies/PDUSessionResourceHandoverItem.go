package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionResourceHandoverItem struct {
	PDUSessionID            int64  `lb:0,ub:255,madatory`
	HandoverCommandTransfer []byte `lb:0,ub:0,madatory`
	// IEExtensions *PDUSessionResourceHandoverItemExtIEs `optional`
}

func (ie *PDUSessionResourceHandoverItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PDUSessionID := NewINTEGER(ie.PDUSessionID, aper.Constraint{Lb: 0, Ub: 255}, false)
	if err = tmp_PDUSessionID.Encode(w); err != nil {
		err = utils.WrapError("Encode PDUSessionID", err)
		return
	}
	tmp_HandoverCommandTransfer := NewOCTETSTRING(ie.HandoverCommandTransfer, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_HandoverCommandTransfer.Encode(w); err != nil {
		err = utils.WrapError("Encode HandoverCommandTransfer", err)
		return
	}
	return
}
func (ie *PDUSessionResourceHandoverItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_PDUSessionID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 255},
		ext: false,
	}
	if err = tmp_PDUSessionID.Decode(r); err != nil {
		err = utils.WrapError("Read PDUSessionID", err)
		return
	}
	ie.PDUSessionID = int64(tmp_PDUSessionID.Value)
	tmp_HandoverCommandTransfer := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_HandoverCommandTransfer.Decode(r); err != nil {
		err = utils.WrapError("Read HandoverCommandTransfer", err)
		return
	}
	ie.HandoverCommandTransfer = tmp_HandoverCommandTransfer.Value
	return
}
