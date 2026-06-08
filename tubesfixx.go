package main

import (
	"bufio" // untuk membaca input string dengan spasi dan menolak input kosong
	"fmt"
	"os" // untuk membaca input dan clear screen
	"os/exec" // untuk clear screen
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
var enter string
var warning bool

func main() {
	clearScreen()
	isiDataDummy()
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
			fmt.Println()
			fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
			fmt.Println("  ┃       !!!  MENU TIDAK TERSEDIA  !!!      ┃")
			fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
			fmt.Print("Tekan Enter untuk melanjutkan...")
			fmt.Scanln(&enter)
			clearScreen()
		}
	}
}

func tambahPeserta() {
	clearScreen()
	var idBaru, status int
	var p Peserta
	var kalender bool
	var hasilKursus, hasilHp, hasilEmail, hasilNama, cekstatus, hasilMinat, hasilUmur, hasilTanggal, hasilTanggalHari, hasilTanggalBulan, hasilTanggalTahun string

	if jumlahPeserta >= maxPeserta {
		fmt.Println("Data peserta sudah penuh!")
		return
	}
	idBaru = generateID() // membuat id dengan id terkecil yang belum terpakai atau terhapus

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

	warning = false
	hasilNama = "tidak_valid"
	for hasilNama == "tidak_valid" {
		fmt.Print("  Masukkan Nama : ")
		p.Nama = bacaString() // masuk ke bufio reader
		if cekString(p.Nama) { // input valid
			if warning {
				hapusBaris(2) // hapus baris input terakhir (yang valid) dan menghapus warning
				warning = false // matikan indikator warning
				fmt.Print("  Nama : ")
				fmt.Println(p.Nama) // keterangan input yang valid
				hasilNama = "valid"
			} else {
				hapusBaris(1) // hapus input terakhir (yang valid)
				warning = false // matikan indikator warning
				fmt.Print("  Nama : ")
				fmt.Println(p.Nama) // keterangan input yang valid
				hasilNama = "valid"
			}
		} else {
			if warning { // sudah ada warning sebelumnya
				hapusBaris(2) // hapus baris input salah + baris warning lama
			} else {
				hapusBaris(1) // hanya menghapus baris input salah
			}
			fmt.Println("  [!] Nama Harus Dimulai Dengan Huruf Kapital!") // cetak warning baru atau menggantikan warning lama
			warning = true // menyalakan indikator warning
		}
	}

	hasilUmur = "tidak_valid"
	for hasilUmur == "tidak_valid" {
		fmt.Print("  Masukkan Umur : ")
		inputString = bacaString()
		p.Umur, valid = stringKeInt(inputString)
		if p.Umur < 7 || p.Umur > 150 {
			valid = false
		}
		if valid {
			if warning {
				hapusBaris(2)
				warning = false
				fmt.Print("  Umur : ")
				fmt.Println(p.Umur)
				hasilUmur = "valid"
			} else {
				hapusBaris(1)
				warning = false
				fmt.Print("  Umur : ")
				fmt.Println(p.Umur)
				hasilUmur = "valid"
			}
		} else {
			if warning {
				hapusBaris(2)
			} else {
				hapusBaris(1)
			}
			fmt.Println("  [!] Umur Harus Berupa lebih dari 7")
			warning = true
		}
	}

	hasilEmail = "tidak_valid"
	for hasilEmail == "tidak_valid" {
		fmt.Print("  Masukkan Email : ")
		p.Email = bacaString()
		if cekEmail(p.Email) {
			if warning {
				hapusBaris(2)
				warning = false
				fmt.Print("  Email : ")
				fmt.Println(p.Email)
				hasilEmail = "valid"
			} else {
				hapusBaris(1)
				warning = false
				fmt.Print("  Email : ")
				fmt.Println(p.Email)
				hasilEmail = "valid"
			}
		} else {
			if warning {
				hapusBaris(2)
			} else {
				hapusBaris(1)
			}

			fmt.Println("  [!] Email tidak valid, Contoh: nama@domain.com")
			warning = true
		}
	}

	hasilHp = "tidak_valid"
	for hasilHp == "tidak_valid" {
		fmt.Print("  Masukkan No HP : ")
		p.NoHP = bacaString()
		if cekNomorHP(p.NoHP) {
			if warning {
				hapusBaris(2)
				warning = false
				fmt.Print("  No HP : ")
				fmt.Println(p.NoHP)
				hasilHp = "valid"
			} else {
				hapusBaris(1)
				warning = false
				fmt.Print("  No HP : ")
				fmt.Println(p.NoHP)
				hasilHp = "valid"
			}
		} else {
			if warning {
				hapusBaris(2)
			} else {
				hapusBaris(1)
			}

			fmt.Println("  [!] Nomor HP tidak valid, Contoh: 081234567890 / +6281234567890")
			warning = true
		}
	}

	hasilMinat = "salah"
	for hasilMinat == "salah" {
		fmt.Print("  Pilihan: Seni / Sains / Olahraga / Prakarya / Sosial\n  Masukkan Bidang Minat : ")
		p.BidangMinat = bacaString()
		if p.BidangMinat == "Seni" || p.BidangMinat == "Sains" || p.BidangMinat == "Olahraga" || p.BidangMinat == "Prakarya" || p.BidangMinat == "Sosial" { // input valid dengan 5 pilihan
			if warning {
				hapusBaris(3) // hapus baris input terakhir (yang valid) dan baris warning lama dan menghapus menu pilihan
				warning = false // matikan indikator warning
				fmt.Print("  Bidang Minat : ")
				fmt.Println(p.BidangMinat) // keterangan input yang valid
				hasilMinat = "valid"
			} else {
				hapusBaris(2) // hapus input terakhir (yang valid) dan menghapus menu pilihan
				warning = false // matikan indikator warning
				fmt.Print("  Bidang Minat : ")
				fmt.Println(p.BidangMinat) // keterangan input yang valid
				hasilMinat = "valid"
			}
		} else {
			if warning {
				hapusBaris(3) // menghapus menu pilihan, baris input salah, dan baris warning lama
			} else {
				hapusBaris(2) // menghapus menu pilihan dan baris input salah
			}

			fmt.Println("  [!] Bidang Minat tidak valid, Contoh: Seni / Sains / Olahraga / Prakarya / Sosial") // cetak warning baru atau menggantikan warning lama
			warning = true // menyalakan indikator warning
		}
	}

	hasilKursus = "tidak_valid"
	for hasilKursus == "tidak_valid" {
		fmt.Print("  Masukkan Kursus : ")
		p.Kursus = bacaString()
		if cekString(p.Kursus) {
			if warning {
				hapusBaris(2)
				warning = false
				fmt.Print("  Kursus : ")
				fmt.Println(p.Kursus)
				hasilKursus = "valid"
			} else {
				hapusBaris(1)
				warning = false
				fmt.Print("  Kursus : ")
				fmt.Println(p.Kursus)
				hasilKursus = "valid"
			}
		} else {
			if warning {
				hapusBaris(2)
			} else {
				hapusBaris(1)
			}

			fmt.Println("  [!] Kursus Harus Dimulai Dengan Huruf Kapital!")
			warning = true
		}
	}

	hasilTanggal = "tidak_valid"
	for hasilTanggal == "tidak_valid" {
		hasilTanggalHari = "tidak_valid"
		if kalender { // tidak memenuhi subprogram cekTanggal, menginput ulanng
			for hasilTanggalHari == "tidak_valid" {
				fmt.Print("  Masukkan Tanggal : ")
				inputString = bacaString()
				inputTanggal.Hari, valid = stringKeInt(inputString)
				if inputTanggal.Hari < 1 || inputTanggal.Hari > 31 {
					valid = false
				}
				if valid {
					if warning {
						hapusBaris(3) // hapus baris input terakhir (yang valid), hapus warning, dan hapus keterangan tanggal tidak valid
						warning = false // matikan indikator warning
						fmt.Print("  Tanggal : ")
						fmt.Println(inputTanggal.Hari) // keterangan input yang valid
						hasilTanggalHari = "valid"
					} else {
						hapusBaris(2) // hapus baris input terakhir (yang valid) dan hapus keterangan tanggal tidak valid
						warning = false
						fmt.Print("  Tanggal : ")
						fmt.Println(inputTanggal.Hari) // keterangan input yang valid
						hasilTanggalHari = "valid"
					}
				} else { // input tidak valid
					if warning {
						hapusBaris(2) // menghapus input salah dan baris warning lama
					} else {
						hapusBaris(1) // menghapus input salah
					}

					fmt.Println("  [!] Tanggal Harus lebih dari 1 dan kurang dari 31!") // Cetak warning baru (menggantikan warning lama jika ada)
					warning = true
				}
			}
		} else { // input pertama sebelum cekTanggal
			for hasilTanggalHari == "tidak_valid" {
				fmt.Print("  Masukkan Tanggal : ")
				inputString = bacaString()
				inputTanggal.Hari, valid = stringKeInt(inputString)
				if inputTanggal.Hari < 1 || inputTanggal.Hari > 31 {
					valid = false
				}
				if valid { // input valid
					if warning {
						hapusBaris(2) // hapus baris input terakhir (yang valid) dan hapus warning
						warning = false // matikan indikator warning
						fmt.Print("  Tanggal : ")
						fmt.Println(inputTanggal.Hari) // keterangan dengan nilai valid
						hasilTanggalHari = "valid"
					} else {
						hapusBaris(1) // hapus input terakhir (yang valid)
						warning = false
						fmt.Print("  Tanggal : ")
						fmt.Println(inputTanggal.Hari) // keterangan dengan nilai valid
						hasilTanggalHari = "valid"
					}
				} else {
					if warning {
						hapusBaris(2) // hapus warning sebelumnya dan hapus input salah
					} else {
						hapusBaris(1) // hapus input salah
					}

					fmt.Println("  [!] Tanggal Harus lebih dari 1 dan kurang dari 31!") // Cetak warning baru (menggantikan warning lama jika ada)
					warning = true // nyalakan indikator warning
				}
			}
		}
		hasilTanggalBulan = "tidak_valid"
		for hasilTanggalBulan == "tidak_valid" {
			fmt.Print("  Masukkan Bulan : ")
			inputString = bacaString()
			inputTanggal.Bulan, valid = stringKeInt(inputString)
			if inputTanggal.Bulan < 1 || inputTanggal.Bulan > 12 {
				valid = false
			}
			if valid {
				if warning {
					hapusBaris(2)
					warning = false
					fmt.Print("  Bulan : ")
					fmt.Println(inputTanggal.Bulan)
					hasilTanggalBulan = "valid"
				} else {
					hapusBaris(1)
					warning = false
					fmt.Print("  Bulan : ")
					fmt.Println(inputTanggal.Bulan)
					hasilTanggalBulan = "valid"
				}
			} else {
				if warning {
					hapusBaris(2)
				} else {
					hapusBaris(1)
				}
				fmt.Println("  [!] Bulan Harus lebih dari 1 dan kurang dari 12!")
				warning = true
			}
		}
		hasilTanggalTahun = "tidak_valid"
		for hasilTanggalTahun == "tidak_valid" {
			fmt.Print("  Masukkan Tahun : ")
			inputString = bacaString()
			inputTanggal.Tahun, valid = stringKeInt(inputString)
			if inputTanggal.Tahun < 2020 || inputTanggal.Tahun > 2026 {
				valid = false
			}
			if valid {
				if warning {
					hapusBaris(2)
					warning = false
					fmt.Print("  Tahun : ")
					fmt.Println(inputTanggal.Tahun)
					hasilTanggalTahun = "valid"
				} else {
					hapusBaris(1)
					warning = false
					fmt.Print("  Tahun : ")
					fmt.Println(inputTanggal.Tahun)
					hasilTanggalTahun = "valid"
					}
			} else {
				if warning {
					hapusBaris(2)
				} else {
					hapusBaris(1)
				}
				fmt.Println("  [!] Tahun Harus lebih dari 2020 dan kurang dari 2026!")
				warning = true
			}
		}
		if cekTanggal(inputTanggal.Hari, inputTanggal.Bulan, inputTanggal.Tahun) {
			p.TanggalDaftar = fmt.Sprintf("%d-%d-%d", inputTanggal.Hari, inputTanggal.Bulan, inputTanggal.Tahun)
			hapusBaris(3) // hapus semua input tanggal yang sudah valid melalu cekTanggal
			fmt.Print("  Tanggal Pendafataran : ", p.TanggalDaftar, "\n")
			hasilTanggal = "valid"
		} else { // cekTanggal tidak valid
			hapusBaris(3) // hapus semua input tanggal yang sudah valid melalu cekTanggal
			fmt.Println("  Tanggal tidak valid! Pastikan format DD-MM-YYYY benar dan pastikan sesuai tahun kabisat.")
			kalender = true // menyalakan indikator kalender bahwa semua input tanggal tidak valid dan harus menginput ulang dari awal
		}
	}
	cekstatus = "salah"
	for cekstatus == "salah" {
		fmt.Print("  Masukkan Status (1=Ya / 0=Tidak) : ")
		inputString = bacaString()
		status, valid = stringKeInt(inputString)
		if valid && (status == 1 || status == 0) { // input valid dan sesuai pilihan 1 atau 0
			if warning {
				hapusBaris(2) // hapus baris input terakhir (yang valid) dan warning
				warning = false // matikan indikator warning
				if status == 1 { // input valid untuk status aktif
					p.StatusAktif = true // masuk ke struct peserta
					fmt.Print("  Status Aktif : Ya\n")
					cekstatus = "benar"
				} else if status == 0 { // input valid untuk status tidak aktif
					p.StatusAktif = false // masuk ke struct peserta
					fmt.Print("  Status Aktif : Tidak\n")
					cekstatus = "benar"
				}
			} else {
				hapusBaris(1) // hapus input terakhir (yang valid)
				warning = false // matikan indikator warning
				if status == 1 { // input valid untuk status aktif
					p.StatusAktif = true // masuk ke struct peserta
					fmt.Print("  Status Aktif : Ya\n")
					cekstatus = "benar"
				} else if status == 0 { // input valid untuk status tidak aktif
					p.StatusAktif = false // masuk ke struct peserta
					fmt.Print("  Status Aktif : Tidak\n")
					cekstatus = "benar"
				}
			}
		} else {
			if warning {
				hapusBaris(2)
			} else {
				hapusBaris(1)
			}

			fmt.Println("  [!] Input tidak valid! Masukkan angka 1 atau 0.")
			warning = true
		}
	}

	daftarPeserta[jumlahPeserta] = p // masukkan data pada struct peserta baru ke array daftarPeserta pada index jumlahPeserta
	jumlahPeserta++
	clearScreen()
	fmt.Println()
	fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("  ┃   !!! PESERTA BERHASIL DITAMBAHKAN !!!   ┃")
	fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	tampilSatuPeserta(jumlahPeserta - 1) // menampilkan data peserta yang baru saja ditambahkan dengan index jumlahPeserta - 1
	fmt.Print("\n  Tekan Enter untuk kembali ke menu utama...") // untuk kembali ke menu dan mengakhiri fungsi tambahPeserta
	fmt.Scanln(&enter)
	clearScreen() // membersihkan layar untuk masuk ke menu utama
} // fungsi tambahPeserta selesai dan otomatis return ke menu utama (func main)

