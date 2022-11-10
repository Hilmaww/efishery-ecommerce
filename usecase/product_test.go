package usecase

import (
	"efishery-ecommerce/entity/response"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	Usecase        ProductUsecase                = ProductUsecase{}
	listSeharusnya []response.GetProductResponse = []response.GetProductResponse{
		{
			Nama:     "Botol biru",
			Foto:     "usr/static/botol.jpg",
			Harga:    20000,
			Kategori: "alat makan",
		},
		{
			Nama:     "Sendok stainless",
			Foto:     "usr/static/sendok.jpg",
			Harga:    8000,
			Kategori: "alat makan",
		},
		{
			Nama:     "Charger Type C",
			Foto:     "usr/static/charger.jpg",
			Harga:    50000,
			Kategori: "elektronik",
		},
	}
	luasSeharusnya     float64 = 96
	kelilingSeharusnya float64 = 48
)

func TestGetProductList(t *testing.T) {

	// assert equality
	assert.Equal(t, Usecase, listSeharusnya, "yh salah")

}
