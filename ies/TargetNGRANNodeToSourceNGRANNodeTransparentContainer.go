package ies

import (
	"bytes"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TargetNGRANNodeToSourceNGRANNodeTransparentContainer struct {
	RRCContainer []byte `lb:0,ub:0,madatory`
	// IEExtensions *TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs `optional`
}

func (ie *TargetNGRANNodeToSourceNGRANNodeTransparentContainer) Encode() (b []byte, err error) {
	var buf bytes.Buffer
	w := aper.NewWriter(&buf)
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_RRCContainer := NewOCTETSTRING(ie.RRCContainer, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_RRCContainer.Encode(w); err != nil {
		err = utils.WrapError("Encode RRCContainer", err)
		return
	}
	err = w.Close()
	b = buf.Bytes()
	return
}
func (ie *TargetNGRANNodeToSourceNGRANNodeTransparentContainer) Decode(wire []byte) (err error) {
	r := aper.NewReader(bytes.NewBuffer(wire))
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_RRCContainer := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_RRCContainer.Decode(r); err != nil {
		err = utils.WrapError("Read RRCContainer", err)
		return
	}
	ie.RRCContainer = tmp_RRCContainer.Value
	return
}
