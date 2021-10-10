package repository

import (
	"pborebuild/sekolah/entity"
	"pborebuild/sekolah/schema"
)

type Biaya interface {
	BiayaPendidikan(input *schema.SchemaBiayaPendidikan) *entity.EntityBiayaPendidikan
}

type BiayaPendidikan struct {
	biaya Biaya
}

func (s *BiayaPendidikan) BiayaPendidikan(input *schema.SchemaBiayaPendidikan) *entity.EntityBiayaPendidikan {
	var entity entity.EntityBiayaPendidikan
	entity.UangPengembangan = input.UangPengembangan
	entity.UangSpp = input.UangSpp
	entity.UangRaport = input.UangRaport
	entity.SubTotal = input.SubTotal

	return &entity
}

func NewBiayaPendidikan(input *schema.SchemaBiayaPendidikan) *entity.EntityBiayaPendidikan {
	var biaya BiayaPendidikan
	return biaya.BiayaPendidikan(input)
}
