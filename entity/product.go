package entity

type Product struct {
	ID        int64  `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Nama      string `json:"nama"`
	Foto      string `json:"foto"`
	Harga     int64  `json:"harga"`
	Kategori  string `json:"kategori"`
	Deskripsi string `json:"deskripsi"`
}

type Cart struct {
	ID        int64  `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Nama      string `json:"nama"`
	Foto      string `json:"foto"`
	Harga     int64  `json:"harga"`
	Kategori  string `json:"kategori"`
	Deskripsi string `json:"deskripsi"`
}
