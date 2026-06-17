package main

import "fmt"

const nmax = 999

type kendaraan struct {
	platnomor, tipekendaraan, tglserviceterakir, nama, kontak, tanggalperbaikan, kerusakan string
	tahunproduksi, biaya, jumlah int
}

type tabkendaraan [nmax]kendaraan
// Handling Eror Tanggal 
func kabisat(tahun int) bool {
	if tahun%400 == 0 || (tahun % 100 != 0 && tahun % 4 == 0)  {
		return true
	}
	return false
}

func cektanggal(tahun, bulan, hari int) bool {
	if bulan < 1 || bulan > 12 {
		return false
	}
	if hari < 1 {
		return false
	}
	switch bulan {
	case 1, 3, 5, 7, 8, 10, 12:
		return hari <= 31
	case 4, 6, 9, 11:
		return hari <= 30
	case 2:
		if kabisat(tahun) {
			return hari <= 29
		}
		return hari <= 28
	}
	return false
}

func cekformattanggal(tgl string) string {
	var tahun, bulan, hari int
	if tgl == "-" {
		return "SALAH"
	}

	// Cek panjang string harus 10 karakter (Tahun-Bulan-Hari)
	if len(tgl) != 10 {
		return "SALAH"
	}

	// Cek posisi strip ('-') di index 4 & 7 
		if tgl[4] != '-' || tgl[7] != '-' {
		return "SALAH"
	}

	// Konversi char jadi angka, - '0' biar jadi angka asli. 
	// contoh tgl[0] = '2' di ascii = 50, 50-48 = 2 dikonversi jadi int 
	
	tahun = int(tgl[0]-'0')*1000 + int(tgl[1]-'0')*100 + int(tgl[2]-'0')*10 + int(tgl[3]-'0')
	bulan = int(tgl[5]-'0')*10 + int(tgl[6]-'0')
	hari = int(tgl[8]-'0')*10 + int(tgl[9]-'0')

	
	if cektanggal(tahun, bulan, hari) {
		return tgl
	}

	return "SALAH"
}	

// CRUD KENDARAAN
// Raden Hasbi Radhitya N.
// Ini procedure crud input, baca, ubah, hapus, tampil KENDARAAN (tampil khusus searching buat nampilin).
func inputkendaraan(daftar *tabkendaraan, n int, total *int) {
	var i int
	var tempTgl, hasilCek string
	var valid bool
	valid = false
	for i = *total; i < n+*total; i++ {
		fmt.Print("Nama Pemilik: ")
		fmt.Scan(&daftar[i].nama)
		fmt.Print("Kontak: ")
		fmt.Scan(&daftar[i].kontak)
		fmt.Print("Plat Nomor: ")
		fmt.Scan(&daftar[i].platnomor)
		fmt.Print("Tipe Kendaraan: ")
		fmt.Scan(&daftar[i].tipekendaraan)
		fmt.Print("Tahun Produksi: ")
		fmt.Scan(&daftar[i].tahunproduksi)
		for !valid {
			fmt.Print("Tanggal Service Terakhir (contoh 2025-01-01): ")
			fmt.Scan(&tempTgl)
			hasilCek = cekformattanggal(tempTgl) 
			if hasilCek != "SALAH" {
				daftar[i].tglserviceterakir = hasilCek
				valid = true 
			} else {
				fmt.Println("Tanggal tidak valid atau format salah!")
			}
		}
		valid = false
		for !valid {
			fmt.Print("Tanggal Perbaikan (contoh 2025-01-01): ")
			fmt.Scan(&tempTgl)
			
	
			hasilCek = cekformattanggal(tempTgl) 
			
			if hasilCek != "SALAH" {
				daftar[i].tanggalperbaikan = hasilCek
				valid = true 
			} else {
				fmt.Println("Tanggal tidak valid atau format salah!")
			}
		}

		fmt.Print("Jenis Kerusakan (Gunakan '_' sebagai pengganti spasi): ")
		fmt.Scan(&daftar[i].kerusakan)
	}
	*total = *total + n
	fmt.Println("Data Kendaraan Berhasil Ditambahkan")
}

