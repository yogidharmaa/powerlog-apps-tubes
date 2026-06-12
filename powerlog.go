package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// STRUCT PERANGKAT UNTUK SIMPAN DATA PERANGKAT
type Perangkat struct {
	namaPerangkat string
	ruangan       string
	dayaW         float64
	waktu         float64
}

// SLICE UNTUK dataPerangkat DARI STRUCT PERANGKAT
var dataPerangkat []Perangkat

// FUNCTION HELPER (UNTUK EFISIENSI LOGIC PROGRAM)

// Validasi sebagai konfirmasi apakah user ingin lanjut untuk execute program atau tidak
func validasi() bool { 
	var confirm string

	for {
		fmt.Print("Apakah anda yakin? [Y/y or N/n]: ")
		fmt.Scanln(&confirm)
		if confirm == "Y" || confirm == "y" {
			return true
		} else if confirm == "N" || confirm == "n" {
			return false
		} else {
			fmt.Println("Input tidak valid, masukkan antara [Y/y or N/n]!")
		}
	}
}

// Mastiin input string engga kosong ("")
func cekInputKosong(input string) bool {
	if input == "" {
		return false
	} else {
		return true
	}
}

// Konversi Watt -> kWh
func hitungKonsumsi(perangkat Perangkat) float64 {
	return (perangkat.dayaW * perangkat.waktu) / 1000
}

// Untuk nampilin data dari slice[] Perangkat
func tampilkanPerangkat() {
	// Melakukan cek apabila isi slice = 0
	if cekAdaPerangkat() {
		return
	}

	fmt.Println("========== DATA PERANGKAT ==========")

	for i := 0; i < len(dataPerangkat); i++ {
		fmt.Println("Data ke-", i+1)
		fmt.Println("Nama Perangkat :", dataPerangkat[i].namaPerangkat)
		fmt.Println("Ruangan        :", dataPerangkat[i].ruangan)
		fmt.Println("Daya (Watt)    :", dataPerangkat[i].dayaW)
		fmt.Println("Durasi (Jam)   :", dataPerangkat[i].waktu)
		fmt.Println("Konsumsi       :", hitungKonsumsi(dataPerangkat[i]), "kWh")
		fmt.Println("------------------------------------")
	}
}

// Untuk membersihkan terminal agar output lebih bersih
func clearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" { // windows
		cmd = exec.Command("cmd", "/c", "cls")
	} else { // mac
		cmd = exec.Command("clear")
	}

	// execute command
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Function penahan agar tidak execute return otomatis
func pause() {
	fmt.Print("\nTekan Enter untuk lanjut...")
	fmt.Scanln()
}

// Untuk cek jika slice data kosong maka program tidak bisa dijalankan
func cekAdaPerangkat() bool {
	if len(dataPerangkat) == 0 {
		fmt.Print("Belum ada data perangkat\n")
		fmt.Print("Silahkan tambahkan perangkat terlebih dahulu di\n")
		fmt.Print("Menu: Kelola Perangkat -> Tambahkan Perangkat")
		pause()
		return true
	}

	return false
}

// DATA DUMMY UNTUK UJI COBA
func DataDummy() {
	dataPerangkat = []Perangkat{
		{
			namaPerangkat: "Kulkas",
			ruangan:       "Dapur",
			dayaW:         150,
			waktu:         24,
		},
		{
			namaPerangkat: "Televisi",
			ruangan:       "Ruang_Tamu",
			dayaW:         100,
			waktu:         5,
		},
		{
			namaPerangkat: "AC",
			ruangan:       "Kamar_Tidur",
			dayaW:         750,
			waktu:         8,
		},
		{
			namaPerangkat: "Laptop",
			ruangan:       "Ruang_Kerja",
			dayaW:         65,
			waktu:         6,
		},
		{
			namaPerangkat: "Dongle",
			ruangan:       "Ruang_Kerja",
			dayaW:         12,
			waktu:         2,
		},
		{
			namaPerangkat: "Proyektor",
			ruangan:       "Ruang_Kerja",
			dayaW:         25,
			waktu:         5,
		},
	}
}

// SORTING FUNCTION
func selectionSortPerangkat() {
	n := len(dataPerangkat)

	for i := 0; i < n-1; i++ {
		maxIdx := i

		for j := i + 1; j < n; j++ {
			if dataPerangkat[j].dayaW > dataPerangkat[maxIdx].dayaW {
				maxIdx = j
			}
		}

		dataPerangkat[i], dataPerangkat[maxIdx] = dataPerangkat[maxIdx], dataPerangkat[i]
	}
}

