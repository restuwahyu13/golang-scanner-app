package repository

import (
	"pborebuild/pinjol/entity"
	"pborebuild/pinjol/schema"
)

type Pinjol interface {
	BiayaPinjol(input *schema.SchemaPinjol) *entity.EntityPinjol
}

type BiayaPinjol struct {
	pinjol Pinjol
}

func (r *BiayaPinjol) BiayaPinjol(input *schema.SchemaPinjol) *entity.EntityPinjol {
	var pinjol entity.EntityPinjol
	pinjol.Name = input.Name
	pinjol.Alamat = input.Alamat
	pinjol.Kk = input.Kk
	pinjol.Pinjaman = input.Pinjaman
	pinjol.Angsuran = input.Angsuran
	pinjol.Salary = input.Salary
	pinjol.Bunga = input.Bunga
	pinjol.LimitPinjaman = pinjol.Salary * 3

	return &pinjol
}

func NewPinjol(input *schema.SchemaPinjol) *entity.EntityPinjol {
	var pinjol BiayaPinjol
	return pinjol.BiayaPinjol(input)
}