func bacakendaraan(daftar tabkendaraan, n int) {
	var i int
	fmt.Println("+++ AutoCare +++")
	for i = 0; i < n; i++ {
		fmt.Println("Nama Pemilik: ", daftar[i].nama)
		fmt.Println("Kontak: ", daftar[i].kontak)
		fmt.Println("Plat Nomor: ", daftar[i].platnomor)
		fmt.Println("Tipe Kendaraan: ", daftar[i].tipekendaraan)
		fmt.Println("Tahun Produksi: ", daftar[i].tahunproduksi)
		fmt.Println("Tanggal Service Terakhir: ", daftar[i].tglserviceterakir)
		fmt.Println("Tanggal Perbaikan: ", daftar[i].tanggalperbaikan)
		fmt.Println("Jenis Kerusakan: ", daftar[i].kerusakan)
		fmt.Println("---")
	}
}

func ubahkendaraan(daftar *tabkendaraan, n int) {
	var target, baru, tempTgl, hasilCek string
	var idx, baruint int
	var valid bool

	fmt.Print("Masukkan plat kendaraan yang ingin diubah:")
	fmt.Scan(&target)
	idx = seqsearchkendaraan(*daftar, n, target)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Print("Masukkan nama pemilik baru ( '-' untuk skip): ")
		fmt.Scan(&baru)
		if baru != "-" {
			daftar[idx].nama = baru
		}

		fmt.Print("Masukkan kontak baru ( '-' untuk skip): ")
		fmt.Scan(&baru)
		if baru != "-" {
			daftar[idx].kontak = baru
		}

		fmt.Print("Masukkan plat nomor baru ( '-' untuk skip): ")
		fmt.Scan(&baru)
		if baru != "-" {
			daftar[idx].platnomor = baru
		}

		fmt.Print("Masukkan merek kendaraan baru ( '-' untuk skip): ")
		fmt.Scan(&baru)
		if baru != "-" {
			daftar[idx].tipekendaraan = baru
		}

		fmt.Print("Masukkan tahun produksi baru ( '0' untuk skip): ")
		fmt.Scan(&baruint)
		if baruint != 0 {
			daftar[idx].tahunproduksi = baruint
		}

		valid = false
		for !valid {
			fmt.Print("Masukkan tgl service terakhir baru ( '-' untuk skip): ")
			fmt.Scan(&tempTgl)
			hasilCek = cekformattanggal(tempTgl)
			if hasilCek != "SALAH" {
				if hasilCek != "-" {
					daftar[idx].tglserviceterakir = hasilCek
				}
			} else {
				fmt.Println("Tanggal tidak valid atau format salah!")
			}
		}

		valid = false
		for !valid {
			fmt.Print("Masukkan tgl perbaikan baru ( '-' untuk skip): ")
			fmt.Scan(&tempTgl)
			
			
			hasilCek = cekformattanggal(tempTgl)
			
			if hasilCek != "SALAH" {
				if hasilCek != "-" {
					daftar[idx].tanggalperbaikan = hasilCek
				}
				valid = true 
			} else {
				fmt.Println("Tanggal tidak valid atau format salah!")
			}
		}

		fmt.Print("Masukkan jenis kerusakan baru ( '-' untuk skip): ")
		fmt.Scan(&baru)
		if baru != "-" {
			daftar[idx].kerusakan = baru
		}
	}
	fmt.Println("Data Kendaraan Berhasil Diubah ")
}

func hapuskendaraan(daftar *tabkendaraan, n *int) {
	var target string
	var idx, i int
	fmt.Print("Masukkan plat kendaraan yang ingin dihapus: ")
	fmt.Scan(&target)
	idx = seqsearchkendaraan(*daftar, *n, target)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		for i = idx; i < *n-1; i++ {
			daftar[i] = daftar[i+1]
		}
		*n--
		fmt.Println("Data kendaraan berhasil dihapus")
	}
}

func tampilkendaraan(daftar tabkendaraan, n int) {
	fmt.Println("+++ AutoCare +++")
	fmt.Println("Nama Pemilik: ", daftar[n].nama)
	fmt.Println("Kontak: ", daftar[n].kontak)
	fmt.Println("Plat Nomor: ", daftar[n].platnomor)
	fmt.Println("Tipe Kendaraan: ", daftar[n].tipekendaraan)
	fmt.Println("Tahun Produksi: ", daftar[n].tahunproduksi)
	fmt.Println("Tanggal Service Terakhir: ", daftar[n].tglserviceterakir)
	fmt.Println("Tanggal Perbaikan: ", daftar[n].tanggalperbaikan)
	fmt.Println("Jenis Kerusakan: ", daftar[n].kerusakan)
	fmt.Println("---")
}

