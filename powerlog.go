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

func validasi() bool{ // Penyempurnaan lagi (baru dipasang ditambahkan perangkat)
	var confirm string

	for {
		clearScreen()

		fmt.Print("Apakah anda yakin? [Y/N]\n")
		fmt.Scanln(&confirm)
		if confirm == "N" || confirm == "n" {
			return false
		} else if confirm == "Y" || confirm == "y" {
			return true
		} else {
			fmt.Print("input tidak valid")
			pause()
		}
	}
}

func hitungKonsumsi(perangkat Perangkat) float64 {
	// Konversi Watt -> kWh
	return (perangkat.dayaW * perangkat.waktu) / 1000
}

func tampilkanPerangkat() {
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

func clearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" { // windows 
		cmd = exec.Command("cmd", "/c", "cls")
	} else { // mac
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func pause() {
	fmt.Print("\nTekan [Enter] untuk lanjut...")
	fmt.Scanln()
}

func cekAdaPerangkat() bool {
	if len(dataPerangkat) == 0 {
		fmt.Println("Belum ada data perangkat, silahkan tambahkan perangkat terlebih dahulu di menu Kelola Perangkat - Tambahkan Perangkat")
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
			ruangan:       "Ruang Tamu",
			dayaW:         100,
			waktu:         5,
		},
		{
			namaPerangkat: "AC",
			ruangan:       "Kamar Tidur",
			dayaW:         750,
			waktu:         8,
		},
		{
			namaPerangkat: "Laptop",
			ruangan:       "Ruang Kerja",
			dayaW:         65,
			waktu:         6,
		},
		{
			namaPerangkat: "Dongle",
			ruangan:       "Ruang Kerja",
			dayaW:         12,
			waktu:         2,
		},
	}
}

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
	fmt.Println("1. Tambahkan Perangkat") // kurang cek valid
	fmt.Println("2. Tampilkan Perangkat")
	fmt.Println("3. Ubah Detail Perangkat") // blum
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
	fmt.Println("3. Perangkat Paling Boros Listrik") // yg paling tinggi
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
			//
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

            fmt.Print("===== DAFTAR PERANGKAT BERDASARKAN ABDJA =====")

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

		fmt.Print("Silahkan pilih menu (0 - 2): ")
		fmt.Scanln(&pilih)

		switch pilih {
		case "1":
			//
		case "2":
			//
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

		fmt.Print("Silahkan pilih menu (0 - 3): ")
		fmt.Scanln(&pilih)

		switch pilih {
		case "1":
			//
		case "2":
			//
		case "3":
			//
		case "0":
			return
		default:
			fmt.Print("Tolong masukkan input berupa angka 0 - 3!")
			pause()
		}
	}
}

// CRUD FEATURE

// ON PROGGRESS
func tambahPerangkat() {

	var tambah Perangkat
	clearScreen()

	fmt.Print("===== TAMBAH PERANGKAT =====\n")

	// Belum ada cek error input
	fmt.Print("Masukkan Nama Perangkat:")
	fmt.Scanln(&tambah.namaPerangkat) // gapapa kalo isi nomor (skenario: Ac-01)
	fmt.Print("Masukkan Nama Ruangan: ")
	fmt.Scanln(&tambah.ruangan) // gapapa kalo isi nomor (skenario: Kamar-01)

	// Validasi duplikat perangkat

	for i:= 0; i < len(dataPerangkat); i++ {

		if tambah.namaPerangkat == dataPerangkat[i].namaPerangkat {
			if tambah.ruangan == dataPerangkat[i].ruangan {
				fmt.Print("Nama perangkat ada diruangan yang sama")
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
	fmt.Println("Nama Perangkat: ",tambah.namaPerangkat)
	fmt.Println("Ruangan: ",tambah.ruangan)
	fmt.Println("Daya (watt): ",tambah.dayaW)
	fmt.Println("Waktu (jam): ",tambah.waktu)

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
    fmt.Print("Pilih data ke- berapa yang ingin dihapus: ")
    fmt.Println("\nTekan Enter untuk batal")
    _,err := fmt.Scanln(&pilihHapus)
    
	if err != nil {
		clearScreen()
        fmt.Println("❌ Error: Input tidak boleh kosong dan harus berupa angka!")
        
        var sampah string
        fmt.Scanln(&sampah)
        
        pause()
        return 
    }

    idx := pilihHapus - 1

	if idx < 0 || idx >= len(dataPerangkat) {
		fmt.Println("Perangkat tidak ditemukan!")
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
