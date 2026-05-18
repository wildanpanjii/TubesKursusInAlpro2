package main

import (
	"fmt"
	"os"
	"os/exec"
)

const maxPeserta = 100

type Tanggal struct {
	Hari int
	Bulan int
	Tahun int
}

type Peserta struct {
	ID, Umur int
	Nama, Email, NoHP, BidangMinat, KatalogKursus, TanggalDaftar string
	StatusAktif bool
}

var daftarPeserta [maxPeserta] Peserta
var jumlahPeserta int = 0
var inputTanggal Tanggal

func main() {
	clearScreen()
	var pilihan int
	var menu bool
	menu = true
	for menu == true {
		fmt.Println("\n------------------------------------")
		fmt.Println(" SISTEM PENDAFTARAN KURSUSIN ")
		fmt.Println("------------------------------------")
		fmt.Println("1. Tambah Peserta")
		fmt.Println("2. Tampilkan Peserta")
		fmt.Println("3. Ubah Peserta")
		fmt.Println("4. Hapus Peserta")
		fmt.Println("5. Cari Peserta")
		fmt.Println("6. Urutkan Peserta")
		fmt.Println("7. Statistik Peserta")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			tambahPeserta()
		case 2:
			tampilPeserta()
		case 3:
			ubahPeserta()
		case 4:
			hapusPeserta()
		case 5:
			menuPencarian()
		case 6:
			menuSorting()
		case 7:
			statistikPeserta()
		case 8:
			fmt.Println("Program selesai.")
			menu = false
		default:
			fmt.Println("Menu tidak tersedia.")
		}
	}
}

