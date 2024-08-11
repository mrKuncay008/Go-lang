# Go Language

[Go-tutor] - Tutorial Go-lang


## Installation
(Untuk Linux)
https://go.dev/doc/install
Download dulu file zip, versi menyesuaikan kebutuhan.
1) Hapus semua instalasi Go sebelumnya dengan menghapus folder /usr/local/go (jika ada), lalu ekstrak arsip yang baru saja Anda unduh ke /usr/local, buatlah path Go baru di /usr/local/go:

```sh
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz
```
Anda mungkin perlu menjalankan perintah sebagai root atau melalui sudo

2) Tambahkan /usr/local/go/bin ke variabel lingkungan PATH 
Anda dapat melakukan ini dengan menambahkan baris berikut pada $HOME/.profile atau /etc/profile (untuk installasi seluruh sistem):

```sh
export PATH=$PATH:/usr/local/go/bin

// Ini juga berlaku jika go not found command 
```

3) Lalu chek Go version
```sh
go version
```
[Go-tutor]: <https://www.w3schools.com/go/>

