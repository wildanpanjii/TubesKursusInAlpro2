package main

import (
	"fmt"
	"os"
	"os/exec"
	"bufio"
)

const maxPeserta = 100

type Tanggal struct {
	Hari int
	Bulan int
	Tahun int
}

type Peserta struct {
	ID, Umur int
	Nama, Email, NoHP, BidangMinat, Kursus, TanggalDaftar string
	StatusAktif bool
}

var daftarPeserta [maxPeserta] Peserta
var jumlahPeserta int = 0
var inputTanggal Tanggal
var reader *bufio.Reader
var inputString string
var valid bool

func main() {
	clearScreen()
	reader = bufio.NewReader(os.Stdin)
	var pilihan int
	var menu bool
	menu = true
	for menu == true {
		fmt.Println()
		fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
		fmt.Println("  ┃                                          ┃")
		fmt.Println("  ┃                 KURSUSIN                 ┃")
		fmt.Println("  ┃                                          ┃")
		fmt.Println("  ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
		fmt.Println("  ┃                                          ┃")
		fmt.Println("  ┃  	1⃣  Tambah Peserta                    ┃")
		fmt.Println("  ┃   	2⃣  Tampilkan Peserta                 ┃")
		fmt.Println("  ┃   	3⃣  Ubah Peserta                      ┃")
		fmt.Println("  ┃   	4⃣  Hapus Peserta                     ┃")
		fmt.Println("  ┃   	5⃣  Cari Peserta                      ┃")
		fmt.Println("  ┃   	6⃣  Urutkan Peserta                   ┃")
		fmt.Println("  ┃   	7⃣  Statistik Peserta                 ┃")
		fmt.Println("  ┃   	8⃣  Keluar                            ┃")
		fmt.Println("  ┃                                          ┃")
		fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
		fmt.Print("  Pilih menu : ")
		inputString = bacaString()
		pilihan, valid = stringKeInt(inputString)
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
			clearScreen()
			fmt.Println()
			fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
			fmt.Println("  ┃         !!!  PROGRAM SELESAI  !!!        ┃")
			fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
			menu = false
		default:
			clearScreen()
			fmt.Println("Menu tidak tersedia.")
		}
	}
}

