package repository

import (
	"pborebuild/sekolah/entity"
	"pborebuild/sekolah/schema"
)

type Biaya interface {
	BiayaSekolah(input *schema.SchemaSekolah) *entity.EntitySekolah
}

type BiayaSekolah struct {
	biaya Biaya
}

func (s *BiayaSekolah) BiayaSekolah(input *schema.SchemaSekolah) *entity.EntitySekolah {
	var entity entity.EntitySekolah
	entity.UangPengembangan = input.UangPengembangan
	entity.UangSpp = input.UangSpp
	entity.UangRaport = input.UangRaport
	entity.SubTotal = input.SubTotal

	return &entity
}

func NewSekolah(input *schema.SchemaSekolah) *entity.EntitySekolah {
	var biaya BiayaSekolah
	return biaya.BiayaSekolah(input)
}
