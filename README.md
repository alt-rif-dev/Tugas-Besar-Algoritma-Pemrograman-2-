AutoCare

AutoCare adalah aplikasi berbasis terminal yang dibuat menggunakan bahasa Go untuk membantu pengelolaan data kendaraan, pemilik kendaraan, dan riwayat servis.
Fitur

1. Pengelolaan Data Kendaraan (CRUD)
- Tambah data kendaraan
- Tampilkan data kendaraan
- Ubah data kendaraan
- Hapus data kendaraan

Data kendaraan meliputi:
- Plat nomor
- Merek kendaraan
- Tahun produksi
- Tanggal servis terakhir

2. Pengelolaan Data Pemilik (CRUD)
- Tambah data pemilik
- Tampilkan data pemilik
- Ubah data pemilik
- Hapus data pemilik

Data pemilik meliputi:
- Nama
- ID pemilik
- Kontak

3. Riwayat Servis Kendaraan
- Tambah data servis
- Tampilkan data servis

Data servis meliputi:
- Plat kendaraan
- Tanggal servis
- Jenis kerusakan
- Biaya servis

4. Pencarian Data Kendaraan
Menggunakan:
- Sequential Search
- Binary Search

Pencarian dilakukan berdasarkan:
- Plat nomor kendaraan

5. Pengurutan Data Kendaraan
Menggunakan:
- Selection Sort (berdasarkan tahun produksi)
- Insertion Sort (berdasarkan tanggal servis terakhir)

6. Statistik
- Statistik jumlah kendaraan yang diservis
- Statistik kategori kerusakan yang paling sering muncul

---

## Struktur Data

### Kendaraan
```go
type kendaraan struct {
    platnomor string
    merek string
    tahunproduksi int
    tglserviceterakir string
}
```

### Pemilik
```go
type pemilik struct {
    nama string
    kontak string
    id int
}
```

### Riwayat Service
```go
type riwayatservice struct {
    plat string
    tanggalservice string
    jeniskerusakan string
    biaya int
}
```

---

## Algoritma yang Digunakan

### Searching
- Sequential Search
- Binary Search

### Sorting
- Selection Sort
- Insertion Sort

### Statistik
- Pengelompokan data servis berdasarkan tanggal
- Penghitungan frekuensi jenis kerusakan


## 🖥️ Menu Utama

```text
+++ AutoCare +++

1. Pengelolaan Data Kendaraan
2. Pengelolaan Data Pemilik
3. Riwayat Servis Kendaraan
4. Pencarian Data Kendaraan
5. Pengurutan Data Kendaraan
6. Statistik Data Kendaraan
0. Keluar
```

---

## Anggota Pengembang

### Raden Hasbi Radhitya N.
Kontribusi:
- CRUD Data Kendaraan
- CRUD Data Pemilik
- CRUD Riwayat Servis

### Rifqi Bhadrika Adwitiya
Kontribusi:
- Searching
- Sorting
- Statistik Kendaraan

---

## Tujuan Proyek

Proyek ini dibuat untuk menerapkan konsep:
- Struct dan Array dalam Go
- Procedure dan Function
- CRUD Operation
- Searching Algorithm
- Sorting Algorithm
- Pengolahan Statistik Data
- Pemrograman Terstruktur



## 📄 Lisensi

Proyek ini dibuat untuk keperluan pembelajaran dan tugas akademik.