func insertionSortPerangkat() {
	n := len(dataPerangkat)

	for i := 1; i < n; i++ {
		key := dataPerangkat[i]
		j := i - 1

		for j >= 0 && dataPerangkat[j].namaPerangkat > key.namaPerangkat {
			dataPerangkat[j+1] = dataPerangkat[j]
			j = j - 1
		}
		dataPerangkat[j+1] = key
	}
}

func sortRuangan() {
	n := len(dataPerangkat)

	for i := 0; i < n-1; i++ {
		minIdx := i

		for j := i + 1; j < n; j++ {
			if dataPerangkat[j].ruangan < dataPerangkat[minIdx].ruangan {
				minIdx = j
			}
		}

		dataPerangkat[i], dataPerangkat[minIdx] = dataPerangkat[minIdx], dataPerangkat[i]
	}
}

func cariNamaPerangkat() {

	var cariberdnama string
	var ketemu bool

	clearScreen()

	fmt.Print("Masukkan nama perangkat yang dicari : ")
	fmt.Scanln(&cariberdnama)
	
	if !cekInputKosong(cariberdnama) {
		fmt.Print("Input tidak boleh kosong!")
		pause()
		return
	}

	fmt.Println("\n===== HASIL PENCARIAN =====")

	// Proses Looping buat nyari perangkat sesuai inputan "cariberdnama"
	for i := 0; i < len(dataPerangkat); i++ {

		// Membandingkan namaPerangkat di slice dengan inputan user
		if dataPerangkat[i].namaPerangkat == cariberdnama {

			fmt.Println("Nama Perangkat :", dataPerangkat[i].namaPerangkat)
			fmt.Println("Ruangan        :", dataPerangkat[i].ruangan)
			fmt.Println("Daya (Watt)    :", dataPerangkat[i].dayaW)
			fmt.Println("Durasi (Jam)   :", dataPerangkat[i].waktu)
			fmt.Println("------------------------------------")

			ketemu = true // Memberikan indikator karena data sudah ketemu
		}
	}

	// Kondisi bool default = false
	if !ketemu {
		fmt.Println("Perangkat tidak ditemukan!")
	}

	pause()
}

