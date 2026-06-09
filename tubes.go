package main

import "fmt"

	const nmax = 999

	type kendaraan struct {
		platnomor, merek, tglserviceterakir,nama,kontak,tanggalperbaikan,kerusakan string
		tahunproduksi,biaya,jumlah int
	}

	type tabkendaraan [nmax]kendaraan


	// CRUD KENDARAAN
	// Raden Hasbi Radhitya N.
	// Ini procedure crud input, baca, ubah, hapus, tampil KENDARAAN (tampil khusus searching buat nampilin). 
	func inputkendaraan(daftar *tabkendaraan,  n int) {
		var i int
		for i = 0; i<n; i++{
			fmt.Print("Plat Nomor: ")
			fmt.Scan(&daftar[i].platnomor)
			fmt.Print("Merek: ")
			fmt.Scan(&daftar[i].merek)
			fmt.Print("Tahun Produksi: ")
			fmt.Scan(&daftar[i].tahunproduksi)
			fmt.Print("Tanggal Servis Terakhir (Tahun-Bulan-Hari): ")
			fmt.Scan(&daftar [i].tglserviceterakir)
			fmt.Print("Tanggal Perbaikan (Tahun-Bulan-Hari): ")
			fmt.Scan(&daftar[i].tanggalperbaikan)
			fmt.Print("Jenis Kerusakan (Gunakan '_' sebagai pengganti spasi): ")
			fmt.Scan(&daftar[i].kerusakan)
			}
			fmt.Println("Data Kendaraan Berhasil Ditambahkan")
	}

	func bacakendaraan(daftar tabkendaraan,n int){
		var i int
		fmt.Println("+++ AutoCare +++")
		for i = 0; i<n; i++{
			fmt.Println("Plat Nomor:", daftar[i].platnomor)
			fmt.Println("Merek:", daftar[i].merek)
			fmt.Println("Tahun Produksi:", daftar[i].tahunproduksi)
			fmt.Println("Tgl Service:", daftar[i].tglserviceterakir)
			fmt.Println("Tanggal Perbaikan:", daftar[i].tanggalperbaikan)
			fmt.Println("Jenis Kerusakan:", daftar[i].kerusakan)
			fmt.Println("---")	
		}
	}

	func ubahkendaraan(daftar *tabkendaraan, n int ){
		var target, baru string
		var idx, baruint int
		fmt.Print("Masukkan plat kendaraan yang ingin diubah:")
		fmt.Scan(&target)
		idx = seqsearchkendaraan(*daftar, n, target)
		if idx == -1{
			fmt.Println("Data tidak ditemukan")
		}else{
			fmt.Print("Masukkan nama pemilik baru ( '-' untuk skip): ")
			fmt.Scan(&baru)
			if baru != "-"{
				daftar[idx].nama = baru
			}

			fmt.Print("Masukkan kontak baru ( '-' untuk skip): ")
			fmt.Scan(&baru)
			if baru != "-"{
				daftar[idx].kontak = baru
			}

			fmt.Print("Masukkan plat nomor baru ( '-' untuk skip): ")
			fmt.Scan(&baru)
			if baru != "-"{
				daftar[idx].platnomor = baru
			}

			fmt.Print("Masukkan merek kendaraan baru ( '-' untuk skip): ")
			fmt.Scan(&baru)
			if baru != "-"{
				daftar[idx].merek = baru
			}
			
			fmt.Print("Masukkan tahun produksi  baru ( '0' untuk skip): ")
			fmt.Scan(&baruint)
			if baruint != 0 {
				daftar[idx].tahunproduksi = baruint
			}

			fmt.Print("Masukkan tgl service terakhir yang baru ( '-' untuk skip): ")
			fmt.Scan(&baru)
			if baru != "-"{
				daftar[idx].tglserviceterakir = baru
			}
			fmt.Print("Masukkan tgl perbaikan baru ( '-' untuk skip): ")
			fmt.Scan(&baru)
			if baru != "-"{
				daftar[idx].tanggalperbaikan = baru
			}
			fmt.Print("Masukkan jenis kerusakan baru ( '-' untuk skip): ")
			fmt.Scan(&baru)
			if baru != "-"{
				daftar[idx].tglserviceterakir = baru
			}
		}
		fmt.Println("Data Kendaraan Berhasil Diubah ")
	}

	func hapuskendaraan(daftar *tabkendaraan, n *int){
		var target string
		var idx,i int
		fmt.Print("Masukkan plat kendaraan yang ingin dihapus: ")
		fmt.Scan(&target)
		idx = seqsearchkendaraan(*daftar,*n,target)
		if idx == -1{
			fmt.Println("Data tidak ditemukan")
		}else{
			for i = idx; i < *n-1; i++{
				daftar[i] = daftar[i+1]
			}
			*n--
			fmt.Println("Data kendaraan berhasil dihapus")
		}
	}
	func tampilkendaraan(daftar tabkendaraan,n int){
		fmt.Println("Nama Pemilik:", daftar[n].nama)
		fmt.Println("Kontak:", daftar[n].kontak)
		fmt.Println("+++ AutoCare +++")
		fmt.Println("Plat Nomor:",daftar[n].platnomor)
		fmt.Println("Merek:",daftar[n].merek)
		fmt.Println("Tahun Produksi:",daftar[n].tahunproduksi)
		fmt.Println("Tanggal Service Terakhir:",daftar[n].tglserviceterakir)
		fmt.Println("---")
	}
	
	func tambahData(daftar *tabkendaraan, n *int, nama, kontak, plat, merek string, tahun int, tglAkhir, tglPerbaikan, rusak string) {
    	var idx int 
		idx = *n
    	daftar[idx].nama = nama
    	daftar[idx].kontak = kontak
    	daftar[idx].platnomor = plat
    	daftar[idx].merek = merek
    	daftar[idx].tahunproduksi = tahun
    	daftar[idx].tglserviceterakir = tglAkhir
    	daftar[idx].tanggalperbaikan = tglPerbaikan
    	daftar[idx].kerusakan = rusak
    	*n++ 
	}

	// END CRUD KENDARAAN


	// ALGORITMA Searching dan Sorting
	// Rifqi Bhadrika Adwitiya

	func selectionsorttahun(daftar *tabkendaraan, n int) {
		// Algoritma Selection Sort dengan pengurutan berdasarkan tahun produksi
		var i, pass, idx int
		var temp kendaraan
		for pass = 0; pass < n-1; pass++ {
			idx = pass
			for i = pass + 1; i < n; i++ {
				if daftar[i].tahunproduksi < daftar[idx].tahunproduksi   {
					idx = i
				}
			}
			temp = daftar[pass]
			daftar[pass] = daftar[idx]
			daftar[idx] = temp
		}
	}

	func selectionsortplat(daftar *tabkendaraan, n int) {
		// Algoritma Selection Sort dengan pengurutan berdasarkan plat kendaraan
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

	func insertionsortservice(daftar *tabkendaraan, n int) {
		// Algoritma Insertion Sort dengan pengurutan berdasarkan tanggal service terakhir
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
	func seqsearchkendaraan(daftar tabkendaraan,n int, target string)int {
		// Algoritma Sequential Search dengan pencarian bedasarkan plat kendaraan
		var i int 
		var found bool
		found = false
		i = 0
		for i < n && !found{
			if daftar[i].platnomor == target{
				found = true
			}else{
				i++
			}
		} 
		if found{
				return i
			}
		return -1
	}

	func binsearchkendaraan(daftar tabkendaraan,n int,target string)int{
		// Algoritma Binary Search dengan pencarian bedasarkan plat kendaraan
		var left,right,mid int
		var found bool
		left = 0
		right = n-1
		found = false
		for left <= right && !found{
			mid = (left+right)/2
			if daftar[mid].platnomor == target{
				found = true
			}else if daftar[mid].platnomor> target{
				right = mid -1
			}else{
				left = mid + 1
			}
		}
		if found{
		return mid
	}
		return  -1 
	}

	// statistik & kerusakan
	// Rifqi Bhadrika Adwitiya
	func statistikperbulan(daftar tabkendaraan,n int){ 
	// Procedure untuk menampilkan statistik service tiap bulan
	// dengan menggunakan format tanggal: YYYY-MM-DD, kita menggunakan slice index 0:4 untuk tahun, 5:7 untuk bulan
		var hasil tabkendaraan
		var jumlahhasil,i,k,found int
		var periode string
		for i = 0; i<n; i++{
			periode = daftar[i].tanggalperbaikan[0:4]+"-" + daftar[i].tanggalperbaikan[5:7]+"-"+daftar[i].tanggalperbaikan[8:10]
			found = -1
			k = 0
			for k < jumlahhasil && found == -1{
				if hasil[k].tanggalperbaikan == periode{
					found = k
				}
			k++
			}
			if found != -1 {
				hasil[found].jumlah++
			}else{
				hasil[jumlahhasil].tanggalperbaikan = periode
				hasil[jumlahhasil].jumlah = 1
				jumlahhasil++
			}
		}
		fmt.Println("+++ AutoCare +++")
		fmt.Println("Statistik Service per bulan:")
		for i = 0; i< jumlahhasil; i++{
			fmt.Println(hasil[i].tanggalperbaikan, hasil[i].jumlah)
		}
	}

	func statistikkerusakan(daftar tabkendaraan,n int){
	// Procedure untuk menampilkan statistik kategori kerusakan yang sering muncul
		var hasil tabkendaraan
		var jumlahhasil,i,k,found int
		var jeniskerusakan string
		for i = 0; i<n; i++{
			jeniskerusakan = daftar[i].kerusakan	
			found = -1
			k = 0
			for k < jumlahhasil && found == -1{
				if hasil[k].kerusakan == jeniskerusakan{
					found = k
				}
			k++
			}
			if found != -1 {
				hasil[found].jumlah++
			}else{
				hasil[jumlahhasil].kerusakan = jeniskerusakan
				hasil[jumlahhasil].jumlah = 1
				jumlahhasil++
			}
		}
		fmt.Println("+++ AutoCare +++")
		fmt.Println("Kategori kerusakan yang paling sering muncul:")
		for i = 0; i< jumlahhasil; i++{
			fmt.Println(hasil[i].kerusakan, hasil[i].jumlah)
		}
	}
	
	func main() {
		var daftarkendaraan tabkendaraan
		var jumlahkendaraan,pilihan,subpilihan,idx int
		var target string
		var selesai bool
		selesai = false
		tambahData(&daftarkendaraan, &jumlahkendaraan, "Agus", "081234", "A1234", "Honda", 2013, "2022-01-01", "2025-01-01", "Ganti_Oli")
		tambahData(&daftarkendaraan, &jumlahkendaraan, "Budi", "089998", "B9999", "Toyota", 2020, "2021-04-13", "2024-06-06", "Rem_Blong")
		tambahData(&daftarkendaraan, &jumlahkendaraan, "Caca", "083210", "B5678", "Honda", 2015, "2023-05-05", "2024-06-06", "Body_Motor_Rusak")
		tambahData(&daftarkendaraan, &jumlahkendaraan, "Duncan", "086767", "D6767", "Toyota", 2006, "2026-06-07", "2026-07-06", "Ganti_Ban")
		for !selesai {
			fmt.Println("+++ AutoCare +++")
			fmt.Println("Selamat Datang Di AutoCare")
			fmt.Println("Silahkan Pilih Menu Yang Kami Sediakan")
			fmt.Println("1. Tambah Data Kendaraan")
			fmt.Println("2. Tampilkan Data Kendaraan")
			fmt.Println("3. Ubah Data Kendaraann")	
			fmt.Println("4. Hapus Data Kendaraann")	
			fmt.Println("5. Pencarian Data Kendaraan")
			fmt.Println("6. Pengurutan Data Kendaraan")
			fmt.Println("7. Statistik Data Kendaraan")
			fmt.Println("0. Keluar")
			fmt.Print("Pilih 0-7: ")
			fmt.Scan(&pilihan)
			switch pilihan{
			case 1:
				fmt.Print("Masukkan Jumlah Kendaraan Yang Ingin Di Inputkan: ")
				fmt.Scan(&jumlahkendaraan)
				inputkendaraan(&daftarkendaraan,jumlahkendaraan)
			case 2:
				bacakendaraan(daftarkendaraan,jumlahkendaraan)
			case 3:
				ubahkendaraan(&daftarkendaraan,jumlahkendaraan)
			case 4:
				hapuskendaraan(&daftarkendaraan,&jumlahkendaraan)
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
			fmt.Println("1. Mengurutkan Data Kendaraan Menggunakan Selection Sort (Tahun)")
			fmt.Println("2. Mengurutkan Data Kendaraan Menggunakan Insertion Sort (Tanggal Service)")
			fmt.Println("0. Kembali")
			fmt.Print("Pilih 0-2: ")
			fmt.Scan(&subpilihan)
			switch subpilihan {
			case 1:
				selectionsorttahun(&daftarkendaraan, jumlahkendaraan)
				fmt.Println("Data berhasil diurutkan berdasarkan tahun produksi")
				bacakendaraan(daftarkendaraan, jumlahkendaraan)
			case 2:
				insertionsortservice(&daftarkendaraan, jumlahkendaraan)
				fmt.Println("Data berhasil diurutkan berdasarkan tanggal service (Terbaru)")
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