func tambahData(daftar *tabkendaraan, n *int, nama, kontak, plat, tipekendaraan string, tahun int, tglAkhir, tglPerbaikan, rusak string) {
	var idx int
	idx = *n
	daftar[idx].nama = nama
	daftar[idx].kontak = kontak
	daftar[idx].platnomor = plat
	daftar[idx].tipekendaraan = tipekendaraan
	daftar[idx].tahunproduksi = tahun
	daftar[idx].tglserviceterakir = tglAkhir
	daftar[idx].tanggalperbaikan = tglPerbaikan
	daftar[idx].kerusakan = rusak
	*n++
}

// END CRUD KENDARAAN

// ALGORITMA Searching dan Sorting
// Rifqi Bhadrika Adwitiya

func selectionsorttahunasc(daftar *tabkendaraan, n int) {
	// Algoritma Selection Sort dengan pengurutan berdasarkan tahun produksi Ascending
	var i, pass, idx int
	var temp kendaraan
	for pass = 0; pass < n-1; pass++ {
		idx = pass
		for i = pass + 1; i < n; i++ {
			if daftar[i].tahunproduksi < daftar[idx].tahunproduksi {
				idx = i
			}
		}
		temp = daftar[pass]
		daftar[pass] = daftar[idx]
		daftar[idx] = temp
	}
}

func selectionsorttahundesc(daftar *tabkendaraan, n int) {
	// Algoritma Selection Sort dengan pengurutan berdasarkan tahun produksi Descending
	var i, pass, idx int
	var temp kendaraan
	for pass = 0; pass < n-1; pass++ {
		idx = pass
		for i = pass + 1; i < n; i++ {
			if daftar[i].tahunproduksi > daftar[idx].tahunproduksi {
				idx = i
			}
		}
		temp = daftar[pass]
		daftar[pass] = daftar[idx]
		daftar[idx] = temp
	}
}

func selectionsortplat(daftar *tabkendaraan, n int) {
	// Algoritma Selection Sort dengan pengurutan berdasarkan plat kendaraan Ascending
	var i, pass, idx int
	var temp kendaraan
	for pass = 0; pass < n-1; pass++ {
		idx = pass
		for i = pass + 1; i < n; i++ {
			if daftar[i].platnomor < daftar[idx].platnomor {
				idx = i
			}
		}
		temp = daftar[pass]
		daftar[pass] = daftar[idx]
		daftar[idx] = temp
	}
}

func insertionsortserviceasc(daftar *tabkendaraan, n int) {
	// Algoritma Insertion Sort dengan pengurutan berdasarkan tanggal service terakhir Ascending
	var pass, i int
	var temp kendaraan
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = daftar[pass]
		for i > 0 && temp.tglserviceterakir < daftar[i-1].tglserviceterakir {
			daftar[i] = daftar[i-1]
			i = i - 1
		}
		daftar[i] = temp
		pass++
	}
}

func insertionsortservicedesc(daftar *tabkendaraan, n int) {
	// Algoritma Insertion Sort dengan pengurutan berdasarkan tanggal service terakhir descending
	var pass, i int
	var temp kendaraan
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = daftar[pass]
		for i > 0 && temp.tglserviceterakir > daftar[i-1].tglserviceterakir {
			daftar[i] = daftar[i-1]
			i = i - 1
		}
		daftar[i] = temp
		pass++
	}
}

func seqsearchkendaraan(daftar tabkendaraan, n int, target string) int {
	// Algoritma Sequential Search dengan pencarian bedasarkan plat kendaraan
	var i int
	var found bool
	found = false
	i = 0
	for i < n && !found {
		if daftar[i].platnomor == target {
			found = true
		} else {
			i++
		}
	}
	if found {
		return i
	}
	return -1
}