func tambahPeserta() {
	clearScreen()
	var i int
	var p Peserta
	var ditemukan bool
	var hasilHp, hasilEmail, hasilNama, cekstatus, hasilMinat, hasilUmur, hasilTanggal, hasilTanggalHari, hasilTanggalBulan, hasilTanggalTahun string

	fmt.Println("\n----- TAMBAH PESERTA -----")
	fmt.Println("Peringatan!!! ID harus berupa digit!!!")
	fmt.Print("Masukkan ID : ")
	fmt.Scan(&p.ID)
	for p.ID <= 0 {
		fmt.Println("ID harus berupa digit positif! Masukkan kembali")
		fmt.Print("Masukkan ID : ")
		fmt.Scan(&p.ID)
	}

	for i = 0; i < jumlahPeserta; i++ {
		if daftarPeserta[i].ID == p.ID {
			ditemukan = true
		}
	}
	if ditemukan {
		fmt.Println("ID sudah digunakan!")
		return
	}

	fmt.Print("Masukkan Nama : ")
	fmt.Scan(&p.Nama)
	hasilNama = "tidak_valid"
	for hasilNama == "tidak_valid" {
		if cekNama(p.Nama) {
			hasilNama = "valid"
		} else {
			fmt.Println("Nama tidak valid!")
			fmt.Print("Masukkan Nama : ")
			fmt.Scan(&p.Nama)
		}
	}

	hasilUmur = "tidak_valid"
	for hasilUmur == "tidak_valid" {
		fmt.Print("Masukkan Umur : ")
		fmt.Scan(&p.Umur)
		if p.Umur >= 7  {
			hasilUmur = "valid"
		} else {
			fmt.Println("Umur harus lebih dari 7!")
		}
	}

	fmt.Print("Masukkan Email : ")
	fmt.Scan(&p.Email)
	hasilEmail = "tidak_valid"
	for hasilEmail == "tidak_valid" {
		if cekEmail(p.Email) {
			hasilEmail = "valid"
		} else {
			fmt.Println("Email tidak valid!")
			fmt.Print("Masukkan Email : ")
			fmt.Scan(&p.Email)
		}
	}

	fmt.Print("Masukkan No HP : ")
	hasilHp = "tidak_valid"
	for hasilHp == "tidak_valid" {
		fmt.Scan(&p.NoHP)
		hasilHp = cekNomorHP(p.NoHP)
		if hasilHp == "tidak_valid" {
			fmt.Println("Nomor HP tidak valid! Contoh: 0812xxxx / +62812xxxx")
			fmt.Print("Masukkan No HP : ")
		}
	}

	hasilMinat = "salah"
	for hasilMinat == "salah" {
		fmt.Print("Pilihan: Seni / Sains / Olahraga / Prakarya / Sosial\nMasukkan Bidang Minat : ")
		fmt.Scan(&p.BidangMinat)
		if p.BidangMinat == "Seni" || p.BidangMinat == "Sains" || p.BidangMinat == "Olahraga" || p.BidangMinat == "Prakarya" || p.BidangMinat == "Sosial" {
			hasilMinat = "benar"
		} else {
			fmt.Println("Bidang minat tidak valid!")
		}
	}

	fmt.Print("Masukkan Katalog Kursus : ")
	fmt.Scan(&p.KatalogKursus)
	
	hasilTanggal = "tidak_valid"
	for hasilTanggal == "tidak_valid" {
		fmt.Println("Format Tanggal: DD-MM-YYYY")
		hasilTanggalHari = "tidak_valid"
		fmt.Print("Masukkan Tanggal Daftar Hari : ")
		fmt.Scan(&inputTanggal.Hari)
		for hasilTanggalHari == "tidak_valid" {
			if inputTanggal.Hari >= 1 && inputTanggal.Hari <= 31 {
				hasilTanggalHari = "valid"
			} else if inputTanggal.Hari < 1 || inputTanggal.Hari > 31 {
				fmt.Println("Hari harus antara 1 dan 31!")
				fmt.Print("Masukkan Tanggal Daftar Hari : ")
				fmt.Scan(&inputTanggal.Hari)
			}
		}
		fmt.Print("Masukkan Tanggal Daftar Bulan : ")
		fmt.Scan(&inputTanggal.Bulan)
		hasilTanggalBulan = "tidak_valid"
		for hasilTanggalBulan == "tidak_valid" {
			if inputTanggal.Bulan >= 1 && inputTanggal.Bulan <= 12 {
				hasilTanggalBulan = "valid"
			} else if inputTanggal.Bulan < 1 || inputTanggal.Bulan > 12 {
				fmt.Println("Bulan harus antara 1 dan 12!")
				fmt.Print("Masukkan Tanggal Daftar Bulan : ")
				fmt.Scan(&inputTanggal.Bulan)
			}
		}
		fmt.Print("Masukkan Tanggal Daftar Tahun : ")
		fmt.Scan(&inputTanggal.Tahun)
		hasilTanggalTahun = "tidak_valid"
		for hasilTanggalTahun == "tidak_valid" {
			if inputTanggal.Tahun >= 2020 && inputTanggal.Tahun <= 2026 {
				hasilTanggalTahun = "valid"
			} else if inputTanggal.Tahun < 2020 || inputTanggal.Tahun > 2026 {
				fmt.Println("Tahun harus antara 2020 dan 2026!")
				fmt.Print("Masukkan Tanggal Daftar Tahun : ")
				fmt.Scan(&inputTanggal.Tahun)
			}
		}
		if cekTanggal(inputTanggal.Hari, inputTanggal.Bulan, inputTanggal.Tahun) {
			p.TanggalDaftar = fmt.Sprintf("%d-%d-%d", inputTanggal.Hari, inputTanggal.Bulan, inputTanggal.Tahun)
			hasilTanggal = "valid"
		} else {
			if inputTanggal.Tahun < 2020 || inputTanggal.Tahun > 2026 {
				fmt.Println("Tahun harus antara 2020 dan 2026!")
			} else {
				fmt.Println("Tanggal tidak valid! Pastikan format DD-MM-YYYY benar dan tanggalnya valid.")
			}
		}
	}

	var status int
	fmt.Print("Status Aktif? (1=Ya / 0=Tidak): ")
	cekstatus = "salah"
	for cekstatus == "salah" {
		fmt.Scan(&status)
		if status == 1 {
			p.StatusAktif = true
			cekstatus = "benar"
		} else if status == 0 {
			p.StatusAktif = false
			cekstatus = "benar"
		} else {
			fmt.Print("Input salah! Masukkan kembali (1/0): ")
		}
	}

	daftarPeserta[jumlahPeserta] = p
	jumlahPeserta++
	fmt.Println("\nPeserta berhasil ditambahkan!")
}