func cariRuangan() {

	var cariberdruangan string

	clearScreen()

	fmt.Print("Masukkan nama ruangan yang dicari : ")
	fmt.Scanln(&cariberdruangan)

	if !cekInputKosong(cariberdruangan) {
		fmt.Print("Input tidak boleh kosong!")
		pause()
		return
	}
	
	sortRuangan()

	// Start batas kiri sama batas kanan 
	kiri := 0
	kanan := len(dataPerangkat) - 1
	ketemu := -1

	for kiri <= kanan {

		tengah := (kiri + kanan) / 2

		if dataPerangkat[tengah].ruangan == cariberdruangan {
			ketemu = tengah
			break
		} else if dataPerangkat[tengah].ruangan < cariberdruangan {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	if ketemu == -1 {
		fmt.Println("\nRuangan tidak ditemukan!")
		pause()
		return
	}

	fmt.Println("\n===== HASIL PENCARIAN =====")

	for i := 0; i < len(dataPerangkat); i++ {

		if dataPerangkat[i].ruangan == cariberdruangan {

			fmt.Println("------------------------------------")
			fmt.Println("Nama Perangkat :", dataPerangkat[i].namaPerangkat)
			fmt.Println("Ruangan        :", dataPerangkat[i].ruangan)
			fmt.Println("Daya (Watt)    :", dataPerangkat[i].dayaW)
			fmt.Println("Durasi (Jam)   :", dataPerangkat[i].waktu)
		}
	}

	pause()

}

func totalKonsumsiHarian() {

	clearScreen()

	var total float64
	var banyakPerangkat int

	fmt.Println("===== TOTAL KONSUMSI ENERGI HARIAN =====")

	for i := 0; i < len(dataPerangkat); i++ {

		// Manggil fungsi konversi ke kWh dengan data perangkat
		konsumsi := hitungKonsumsi(dataPerangkat[i])
		total += konsumsi
		banyakPerangkat++
	}

	tampilkanPerangkat()

	fmt.Println("Total Perangkat:", banyakPerangkat, "Perangkat")
	fmt.Println("Total Konsumsi Harian:", total, "kWh")

	pause()
}

func sortKonsumsiTertinggi() {

	n := len(dataPerangkat)

	for i := 0; i < n-1; i++ {

		maxIdx := i

		for j := i + 1; j < n; j++ {

			if hitungKonsumsi(dataPerangkat[j]) > hitungKonsumsi(dataPerangkat[maxIdx]) {
				maxIdx = j
			}
		}

		dataPerangkat[i], dataPerangkat[maxIdx] = dataPerangkat[maxIdx], dataPerangkat[i]
	}
}

func daftarKonsumsiTertinggi() {

	clearScreen()

	sortKonsumsiTertinggi()

	fmt.Println("===== DAFTAR PERANGKAT KONSUMSI TERTINGGI =====")
	tampilkanPerangkat()

	pause()
}

func perangkatPalingBoros() {

	clearScreen()

	maxIdx := 0

	for i := 1; i < len(dataPerangkat); i++ {

		if hitungKonsumsi(dataPerangkat[i]) > hitungKonsumsi(dataPerangkat[maxIdx]) {
			maxIdx = i
		}
	}

	fmt.Println("===== PERANGKAT PALING BOROS LISTRIK =====")

	fmt.Println("Nama Perangkat :", dataPerangkat[maxIdx].namaPerangkat)
	fmt.Println("Ruangan        :", dataPerangkat[maxIdx].ruangan)
	fmt.Println("Durasi (Jam)   :", dataPerangkat[maxIdx].waktu)
	fmt.Println("Konsumsi       :", hitungKonsumsi(dataPerangkat[maxIdx]), "kWh")

	pause()
}

// FUNCTION DISPLAY MENU

func listMainMenu() {
	fmt.Println(">>> WELCOME TO POWERLOG APPS <<<")
	fmt.Println("=============== MAIN MENU ================")
	fmt.Println("1. Kelola Perangkat")
	fmt.Println("2. Cari Perangkat")
	fmt.Println("3. Statistik Konsumsi Listrik")
	fmt.Println("0. Keluar")
	fmt.Println("==========================================")
}

func MenuKelolaPerangkat() {
	fmt.Println("=============== KELOLA PERANGKAT ================")
	fmt.Println("1. Tambahkan Perangkat") 
	fmt.Println("2. Tampilkan Perangkat")
	fmt.Println("3. Ubah Detail Perangkat") 
	fmt.Println("4. Hapus Perangkat")
	fmt.Println("0. Kembali")
	fmt.Println("=================================================")
}

func MenuCariPerangkat() {
	fmt.Println("=============== CARI PERANGKAT ================")
	fmt.Println("1. Cari Berdasarkan Jenis Ruangan")
	fmt.Println("2. Cari Berdasarkan Nama Perangkat")
	fmt.Println("0. Kembali")
	fmt.Println("================================================")
}

func MenuStatistikPerangkat() {
	fmt.Println("=============== STATISTIK PERANGKAT ================")
	fmt.Println("1. Total Konsumsi Energi Harian")
	fmt.Println("2. Daftar Perangkat Konsumsi Tertinggi")
	fmt.Println("3. Perangkat Paling Boros Listrik") 
	fmt.Println("0. Kembali")
	fmt.Println("====================================================")
}

func menuUrutkanPerangkat() {
	fmt.Println("=============== TAMPILKAN PERANGKAT ===============")
	fmt.Println("1. Tampilkan Berdasarkan Konsumsi Energi Tertinggi")
	fmt.Println("2. Tampilkan Berdasarkan Abjad Perangkat (A-Z)")
	fmt.Println("0. Kembali")
	fmt.Println("=================================================")
}

// FUNCTION INTERAKSI MENU DARI MAIN MENU
func inputMainMenu() {
	var pilih string
	clearScreen()

	for {
		clearScreen()
		listMainMenu()

		pilih = ""
		fmt.Print("Silahkan pilih menu (0 - 3): ")
		fmt.Scanln(&pilih)

		switch pilih {
		case "1":
			clearScreen()
			pilihKelolaPerangkat()
		case "2":
			clearScreen()
			pilihCariPerangkat()
		case "3":
			clearScreen()
			pilihStatistikPerangkat()
		case "0":
			clearScreen()
			fmt.Print("Terima kasih telah menggunakan aplikasi Powerlog")
			return
		default:
			fmt.Print("Tolong masukkan input berupa angka 0 - 3!")
			pause()
		}
	}
}

func pilihKelolaPerangkat() {
	var pilih string

	for {

		clearScreen()
		MenuKelolaPerangkat()

		pilih = ""
		fmt.Print("Silahkan pilih menu (0 - 4): ")
		fmt.Scanln(&pilih)

		switch pilih {
		case "1":
			tambahPerangkat()
		case "2":
			pilihUrutkanPerangkat()
		case "3":
			ubahPerangkat()
		case "4":
			hapusPerangkat()
		case "0":
			return
		default:
			fmt.Print("Tolong masukkan input berupa angka 0 - 4!")
			pause()
		}
	}
}

func pilihUrutkanPerangkat() {
	var pilih string

	clearScreen()

	if cekAdaPerangkat() {
		return
	}

	for {
		clearScreen()
		menuUrutkanPerangkat()

		pilih = ""
		fmt.Print("Silahkan pilih menu (0 - 2): ")
		fmt.Scanln(&pilih)

		switch pilih {
		case "1":
			clearScreen()

			selectionSortPerangkat()
			tampilkanPerangkat()

			fmt.Print("===== DAFTAR PERANGKAT KONSUMSI DAYA TERTINGGI =====")

			pause()
		case "2":
			clearScreen()

			insertionSortPerangkat()
			tampilkanPerangkat()

			fmt.Print("===== DAFTAR PERANGKAT BERDASARKAN ABJAD =====")

			pause()
		case "0":
			return
		default:
			fmt.Print("Tolong masukkan input berupa angka 0 - 2!")
			pause()
		}
	}
}

func pilihCariPerangkat() {
	var pilih string

	if cekAdaPerangkat() {
		return
	}

	for {

		clearScreen()
		MenuCariPerangkat()

		pilih = ""

		fmt.Print("Silahkan pilih menu (0 - 2): ")
		fmt.Scanln(&pilih)

		switch pilih {
		case "1":
			cariRuangan()
		case "2":
			cariNamaPerangkat()
		case "0":
			return
		default:
			fmt.Print("Tolong masukkan input berupa angka 0 - 2!")
			pause()
		}
	}
}

func pilihStatistikPerangkat() {
	var pilih string

	if cekAdaPerangkat() {
		return
	}

	for {

		clearScreen()
		MenuStatistikPerangkat()

		pilih = ""

		fmt.Print("Silahkan pilih menu (0 - 3): ")
		fmt.Scanln(&pilih)

		switch pilih {
		case "1":
			totalKonsumsiHarian()
		case "2":
			daftarKonsumsiTertinggi()
		case "3":
			perangkatPalingBoros()
		case "0":
			return
		default:
			fmt.Print("Tolong masukkan input berupa angka 0 - 3!")
			pause()
		}
	}
}

// CRUD FEATURE
func tambahPerangkat() {

	var tambah Perangkat
	clearScreen()

	fmt.Print("===== TAMBAH PERANGKAT =====\n")

	fmt.Print("Masukkan Nama Perangkat:")
	fmt.Scanln(&tambah.namaPerangkat)
	
	if !cekInputKosong(tambah.namaPerangkat) {
		fmt.Print("Input tidak boleh kosong!")
		pause()
		return
	}

	fmt.Print("Masukkan Nama Ruangan: ")
	fmt.Scanln(&tambah.ruangan) 

	if !cekInputKosong(tambah.ruangan) {
		fmt.Print("Input tidak boleh kosong!")
		pause()
		return
	}

	// Validasi duplikat perangkat
	for i := 0; i < len(dataPerangkat); i++ {

		if tambah.namaPerangkat == dataPerangkat[i].namaPerangkat {
			if tambah.ruangan == dataPerangkat[i].ruangan {
				fmt.Print("Nama perangkat sudah ada diruangan yang sama")
				pause()
				return
			}
		}
	}

	fmt.Print("Masukkan Daya Perangkat (watt): ")
	fmt.Scanln(&tambah.dayaW)

	if tambah.dayaW <= 0 {
		fmt.Print("Daya harus lebih dari 0")
		pause()
		return
	}

	fmt.Print("Masukkan Lama Pemakaian (jam): ")
	fmt.Scanln(&tambah.waktu)

	if tambah.waktu <= 0 {
		fmt.Print("Waktu harus lebih dari 0")
		pause()
		return
	}

	if !validasi() {
		clearScreen()
		fmt.Print("Input dibatalkan")
		pause()
		return
	}

	// Output
	clearScreen()
	dataPerangkat = append(dataPerangkat, tambah)

	fmt.Print("===== INPUT TELAH DITAMBAHKAN =====\n")
	fmt.Print("===== DETAIL INPUT =====\n")
	fmt.Println("Nama Perangkat : ", tambah.namaPerangkat)
	fmt.Println("Ruangan		: ", tambah.ruangan)
	fmt.Println("Daya (watt)	: ", tambah.dayaW)
	fmt.Println("Waktu (jam)	: ", tambah.waktu)

	pause()
}

func ubahPerangkat() {

	var pilihUbah int
	var namaBaru, ruanganBaru string
	var dayaBaru, waktuBaru float64

	clearScreen()

	if cekAdaPerangkat() {
		return
	}

	tampilkanPerangkat()

	fmt.Println("===== UBAH DETAIL PERANGKAT =====")
	fmt.Print("Pilih data ke- berapa yang ingin diubah: ")

	_, err := fmt.Scanln(&pilihUbah)

	if err != nil {
		fmt.Println("Input harus berupa angka!")
		pause()
		return
	}

	idx := pilihUbah - 1

	if idx < 0 || idx >= len(dataPerangkat) {
		fmt.Println("Data tidak ditemukan!")
		pause()
		return
	}

	fmt.Println("\n===== DATA BARU =====")

	fmt.Print("Nama Perangkat : ")
	fmt.Scanln(&namaBaru)

	if !cekInputKosong(namaBaru) {
		fmt.Print("Input tidak boleh kosong!")
		pause()
		return
	}

	fmt.Print("Ruangan : ")
	fmt.Scanln(&ruanganBaru)

	if !cekInputKosong(ruanganBaru) {
		fmt.Print("Input tidak boleh kosong!")
		pause()
		return
	}

	// Cek duplikat data 
	for i := 0; i < len(dataPerangkat); i++ {

		if namaBaru == dataPerangkat[i].namaPerangkat {
			if ruanganBaru == dataPerangkat[i].ruangan {
				fmt.Print("Nama perangkat sudah ada diruangan yang sama")
				pause()
				return
			}
		}
	}

	fmt.Print("Daya (Watt) : ")
	fmt.Scanln(&dayaBaru)

	if dayaBaru <= 0 {
		fmt.Println("Daya harus lebih dari 0")
		pause()
		return
	}

	fmt.Print("Lama Pemakaian (Jam) : ")
	fmt.Scanln(&waktuBaru)

	if waktuBaru <= 0 {
		fmt.Println("Waktu harus lebih dari 0")
		pause()
		return
	}

	if !validasi() {
		fmt.Println("Perubahan dibatalkan")
		pause()
		return
	}

	dataPerangkat[idx].namaPerangkat = namaBaru
	dataPerangkat[idx].ruangan = ruanganBaru
	dataPerangkat[idx].dayaW = dayaBaru
	dataPerangkat[idx].waktu = waktuBaru

	fmt.Println("\nData berhasil diubah!")
	pause()
}

func hapusPerangkat() {

	var pilihHapus int
	clearScreen()

	if cekAdaPerangkat() {
		return
	}

	tampilkanPerangkat()

	fmt.Println("===== HAPUS PERANGKAT =====")
	fmt.Print("Pilih data ke- berapa yang ingin dihapus : ")

	_, err := fmt.Scanln(&pilihHapus)

	if err != nil {
		fmt.Println("Masukkan nomor yang valid!")
		pause()
		return
	}

	idx := pilihHapus - 1

	if idx < 0 || idx >= len(dataPerangkat) {
		fmt.Println("Data tidak ditemukan!")
		pause()
		return
	}

	if !validasi() {
		fmt.Println("Penghapusan dibatalkan")
		pause()
		return
	}

	dataPerangkat = append(dataPerangkat[:idx], dataPerangkat[idx+1:]...)

	clearScreen()
	fmt.Println("\nData perangkat berhasil dihapus!")
	pause()
}

func main() {
	DataDummy()
	inputMainMenu()
}