func binsearchkendaraan(daftar tabkendaraan, n int, target string) int {
	// Algoritma Binary Search dengan pencarian bedasarkan plat kendaraan
	var left, right, mid int
	var found bool
	left = 0
	right = n - 1
	found = false
	for left <= right && !found {
		mid = (left + right) / 2
		if daftar[mid].platnomor == target {
			found = true
		} else if daftar[mid].platnomor > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if found {
		return mid
	}
	return -1
}

// statistik & kerusakan
// Rifqi Bhadrika Adwitiya
func statistikperbulan(daftar tabkendaraan, n int) {
	// Procedure untuk menampilkan statistik service tiap bulan
	// karena format tanggal: Tahun-Bulan-Hari, pake slice index 0:4 untuk tahun, 5:7 untuk bulan
	var hasil tabkendaraan
	var jumlahhasil, i, k, found int
	var periode, tipe string
	for i = 0; i < n; i++ {
		if len(daftar[i].tanggalperbaikan) >= 7 && len(daftar[i].tanggalperbaikan) <= 10{
			periode = daftar[i].tanggalperbaikan[0:4] + "-" + daftar[i].tanggalperbaikan[5:7]
			tipe = daftar[i].tipekendaraan
			found = -1
			k = 0
			for k < jumlahhasil && found == -1 {
				if hasil[k].tanggalperbaikan == periode && hasil[k].tipekendaraan == tipe {
					found = k
				}
				k++
			}
			if found != -1 {
				hasil[found].jumlah++
			} else {
				hasil[jumlahhasil].tanggalperbaikan = periode
				hasil[jumlahhasil].tipekendaraan = tipe
				hasil[jumlahhasil].jumlah = 1
				jumlahhasil++
			}
		}
	}
	fmt.Println("+++ AutoCare +++")
	fmt.Println("Statistik Service per bulan:")
	for i = 0; i < jumlahhasil; i++ {
		fmt.Println(hasil[i].tanggalperbaikan, hasil[i].jumlah, hasil[i].tipekendaraan)
	}
}

func statistikkerusakan(daftar tabkendaraan, n int) {
	// Procedure untuk menampilkan statistik kategori kerusakan yang sering muncul
	var hasil tabkendaraan
	var jumlahhasil, i, k, found int
	var jeniskerusakan, tipe string
	for i = 0; i < n; i++ {
		jeniskerusakan = daftar[i].kerusakan
		tipe = daftar[i].tipekendaraan
		found = -1
		k = 0
		for k < jumlahhasil && found == -1 {
			if hasil[k].kerusakan == jeniskerusakan && hasil[k].tipekendaraan == tipe {
				found = k
			}
			k++
		}
		if found != -1 {
			hasil[found].jumlah++
		} else {
			hasil[jumlahhasil].kerusakan = jeniskerusakan
			hasil[jumlahhasil].tipekendaraan = tipe
			hasil[jumlahhasil].jumlah = 1
			jumlahhasil++
		}
	}
	fmt.Println("+++ AutoCare +++")
	fmt.Println("Kategori kerusakan yang paling sering muncul:")
	for i = 0; i < jumlahhasil; i++ {
		fmt.Println(hasil[i].kerusakan, hasil[i].jumlah, hasil[i].tipekendaraan)
	}
}

func main() {
	var daftarkendaraan tabkendaraan
	var jumlahkendaraan, pilihan, subpilihan, idx, jumlahbaru int
	var target string
	var selesai bool
	selesai = false

	tambahData(&daftarkendaraan, &jumlahkendaraan, "Agus", "081234", "A1234", "Motor", 2013, "2022-01-01", "2025-01-01", "Ganti_Oli")
	tambahData(&daftarkendaraan, &jumlahkendaraan, "Budi", "089998", "B9999", "Mobil", 2020, "2021-04-13", "2024-06-06", "Rem_Blong")
	tambahData(&daftarkendaraan, &jumlahkendaraan, "Caca", "083210", "B5678", "Motor", 2015, "2023-05-05", "2024-06-06", "Body_Motor_Rusak")
	tambahData(&daftarkendaraan, &jumlahkendaraan, "Duncan", "086767", "D6767", "Motor", 2006, "2026-06-07", "2026-07-06", "Ganti_Ban")
	tambahData(&daftarkendaraan, &jumlahkendaraan, "Enzen", "087676", "D7676", "Mobil", 2007, "2026-05-12", "2026-07-07", "Ganti_Ban")

	for !selesai {
		fmt.Println("\n+++ AutoCare +++")
		fmt.Println("Selamat Datang Di AutoCare")
		fmt.Println("Silahkan Pilih Menu Yang Kami Sediakan")
		fmt.Println("1. Tambah Data Kendaraan")
		fmt.Println("2. Tampilkan Data Kendaraan")
		fmt.Println("3. Ubah Data Kendaraan")
		fmt.Println("4. Hapus Data Kendaraan")
		fmt.Println("5. Pencarian Data Kendaraan")
		fmt.Println("6. Pengurutan Data Kendaraan")
		fmt.Println("7. Statistik Data Kendaraan")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih 0-7: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Print("Masukkan Jumlah Kendaraan Yang Ingin Di Inputkan: ")
			fmt.Scan(&jumlahbaru)
			inputkendaraan(&daftarkendaraan, jumlahbaru, &jumlahkendaraan)
		case 2:
			bacakendaraan(daftarkendaraan, jumlahkendaraan)
		case 3:
			ubahkendaraan(&daftarkendaraan, jumlahkendaraan)
		case 4:
			hapuskendaraan(&daftarkendaraan, &jumlahkendaraan)
		case 5:
			fmt.Println("+++ AutoCare +++")
			fmt.Println("1. Cari Data Kendaraan Menggunakan Sequential Search")
			fmt.Println("2. Cari Data Kendaraan Menggunakan Binary Search")
			fmt.Println("0. Kembali")
			fmt.Print("Pilih 0-2: ")
			fmt.Scan(&subpilihan)
			switch subpilihan {
			case 1:
				fmt.Print("Masukkan Nomor Plat Kendaraan Yang Ingin Dicari: ")
				fmt.Scan(&target)
				idx = seqsearchkendaraan(daftarkendaraan, jumlahkendaraan, target)
				if idx == -1 {
					fmt.Println("Data Tidak Ditemukan")
				} else {
					fmt.Println("Data Ditemukan:")
					tampilkendaraan(daftarkendaraan, idx)
				}
			case 2:
				fmt.Print("Masukkan Nomor Plat Kendaraan Yang Ingin Dicari: ")
				fmt.Scan(&target)
				selectionsortplat(&daftarkendaraan, jumlahkendaraan)
				idx = binsearchkendaraan(daftarkendaraan, jumlahkendaraan, target)
				if idx == -1 {
					fmt.Println("Data Tidak Ditemukan")
				} else {
					fmt.Println("Data Ditemukan:")
					tampilkendaraan(daftarkendaraan, idx)
				}
			}
		case 6:
			fmt.Println("+++ AutoCare +++")
			fmt.Println("1. Mengurutkan Data Kendaraan Menggunakan Selection Sort Ascending (Tahun Produksi)")
			fmt.Println("2. Mengurutkan Data Kendaraan Menggunakan Selection Sort Descending (Tahun Produksi)")
			fmt.Println("3. Mengurutkan Data Kendaraan Menggunakan Insertion Sort Ascending (Tanggal Service Terakhir)")
			fmt.Println("4. Mengurutkan Data Kendaraan Menggunakan Insertion Sort Descending (Tanggal Service Terkahir)")
			fmt.Println("0. Kembali")
			fmt.Print("Pilih 0-4: ")
			fmt.Scan(&subpilihan)
			switch subpilihan {
			case 1:
				selectionsorttahunasc(&daftarkendaraan, jumlahkendaraan)
				fmt.Println("Data berhasil diurutkan secara ascending berdasarkan tahun produksi")
				bacakendaraan(daftarkendaraan, jumlahkendaraan)
			case 2:
				selectionsorttahundesc(&daftarkendaraan, jumlahkendaraan)
				fmt.Println("Data berhasil diurutkan secara descending berdasarkan tahun produksi")
				bacakendaraan(daftarkendaraan, jumlahkendaraan)
			case 3:
				insertionsortserviceasc(&daftarkendaraan, jumlahkendaraan)
				fmt.Println("Data berhasil diurutkan secara ascending berdasarkan tanggal service (Terdahulu ke Terbaru)")
				bacakendaraan(daftarkendaraan, jumlahkendaraan)
			
			case 4:
				insertionsortservicedesc(&daftarkendaraan, jumlahkendaraan)
				fmt.Println("Data berhasil diurutkan secara descending berdasarkan tanggal service (Terdahulu ke Terbaru)")
				bacakendaraan(daftarkendaraan, jumlahkendaraan)
			}
		case 7:
			fmt.Println("+++ AutoCare +++")
			fmt.Println("1. Menampilkan Statistik Jumlah Kendaraan Yang Diservice Per Bulan")
			fmt.Println("2. Menampilkan Statistik Kategori Kerusakan Yang Paling Sering Muncul")
			fmt.Println("0. Kembali")
			fmt.Print("Pilih 0-2: ")
			fmt.Scan(&subpilihan)
			switch subpilihan {
			case 1:
				statistikperbulan(daftarkendaraan, jumlahkendaraan)
			case 2:
				statistikkerusakan(daftarkendaraan, jumlahkendaraan)
			}
		case 0:
			fmt.Print("Terima Kasih")
			selesai = true
		}
	}
}