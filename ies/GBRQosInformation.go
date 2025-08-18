package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GBRQosInformation struct {
	MaximumFlowBitRateDL    int64                `lb:0,ub:4000000000000,madatory,valExt`
	MaximumFlowBitRateUL    int64                `lb:0,ub:4000000000000,madatory,valExt`
	GuaranteedFlowBitRateDL int64                `lb:0,ub:4000000000000,madatory,valExt`
	GuaranteedFlowBitRateUL int64                `lb:0,ub:4000000000000,madatory,valExt`
	NotificationControl     *NotificationControl `optional`
	MaximumPacketLossRateDL *int64               `lb:0,ub:1000,optional,valExt`
	MaximumPacketLossRateUL *int64               `lb:0,ub:1000,optional,valExt`
	// IEExtensions *GBRQosInformationExtIEs `optional`
}

func (ie *GBRQosInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.NotificationControl != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.MaximumPacketLossRateDL != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.MaximumPacketLossRateUL != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)
	tmp_MaximumFlowBitRateDL := NewINTEGER(ie.MaximumFlowBitRateDL, aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
	if err = tmp_MaximumFlowBitRateDL.Encode(w); err != nil {
		err = utils.WrapError("Encode MaximumFlowBitRateDL", err)
		return
	}
	tmp_MaximumFlowBitRateUL := NewINTEGER(ie.MaximumFlowBitRateUL, aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
	if err = tmp_MaximumFlowBitRateUL.Encode(w); err != nil {
		err = utils.WrapError("Encode MaximumFlowBitRateUL", err)
		return
	}
	tmp_GuaranteedFlowBitRateDL := NewINTEGER(ie.GuaranteedFlowBitRateDL, aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
	if err = tmp_GuaranteedFlowBitRateDL.Encode(w); err != nil {
		err = utils.WrapError("Encode GuaranteedFlowBitRateDL", err)
		return
	}
	tmp_GuaranteedFlowBitRateUL := NewINTEGER(ie.GuaranteedFlowBitRateUL, aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
	if err = tmp_GuaranteedFlowBitRateUL.Encode(w); err != nil {
		err = utils.WrapError("Encode GuaranteedFlowBitRateUL", err)
		return
	}
	if ie.NotificationControl != nil {
		if err = ie.NotificationControl.Encode(w); err != nil {
			err = utils.WrapError("Encode NotificationControl", err)
			return
		}
	}
	if ie.MaximumPacketLossRateDL != nil {
		tmp_MaximumPacketLossRateDL := NewINTEGER(*ie.MaximumPacketLossRateDL, aper.Constraint{Lb: 0, Ub: 1000}, true)
		if err = tmp_MaximumPacketLossRateDL.Encode(w); err != nil {
			err = utils.WrapError("Encode MaximumPacketLossRateDL", err)
			return
		}
	}
	if ie.MaximumPacketLossRateUL != nil {
		tmp_MaximumPacketLossRateUL := NewINTEGER(*ie.MaximumPacketLossRateUL, aper.Constraint{Lb: 0, Ub: 1000}, true)
		if err = tmp_MaximumPacketLossRateUL.Encode(w); err != nil {
			err = utils.WrapError("Encode MaximumPacketLossRateUL", err)
			return
		}
	}
	return
}
func (ie *GBRQosInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	tmp_MaximumFlowBitRateDL := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
		ext: true,
	}
	if err = tmp_MaximumFlowBitRateDL.Decode(r); err != nil {
		err = utils.WrapError("Read MaximumFlowBitRateDL", err)
		return
	}
	ie.MaximumFlowBitRateDL = int64(tmp_MaximumFlowBitRateDL.Value)
	tmp_MaximumFlowBitRateUL := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
		ext: true,
	}
	if err = tmp_MaximumFlowBitRateUL.Decode(r); err != nil {
		err = utils.WrapError("Read MaximumFlowBitRateUL", err)
		return
	}
	ie.MaximumFlowBitRateUL = int64(tmp_MaximumFlowBitRateUL.Value)
	tmp_GuaranteedFlowBitRateDL := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
		ext: true,
	}
	if err = tmp_GuaranteedFlowBitRateDL.Decode(r); err != nil {
		err = utils.WrapError("Read GuaranteedFlowBitRateDL", err)
		return
	}
	ie.GuaranteedFlowBitRateDL = int64(tmp_GuaranteedFlowBitRateDL.Value)
	tmp_GuaranteedFlowBitRateUL := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
		ext: true,
	}
	if err = tmp_GuaranteedFlowBitRateUL.Decode(r); err != nil {
		err = utils.WrapError("Read GuaranteedFlowBitRateUL", err)
		return
	}
	ie.GuaranteedFlowBitRateUL = int64(tmp_GuaranteedFlowBitRateUL.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp := new(NotificationControl)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read NotificationControl", err)
			return
		}
		ie.NotificationControl = tmp
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_MaximumPacketLossRateDL := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 1000},
			ext: true,
		}
		if err = tmp_MaximumPacketLossRateDL.Decode(r); err != nil {
			err = utils.WrapError("Read MaximumPacketLossRateDL", err)
			return
		}
		ie.MaximumPacketLossRateDL = (*int64)(&tmp_MaximumPacketLossRateDL.Value)
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_MaximumPacketLossRateUL := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 1000},
			ext: true,
		}
		if err = tmp_MaximumPacketLossRateUL.Decode(r); err != nil {
			err = utils.WrapError("Read MaximumPacketLossRateUL", err)
			return
		}
		ie.MaximumPacketLossRateUL = (*int64)(&tmp_MaximumPacketLossRateUL.Value)
	}
	return
}
