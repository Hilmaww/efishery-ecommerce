package response

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type GetProductResponse struct {
	Nama     string `json:"nama"`
	Foto     string `json:"foto"`
	Harga    int    `json:"harga"`
	Kategori string `json:"kategori"`
}
type GetProductDetailResponse struct {
	Nama      string `json:"nama"`
	Foto      string `json:"foto"`
	Harga     int    `json:"harga"`
	Deskripsi string `json:"deskripsi"`
}

type GetProductFilteredResponse struct {
	Nama      string `json:"nama"`
	Foto      string `json:"foto"`
	Harga     int    `json:"harga"`
	Deskripsi string `json:"deskripsi"`
	Kategori  string `json:"kategori"`
}

type AddToCartRequest struct {
	Nama      string `json:"nama"`
	Foto      string `json:"foto"`
	Harga     int    `json:"harga"`
	Deskripsi string `json:"deskripsi"`
	Kategori  string `json:"kategori"`
}
type GetCartResponse struct {
	Nama     string `json:"nama"`
	Foto     string `json:"foto"`
	Harga    int    `json:"harga"`
	Kategori string `json:"kategori"`
}

type PostProofResponse struct {
	Name         string `json:"name"`
	Bukti        string `json:"bukti"`
	IsCheckedOut bool   `json:"is_checked_out"`
}
