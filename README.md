## BOOK LIST APP
## Training assessment Batch Excellent Echo

#

<span style="color:red; font-weight:bold"> WARNING</span>

**DILARANG MELAKUKAN TINDAKAN ANTARA LAIN :**
- Dilarang me-copy code baik di website atau di project laptop, wajib mengerjakan dari awal
- Dilarang bekerja sama atau meminta code ke teman / murid lain
- tidak mengerjakan project sesuai waktu yang ditentukan (selain jam 09.15 - 17.00)

**Jika ketahuan melakukan salah satu tindakan diatas, maka dianggap gagal dan mendapat score 0**

#
<span style="font-weight:bold">DOCUMENTATION</span>
<P>Berikut dokumentasi pendukung dari aplikasi ini, antara lain:</p>

```
1. database_diagram.png 
berisi gambaran table-tabel yang perlu di buat dan juga relasinya.

2. api-doc.md 
berisi bentuk response REST API dari setiap routing yang dibuat untuk aplikasi ini

3. example.env
berisi contoh isi file .env (environment variable) yang dibutuhkan untuk aplikasi ini

```

**NB : Silahkan pakai dokumentasi pendukung tersebut untuk memudahkan mengerjakan book list app**

## Release 1

Terdapat 2 direktori / folder :
- folder **backend** yang akan digunakan untuk code REST API / develop back end service
- folder **frontend** yang akan digunakan untuk code front end / client side

<br/>Kemudian lakukan inisiaiisasi project di folder backend :
```
go mod init <name project>
```

kemudian lakukan instalasi project yang diperlukan, antara lain:
1. GORM
2. GORM driver mysql
3. GIN - GORM

Buatlah database yang diperlukan dan code untuk membuat koneksi ke database menggunakan GORM dengan dua tabel yaitu user dan book dengan kolom yang sudah diatur. Relasi antara kedua tabel adalah user dapat memiliki banyak user tetapi book hanya punya satu user. Gambar dapat dilihat di file : 
```
database_diagram.png
```

Simpan data ***credential*** atau data sensitif dengan menerapkan environment variable. Contoh bentuk environment variable dapat dilihat di file :
```
example.env
```

## Release 2

Buatlah rest API yang dibutuhkan untuk menggunakan GIN - GONIC.
- REST API yang harus dibuat adalah CRUD untuk tabel users dan CRUD untuk tabel books.
- Routing yang harus dibuat antara lain
```
REST API users
GET /users
POST /users/register
POST /users/login
GET /users/:user_id
PUT /users/:user_id
DELETE /users/:user_id

REST API books
GET /books
POST /books
GET /books/:book_id
PUT /books/:book_id
DELETE /books/:book_id
```

SIlahkan menerapkan clean architecture dengan menggunakan layer:
- repository sebagai database commnucation dan ORM
- service sebagai bussiness logic
- handler untuk menerima request dan mengembalikan response

Anda wajib membuat bentuk request, response dan juga header sesuai file :  
```
api-doc.md
```

## Release 3

1. Buatlah authentication dan authorization menggunakan JWT dengan menginstall package jwt go terlebih dahulu
```
github.com/dgrijalva/jwt-go
```
2. Buatlah middleware yang diperlukan untuk mengecek apakah user sudah login dan juga mendapatkan user id yang sudah login.

## Release 4

1. gunakan package untuk menghandle CORS GIN menggunakan :
```
github.com/gin-contrib/cors
```
silahkan buat sebuah fungsi dengan nama CORSMiddleware untuk menghandle cors gin middleware :
```go
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

```

Kemudian pakai di main aplikasi kalian dengan cara menggunakan fungsi .use() dari package GIN-GONIC
<br/>
2. (OPTIONAL / POINT PLUS) Silahkan melakukan deployment database mysql menggunakan website yang bisa dimanfaatkan seperti remotemysql.com dan juga lakukan deployment REST API menggunakan heroku.

## Release 5
Lakukan pengerjaan mulai dari Release 5 di folder **frontend** :
1. Buatlah project baru front end apllication menggunakan ReactJS
2. Buatlah page antara lain:
- Home
- Register user
- Login user
- Dashboard (show all book, button update book, and delete book)
- Profile
- Book detail
- Update Book

## Release 6

1. Gunakan Routing untuk mengarahkan tiap - tiap page dengan ketentukan page sebagai berikut:
- public route : Home, Register, Login
- Private route : Dashboard, Profile, Book Detail, Book Update 

2. Silahkan menerapkan authentication dan authorization sebagai berikut:
- User tetap dalam kondisi login ketika website / aplikasi di refresh. 
- terdapat fitur logout.
- private route hanya bisa di akses jika user sudah login.
- Book list yang berada di dashboard hanya memunculkan list yang di create oleh user yang sedang login.
- Book detail hanya bisa dilihat jika book tersebut di create oleh user yang sedang login.

## Release 7
Lakukan wiring atau penggabungan antara back end dan front end menggunakan library AXIOS dan state management REDUX

## Release 8

1. Buatlah styling yang diperlukan, tidak wajib terlalu bagus tetapi minimal website tersebut enak dipakai.
2. (OPTIONAL / NILAI PLUS) Silahkan menerapkan responsive web design untuk setiap page yang dibuat
3. (OPTIONAL / NILAI PLUS) Silahkan melakukan deployment untuk aplikasi front end yang dibuat menggunakan netlify


> NB : Semua ketentuan dari releasi 1 sampai relase 8 wajib dibuat. Kecuali untuk yang OPTIONAL, boleh dikerjakan atau tidak  
