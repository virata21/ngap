package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

// srv6: inserted by shchoi
// SID가 한개의 string으로 만들어진다.
type SegmentIdList struct {
	Value string `aper:"sizeExt,sizeLB:1,sizeUB:4096"`
	//Value []aper.BitString `aper:"sizeExt,sizeLB:1,sizeUB:160"`
}

// SegmentIdList []SegmentIdList support
// SegmentIdList *SegmentIdList support
// 위 두 조건 모두 사용 가능
func (s *SegmentIdList) Encode(w *aper.AperWriter) (err error) {
	b := []byte(s.Value)
	constraint := &aper.Constraint{Lb: 1, Ub: 4096}
	if err = w.WriteOctetString(b, constraint, true); err != nil {
		return utils.WrapError("Encode SegmentIdList", err)
	}
	return nil
}

func (s *SegmentIdList) Decode(r *aper.AperReader) (err error) {
	constraint := &aper.Constraint{Lb: 1, Ub: 4096}
	b, err := r.ReadOctetString(constraint, true)
	if err != nil {
		return utils.WrapError("Decode SegmentIdList", err)
	}
	s.Value = string(b)
	return nil
}