func tambahPeserta() {
	clearScreen()
	var idBaru int
	var p Peserta
	var hasilKursus, hasilHp, hasilEmail, hasilNama, cekstatus, hasilMinat, hasilUmur, hasilTanggal, hasilTanggalHari, hasilTanggalBulan, hasilTanggalTahun string

	if jumlahPeserta >= maxPeserta {
		fmt.Println("Data peserta sudah penuh!")
		return
	}
	idBaru = generateID()

	if idBaru == -1 {
		fmt.Println("Tidak ada ID yang tersedia! Kapasitas penuh (1-100).")
		return
	}
 
	p.ID = idBaru

	fmt.Println()
	fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("  ┃            TAMBAH PESERTA                ┃")
	fmt.Println("  ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	fmt.Printf("  ┃  ID Peserta    : %-24d┃\n", idBaru)
	fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	fmt.Println()

	hasilNama = "tidak_valid"
	for hasilNama == "tidak_valid" {
		fmt.Print("  Masukkan Nama : ")
		p.Nama = bacaString()
		if cekString(p.Nama) {
			hasilNama = "valid"
		} else {
			fmt.Println("  Nama Harus Dimulai Dengan Huruf Kapital!")
		}
	}

	hasilUmur = "tidak_valid"
	for hasilUmur == "tidak_valid" {
		valid = false
		fmt.Print("  Masukkan Umur : ")
		inputString = bacaString()
		p.Umur, valid = stringKeInt(inputString)
		if !valid {
			fmt.Println("  Input umur tidak valid! Masukkan angka.")
		} else if p.Umur >= 7 && p.Umur <= 150 {
			hasilUmur = "valid"
		} else {
			fmt.Println("  Umur harus lebih dari 7 atau umur tidak masuk akal!")
		}
	}

	hasilEmail = "tidak_valid"
	for hasilEmail == "tidak_valid" {
		fmt.Print("  Masukkan Email : ")
		p.Email = bacaString()
		if cekEmail(p.Email) {
			hasilEmail = "valid"
		} else {
			fmt.Println("  Email tidak valid!")
		}
	}

	hasilHp = "tidak_valid"
	for hasilHp == "tidak_valid" {
		fmt.Print("  Masukkan No HP : ")
		p.NoHP = bacaString()
		if cekNomorHP(p.NoHP) {
			hasilHp = "valid"
		} else {
			fmt.Println("  Nomor HP tidak valid! Contoh: 081234567890 / +6281234567890")
		}
	}

	hasilMinat = "salah"
	for hasilMinat == "salah" {
		fmt.Print("  Pilihan: Seni / Sains / Olahraga / Prakarya / Sosial\n  Masukkan Bidang Minat : ")
		p.BidangMinat = bacaString()
		if p.BidangMinat == "Seni" || p.BidangMinat == "Sains" || p.BidangMinat == "Olahraga" || p.BidangMinat == "Prakarya" || p.BidangMinat == "Sosial" {
			hasilMinat = "benar"
		} else {
			fmt.Println("  Bidang minat tidak valid!")
		}
	}

	hasilKursus = "tidak_valid"
	for hasilKursus == "tidak_valid" {
		fmt.Print("  Masukkan Kursus : ")
		p.Kursus = bacaString()
		if cekString(p.Kursus) {
			hasilKursus = "valid"
		} else {
			fmt.Println("  Kursus tidak valid (harus diawali kapital)!")
		}
	}

	hasilTanggal = "tidak_valid"
	for hasilTanggal == "tidak_valid" {
		fmt.Println("  Format Tanggal : DD-MM-YYYY")
		hasilTanggalHari = "tidak_valid"
		for hasilTanggalHari == "tidak_valid" {
			fmt.Print("  Masukkan Tanggal Daftar Hari : ")
			inputString = bacaString()
			inputTanggal.Hari, valid = stringKeInt(inputString)
			if inputTanggal.Hari >= 1 && inputTanggal.Hari <= 31 {
				hasilTanggalHari = "valid"
			} else if inputTanggal.Hari < 1 || inputTanggal.Hari > 31 {
				fmt.Println("  Hari harus antara 1 dan 31!")
			}
		}
		hasilTanggalBulan = "tidak_valid"
		for hasilTanggalBulan == "tidak_valid" {
			fmt.Print("  Masukkan Tanggal Daftar Bulan : ")
			inputString = bacaString()
			inputTanggal.Bulan, valid = stringKeInt(inputString)
			if inputTanggal.Bulan >= 1 && inputTanggal.Bulan <= 12 {
				hasilTanggalBulan = "valid"
			} else if inputTanggal.Bulan < 1 || inputTanggal.Bulan > 12 {
				fmt.Println("  Bulan harus antara 1 dan 12!")
			}
		}
		hasilTanggalTahun = "tidak_valid"
		for hasilTanggalTahun == "tidak_valid" {
			fmt.Print("  Masukkan Tanggal Daftar Tahun : ")
			inputString = bacaString()
			inputTanggal.Tahun, valid = stringKeInt(inputString)
			if inputTanggal.Tahun >= 2020 && inputTanggal.Tahun <= 2026 {
				hasilTanggalTahun = "valid"
			} else if inputTanggal.Tahun < 2020 || inputTanggal.Tahun > 2026 {
				fmt.Println(  "Tahun harus antara 2020 dan 2026!")
			}
		}
		if cekTanggal(inputTanggal.Hari, inputTanggal.Bulan, inputTanggal.Tahun) {
			p.TanggalDaftar = fmt.Sprintf("%d-%d-%d", inputTanggal.Hari, inputTanggal.Bulan, inputTanggal.Tahun)
			hasilTanggal = "valid"
		} else {
			if inputTanggal.Tahun < 2020 || inputTanggal.Tahun > 2026 {
				fmt.Println("  Tahun harus antara 2020 dan 2026!")
			} else {
				fmt.Println("  Tanggal tidak valid! Pastikan format DD-MM-YYYY benar dan tanggalnya valid.")
			}
		}
	}
	var status int
	cekstatus = "salah"
	for cekstatus == "salah" {
		fmt.Print("  Status Aktif? (1=Ya / 0=Tidak): ")
		inputString = bacaString()
		status, valid = stringKeInt(inputString)
		if !valid {
			fmt.Println("  Input tidak valid! Masukkan angka 1 atau 0.")
		} else if status == 1 {
			p.StatusAktif = true
			cekstatus = "benar"
		} else if status == 0 {
			p.StatusAktif = false
			cekstatus = "benar"
		} else {
			fmt.Print("  Input salah! Masukkan kembali (1/0): ")
		}
	}

	daftarPeserta[jumlahPeserta] = p
	jumlahPeserta++
	clearScreen()
	fmt.Println("\n  Peserta berhasil ditambahkan!")
}

func bacaString() string {
	var input string
	var err   error
	var i     int

	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Printf("\n  Oh udh nih?!\n  ywdh... ")
		return ""
	}

	for i = len(input) - 1; i >= 0; i-- {
		if input[i] == '\n' || input[i] == '\r' || input[i] == ' ' {
			input = input[:i]
		} else {
			break
		}
	}

	for len(input) > 0 && input[0] == ' ' {
		input = input[1:]
	}
	return input
}

