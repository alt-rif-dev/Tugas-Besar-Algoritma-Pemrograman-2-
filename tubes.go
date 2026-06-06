package main

import "fmt"

	const nmax = 999

	type kendaraan struct {
		platnomor, merek, model, tglserviceterakir string
		tahunproduksi int
	}
	type pemilik struct {
		nama, kontak string
		id int
	}
	type riwayatservice struct{
		plat,tanggalservice,jeniskerusakan,keterangan string
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
	// Hasbi: Ini procedure2 crud input, baca, ubah, hapus, tampil KENDARAAN (tampil khusus searching buat nampilin). 
	// Klo ada tambahan tambah aja \\ 25-05-2026
	func inputkendaraan(daftar *tabkendaraan,  n *int) {
		var k kendaraan
		if *n >= nmax {
			fmt.Println("Data Penuh")
		} else {
			fmt.Print("Plat Nomor: ")
			fmt.Scan(&k.platnomor)
			fmt.Print("Merek: ")
			fmt.Scan(&k.merek)
			fmt.Print("Model: ")
			fmt.Scan(&k.model)
			fmt.Print("Tahun Produksi: ")
			fmt.Scan(&k.tahunproduksi)
			fmt.Print("Tanggal Servis Terakhir (Tahun-Bulan-Hari): ")
			fmt.Scan(&k.tglserviceterakir)
			daftar[*n] = k
			*n++
			fmt.Println("Kendaraan berhasil ditambahkan")
			}
	}

	func bacakendaraan(daftar tabkendaraan,n int){
		var i int
		if  n == 0{
			fmt.Println("Data Kosong")
			}else{
				fmt.Println("+++ AutoCare +++")
				for i = 0; i<n; i++{
					fmt.Println("Plat Nomor:", daftar[i].platnomor)
					fmt.Println("Merek:", daftar[i].merek)
					fmt.Println("Model:", daftar[i].model)
					fmt.Println("Tahun Produksi:", daftar[i].tahunproduksi)
					fmt.Println("Tgl Service:", daftar[i].tglserviceterakir)
					fmt.Println("---")	
					}
				}
	}

	func ubahkendaraan(daftar *tabkendaraan, n int ){
		var target,baru string
		var idx,baruint int
		fmt.Print("Masukkan plat kendaraan yang ingin diubah:")
		fmt.Scan(&target)
		idx = seqsearchkendaraan(*daftar,n,target)
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
			
			fmt.Print("Masukkan model kendaraan baru ( '-' untuk skip): ")
			fmt.Scan(&baru)
			if baru != "-"{
				daftar[idx].model = baru
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
		fmt.Println("Model:",daftar[n].model)
		fmt.Println("Tahun Produksi:",daftar[n].tahunproduksi)
		fmt.Println("Tanggal Service Terakhir:",daftar[n].tglserviceterakir)
		fmt.Println("---")
	}
	// END CRUD KENDARAAN

	// CRUD PEMILIK
	// 
	func inputpemilik(daftar *tabpemilik,  n *int) {
		var k pemilik
		if *n >= nmax {
			fmt.Println("Data Penuh")
		}else{
			fmt.Print("Nama: ")
			fmt.Scan(&k.nama)
			fmt.Print("Id: ")
			fmt.Scan(&k.id)
			fmt.Print("Kontak (No Hp): ")
			fmt.Scan(&k.kontak)
			daftar[*n] = k
			*n++
			fmt.Println("Data pemilik berhasil ditambahkan")
		}
	}

	func bacapemilik(daftar tabpemilik, n int){
		var i int
		if  n == 0{
			fmt.Println("Data Kosong")
		}else{
			fmt.Println("+++ AutoCare +++")
			for i = 0; i<n; i++{
				fmt.Println("Nama:", daftar[i].nama)
				fmt.Println("Id:", daftar[i].id)
				fmt.Println("Kontak (Nomor Hp):", daftar[i].kontak)
				fmt.Println("---")
			}
		}
	}

	func ubahpemilik(daftar *tabpemilik, n int ){
		var target,baru string
		var idx,baruint int
		fmt.Print("Masukkan nama pemilik yang ingin diubah:")
		fmt.Scan(&target)
		idx = seqsearchpemilik(*daftar,n,target)
		if idx == -1{
			fmt.Println("Data tidak ditemukan")
		}else{
			fmt.Print("Masukkan nama pemilik baru ( '-' untuk skip): ")
			fmt.Scan(&baru)
			if baru != "-"{
				daftar[idx].nama = baru
			}

			fmt.Print("Masukkan id pemilik baru ( '0' untuk skip): ")
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
		var target string
		var idx,i int
		fmt.Print("Masukkan nama pemilik yang ingin dihapus: ")
		fmt.Scan(&target)
		idx = seqsearchpemilik(*daftar,*n,target)
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
	func inputservice(daftar *tabservice,n *int){
		var k riwayatservice
		if *n >= nmax{
			fmt.Println("Data Penuh")
		}else{
			fmt.Print("Plat Nomor: ")
			fmt.Scan(&k.plat)
			fmt.Print("Tanggal Service: ")
			fmt.Scan(&k.tanggalservice)
			fmt.Print("Jenis Kerusakan: ")
			fmt.Scan(&k.jeniskerusakan)
			fmt.Print("Keterangan : ")
			fmt.Scan(&k.keterangan)
			fmt.Print("Biaya: ")
			fmt.Scan(&k.biaya)
			daftar[*n]=k
			*n++
			fmt.Println("Riwayat service berhasil ditambahkan")
		}
	}

	func bacaservice(daftar tabservice,n int){
		var i int
		if n == 0{
			fmt.Println("Data Kosong")
		}else{
			fmt.Println("+++ AutoCare +++")
			for i = 0; i<n; i++{
				fmt.Println("Plat Nomor:",daftar[i].plat)
				fmt.Println("Tanggal Service:", daftar[i].tanggalservice)
				fmt.Println("Jenis Kerusakan:", daftar[i].jeniskerusakan)
				fmt.Println("Keterangan:",daftar[i].keterangan)
				fmt.Println("Biaya:",daftar[i].biaya)
			}
		}
	}
	// END CR SERVICE

	// ALGORITHM
	func selectionsorttahun(daftar *tabkendaraan, n int) {
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
	func seqsearchpemilik(daftar tabpemilik,n int, target string)int {
		var i int 
		var found bool
		found = false
		i = 0
		for i < n && !found{
			if daftar[i].nama == target{
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

	// stat & kerusakan
	func statistikperbulan(daftar tabservice,n int){
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
					inputkendaraan(&daftarkendaraan,&jumlahkendaraan)
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
					inputpemilik(&daftarpemilik,&jumlahpemilik)
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
					inputservice(&daftarservice,&jumlahservice)
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
