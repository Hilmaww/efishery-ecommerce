# EFISHERY-ECOMMERCE API SYSTEM #
- Expected Finish Date: 2022-11-11
- Author: Muhammad Hilmi Fawwaz
- Approvers: Efishery Academy Mentors

![Gambar Meme RFC](static_readme/RFC_Meme.JPG)

## Summary ##
    Efishery Ecommerce adalah sebuah layanan API untuk membantu keberjalanan pemilihan produk dan penunjukkan bukti pembayaran. 
    Layanan ini dibuat dalam bahasa Go dengan database Postgresql.

## Problem & Motivation ##
    Sebuah UMKM ingin melayani pelanggan secara digital melalui internet. Pelayanan yang ingin dicakup dalam sistem ecommerce ini adalah pengecekan barang, pemilihan barang yang ingin dibeli, serta pengiriman bukti pembayaran. 
    Desain arsitektur dibuat dengan sederhana supaya para Front End Engineer dapat menampilkan layanan dengan mudah.

## Detailed Design ##
### Database Design ###
Database akan dibuat sebagai berikut,

![Gambar Database](static_readme/ERD.png "Entity Relationship Diagram")

- Tabel Produk
![Gambar tabel produk](static_readme/Tabel_Produk.png)
- Tabel Cart
![Gambar tabel cart](static_readme/Tabel_Cart.png)
- Tabel Proof
![Gambar tabel proof](static_readme/Tabel_Proof.png)
### API Contract ###
* Dokumentasi API dapat dibaca [disini](https://app.swaggerhub.com/apis-docs/HILMIAMBARI27/efishery-ecommerce/1.0.0#/developers/GetProductList)
### Flowchart ###

![Gambar Flowchart sistem Ecommerce](static_readme/efishery-ecommerce-flow-app@2x.png)
## Dependencies ##
- Echo Framework & middleware
- Postgresql
- GORM

## Milestone/Deployment Strategy ##
### Adoption and Migration ###

![Gambar adaptation strategy](static_readme/Adaptation_flow_for_efishery-ecommerce.png)

## Drawbacks/Risks ##
- Server Down
- Database Breached

## Alternatives ##
- Open shop at other ecommerce platform
- Make own platform using Shopify

## Development Possibilities ##
- Membuat docker container dengan multistage
- learn ci/cd deployment
- Menambahkan User management
- Menambahkan Unit Testing

## Link to Deployment ##
* https://efishery-ecommerce-web.herokuapp.com/