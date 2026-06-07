package main

import "fmt"

	const nmax = 999

	type kendaraan struct {
		platnomor, merek, tglserviceterakir string
		tahunproduksi int
	}
	type pemilik struct {
		nama, kontak string
		id int
	}
	type riwayatservice struct{
		plat,tanggalservice,jeniskerusakan string
		biaya int
	}

	type statbulan struct {
		periode string  
		jumlah  int
	}

	type statkerusakan struct {
		jenis  string  
		jumlah int
	}

	type tabkendaraan [nmax]kendaraan
	type tabpemilik [nmax]pemilik
	type tabservice [nmax]riwayatservice
	type tabstat [nmax]statbulan
	type tabkerusakan [nmax]statkerusakan

	// CRUD KENDARAAN
	// Raden Hasbi Radhitya N.
	// Ini procedure crud input, baca, ubah, hapus, tampil KENDARAAN (tampil khusus searching buat nampilin). 
	func inputkendaraan(daftar *tabkendaraan,  n int) {
		var i int
		for i = 0; i < n; i++{
			fmt.Print("Plat Nomor: ")
			fmt.Scan(&daftar[i].platnomor)
			fmt.Print("Merek: ")
			fmt.Scan(&daftar[i].merek)
			fmt.Print("Tahun Produksi: ")
			fmt.Scan(&daftar[i].tahunproduksi)
			fmt.Print("Tanggal Servis Terakhir (Tahun-Bulan-Hari): ")
			fmt.Scan(&daftar[i].tglserviceterakir)
			}
			fmt.Println("Kendaraan berhasil ditambahkan")
	}

	func bacakendaraan(daftar tabkendaraan,n int){
		var i int
		fmt.Println("+++ AutoCare +++")
		for i = 0; i < n; i++{
			fmt.Println("Plat Nomor:", daftar[i].platnomor)
			fmt.Println("Merek:", daftar[i].merek)
			fmt.Println("Tahun Produksi:", daftar[i].tahunproduksi)
			fmt.Println("Tgl Service:", daftar[i].tglserviceterakir)
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

			fmt.Print("Data Kendaraan Berhasil Diubah ")
			fmt.Print("Masukkan tgl service baru ( '-' untuk skip): ")
			fmt.Scan(&baru)
			if baru != "-"{
				daftar[idx].tglserviceterakir = baru
			}
		}
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
		fmt.Println("+++ AutoCare +++")
		fmt.Println("Plat Nomor:",daftar[n].platnomor)
		fmt.Println("Merek:",daftar[n].merek)
		fmt.Println("Tahun Produksi:",daftar[n].tahunproduksi)
		fmt.Println("Tanggal Service Terakhir:",daftar[n].tglserviceterakir)
		fmt.Println("---")
	}
	// END CRUD KENDARAAN

	// CRUD PEMILIK
	// Raden Hasbi Radhitya N.
	// Ini procedure crud input, baca, ubah, hapus untuk data pemilik. 
	func inputpemilik(daftar *tabpemilik,  n int) {
	var i int
	for i = 0; i < n; i++{
			fmt.Print("Nama: ")
			fmt.Scan(&daftar[i].nama)
			fmt.Print("Id: ")
			fmt.Scan(&daftar[i].id)
			fmt.Print("Kontak (No Hp): ")
			fmt.Scan(&daftar[i].kontak)
		}
		fmt.Println("Data pemilik berhasil ditambahkan")
	}

	func bacapemilik(daftar tabpemilik, n int){
		var i int
		fmt.Println("+++ AutoCare +++")
		for i = 0; i < n; i++{
			fmt.Println("Nama:", daftar[i].nama)
			fmt.Println("Id:", daftar[i].id)
			fmt.Println("Kontak (Nomor Hp):", daftar[i].kontak)
			fmt.Println("---")
		}
	}

	func ubahpemilik(daftar *tabpemilik, n int ){
		var baru string
		var idx,baruint,target int
		fmt.Print("Masukkan ID pemilik yang ingin diubah:")
		fmt.Scan(&target)
		idx = seqsearchpemilik(*daftar, n, target)
		if idx == -1{
			fmt.Println("Data tidak ditemukan")
		}else{
			fmt.Print("Masukkan nama pemilik baru ( '-' untuk skip): ")
			fmt.Scan(&baru)
			if baru != "-"{
				daftar[idx].nama = baru
			}

			fmt.Print("Masukkan ID pemilik baru ( '0' untuk skip): ")
			fmt.Scan(&baruint)
			if baruint != 0{
				daftar[idx].id = baruint
			}

			fmt.Print("Masukkan kontak pemilik baru ( '-' untuk skip): ")
			fmt.Scan(&baru)
			if baru != "-"{
				daftar[idx].kontak = baru
			}
			
		}
	}

	func hapuspemilik(daftar *tabpemilik,n *int){
		var idx, i, target int
		fmt.Print("Masukkan ID pemilik yang ingin dihapus: ")
		fmt.Scan(&target)
		idx = seqsearchpemilik(*daftar, *n, target)
		if idx == -1{
			fmt.Println("Data tidak ditemukan")
		}else{
			for i = idx; i < *n-1; i++{
				daftar[i] = daftar[i+1]
			}
			*n--
			fmt.Println("Data pemilik berhasil dihapus")
		}
	}
	// END CRUD PEMILIK

	// CR SERVICE
	// Raden Hasbi Radhitya N.
	// Ini procedure CREATE dan READ untuk riwayat service. 
	func inputservice(daftar *tabservice, n int){
		var i int 
		for i = 0; i<n; i++{
			fmt.Print("Plat Nomor: ")
			fmt.Scan(&daftar[i].plat)
			fmt.Print("Tanggal Service: ")
			fmt.Scan(&daftar[i].tanggalservice)
			fmt.Print("Jenis Kerusakan: ")
			fmt.Scan(&daftar[i].jeniskerusakan)
			fmt.Print("Biaya: ")
			fmt.Scan(&daftar[i].biaya)
		}
		fmt.Println("Riwayat service berhasil ditambahkan")
	}

	func bacaservice(daftar tabservice, n int){
		var i int
		if n == 0{
			fmt.Println("Data Kosong")
		}else{
			fmt.Println("+++ AutoCare +++")
			for i = 0; i < n; i++{
				fmt.Println("Plat Nomor:", daftar[i].plat)
				fmt.Println("Tanggal Service:", daftar[i].tanggalservice)
				fmt.Println("Jenis Kerusakan:", daftar[i].jeniskerusakan)
				fmt.Println("Biaya:", daftar[i].biaya)
			}
		}
	}
	// END CR SERVICE

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
	func seqsearchpemilik(daftar tabpemilik,n int, target int)int {
		// Algoritma Sequential Search dengan pencarian bedasarkan Nama Pemilik
		var i int 
		var found bool
		found = false
		i = 0
		for i < n && !found{
			if daftar[i].id == target{
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
	func statistikperbulan(daftar tabservice,n int){ 
	// Procedure untuk menampilkan statistik service tiap bulan
	// dengan menggunakan format tanggal: YYYY-MM-DD, kita menggunakan slice index 0:4 untuk tahun, 5:7 untuk bulan
		var hasil tabstat
		var jumlahhasil,i,k,found int
		var periode string
		for i = 0; i<n; i++{
			periode = daftar[i].tanggalservice[0:4]+"-" + daftar[i].tanggalservice[5:7]+"-"+daftar[i].tanggalservice[8:10]
			found = -1
			k = 0
			for k < jumlahhasil && found == -1{
				if hasil[k].periode == periode{
					found = k
				}
			k++
			}
			if found != -1 {
				hasil[found].jumlah++
			}else{
				hasil[jumlahhasil].periode = periode
				hasil[jumlahhasil].jumlah = 1
				jumlahhasil++
			}
		}
		fmt.Println("+++ AutoCare +++")
		fmt.Println("Statistik Service per bulan:")
		for i = 0; i< jumlahhasil; i++{
			fmt.Println(hasil[i].periode, hasil[i].jumlah)
		}
	}

	func statistikkerusakan(daftar tabservice,n int){
	// Procedure untuk menampilkan statistik kategori kerusakan yang sering muncul
		var hasil tabkerusakan
		var jumlahhasil,i,k,found int
		var jeniskerusakan string
		for i = 0; i<n; i++{
			jeniskerusakan = daftar[i].jeniskerusakan	
			found = -1
			k = 0
			for k < jumlahhasil && found == -1{
				if hasil[k].jenis == jeniskerusakan{
					found = k
				}
			k++
			}
			if found != -1 {
				hasil[found].jumlah++
			}else{
				hasil[jumlahhasil].jenis = jeniskerusakan
				hasil[jumlahhasil].jumlah = 1
				jumlahhasil++
			}
		}
		fmt.Println("+++ AutoCare +++")
		fmt.Println("Kategori kerusakan yang paling sering muncul:")
		for i = 0; i< jumlahhasil; i++{
			fmt.Println(hasil[i].jenis, hasil[i].jumlah)
		}
	}
	
	func main() {
		var daftarkendaraan tabkendaraan
		var daftarpemilik tabpemilik
		var daftarservice tabservice 
		var jumlahkendaraan,jumlahpemilik,jumlahservice,pilihan,subpilihan,idx int
		var target string
		var selesai bool
		selesai = false
		for !selesai {
			fmt.Println("+++ AutoCare +++")
			fmt.Println("Selamat Datang Di AutoCare")
			fmt.Println("Silahkan Pilih Menu Yang Kami Sediakan")
			fmt.Println("1. Pengelolaan Data Kendaraan")
			fmt.Println("2. Pengelolaan Data Pemilik")
			fmt.Println("3. Riwayat Servis Kendaraann")	
			fmt.Println("4. Pencarian Data Kendaraan")
			fmt.Println("5. Pengurutan Data Kendaraan")
			fmt.Println("6. Statistik Data Kendaraan")
			fmt.Println("0. Keluar")
			fmt.Print("Pilih 0-6: ")
			fmt.Scan(&pilihan)
			switch pilihan{
			case 1:
				fmt.Println("+++ AutoCare +++")
				fmt.Println("1. Tambah Data Kendaraan")
				fmt.Println("2. Tampilkan Data Kendaraan")
				fmt.Println("3. Ubah Data Kendaraan")
				fmt.Println("4. Hapus Data Kendaraan")
				fmt.Println("0. Kembali")
				fmt.Print("Pilih 0-4: ")
				fmt.Scan(&subpilihan)
				switch subpilihan{
				case 1:
					fmt.Print("Masukkan Jumlah Data Kendaraan Yang Ingin Di Input: ")
					fmt.Scan(&jumlahkendaraan)
					inputkendaraan(&daftarkendaraan,jumlahkendaraan)
				case 2:
					bacakendaraan(daftarkendaraan,jumlahkendaraan)
				case 3:
					ubahkendaraan(&daftarkendaraan,jumlahkendaraan)
				case 4:
					hapuskendaraan(&daftarkendaraan,&jumlahkendaraan)
				}
			case 2:
				fmt.Println("+++ AutoCare +++")
				fmt.Println("1. Tambah Data Pemilik ")
				fmt.Println("2. Tampilkan Data Pemilik")
				fmt.Println("3. Ubah Data Pemilik")
				fmt.Println("4. Hapus Data Pemilik")
				fmt.Println("0. Kembali")
				fmt.Print("Pilih 0-4: ")
				fmt.Scan(&subpilihan)
				switch subpilihan{
				case 1:
					fmt.Print("Masukkan Jumlah Data Pemilik Yang Ingin Di Input: ")
					fmt.Scan(&jumlahpemilik)
					inputpemilik(&daftarpemilik,jumlahpemilik)
				case 2:
					bacapemilik(daftarpemilik,jumlahpemilik)
				case 3:
					ubahpemilik(&daftarpemilik,jumlahpemilik)
				case 4:
					hapuspemilik(&daftarpemilik,&jumlahpemilik)
				}
			case 3:
				fmt.Println("+++ AutoCare +++")
				fmt.Println("1. Tambah Data Riwayat Service ")
				fmt.Println("2. Tampilkan Data Riwayat Service")
				fmt.Println("0. Kembali")
				fmt.Print("Pilih 0-2: ")
				fmt.Scan(&subpilihan)
				switch subpilihan{
				case 1:
					fmt.Print("Masukkan Jumlah Data Pemilik Yang Ingin Di Input: ")
					fmt.Scan(&jumlahservice)
					inputservice(&daftarservice,jumlahservice)
				case 2:
					bacaservice(daftarservice, jumlahservice)
				}
			case 4:
				fmt.Println("+++ AutoCare +++")
				fmt.Println("1. Cari Data Kendaraan Menggunakan Sequential Search ")
				fmt.Println("2. Cari Data Kendaraan Menggunakan Binary Search")
				fmt.Println("0. Kembali")
				fmt.Print("Pilih 0-2: ")
				fmt.Scan(&subpilihan)
				switch subpilihan{
				case 1:
					fmt.Print("Masukkan Nomor Plat Kendaraan Yang Ingin Dicari: ")
					fmt.Scan(&target)
					idx = seqsearchkendaraan(daftarkendaraan,jumlahkendaraan,target)
					if idx == -1{
						fmt.Println("Data Tidak Ditemukan")
						}else{
							fmt.Println("Data Ditemukan:")
							tampilkendaraan(daftarkendaraan,idx)
						}
					case 2:
						fmt.Print("Masukkan Nomor Plat Kendaraan Yang Ingin Dicari :")
						fmt.Scan(&target)
						selectionsortplat(&daftarkendaraan,jumlahkendaraan)
						idx = binsearchkendaraan(daftarkendaraan,jumlahkendaraan,target)
						if idx == -1{
							fmt.Println("Data Tidak Ditemukan")
							}else{
								fmt.Println("Data Ditemukan:")
								tampilkendaraan(daftarkendaraan,idx)
							}
						}
				case 5:
						fmt.Println("+++ AutoCare +++")
						fmt.Println("1. Mengurutkan Data Kendaraan Menggunakan Selection Sort")
						fmt.Println("2. Mengurutkan Data Kendaraan Menggunakan Insertion Sort")
						fmt.Println("0. Kembali")
						fmt.Print("Pilih 0-2: ")
						fmt.Scan(&subpilihan)
						switch subpilihan{
						case 1:
							selectionsorttahun(&daftarkendaraan,jumlahkendaraan)
							fmt.Println("Data berhasil diurutkan berdasarkan tahun produksi")
							bacakendaraan(daftarkendaraan,jumlahkendaraan)
						case 2:
							insertionsortservice(&daftarkendaraan,jumlahkendaraan)
							fmt.Println("Data berhasil diurutkan berdasarkan tanggal service")
							bacakendaraan(daftarkendaraan,jumlahkendaraan)
						}
				case 6:
						fmt.Println("+++ AutoCare +++")
						fmt.Println("1. Menampilkan Statistik Jumlah Kendaraan Yang Diservice  ")
						fmt.Println("2. Menampilkan Statistik Kategori Kerusakan Yang Paling Sering Muncul")
						fmt.Println("0. Kembali")
						fmt.Print("Pilih 0-2: ")
						fmt.Scan(&subpilihan)
						switch subpilihan{
						case 1:
							statistikperbulan(daftarservice,jumlahservice)
						case 2:
							statistikkerusakan(daftarservice,jumlahservice)
						}
				case 0:
					fmt.Print("Terima Kasih")
					selesai = true
		}
	}
}