func cekTanggal(hari, bulan, tahun int) bool {
	if tahun >= 2020 && tahun <= 2026 {
		switch bulan {
		case 2:
			if tahun == 2020 || tahun == 2024 {
				if hari > 0 && hari <= 29 {
					return true
				} else {
					return false
				}
			} else {
				if hari > 0 && hari <= 28 {
					return true
				} else {
					return false
				}
			}
		case 4, 6, 9, 11:
			if hari > 0 && hari <= 30 {
				return true
			} else {
				return false
			}
		case 1, 3, 5, 7, 8, 10, 12:
			if hari > 0 && hari <= 31 {
				return true
			} else {
				return false
			}
		}
		return false
	}
	return false
}

func cekNomorHP(noHP string) string {
	var panjang int
	panjang = 0
	for range noHP {
		panjang++
	}
	if panjang >= 2 && noHP[0:2] == "08" {
		if panjang == 12 || panjang == 13 {
			return noHP
		}
	} else if panjang >= 3 && noHP[0:3] == "+62" {
		if panjang == 14 || panjang == 15 {
			return noHP
		}
	}
	return "tidak_valid"
}

func cekEmail(email string) bool {
	var cekAt, cekDot bool
	var i int
	cekAt = false
	cekDot = false
	for i = 0; i < len(email); i++ {
		if email[i] == '@' {
			cekAt = true
		}
		if email[i] == '.' {
			cekDot = true
		}
	}
	if cekAt && cekDot {
		return true
	}
	return false
}

func cekNama(nama string) bool {
	var cekDigit bool
	var i int
	cekDigit = true
	for i = 0; i < len(nama); i++ {
		if nama[i] >= '0' && nama[i] <= '9' {
			cekDigit = false
		}
	}
	return cekDigit
}

func tampilPeserta() {
	var i int
	clearScreen()
	fmt.Println("\n----- DATA PESERTA -----")
	if jumlahPeserta == 0 {
		fmt.Println("Belum ada data peserta.")
		return
	}
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("%-4s%-20s%-5s%-20s%-20s%-20s%-20s%-20s%-20s\n", "ID", "Nama", "Umur", "Email", "HP", "Minat", "Kursus", "Tanggal", "Status")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------")

	for i = 0; i < jumlahPeserta; i++ {
		var status string
		if daftarPeserta[i].StatusAktif {
			status = "Aktif"
		} else {
			status = "Tidak"
		}
		fmt.Printf("%-4d%-20s%-5d%-20s%-20s%-20s%-20s%-20s%-20s\n",
			daftarPeserta[i].ID, daftarPeserta[i].Nama, daftarPeserta[i].Umur,
			daftarPeserta[i].Email, daftarPeserta[i].NoHP, daftarPeserta[i].BidangMinat,
			daftarPeserta[i].KatalogKursus, daftarPeserta[i].TanggalDaftar, status)
	}
}

