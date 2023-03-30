# Online Store Application

Online Store Application adalah aplikasi backend yang menggunakan RESTFUL API. Pengguna bisa:

1. Melihat produk yang dijual
2. Menaruh Produk ke keranjang
3. Melihat daftar produk yang sudah dimasukkan ke keranjang
4. Menghapus daftar produk di keranjang
5. Melakukan Pembayaran dari daftar produk di keranjang
6. Login dan register sebagai pelanggan

## Cara Menjalankan server backend

### Menjalankan server backend secara lokal

1. Buka terminal CLI (powershell/cmd) dan pindahkan ke repository "Backend-Online-Store"
2. Ketikkan "go mod download" untuk mendownload dependency yang digunakan
3. Hidupkan server database Postgresql
4. Ketikkan "go run migrate/migrate.go" untuk membuat database serta tabel yang dibutuhkan ke database
5. Ketikkan "go build" untuk membangun aplikasi
6. Ketikkan "go run ." untuk menjalankan server

### Menjalankan server dengan docker image

1. Buka terminal CLI (powershell/cmd) dan pindahkan ke repository "Backend-Online-Store"
2. Ketikkan "docker-compose up -build" untuk membuild image dan menjalankan server melalui container

## Endpoint API

1.  Login : POST /login
    Parameter JSON
    {
    "Email" :
    "Password" :
    }
    Response
    200 Status OK
    404 "Email or Password are incorrect"
    404 "Please insert your Email and Password"
2.  Register : POST /register
    Parameter JSON
    {
    "Name" :
    "Email" :
    "Password" :
    }
    Response
    200 Status OK
    400
3.  Add to Cart : POST /carts
    Parameter JSON
    {
    "UserID" :
    "ProductID" :
    "Quantity" :
    }
    Response
    200 Status OK
    400
4.  Checkout and Payment : PUT /carts/payment/:user_id
    Parameter JSON
    {
    "Money" :
    }
    Response
    200 Status OK
    400
5.  List of Product in Cart : GET /carts/:user_id
    Response
    200 Status OK
    400
6.  Delete list of product in cart : DELETE /carts/:user_id
    Response
    200 Status OK
7.  View Product list : GET /products
    Response
    200 Status OK
    400
