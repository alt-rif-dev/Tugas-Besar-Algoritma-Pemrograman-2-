# AutoCare

AutoCare adalah aplikasi untuk mengelola data kendaraan dan pemeliharaan teknis secara berkala. Data utama yang digunakan adalah data kendaraan, data pemilik, dan data riwayat servis. Pengguna aplikasi adalah admin bengkel atau manajer operasional armada.

## Fitur

### 1. Pengelolaan Data Kendaraan (CRUD)
- Tambah data kendaraan
- Tampilkan data kendaraan
- Ubah data kendaraan
- Hapus data kendaraan

Data kendaraan meliputi:
- Nama Pemilik
- Kontak
- Plat nomor
- Merek kendaraan
- Tahun produksi
- Tanggal servis terakhir
- Tanggal Perbaikan
- Jenis kerusakan

### 2. Pencarian Data Kendaraan
Menggunakan:
- Sequential Search
- Binary Search

Pencarian dilakukan berdasarkan:
- Plat nomor kendaraan

### 3. Pengurutan Data Kendaraan
Menggunakan:
- Selection Sort (berdasarkan tahun produksi)
- Insertion Sort (berdasarkan tanggal servis terakhir)

### 4. Statistik
- Statistik jumlah kendaraan yang diservis per bulan
- Statistik kategori kerusakan yang paling sering muncul

---

## Struktur Data

### Kendaraan
```go
type kendaraan struct {
   platnomor, merek, tglserviceterakir,nama,kontak,tanggalperbaikan,kerusakan string
   tahunproduksi,biaya,jumlah int
}
```

---

##  Algoritma yang Digunakan

### Searching
- Sequential Search
- Binary Search

### Sorting
- Selection Sort
- Insertion Sort

### Statistik
- Pengelompokan data servis berdasarkan tanggal
- Penghitungan frekuensi jenis kerusakan

---

##  Menu Utama

```text
+++ AutoCare +++

1. Input Data Kendaraan
2. Tampilkan Data Kendaraan 
3. Hapus Data Kendaraan
4. Ubah Data Kendaraan
5. Pencarian Data Kendaraan
6. Pengurutan Data Kendaraan
7. Statistik Data Kendaraan
0. Keluar
```

---

##  Pembuat

### Raden Hasbi Radhitya N. (103012530017)
Kontribusi:
- CRUD Data Kendaraan
- Main
### Rifqi Bhadrika Adwitiya (103012500339)
Kontribusi:
- Searching
- Sorting
- Statistik Kendaraan

---