func ubahPeserta() {
	clearScreen()
	var id, index int
	var hasilEmail, hasilHp, hasilNama, hasilUmur, hasilMinat string
	fmt.Println("\n----- UBAH PESERTA -----")
	fmt.Print("Masukkan ID peserta: ")
	fmt.Scan(&id)
	index = cariIndexByID(id)
	if index == -1 {
		fmt.Println("Peserta tidak ditemukan!")
		return
	}
	fmt.Println("Data ditemukan! Masukkan data baru:")
	fmt.Print("Nama Baru : ")
	fmt.Scan(&daftarPeserta[index].Nama)
	hasilNama = "tidak_valid"
	for hasilNama == "tidak_valid" {
		if cekNama(daftarPeserta[index].Nama) {
			hasilNama = "valid"
		} else {
			fmt.Println("Nama tidak valid!")
			fmt.Print("Masukkan Nama : ")
			fmt.Scan(&daftarPeserta[index].Nama)
		}
	}
	fmt.Print("Umur Baru : ")
	fmt.Scan(&daftarPeserta[index].Umur)
	hasilUmur = "tidak_valid"
	for hasilUmur == "tidak_valid" {
		if daftarPeserta[index].Umur > 7 {
			hasilUmur = "valid"
		}
		fmt.Println("Umur harus lebih dari 7!")
		fmt.Print("Masukkan Umur : ")
		fmt.Scan(&daftarPeserta[index].Umur)
	}
	fmt.Print("Email Baru : ")
	fmt.Scan(&daftarPeserta[index].Email)
	hasilEmail = "tidak_valid"
	for hasilEmail == "tidak_valid" {
		if cekEmail(daftarPeserta[index].Email) {
			hasilEmail = "valid"
		} else {
			fmt.Println("Email tidak valid!")
			fmt.Print("Masukkan Email : ")
			fmt.Scan(&daftarPeserta[index].Email)
		}
	}
	fmt.Print("No HP Baru : ")
	fmt.Scan(&daftarPeserta[index].NoHP)
	hasilHp = "tidak_valid"
	for hasilHp == "tidak_valid" {
		hasilHp = cekNomorHP(daftarPeserta[index].NoHP)
		if hasilHp == "tidak_valid" {
			fmt.Println("Nomor HP tidak valid! Contoh: 0812xxxx / +62812xxxx")
			fmt.Print("Masukkan No HP : ")
			fmt.Scan(&daftarPeserta[index].NoHP)
		}
	}
	fmt.Print("Pilihan: Seni / Sains / Olahraga / Prakarya / Sosial\nMasukkan Bidang Minat : ")
	fmt.Scan(&daftarPeserta[index].BidangMinat)
	hasilMinat = "tidak_valid"
	for hasilMinat == "tidak_valid" {
		if daftarPeserta[index].BidangMinat == "Seni" || daftarPeserta[index].BidangMinat == "Sains" || daftarPeserta[index].BidangMinat == "Olahraga" || daftarPeserta[index].BidangMinat == "Prakarya" || daftarPeserta[index].BidangMinat == "Sosial" {
			hasilMinat = "valid"
		} else {
			fmt.Println("Bidang minat tidak valid!")
			fmt.Print("Pilihan: Seni / Sains / Olahraga / Prakarya / Sosial\nMasukkan Bidang Minat : ")
			fmt.Scan(&daftarPeserta[index].BidangMinat)
		}
	}
	fmt.Print("Katalog Kursus Baru : ")
	fmt.Scan(&daftarPeserta[index].KatalogKursus)
	fmt.Println("Data berhasil diubah!")
}

func cariIndexByID(id int) int {
	var i int
	for i = 0; i < jumlahPeserta; i++ {
		if daftarPeserta[i].ID == id {
			return i
		}
	}
	return -1
}

func hapusPeserta() {
	clearScreen()
	var id, index, i int
	fmt.Println("\n----- HAPUS PESERTA -----")
	fmt.Print("Masukkan ID peserta: ")
	fmt.Scan(&id)
	index = cariIndexByID(id)
	if index == -1 {
		fmt.Println("Data tidak ditemukan!!!")
		return
	}
	for i = index; i < jumlahPeserta-1; i++ {
		daftarPeserta[i] = daftarPeserta[i+1]
	}
	jumlahPeserta--
	fmt.Println("Data berhasil dihapus!")
}

