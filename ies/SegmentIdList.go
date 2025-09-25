package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

//SegmentIdList struct의 string 타입은 ueransim gNB에 맞춘 타입이다

// srv6: inserted by shchoi
// SID가 한개의 string으로 만들어진다.
type SegmentIdList struct {
	Value string `aper:"sizeExt,sizeLB:1,sizeUB:4096"`
	//Value aper.BitString `aper:"sizeExt,sizeLB:1,sizeUB:4096"`
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

// srv6: inserted by shchoi
/*
// 아래 코드는 aper.BitString처리 코드임
// BIT STRING(SIZE(1..4096)) 가정
type SegmentIdList struct {
	// 프로젝트 스타일에 따라 태그는 조정하세요.
	// 예: `aper:"sizeExt,sizeLB:1,sizeUB:4096"`
	Value *aper.BitString `lb:1,ub:4096`
}

func (s *SegmentIdList) Encode(w *aper.AperWriter) (err error) {
	if s == nil || s.Value == nil {
		return utils.WrapError("Encode SegmentIdList", err)
	}
	// ASN.1 정의에 따라 확장 가능 여부 설정
	ext := true // 필요 시 false로 바꾸세요
	tmp := NewBITSTRING(*s.Value, aper.Constraint{Lb: 1, Ub: 4096}, ext)
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SegmentIdList", err)
	}
	return nil
}

func (s *SegmentIdList) Decode(r *aper.AperReader) (err error) {
	ext := true
	tmp := BITSTRING{c: aper.Constraint{Lb: 1, Ub: 4096}, ext: ext}
	if err = tmp.Decode(r); err != nil {
		return utils.WrapError("Decode SegmentIdList", err)
	}
	s.Value = &aper.BitString{
		Bytes:   tmp.Value.Bytes,
		NumBits: tmp.Value.NumBits,
	}
	return nil
}
*/
