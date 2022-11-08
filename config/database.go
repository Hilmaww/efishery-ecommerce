package config

import (
	"efishery-ecommerce/entity"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB
var err error

func Database() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

//func Seeding() error {
//	seed := []entity.Product{
//		{
//			Nama:      "Botol biru",
//			Foto:      "usr/static/botol.jpg",
//			Harga:     20000,
//			Kategori:  "alat makan",
//			Deskripsi: "sebuah alat yang dapat membantu manusia minum",
//		},
//		{
//			Nama:      "Sendok stainless",
//			Foto:      "usr/static/sendok.jpg",
//			Harga:     8000,
//			Kategori:  "alat makan",
//			Deskripsi: "sebuah alat yang dapat membantu manusia makan",
//		},
//		{
//			Nama:      "Charger Type C",
//			Foto:      "usr/static/charger.jpg",
//			Harga:     50000,
//			Kategori:  "elektronik",
//			Deskripsi: "sebuah alat yang dapat membantu elektronik bekerja",
//		},
//	}
//	if err := DB.Create(&seed).Error; err != nil {
//		return err
//	}
//	return nil
//}

func AutoMigrate() {
	DB.AutoMigrate(&entity.Product{}, &entity.Cart{})
}