func menuPencarian() {
	clearScreen()
	var pilih int
	fmt.Println("\n----- MENU PENCARIAN -----")
	fmt.Println("1. Sequential Search Nama")
	fmt.Println("2. Sequential Search Bidang Minat")
	fmt.Println("3. Binary Search ID")
	fmt.Println("4. Kembali ke Menu Utama")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilih)
	for pilih != 4 {
		switch pilih {
		case 1:
			sequentialSearchNama()
		case 2:
			sequentialSearchMinat()
		case 3:
			binarySearchID()
		case 4:
			return
		default:
			fmt.Println("Menu tidak tersedia.")
		}
		fmt.Println("\n----- MENU PENCARIAN -----")
		fmt.Println("1. Sequential Search Nama")
		fmt.Println("2. Sequential Search Bidang Minat")
		fmt.Println("3. Binary Search ID")
		fmt.Println("4. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilih)
	}
}

func sequentialSearchNama() {
	clearScreen()
	var nama string
	var ditemukan bool
	var i int
	fmt.Println("\n----- SEARCH NAMA -----")
	fmt.Print("Masukkan nama: ")
	fmt.Scan(&nama)
	for i = 0; i < jumlahPeserta; i++ {
		if daftarPeserta[i].Nama == nama {
			fmt.Println("Data ditemukan!")
			tampilSatuPeserta(i)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Data tidak ditemukan.")
	}
}

func sequentialSearchMinat() {
	clearScreen()
	var minat string
	var ditemukan bool
	var i int
	fmt.Println("\n----- SEARCH BIDANG MINAT -----")
	fmt.Print("Masukkan bidang minat: ")
	fmt.Scan(&minat)
	for i = 0; i < jumlahPeserta; i++ {
		if daftarPeserta[i].BidangMinat == minat {
			tampilSatuPeserta(i)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Data tidak ditemukan.")
	}
}

func binarySearchID() {
	clearScreen()
	var id, low, high, mid int
	fmt.Println("\n----- BINARY SEARCH -----")
	selectionSortID()
	fmt.Print("Masukkan ID yang dicari: ")
	fmt.Scan(&id)
	low = 0
	high = jumlahPeserta - 1
	for low <= high {
		mid = (low + high) / 2
		if daftarPeserta[mid].ID == id {
			fmt.Println("Data ditemukan!")
			tampilSatuPeserta(mid)
			return
		} else if daftarPeserta[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Data tidak ditemukan!")
}

func menuSorting() {
	clearScreen()
	var pilih int
	fmt.Println("\n----- MENU SORTING -----")
	fmt.Println("1. Selection Sort ID")
	fmt.Println("2. Insertion Sort Nama")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		selectionSortID()
	case 2:
		insertionSortNama()
	default:
		fmt.Println("Menu tidak tersedia.")
	}
}

func selectionSortID() {
	clearScreen()
	var min, j, i int
	var temp Peserta
	for i = 0; i < jumlahPeserta-1; i++ {
		min = i
		for j = i + 1; j < jumlahPeserta; j++ {
			if daftarPeserta[j].ID < daftarPeserta[min].ID {
				min = j
			}
		}
		temp = daftarPeserta[i]
		daftarPeserta[i] = daftarPeserta[min]
		daftarPeserta[min] = temp
	}
	fmt.Println("Data berhasil diurutkan berdasarkan ID!")
}

func insertionSortNama() {
	clearScreen()
	var temp Peserta
	var j, i int
	for i = 1; i < jumlahPeserta; i++ {
		temp = daftarPeserta[i]
		j = i - 1
		for j >= 0 && daftarPeserta[j].Nama > temp.Nama {
			daftarPeserta[j+1] = daftarPeserta[j]
			j--
		}
		daftarPeserta[j+1] = temp
	}
	fmt.Println("Data berhasil diurutkan berdasarkan Nama!")
}

func statistikPeserta() {
	clearScreen()
	var aktif, idTerbesar, idTerkecil, minatSeni, minatSains, minatOlahraga, minatPrakarya, minatSosial int
	var i int
	if jumlahPeserta == 0 {
		fmt.Println("\nBelum ada data peserta.")
		return
	}
	idTerbesar = daftarPeserta[0].ID
	idTerkecil = daftarPeserta[0].ID
	for i = 0; i < jumlahPeserta; i++ {
		if daftarPeserta[i].StatusAktif {
			aktif++
		}
		if daftarPeserta[i].ID > idTerbesar {
			idTerbesar = daftarPeserta[i].ID
		}
		if daftarPeserta[i].ID < idTerkecil {
			idTerkecil = daftarPeserta[i].ID
		}
		if daftarPeserta[i].BidangMinat == "Seni" {
			minatSeni++
		} else if daftarPeserta[i].BidangMinat == "Sains" {
			minatSains++
		} else if daftarPeserta[i].BidangMinat == "Olahraga" {
			minatOlahraga++
		} else if daftarPeserta[i].BidangMinat == "Prakarya" {
			minatPrakarya++
		} else if daftarPeserta[i].BidangMinat == "Sosial" {
			minatSosial++
		}
	}
	fmt.Println("\n===== STATISTIK PESERTA =====")
	fmt.Println("Total Peserta Aktif :", aktif)
	fmt.Println("Jumlah Minat Seni :", minatSeni)
	fmt.Println("Jumlah Minat Sains :", minatSains)
	fmt.Println("Jumlah Minat Olahraga :", minatOlahraga)
	fmt.Println("Jumlah Minat Prakarya :", minatPrakarya)
	fmt.Println("Jumlah Minat Sosial :", minatSosial)
	fmt.Println("ID Terbesar :", idTerbesar)
	fmt.Println("ID Terkecil :", idTerkecil)

	if minatSeni > minatSains && minatSeni > minatOlahraga && minatSeni > minatPrakarya && minatSeni > minatSosial {

		fmt.Println("Bidang Terpopuler : Seni")

	} else if minatSains > minatSeni && minatSains > minatOlahraga && minatSains > minatPrakarya && minatSains > minatSosial {

		fmt.Println("Bidang Terpopuler : Sains")

	} else if minatOlahraga > minatSeni && minatOlahraga > minatSains && minatOlahraga > minatPrakarya && minatOlahraga > minatSosial {

		fmt.Println("Bidang Terpopuler : Olahraga")

	} else if minatPrakarya > minatSeni && minatPrakarya > minatSains && minatPrakarya > minatOlahraga && minatPrakarya > minatSosial {

		fmt.Println("Bidang Terpopuler : Prakarya")

	} else if minatSosial > minatSeni && minatSosial > minatSains && minatSosial > minatOlahraga && minatSosial > minatPrakarya {

		fmt.Println("Bidang Terpopuler : Sosial")

	} else {

		fmt.Println("Bidang Terpopuler : Sama Banyak")
	}
}

func tampilSatuPeserta(i int) {
	var status string
	if daftarPeserta[i].StatusAktif {
		status = "Aktif"
	} else {
		status = "Tidak"
	}
	fmt.Println("--------------------------------")
	fmt.Println("ID :", daftarPeserta[i].ID)
	fmt.Println("Nama :", daftarPeserta[i].Nama)
	fmt.Println("Umur :", daftarPeserta[i].Umur)
	fmt.Println("Email :", daftarPeserta[i].Email)
	fmt.Println("No HP :", daftarPeserta[i].NoHP)
	fmt.Println("Bidang Minat :", daftarPeserta[i].BidangMinat)
	fmt.Println("Katalog Kursus :", daftarPeserta[i].KatalogKursus)
	fmt.Println("Tanggal Daftar :", daftarPeserta[i].TanggalDaftar)
	fmt.Println("Status :", status)
	fmt.Println("--------------------------------")
}

func clearScreen() {
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}