func stringKeInt(s string) (int, bool) {
	var i, hasil int
	var negatif  bool
 
	if len(s) == 0 {
		return 0, false
	}
 
	i       = 0
	negatif = false
	hasil   = 0
 
	if s[0] == '-' {
		negatif = true
		i       = 1
	}
 
	if i >= len(s) {
		return 0, false
	}
 
	for i < len(s) {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		hasil = hasil*10 + int(s[i]-'0')
		i++
	}
 
	if negatif {
		hasil = -hasil
	}
 
	return hasil, true
}

func generateID() int {
    var i, j int
    var idDipakai bool
    var idHasil int
    idHasil = -1
    for i = 1; i <= maxPeserta; i++ {
        idDipakai = false
        for j = 0; j < jumlahPeserta; j++ {
            if daftarPeserta[j].ID == i {
                idDipakai = true
            }
        }
        if !idDipakai {
            idHasil = i
            i = maxPeserta + 1
        }
    }
    return idHasil
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

func cekNomorHP(noHP string) bool {
	var panjang, i int
	panjang = 0
	for i = 0; i < len(noHP); i++ {
		panjang++
	}
	if panjang >= 2 && noHP[0:2] == "08" {
		if panjang == 12 || panjang == 13 {
			for i = 0; i < panjang; i++ {
				if noHP[i] < '0' || noHP[i] > '9' {
					return false
				}
			}
			return true
		}
	} else if panjang >= 3 && noHP[0:3] == "+62" {
		if panjang == 14 || panjang == 15 {
			for i = 1; i < panjang; i++ {
				if noHP[i] < '0' || noHP[i] > '9' {
					return false
				}
			}
			return true
		}
	}
	return false
}

func cekEmail(email string) bool {
	var i, posAt, posTitik, jumlahAt int
	posAt    = -1
	posTitik = -1
	jumlahAt = 0
	if len(email) == 0 {
		return false
	}
	for i = 0; i < len(email); i++ {
		if email[i] == '@' {
			jumlahAt++
			posAt = i
		}
		if email[i] == '.' && posAt != -1 {
			posTitik = i
		}
	}
	if jumlahAt != 1 {
		return false
	}
	if posTitik == -1 {
		return false
	}
	if posAt == 0 {
		return false
	}
	if posTitik == posAt+1 {
		return false
	}
	if posTitik == len(email)-1 {
		return false
	}
	return true
}

func cekString(nama string) bool {
	var str bool
	var i int
	str = false
	if len(nama) == 0 {
		return str
	}
	if nama[0] >= 'A' && nama[0] <= 'Z' {
		for i = 1; i < len(nama); i++ {
			if nama[i] >= 'a' && nama[i] <= 'z' {
				str = true
			}
		}
	}
	return str
}

func tampilPeserta() {
	var i int
	clearScreen()
	fmt.Println()
	fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("  ┃             DATA PESERTA                 ┃")
	fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	if jumlahPeserta == 0 {
		fmt.Println()
		fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
		fmt.Println("  ┃     !!! BELUM ADA DATA PESERTA !!!       ┃")
		fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
		return
	}
	fmt.Println("  ╭────┬────────────────────┬──────┬────────────────────────────┬──────────────────┬──────────────┬────────────────────┬────────────┬────────╮")
	fmt.Printf("  │ %-2s │ %-18s │ %-4s │ %-26s │ %-16s │ %-12s │ %-18s │ %-10s │ %-6s │\n", "ID", "Nama", "Umur", "Email", "No HP", "Minat", "Kursus", "Tanggal", "Status")
	fmt.Println("  ├────┼────────────────────┼──────┼────────────────────────────┼──────────────────┼──────────────┼────────────────────┼────────────┼────────┤")

	for i = 0; i < jumlahPeserta; i++ {
		var status string
		if daftarPeserta[i].StatusAktif {
			status = "Aktif"
		} else {
			status = "Tidak"
		}
		fmt.Printf("  │ %-2d │ %-18s │ %-4d │ %-26s │ %-16s │ %-12s │ %-18s │ %-10s │ %-6s │\n", daftarPeserta[i].ID, daftarPeserta[i].Nama, daftarPeserta[i].Umur, daftarPeserta[i].Email, daftarPeserta[i].NoHP, daftarPeserta[i].BidangMinat, daftarPeserta[i].Kursus, daftarPeserta[i].TanggalDaftar, status)
	}
	fmt.Println("  ╰────┴────────────────────┴──────┴────────────────────────────┴──────────────────┴──────────────┴────────────────────┴────────────┴────────╯")
	fmt.Printf("  Total peserta : %d\n", jumlahPeserta)
}

func ubahPeserta() {
	clearScreen()
	var id, index int
	var p Peserta
	var hasilKursus, hasilEmail, hasilHp, hasilNama, hasilUmur, hasilMinat string
	fmt.Println()
	fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("  ┃             DATA PESERTA                 ┃")
	fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	fmt.Print("  Masukkan ID peserta : ")
	inputString = bacaString()
	id, valid = stringKeInt(inputString)
	for !valid {
		fmt.Println()
		fmt.Println("  ╭──────────────────────────────────────────╮")
		fmt.Println("  │       !!! ID HARUS BERUPA ANGKA !!!      │")
		fmt.Println("  ╰──────────────────────────────────────────╯")
		fmt.Print("  Masukkan ID peserta : ")
		inputString = bacaString()
		id, valid = stringKeInt(inputString)
	}
	index = cariIndexByID(id)
	if index == -1 {
		fmt.Println()
		fmt.Println("  ╭──────────────────────────────────────────╮")
		fmt.Println("  │     !!! DATA TIDAK DITEMUKAN !!!         │")
		fmt.Println("  ╰──────────────────────────────────────────╯")
		return
	}
	fmt.Println()
	fmt.Println("  ╭──────────────────────────────────────────╮")
	fmt.Println("  │              DATA DITEMUKAN              │")
	fmt.Println("  │            MASUKKAN DATA BARU            │")
	fmt.Println("  ╰──────────────────────────────────────────╯")
	fmt.Print("  Nama Depan Baru : ")
	daftarPeserta[index].Nama = bacaString()
	hasilNama = "tidak_valid"
	for hasilNama == "tidak_valid" {
		if cekString(daftarPeserta[index].Nama) {
			hasilNama = "valid"
		} else {
			fmt.Println("  Nama Harus Dimulai Dengan Huruf Kapital!")
			fmt.Print("  Masukkan Nama Depan : ")
			daftarPeserta[index].Nama = bacaString()
		}
	}
	hasilUmur = "tidak_valid"
	for hasilUmur == "tidak_valid" {
		fmt.Print("  Masukkan Umur : ")
		inputString = bacaString()
		p.Umur, valid = stringKeInt(inputString)
		if !valid {
			fmt.Println("  Input umur tidak valid! Masukkan angka.")
		} else if p.Umur >= 7 && p.Umur <= 150 {
			daftarPeserta[index].Umur = p.Umur
			hasilUmur = "valid"
		} else {
			fmt.Println("  Umur harus lebih dari 7 atau umur tidak masuk akal!")
		}
	}
	hasilEmail = "tidak_valid"
	for hasilEmail == "tidak_valid" {
		fmt.Print("  Email Baru : ")
		daftarPeserta[index].Email = bacaString()
		if cekEmail(daftarPeserta[index].Email) {
			hasilEmail = "valid"
		} else {
			fmt.Println("  Email tidak valid!")
		}
	}
	hasilHp = "tidak_valid"
	for hasilHp == "tidak_valid" {
		fmt.Print("  No HP Baru : ")
		daftarPeserta[index].NoHP = bacaString()
		if cekNomorHP(daftarPeserta[index].NoHP) {
			hasilHp = "valid"
		} else {
			fmt.Println("  Nomor HP tidak valid! Contoh: 081234567890 / +6281234567890")
		}
	}
	hasilMinat = "tidak_valid"
	for hasilMinat == "tidak_valid" {
		fmt.Print("  Pilihan: Seni / Sains / Olahraga / Prakarya / Sosial\n  Masukkan Bidang Minat : ")
		daftarPeserta[index].BidangMinat = bacaString()
		if daftarPeserta[index].BidangMinat == "Seni" || daftarPeserta[index].BidangMinat == "Sains" || daftarPeserta[index].BidangMinat == "Olahraga" || daftarPeserta[index].BidangMinat == "Prakarya" || daftarPeserta[index].BidangMinat == "Sosial" {
			hasilMinat = "valid"
		} else {
			fmt.Println("  Bidang minat tidak valid!")
		}
	}
	hasilKursus = "tidak_valid"
	for hasilKursus == "tidak_valid" {
		fmt.Print("  Masukkan Kursus : ")
		daftarPeserta[index].Kursus = bacaString()
		if cekString(daftarPeserta[index].Kursus) {
			hasilKursus = "valid"
		} else {
			fmt.Println("  Kursus tidak valid (harus diawali kapital)!")
		}
	}
	clearScreen()
	fmt.Println()
	fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("  ┃       !!! DATA BERHASIL DIUBAH !!!       ┃")
	fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
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
	fmt.Println()
	fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("  ┃             DATA PESERTA                 ┃")
	fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	fmt.Print("  Masukkan ID peserta : ")
	inputString = bacaString()
	id, valid = stringKeInt(inputString)
	for !valid {
		fmt.Println("  Input ID tidak valid!")
		fmt.Print("  Masukkan ID peserta : ")
		inputString = bacaString()
		id, valid = stringKeInt(inputString)
	}
	index = cariIndexByID(id)
	if index == -1 {
		fmt.Println()
		fmt.Println("  ╭──────────────────────────────────────────╮")
		fmt.Println("  │     !!! DATA TIDAK DITEMUKAN !!!         │")
		fmt.Println("  ╰──────────────────────────────────────────╯")
		return
	}
	fmt.Println("  ╭────┬────────────────────┬─────┬────────────────────────────┬──────────────────┬──────────────┬────────────────────┬────────────╮")
	fmt.Printf("  │ %-2s │ %-18s │ %-3s │ %-26s │ %-16s │ %-12s │ %-18s │ %-10s │\n", "ID", "Nama", "Umr", "Email", "No HP", "Minat", "Kursus", "Tanggal")
	fmt.Println("  ├────┼────────────────────┼─────┼────────────────────────────┼──────────────────┼──────────────┼────────────────────┼────────────┤")
	fmt.Printf("  │ %-2d │ %-18s │ %-3d │ %-26s │ %-16s │ %-12s │ %-18s │ %-10s │\n", daftarPeserta[index].ID, daftarPeserta[index].Nama, daftarPeserta[index].Umur, daftarPeserta[index].Email, daftarPeserta[index].NoHP, daftarPeserta[index].BidangMinat, daftarPeserta[index].Kursus, daftarPeserta[index].TanggalDaftar)
	fmt.Println("  ╰────┴────────────────────┴─────┴────────────────────────────┴──────────────────┴──────────────┴────────────────────┴────────────╯")
	for i = index; i < jumlahPeserta-1; i++ {
		daftarPeserta[i] = daftarPeserta[i+1]
	}
	jumlahPeserta--
	fmt.Printf("  Data dengan ID %d berhasil dihapus!\n", id)
}

func menuPencarian() {
	clearScreen()
	var pilih int
	fmt.Println()
	fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("  ┃             MENU PENCARIAN                ┃")
	fmt.Println("  ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	fmt.Println("  ┃                                           ┃")
	fmt.Println("  ┃   	1⃣  Sequential Search  —  Nama         ┃")
	fmt.Println("  ┃   	2⃣  Sequential Search  —  Bidang Minat ┃")
	fmt.Println("  ┃   	3⃣  Binary Search      —  ID           ┃")
	fmt.Println("  ┃   	4⃣  Kembali ke Menu Utama              ┃")
	fmt.Println("  ┃                                           ┃")
	fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	fmt.Print("  Pilih menu : ")
	inputString = bacaString()
	pilih, valid = stringKeInt(inputString)
	for !valid {
		fmt.Println("  Input tidak valid!")
		fmt.Print("  Masukkan  pilih menu: ")
		inputString = bacaString()
		pilih, valid = stringKeInt(inputString)
	}
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
			fmt.Println("  Menu tidak tersedia.")
		}
		fmt.Println()
		fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
		fmt.Println("  ┃             MENU PENCARIAN                ┃")
		fmt.Println("  ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
		fmt.Println("  ┃                                           ┃")
		fmt.Println("  ┃   	1⃣  Sequential Search  —  Nama         ┃")
		fmt.Println("  ┃   	2⃣  Sequential Search  —  Bidang Minat ┃")
		fmt.Println("  ┃   	3⃣  Binary Search      —  ID           ┃")
		fmt.Println("  ┃   	4⃣  Kembali ke Menu Utama              ┃")
		fmt.Println("  ┃                                           ┃")
		fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
		fmt.Print("  Pilih menu : ")
		inputString = bacaString()
		pilih, valid = stringKeInt(inputString)
		for !valid {
			fmt.Println("  Input tidak valid!")
			fmt.Print("  Masukkan  pilih menu: ")
			inputString = bacaString()
			pilih, valid = stringKeInt(inputString)
		}
	}
}

func sequentialSearchNama() {
	clearScreen()
	var nama, hasilNama string
	var ditemukan bool
	var i int
	fmt.Println()
	fmt.Println("  ╭──────────────────────────────────────────╮")
	fmt.Println("  │             SEARCH NAMA                  │")
	fmt.Println("  ╰──────────────────────────────────────────╯")
	hasilNama = "tidak_valid"
	for hasilNama == "tidak_valid" {
		fmt.Print("  Masukkan Nama : ")
		nama = bacaString()
		if cekString(nama) {
			hasilNama = "valid"
		} else {
			fmt.Println("  Nama Harus Dimulai Dengan Huruf Kapital!")
		}
	}
	for i = 0; i < jumlahPeserta; i++ {
		if daftarPeserta[i].Nama == nama {
			fmt.Println()
			fmt.Println("  ╭──────────────────────────────────────────╮")
			fmt.Println("  │          !!! DATA DITEMUKAN !!!          │")
			fmt.Println("  ╰──────────────────────────────────────────╯")
			tampilSatuPeserta(i)
			ditemukan = true
		}
	}
	if !ditemukan {
		clearScreen()
		fmt.Println()
		fmt.Println("  ╭──────────────────────────────────────────╮")
		fmt.Println("  │     !!! DATA TIDAK DITEMUKAN !!!         │")
		fmt.Println("  ╰──────────────────────────────────────────╯")
	}
}

func sequentialSearchMinat() {
	clearScreen()
	var minat, hasilMinat string
	var ditemukan bool
	var i int
	fmt.Println()
	fmt.Println("  ╭──────────────────────────────────────────╮")
	fmt.Println("  │         SEARCH BIDANG MINAT              │")
	fmt.Println("  ╰──────────────────────────────────────────╯")
	hasilMinat = "salah"
	for hasilMinat == "salah" {
		fmt.Print("  Pilihan: Seni / Sains / Olahraga / Prakarya / Sosial\n  Masukkan Bidang Minat : ")
		minat = bacaString()
		if minat == "Seni" || minat == "Sains" || minat == "Olahraga" || minat == "Prakarya" || minat == "Sosial" {
			hasilMinat = "benar"
		} else {
			fmt.Println("  Bidang minat tidak valid!")
		}
	}
	for i = 0; i < jumlahPeserta; i++ {
		if daftarPeserta[i].BidangMinat == minat {
			fmt.Println()
			fmt.Println("  ╭──────────────────────────────────────────╮")
			fmt.Println("  │          !!! DATA DITEMUKAN !!!          │")
			fmt.Println("  ╰──────────────────────────────────────────╯")
			tampilSatuPeserta(i)
			ditemukan = true
		}
	}
	if !ditemukan {
		clearScreen()
		fmt.Println()
		fmt.Println("  ╭──────────────────────────────────────────╮")
		fmt.Println("  │     !!! DATA TIDAK DITEMUKAN !!!         │")
		fmt.Println("  ╰──────────────────────────────────────────╯")
	}
}

func binarySearchID() {
	clearScreen()
	var id, low, high, mid int
	fmt.Println()
	fmt.Println("  ╭──────────────────────────────────────────╮")
	fmt.Println("  │           BINARY SEARCH                  │")
	fmt.Println("  ╰──────────────────────────────────────────╯")
	selectionSortID()
	fmt.Print("  Masukkan ID yang dicari: ")
	inputString = bacaString()
	id, valid = stringKeInt(inputString)
	for !valid {
		fmt.Println("  ID tidak valid!")
		fmt.Print("  Masukkan ID yang dicari: ")
		inputString = bacaString()
		id, valid = stringKeInt(inputString)
	}
	low = 0
	high = jumlahPeserta - 1
	for low <= high {
		mid = (low + high) / 2
		if daftarPeserta[mid].ID == id {
			clearScreen()
			fmt.Println()
			fmt.Println("  ╭──────────────────────────────────────────╮")
			fmt.Println("  │          !!! DATA DITEMUKAN !!!          │")
			fmt.Println("  ╰──────────────────────────────────────────╯")
			tampilSatuPeserta(mid)
			return
		} else if daftarPeserta[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	clearScreen()
	fmt.Println()
	fmt.Println("  ╭──────────────────────────────────────────╮")
	fmt.Println("  │     !!! DATA TIDAK DITEMUKAN !!!         │")
	fmt.Println("  ╰──────────────────────────────────────────╯")
}

func menuSorting() {
	clearScreen()
	var pilih int
	fmt.Println()
	fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("  ┃             MENU SORTING                  ┃")
	fmt.Println("  ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	fmt.Println("  ┃                                           ┃")
	fmt.Println("  ┃   	1⃣  Selection Sort  —  ID              ┃")
	fmt.Println("  ┃   	2⃣  Insertion Sort  —  Nama            ┃")
	fmt.Println("  ┃                                           ┃")
	fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	fmt.Print("  Pilih menu: ")
	inputString = bacaString()
	pilih, valid = stringKeInt(inputString)
	for !valid {
		fmt.Println("  Input tidak valid!")
		fmt.Print("  Pilih menu: ")
		inputString = bacaString()
		pilih, valid = stringKeInt(inputString)
	}
	switch pilih {
	case 1:
		selectionSortID()
	case 2:
		insertionSortNama()
	default:
		fmt.Println()
		fmt.Println("  ╭──────────────────────────────────────────╮")
		fmt.Println("  │       !!! MENU TIDAK TERSEDIA !!!        │")
		fmt.Println("  ╰──────────────────────────────────────────╯")
	}
}

func selectionSortID() {
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
	fmt.Println()
	fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("  ┃         DATA BERHASIL DIURUTKAN          ┃")
	fmt.Println("  ┃        !!!   BERDASARKAN ID  !!!         ┃")
	fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
}

func insertionSortNama() {
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
	fmt.Println()
	fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("  ┃         DATA BERHASIL DIURUTKAN          ┃")
	fmt.Println("  ┃       !!!   BERDASARKAN NAMA  !!!        ┃")
	fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
}

func statistikPeserta() {
	clearScreen()
	var aktif, idTerbesar, idTerkecil, minatSeni, minatSains, minatOlahraga, minatPrakarya, minatSosial int
	var i int
	if jumlahPeserta == 0 {
		fmt.Println()
		fmt.Println("  ╭──────────────────────────────────────────╮")
		fmt.Println("  │      !!! BELUM ADA DATA PESERTA !!!      │")
		fmt.Println("  ╰──────────────────────────────────────────╯")
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
	fmt.Println()
	fmt.Println("  ╔══════════════════════════════════════════╗")
	fmt.Println("  ║           STATISTIK PESERTA              ║")
	fmt.Println("  ╠══════════════════════════════════════════╣")
	fmt.Println("  ║                                          ║")
	fmt.Printf("  ║   Total Peserta Aktif    : %-14d║\n", aktif)
	fmt.Println("  ║                                          ║")
	fmt.Println("  ╠═══════════ Bidang Minat ═════════════════╣")
	fmt.Println("  ║                                          ║")
	fmt.Printf("  ║   Seni                   : %-14d║\n", minatSeni)
	fmt.Printf("  ║   Sains                  : %-14d║\n", minatSains)
	fmt.Printf("  ║   Olahraga               : %-14d║\n", minatOlahraga)
	fmt.Printf("  ║   Prakarya               : %-14d║\n", minatPrakarya)
	fmt.Printf("  ║   Sosial                 : %-14d║\n", minatSosial)
	fmt.Println("  ║                                          ║")
	fmt.Println("  ╠══════════════════════════════════════════╣")
	fmt.Println("  ║                                          ║")
	fmt.Printf("  ║   ID Terbesar            : %-14d║\n", idTerbesar)
	fmt.Printf("  ║   ID Terkecil            : %-14d║\n", idTerkecil)
	fmt.Println("  ║                                          ║")
	fmt.Println("  ╚══════════════════════════════════════════╝")

	if minatSeni > minatSains && minatSeni > minatOlahraga && minatSeni > minatPrakarya && minatSeni > minatSosial {

		fmt.Println("  Bidang Terpopuler : Seni")

	} else if minatSains > minatSeni && minatSains > minatOlahraga && minatSains > minatPrakarya && minatSains > minatSosial {

		fmt.Println("  Bidang Terpopuler : Sains")

	} else if minatOlahraga > minatSeni && minatOlahraga > minatSains && minatOlahraga > minatPrakarya && minatOlahraga > minatSosial {

		fmt.Println("  Bidang Terpopuler : Olahraga")

	} else if minatPrakarya > minatSeni && minatPrakarya > minatSains && minatPrakarya > minatOlahraga && minatPrakarya > minatSosial {

		fmt.Println("  Bidang Terpopuler : Prakarya")

	} else if minatSosial > minatSeni && minatSosial > minatSains && minatSosial > minatOlahraga && minatSosial > minatPrakarya {

		fmt.Println("  Bidang Terpopuler : Sosial")

	} else {

		fmt.Println("  Bidang Terpopuler : Sama Banyak")
	}
}

func tampilSatuPeserta(i int) {
	var status string
	if daftarPeserta[i].StatusAktif {
		status = "Aktif"
	} else {
		status = "Tidak"
	}
	fmt.Println()
	fmt.Println("  ╭───────────────────────────────────────────────────────────╮")
	fmt.Println("  │                    DETAIL PESERTA                         │")
	fmt.Println("  ├──────────────────────────┬────────────────────────────────┤")
	fmt.Printf("  │  ID Peserta              │ %-30d │\n", daftarPeserta[i].ID)
	fmt.Println("  ├──────────────────────────┼────────────────────────────────┤")
	fmt.Printf("  │  Nama                    │ %-30s │\n", daftarPeserta[i].Nama)
	fmt.Println("  ├──────────────────────────┼────────────────────────────────┤")
	fmt.Printf("  │  Umur                    │ %-30d │\n", daftarPeserta[i].Umur)
	fmt.Println("  ├──────────────────────────┼────────────────────────────────┤")
	fmt.Printf("  │  Email                   │ %-30s │\n", daftarPeserta[i].Email)
	fmt.Println("  ├──────────────────────────┼────────────────────────────────┤")
	fmt.Printf("  │  No HP                   │ %-30s │\n", daftarPeserta[i].NoHP)
	fmt.Println("  ├──────────────────────────┼────────────────────────────────┤")
	fmt.Printf("  │  Bidang Minat            │ %-30s │\n", daftarPeserta[i].BidangMinat)
	fmt.Println("  ├──────────────────────────┼────────────────────────────────┤")
	fmt.Printf("  │  Kursus                  │ %-30s │\n", daftarPeserta[i].Kursus)
	fmt.Println("  ├──────────────────────────┼────────────────────────────────┤")
	fmt.Printf("  │  Tanggal Daftar          │ %-30s │\n", daftarPeserta[i].TanggalDaftar)
	fmt.Println("  ├──────────────────────────┼────────────────────────────────┤")
	fmt.Printf("  │  Status                  │ %-30s │\n", status)
	fmt.Println("  ╰───────────────────────────────────────────────────────────╯")
}

func clearScreen() {
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}