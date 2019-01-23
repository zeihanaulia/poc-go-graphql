# POC GO language GraphQl

Proyek ini pembuktian konsep golang dengan graphql. Data yang digunakan menggunakan contoh data dari [postgresqltutorial](http://www.postgresqltutorial.com/postgresql-sample-database/).

Note:
2019-01-23:

Masih kurang nyaman kode untuk penyelesaian n+1. Jadi Graphql ini kalau ingin mengeluarkan data yang berrelasi. Dia akan melakukan perulangan query.
Misal pada contoh disini.
Gw mau mengquery data seluruh staff dan alamatnya. Hasilnya, dari total 10 staff akan mengambil data 10 kali 

```bash
select  address_id, address from address where address_id = staff.address_id (10x)
```

sebenernya ada solusinya dengan menggunakan data loader. tapi kurang nyaman dalam proses developmentnya

TODO:
[x] Query data
[x] Mytation data
[] Authorize resolver
[] Rate Limiting resolver

Library yang digunakan:
[] [GQLGen](https://gqlgen.com/)

## Catatan Library

### GQLGEN

Gqlgen menawarkan generator resolver untuk pertama kali. menarik,
Tapi ...
Gqlgen hanya mengenerate schema pertama kali. setelahnya harus dibuat manual.
Misal:

    - Step 1: coba gqlgen gq ikuti gettinstarted dari gqlgen. setelah itu gw coba generate dan berhasil.
    - Step2: gw coba membuat schema baru Staff dimana gw mau punya resolver createStaff. Setelah digenerate gw dapet error code. karena method method resolver belum didefine dan gqlgen tidak dapat memodifikasi hanya untuk menambahkan resolver untuk staff.
    - Step3: Mencoba menambahkan secara manual RegisterStaff mutation dan query resolver, sesuai dengan schema yang kita tulis

```go
    func (r *mutationResolver) RegisterStaff(ctx context.Context, email string, password string) (models.Staff, error) {
        panic("not implemented")
    }

    func (r *queryResolver) Staffs(ctx context.Context) ([]models.Staff, error) {
	    return r.staffs, nil
    }
```
