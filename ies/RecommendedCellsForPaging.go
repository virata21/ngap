package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type RecommendedCellsForPaging struct {
	RecommendedCellList []RecommendedCellItem `lb:1,ub:maxnoofRecommendedCells,madatory`
	// IEExtensions *RecommendedCellsForPagingExtIEs `optional`
}

func (ie *RecommendedCellsForPaging) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if len(ie.RecommendedCellList) > 0 {
		tmp := Sequence[*RecommendedCellItem]{
			Value: []*RecommendedCellItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofRecommendedCells},
			ext:   false,
		}
		for _, i := range ie.RecommendedCellList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode RecommendedCellList", err)
			return
		}
	} else {
		err = utils.WrapError("RecommendedCellList is nil", err)
		return
	}
	return
}
func (ie *RecommendedCellsForPaging) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_RecommendedCellList := Sequence[*RecommendedCellItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofRecommendedCells},
		ext: false,
	}
	fn := func() *RecommendedCellItem { return new(RecommendedCellItem) }
	if err = tmp_RecommendedCellList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read RecommendedCellList", err)
		return
	}
	ie.RecommendedCellList = []RecommendedCellItem{}
	for _, i := range tmp_RecommendedCellList.Value {
		ie.RecommendedCellList = append(ie.RecommendedCellList, *i)
	}
	return
}
