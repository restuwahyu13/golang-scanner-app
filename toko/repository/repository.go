package repository

import (
	"pborebuild/toko/entity"
	"pborebuild/toko/schema"
)

type Toko interface {
	BiayaToko(input *schema.SchemaToko) *entity.EntityToko
}

type BiayaToko struct {
	toko Toko
}

func (r *BiayaToko) BiayaToko(input *schema.SchemaToko) *entity.EntityToko {
	var toko entity.EntityToko

	toko.Modal = input.Modal + input.TotalPenjualan
	toko.Keuntungan = (input.Jual - input.Modal) * input.TotalPenjualan
	toko.TotalPenjualan = input.TotalPenjualan

	return &toko
}

func NewToko(input *schema.SchemaToko) *entity.EntityToko {
	var toko BiayaToko
	return toko.BiayaToko(input)
}