func hapusBaris(n int) { // menggunakan tabel ANSI escape code untuk menghapus n baris terakhir di terminal dan beberapa baris sebelumnya jika ada warning, untuk menghindari penumpukan baris input dan warning yang membuat tampilan berantakan
	var i int
	for i = 0; i < n; i++ {
		fmt.Print("\033[1A") // naik 1 baris
		fmt.Print("\033[2K") // hapus baris tersebut lalu mengoutput dengan output yang sudah dibutuhkan
	}
}

func bacaString() string { // menggunakan bufio reader untuk membaca input string dengan spasi dan menolak input kosong
	var input string
	var err   error
	var i     int

	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Printf("\n  Oh udh nih?!\n  ywdh... ")
		return ""
	}

	i = len(input) - 1
	for i != -1 {
		if input[i] == '\n' || input[i] == '\r' || input[i] == ' ' {
			input = input[:i]
		} else {
			break
		}
		i--
	}

	for len(input) > 0 && input[0] == ' ' {
		input = input[1:]
	}
	return input
}

func stringKeInt(s string) (int, bool) { // subprogram untuk mengubah string ke integer setelah dicek menggunakan bufio atau bacaString, dengan mengembalikan nilai integer dan boolean validasi apakah string tersebut valid untuk diubah ke integer atau tidak dan apakah negatif
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

func generateID() int { // membuat ID dengan mencari ID terkecil yang kosong
    var i, j int
    var idDipakai bool
    var idHasil int
    idHasil = -1 // nilai dasar jika tidak ditemukan
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

func cekTanggal(hari, bulan, tahun int) bool { // subprogram untuk mengecek apakah tanggal yang diinput valid atau tidak dengan memperhatikan jumlah hari pada setiap bulan dan tahun kabisat, dengan batasan tahun 2020-2026
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

func cekNomorHP(noHP string) bool { // subprogram untuk mengecek apakah nomor HP yang diinput valid atau tidak dengan memperhatikan format nomor yang dimulai dengan "08" atau "+62" dan diikuti oleh 10-11 digit angka, dengan total panjang nomor HP 12-15 karakter
	var panjang, i int
	panjang = 0
	for i = 0; i < len(noHP); i++ { // untuk menghitung panjang nomor HP yang diinput
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
			for i = 1; i < panjang; i++ { // diawali dari 1 dikarenakan pada array 0 adalah '+'
				if noHP[i] < '0' || noHP[i] > '9' {
					return false
				}
			}
			return true
		}
	}
	return false
}

func cekEmail(email string) bool { // subprogram untuk mengecek apakah email yang diinput valid atau tidak dengan memperhatikan format email yang mengandung satu karakter '@' dan setidaknya satu karakter '.' setelah '@', dengan beberapa aturan tambahan seperti posisi '@' dan '.' yang tidak boleh bersebelahan atau berada di awal atau akhir email
	var i, posAt, posTitik, jumlahAt int
	posAt    = -1
	posTitik = -1
	jumlahAt = 0
	if len(email) == 0 { // jika email kosong, langsung return false
		return false
	}
	for i = 0; i < len(email); i++ {
		if email[i] == '@' { 
			jumlahAt++
			posAt = i // menyimpan posisi '@' terakhir yang ditemukan
		}
		if email[i] == '.' && posAt != -1 { // mencari posisi '.' setelah '@'
			posTitik = i // menyimpan posisi '.' terakhir yang ditemukan setelah '@'
		}
	}
	if jumlahAt != 1 { // jika '@' tidak ditemukan atau lebih dari satu, return false
		return false
	}
	if posTitik == -1 { // jika '.' tidak ditemukan setelah '@', return false
		return false
	}
	if posAt == 0 { // jika '@' berada di awal email, return false
		return false
	}
	if posTitik == posAt+1 { // jika '.' berada tepat setelah '@', return false
		return false
	}
	if posTitik == len(email) - 1 { // jika '.' berada di akhir email, return false
		return false
	}
	return true // jika semua kondisi di atas terpenuhi, email dianggap valid
}

func cekString(nama string) bool { // subprogram untuk mengecek apakah string yang diinput valid atau tidak dengan memperhatikan format string yang harus dimulai dengan huruf kapital dan diikuti oleh huruf kecil, dengan panjang maksimal 18 karakter dan tidak mengandung angka atau karakter khusus lainnya
	var str bool
	var i int
	str = false
	if len(nama) == 0 {
		return str
	}
	if nama[0] >= 'A' && nama[0] <= 'Z' {
		for i = 1; i < len(nama); i++ {
			if nama[i] >= 'a' && nama[i] <= 'z' || nama[i] == 'A' && nama[i] == 'Z'{
				str = true
			}
		}
	}
	if len(nama) > 18 {
		fmt.Println("  Nama terlalu panjang! Silahkan input ulang")
		str = false
	}
	return str
}

func tampilPeserta() {
	var status string
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
		if daftarPeserta[i].StatusAktif { // menentukan output pada status melalui  cek boolean pada truct peserta
			status = "Aktif"
		} else {
			status = "Tidak"
		}
		fmt.Printf("  │ %-2d │ %-18s │ %-4d │ %-26s │ %-16s │ %-12s │ %-18s │ %-10s │ %-6s │\n", daftarPeserta[i].ID, daftarPeserta[i].Nama, daftarPeserta[i].Umur, daftarPeserta[i].Email, daftarPeserta[i].NoHP, daftarPeserta[i].BidangMinat, daftarPeserta[i].Kursus, daftarPeserta[i].TanggalDaftar, status)
	}
	fmt.Println("  ╰────┴────────────────────┴──────┴────────────────────────────┴──────────────────┴──────────────┴────────────────────┴────────────┴────────╯")
	fmt.Printf("  Total peserta : %d\n", jumlahPeserta)
	fmt.Print("\n  Tekan Enter untuk kembali ke menu utama...")
	fmt.Scanln(&enter)
	clearScreen()
}

func ubahPeserta() { // subprogram untuk mengubah data peserta berdasarkan ID yang diinput
	clearScreen()
	var id, index, status int
	var hasilStatus, hasilKursus, hasilEmail, hasilHp, hasilNama, hasilUmur, hasilMinat string
	fmt.Println()
	fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("  ┃            UBAH DATA PESERTA             ┃")
	fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	fmt.Print("  Masukkan ID peserta : ")
	inputString = bacaString()
	id, valid = stringKeInt(inputString)
	for !valid {
		if valid { // input valid
			if warning {
				hapusBaris(4) // menghapus baris input terakhir (yang valid) dan tabel warning
			} else {
				hapusBaris(1) // hanya menghapus baris input terakhir (yang valid)
			}
		} else {
			if warning {
				hapusBaris(4) // menghapus baris input salah dan tabel warning lama
			} else {
				hapusBaris(1) // hanya menghapus baris input salah
			}

			// Cetak warning baru (menggantikan warning lama jika ada)
			fmt.Println("  ╭──────────────────────────────────────────╮")
			fmt.Println("  │           !!! ID TIDAK VALID !!!         │")
			fmt.Println("  ╰──────────────────────────────────────────╯")
			warning = true // menyalakan indikator warning
			fmt.Print("  Masukkan ID peserta : ") // untuk menginput ulang ID peserta
			inputString = bacaString()
			id, valid = stringKeInt(inputString)
		}
	}
	index = cariIndexByID(id)
	if index == -1 {
		if warning {
			hapusBaris(4)
		} else {
			hapusBaris(1)
		}
		fmt.Println("  ╭──────────────────────────────────────────╮")
		fmt.Println("  │       !!! DATA TIDAK DITEMUKAN !!!       │")
		fmt.Println("  ╰──────────────────────────────────────────╯")
		fmt.Print("\n  Tekan Enter untuk kembali ke menu utama...")
		fmt.Scanln(&enter)
		clearScreen()
		return
	}
	if warning {
		hapusBaris(4)
	} else {
		hapusBaris(1)
	}
	tampilSatuPeserta(index)
	fmt.Println()
	fmt.Println("  ╭──────────────────────────────────────────╮")
	fmt.Println("  │              DATA DITEMUKAN              │")
	fmt.Println("  │            MASUKKAN DATA BARU            │")
	fmt.Println("  ╰──────────────────────────────────────────╯")
	warning = false
	hasilNama = "tidak_valid"
	for hasilNama == "tidak_valid" {
		fmt.Print("  Masukkan Nama Baru : ")
		daftarPeserta[index].Nama = bacaString()
		if cekString(daftarPeserta[index].Nama) {
			if warning {
				hapusBaris(2) // hapus baris input terakhir (yang valid) dan baris warning
				warning = false
				fmt.Print("  Nama : ")
				fmt.Println(daftarPeserta[index].Nama) // keterangan dengan nilai valid
				hasilNama = "valid"
			} else {
				hapusBaris(1) // hapus baris input terakhir (yang valid)
				warning = false
				fmt.Print("  Nama : ")
				fmt.Println(daftarPeserta[index].Nama) // keterangan dengan nilai valid
				hasilNama = "valid"
			}
		} else {
			if warning {
				hapusBaris(2)
			} else {
				hapusBaris(1)
			}

			fmt.Println("  [!] Nama Harus Dimulai Dengan Huruf Kapital!")
			warning = true
		}
	}
	hasilUmur = "tidak_valid"
	for hasilUmur == "tidak_valid" {
		fmt.Print("  Masukkan Umur : ")
		inputString = bacaString()
		daftarPeserta[index].Umur, valid = stringKeInt(inputString)
		if daftarPeserta[index].Umur < 7 || daftarPeserta[index].Umur > 150 {
			valid = false
		}
		if valid {
			if warning {
				hapusBaris(2)
				warning = false
				fmt.Print("  Umur : ")
				fmt.Println(daftarPeserta[index].Umur)
				hasilUmur = "valid"
			} else {
				hapusBaris(1)
				warning = false
				fmt.Print("  Umur : ")
				fmt.Println(daftarPeserta[index].Umur)
				hasilUmur = "valid"
			}
		} else {
			if warning {
				hapusBaris(2)
			} else {
				hapusBaris(1)
			}
			fmt.Println("  [!] Umur Harus Berupa Angka lebih dari 7")
			warning = true
		}
	}
	hasilEmail = "tidak_valid"
	for hasilEmail == "tidak_valid" {
		fmt.Print("  Masukkan Email : ")
		daftarPeserta[index].Email = bacaString()
		if cekEmail(daftarPeserta[index].Email) {
			if warning {
				hapusBaris(2)
				warning = false
				fmt.Print("  Email : ")
				fmt.Println(daftarPeserta[index].Email)
				hasilEmail = "valid"
			} else {
				hapusBaris(1)
				warning = false
				fmt.Print("  Email : ")
				fmt.Println(daftarPeserta[index].Email)
				hasilEmail = "valid"
			}
		} else {
			if warning {
				hapusBaris(2)
			} else {
				hapusBaris(1)
			}

			fmt.Println("  [!] Email tidak valid, Contoh: nama@domain.com")
			warning = true
		}
	}
	hasilHp = "tidak_valid"
	for hasilHp == "tidak_valid" {
		fmt.Print("  Masukkan No HP : ")
		daftarPeserta[index].NoHP = bacaString()
		if cekNomorHP(daftarPeserta[index].NoHP) {
			if warning {
				hapusBaris(2)
				warning = false
				fmt.Print("  No HP : ")
				fmt.Println(daftarPeserta[index].NoHP)
				hasilHp = "valid"
			} else {
				hapusBaris(1)
				warning = false
				fmt.Print("  No HP : ")
				fmt.Println(daftarPeserta[index].NoHP)
				hasilHp = "valid"
			}
		} else {
			if warning {
				hapusBaris(2)
			} else {
				hapusBaris(1)
			}
			fmt.Println("  [!] Nomor HP tidak valid, Contoh: 081234567890 / +6281234567890")
			warning = true
		}
	}
	hasilMinat = "tidak_valid"
	for hasilMinat == "tidak_valid" {
		fmt.Print("  Pilihan: Seni / Sains / Olahraga / Prakarya / Sosial\n  Masukkan Bidang Minat : ")
		daftarPeserta[index].BidangMinat = bacaString()
		if daftarPeserta[index].BidangMinat == "Seni" || daftarPeserta[index].BidangMinat == "Sains" || daftarPeserta[index].BidangMinat == "Olahraga" || daftarPeserta[index].BidangMinat == "Prakarya" || daftarPeserta[index].BidangMinat == "Sosial" {
			if warning {
				hapusBaris(3)
				warning = false
				fmt.Print("  Bidang Minat : ")
				fmt.Println(daftarPeserta[index].BidangMinat)
				hasilMinat = "valid"
			} else {
				hapusBaris(2)
				warning = false
				fmt.Print("  Bidang Minat : ")
				fmt.Println(daftarPeserta[index].BidangMinat)
				hasilMinat = "valid"
			}
		} else {
			if warning {
				hapusBaris(3)
			} else {
				hapusBaris(2)
			}
			fmt.Println("  [!] Bidang Minat tidak valid, Contoh: Seni / Sains / Olahraga / Prakarya / Sosial")
			warning = true
		}
	}
	hasilKursus = "tidak_valid"
	for hasilKursus == "tidak_valid" {
		fmt.Print("  Masukkan Kursus : ")
		daftarPeserta[index].Kursus = bacaString()
		if cekString(daftarPeserta[index].Kursus) {
			if warning {
				hapusBaris(2)
				warning = false
				fmt.Print("  Kursus : ")
				fmt.Println(daftarPeserta[index].Kursus)
				hasilKursus = "valid"
			} else {
				hapusBaris(1)
				warning = false
				fmt.Print("  Kursus : ")
				fmt.Println(daftarPeserta[index].Kursus)
				hasilKursus = "valid"
			}
		} else {
			if warning {
				hapusBaris(2)
			} else {
				hapusBaris(1)
			}
			fmt.Println("  [!] Kursus Harus Dimulai Dengan Huruf Kapital!")
			warning = true
		}
	}
	hasilStatus = "salah"
	for hasilStatus == "salah" {
		fmt.Print("  Masukkan Status (1=Ya / 0=Tidak) : ")
		inputString = bacaString()
		status, valid = stringKeInt(inputString)
		if valid && (status == 1 || status == 0) {
			if warning {
				hapusBaris(2)
				warning = false
				if status == 1 {
					daftarPeserta[index].StatusAktif = true
					fmt.Print("  Status Aktif : Ya\n")
					hasilStatus = "benar"
				} else if status == 0 {
					daftarPeserta[index].StatusAktif = false
					fmt.Print("  Status Aktif : Tidak\n")
					hasilStatus = "benar"
				}
			} else {
				hapusBaris(1)
				warning = false
				if status == 1 {
					daftarPeserta[index].StatusAktif = true
					fmt.Print("  Status Aktif : Ya\n")
					hasilStatus = "benar"
				} else if status == 0 {
					daftarPeserta[index].StatusAktif = false
					fmt.Print("  Status Aktif : Tidak\n")
					hasilStatus = "benar"
				}
			}
		} else {
			if warning {
				hapusBaris(2)
			} else {
				hapusBaris(1)
			}
			fmt.Println("  [!] Input tidak valid! Masukkan angka 1 atau 0.")
			warning = true
		}
	}
	clearScreen()
	fmt.Println()
	fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("  ┃       !!! DATA BERHASIL DIUBAH !!!       ┃")
	fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	tampilSatuPeserta(index)
	fmt.Print("\n  Tekan Enter untuk kembali ke menu utama...")
	fmt.Scanln(&enter)
	clearScreen()
}

func cariIndexByID(id int) int { // subprogram untuk mencari index peserta pada array daftarPeserta berdasarkan ID yang diinput, dengan mengembalikan nilai index jika ditemukan dan -1 jika tidak ditemukan
	var i int
	for i = 0; i < jumlahPeserta; i++ {
		if daftarPeserta[i].ID == id {
			return i
		}
	}
	return -1
}

func hapusPeserta() { // subprogram untuk menghapus data peserta berdasarkan ID yang diinput, dengan menggeser data peserta setelah index yang dihapus ke kiri untuk menutup celah pada array daftarPeserta
	clearScreen()
	var id, index, i int
	fmt.Println()
	fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("  ┃            HAPUS DATA PESERTA            ┃")
	fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	fmt.Print("  Masukkan ID peserta : ")
	inputString = bacaString()
	id, valid = stringKeInt(inputString)
	for !valid {
		if valid {
			if warning {
				hapusBaris(4)
			} else {
				hapusBaris(1)
			}
		} else {
			if warning {
				hapusBaris(4)
			} else {
				hapusBaris(1)
			}
			fmt.Println("  ╭──────────────────────────────────────────╮")
			fmt.Println("  │           !!! ID TIDAK VALID !!!         │")
			fmt.Println("  ╰──────────────────────────────────────────╯")
			warning = true
			fmt.Print("  Masukkan ID peserta : ")
			inputString = bacaString()
			id, valid = stringKeInt(inputString)
		}
	}
	index = cariIndexByID(id)
	if index == -1 {
		if warning {
			hapusBaris(4)
		} else {
			hapusBaris(1)
		}
		fmt.Println("  ╭──────────────────────────────────────────╮")
		fmt.Println("  │     !!! DATA TIDAK DITEMUKAN !!!         │")
		fmt.Println("  ╰──────────────────────────────────────────╯")
		fmt.Print("\n  Tekan Enter untuk kembali ke menu utama...")
		fmt.Scanln(&enter)
		clearScreen()
		return
	}
	if warning {
		hapusBaris(4)
	} else {
		hapusBaris(1)
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
	fmt.Print("\n  Tekan Enter untuk kembali ke menu utama...")
	fmt.Scanln(&enter)
	clearScreen()
}

func menuPencarian() { // subprogram untuk menampilkan menu pencarian dan memproses pilihan pencarian berdasarkan nama, bidang minat, atau ID dengan menggunakan sequential search untuk nama dan bidang minat, serta binary search untuk ID, dengan validasi input pilihan menu dan penanganan kasus data tidak ditemukan
	clearScreen()
	var pilih int
	var cari bool
	cari = true
	for cari {
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
			if valid {
				if warning {
					hapusBaris(4)
				} else {
					hapusBaris(1)
				}
			} else {
				if warning {
					hapusBaris(4)
				} else {
					hapusBaris(1)
				}
				// Cetak warning baru (menggantikan warning lama jika ada)
				fmt.Println("  ╭──────────────────────────────────────────╮")
				fmt.Println("  │        !!! INPUT TIDAK VALID !!!         │")
				fmt.Println("  ╰──────────────────────────────────────────╯")
				warning = true
				fmt.Print("  Pilih Menu : ")
				inputString = bacaString()
				pilih, valid = stringKeInt(inputString)
			}
		}
		switch pilih {
		case 1:
			sequentialSearchNama()
		case 2:
			sequentialSearchMinat()
		case 3:
			binarySearchID()
		case 4:
			clearScreen()
			cari = false
		default:
			clearScreen()
			fmt.Println("  ╭──────────────────────────────────────────╮")
			fmt.Println("  │        !!! MENU TIDAK TERSEDIA !!!       │")
			fmt.Println("  ╰──────────────────────────────────────────╯")
			fmt.Print("\n  Tekan Enter untuk kembali ke menu pencarian...")
			fmt.Scanln(&enter)
			clearScreen()
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
	fmt.Println("  │               SEARCH NAMA                │")
	fmt.Println("  ╰──────────────────────────────────────────╯")
	hasilNama = "tidak_valid"
	for hasilNama == "tidak_valid" {
		fmt.Print("  Masukkan Nama : ")
		nama = bacaString()
		if cekString(nama) { // input valid
			if warning {
				hapusBaris(2)
				warning = false
				fmt.Print("  Nama : ")
				fmt.Println(nama) // keterangan dengan nilai valid
				hasilNama = "valid"
			} else {
				hapusBaris(1)
				warning = false
				fmt.Print("  Nama : ")
				fmt.Println(nama) // keterangan dengan nilai valid
				hasilNama = "valid"
			}
		} else {
			if warning {
				hapusBaris(2)
			} else {
				hapusBaris(1)
			}
			fmt.Println("  [!] Nama Harus Dimulai Dengan Huruf Kapital!")
			warning = true
		}
	}
	for i = 0; i < jumlahPeserta; i++ {
		if daftarPeserta[i].Nama == nama { // nama harus sama persis dengan data yang ada
			if warning {
				hapusBaris(4)
			} else {
				hapusBaris(1)
			}
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
	fmt.Print("\n  Tekan Enter untuk mencari data lain...")
	fmt.Scanln(&enter)
	clearScreen()
} // prosedur selesai dan kembali ke menu search

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
			if warning {
				hapusBaris(3)
				warning = false
				fmt.Print("  Bidang Minat : ")
				fmt.Println(minat)
				hasilMinat = "valid"
			} else {
				hapusBaris(2)
				warning = false
				fmt.Print("  Bidang Minat : ")
				fmt.Println(minat)
				hasilMinat = "valid"
			}
		} else {
			if warning {
				hapusBaris(3)
			} else {
				hapusBaris(2)
			}
			fmt.Println("  [!] Bidang Minat tidak valid, Contoh: Seni / Sains / Olahraga / Prakarya / Sosial")
			warning = true
		}
	}
	for i = 0; i < jumlahPeserta; i++ {
		if daftarPeserta[i].BidangMinat == minat {
			if warning {
				hapusBaris(4)
			} else {
				hapusBaris(1)
			}
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
	fmt.Print("\n  Tekan Enter untuk mencari data lain...")
	fmt.Scanln(&enter)
	clearScreen()
}

func binarySearchID() {
	clearScreen()
	var id, low, high, mid int
	fmt.Println()
	fmt.Println("  ╭──────────────────────────────────────────╮")
	fmt.Println("  │           BINARY SEARCH                  │")
	fmt.Println("  ╰──────────────────────────────────────────╯")
	selectionSortIDAscending()
	fmt.Println()
	fmt.Println("  ╭──────────────────────────────────────────╮")
	fmt.Println("  │           BINARY SEARCH                  │")
	fmt.Println("  ╰──────────────────────────────────────────╯")
	fmt.Print("  Masukkan ID yang dicari: ")
	inputString = bacaString()
	id, valid = stringKeInt(inputString)
	for !valid {
		if valid {
			if warning {
				hapusBaris(4)
			} else {
				hapusBaris(1)
			}
		} else {
			if warning {
				hapusBaris(4)
			} else {
				hapusBaris(1)
			}
			fmt.Println("  ╭──────────────────────────────────────────╮")
			fmt.Println("  │           !!! ID TIDAK VALID !!!         │")
			fmt.Println("  ╰──────────────────────────────────────────╯")
			warning = true
			fmt.Print("  Masukkan ID peserta : ")
			inputString = bacaString()
			id, valid = stringKeInt(inputString)
		}
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
			fmt.Print("\n  Tekan Enter untuk kembali ke menu pencarian...")
			fmt.Scanln(&enter)
			clearScreen()
			return // keluar dari prosedur jika data ditemukan
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
	fmt.Print("\n  Tekan Enter untuk kembali ke menu pencarian...")
	fmt.Scanln(&enter)
	clearScreen()
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
		if valid {
			if warning {
				hapusBaris(4)
			} else {
				hapusBaris(1)
			}
		} else {
			if warning {
				hapusBaris(4)
			} else {
				hapusBaris(1)
			}
			fmt.Println("  ╭──────────────────────────────────────────╮")
			fmt.Println("  │        !!! INPUT TIDAK VALID !!!         │")
			fmt.Println("  ╰──────────────────────────────────────────╯")
			warning = true
			fmt.Print("  Pilih Menu : ")
			inputString = bacaString()
			pilih, valid = stringKeInt(inputString)
		}
	}
	switch pilih {
	case 1:
		menuSortingID()
	case 2:
		menuSortingNama()
	default:
		fmt.Println()
		fmt.Println("  ╭──────────────────────────────────────────╮")
		fmt.Println("  │       !!! MENU TIDAK TERSEDIA !!!        │")
		fmt.Println("  ╰──────────────────────────────────────────╯")
		fmt.Print("\n  Tekan Enter untuk kembali ke menu utama...")
		fmt.Scanln(&enter)
		clearScreen()
	}
}

func menuSortingID() {
	clearScreen()
	var pilih int
	fmt.Println()
	fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("  ┃             MENU SORTING ID               ┃")
	fmt.Println("  ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	fmt.Println("  ┃                                           ┃")
	fmt.Println("  ┃   	1⃣  Ascending                          ┃")
	fmt.Println("  ┃   	2⃣  Descending                         ┃")
	fmt.Println("  ┃                                           ┃")
	fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	fmt.Print("  Pilih menu: ")
	inputString = bacaString()
	pilih, valid = stringKeInt(inputString)
	for !valid {
		if valid {
			if warning {
				hapusBaris(4)
			} else {
				hapusBaris(1)
			}
		} else {
			if warning {
				hapusBaris(4)
			} else {
				hapusBaris(1)
			}
			fmt.Println("  ╭──────────────────────────────────────────╮")
			fmt.Println("  │        !!! INPUT TIDAK VALID !!!         │")
			fmt.Println("  ╰──────────────────────────────────────────╯")
			warning = true
			fmt.Print("  Pilih Menu : ")
			inputString = bacaString()
			pilih, valid = stringKeInt(inputString)
		}
	}
	switch pilih {
	case 1:
		selectionSortIDAscending()
	case 2:
		selectionSortIDDescending()
	default:
		fmt.Println()
		fmt.Println("  ╭──────────────────────────────────────────╮")
		fmt.Println("  │       !!! MENU TIDAK TERSEDIA !!!        │")
		fmt.Println("  ╰──────────────────────────────────────────╯")
		fmt.Print("\n  Tekan Enter untuk kembali ke menu utama...")
		fmt.Scanln(&enter)
		clearScreen()
	}
}

func menuSortingNama() {
	clearScreen()
	var pilih int
	fmt.Println()
	fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("  ┃             MENU SORTING NAMA             ┃")
	fmt.Println("  ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	fmt.Println("  ┃                                           ┃")
	fmt.Println("  ┃   	1⃣  Ascending                          ┃")
	fmt.Println("  ┃   	2⃣  Descending                         ┃")
	fmt.Println("  ┃                                           ┃")
	fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	fmt.Print("  Pilih menu: ")
	inputString = bacaString()
	pilih, valid = stringKeInt(inputString)
	for !valid {
		if valid {
			if warning {
				hapusBaris(4)
			} else {
				hapusBaris(1)
			}
		} else {
			if warning {
				hapusBaris(4)
			} else {
				hapusBaris(1)
			}
			fmt.Println("  ╭──────────────────────────────────────────╮")
			fmt.Println("  │        !!! INPUT TIDAK VALID !!!         │")
			fmt.Println("  ╰──────────────────────────────────────────╯")
			warning = true
			fmt.Print("  Pilih Menu : ")
			inputString = bacaString()
			pilih, valid = stringKeInt(inputString)
		}
	}
	switch pilih {
	case 1:
		selectionSortIDAscending()
	case 2:
		selectionSortIDDescending()
	default:
		fmt.Println()
		fmt.Println("  ╭──────────────────────────────────────────╮")
		fmt.Println("  │       !!! MENU TIDAK TERSEDIA !!!        │")
		fmt.Println("  ╰──────────────────────────────────────────╯")
		fmt.Print("\n  Tekan Enter untuk kembali ke menu utama...")
		fmt.Scanln(&enter)
		clearScreen()
	}
}

func selectionSortIDAscending() {
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
	fmt.Println("  ┃                 ASCENDING                ┃")
	fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	fmt.Print("\n  Tekan Enter untuk kembali ke menu utama...")
	fmt.Scanln(&enter)
	clearScreen()
}

func selectionSortIDDescending() {
	var min, j, i int
	var temp Peserta
	for i = 0; i < jumlahPeserta-1; i++ {
		min = i
		for j = i + 1; j < jumlahPeserta; j++ {
			if daftarPeserta[j].ID > daftarPeserta[min].ID {
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
	fmt.Println("  ┃                DESCENDING                ┃")
	fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	fmt.Print("\n  Tekan Enter untuk kembali ke menu utama...")
	fmt.Scanln(&enter)
	clearScreen()
}

func insertionSortNamaAscending() {
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
	fmt.Println("  ┃                 ASCENDING                ┃")
	fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	fmt.Print("\n  Tekan Enter untuk kembali ke menu utama...")
	fmt.Scanln(&enter)
	clearScreen()
}

func insertionSortNamaDescending() {
	var temp Peserta
	var j, i int
	for i = 1; i < jumlahPeserta; i++ {
		temp = daftarPeserta[i]
		j = i - 1
		for j >= 0 && daftarPeserta[j].Nama < temp.Nama {
			daftarPeserta[j+1] = daftarPeserta[j]
			j--
		}
		daftarPeserta[j+1] = temp
	}
	fmt.Println()
	fmt.Println("  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("  ┃         DATA BERHASIL DIURUTKAN          ┃")
	fmt.Println("  ┃       !!!   BERDASARKAN NAMA  !!!        ┃")
	fmt.Println("  ┃                DESCENDING                ┃")
	fmt.Println("  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	fmt.Print("\n  Tekan Enter untuk kembali ke menu utama...")
	fmt.Scanln(&enter)
	clearScreen()
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
	fmt.Print("\n  Tekan Enter untuk kembali ke menu utama...")
	fmt.Scanln(&enter)
	clearScreen()
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
	fmt.Println("  ╰──────────────────────────┴────────────────────────────────╯")
}

func clearScreen() {
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}