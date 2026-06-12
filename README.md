<h1 align="center">⚡ PowerLog Apps</h1>

<p align="center">
  Aplikasi CLI berbasis Go untuk manajemen dan analisis konsumsi energi listrik perangkat elektronik.
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/Status-Active-brightgreen?style=for-the-badge" alt="Status">
</p>

---

## 📖 Tentang Proyek
**PowerLog** adalah aplikasi berbasis terminal (*Command Line Interface*) yang dirancang untuk membantu efisiensi penggunaan listrik. Aplikasi ini memungkinkan pengguna untuk melakukan manajemen data perangkat elektronik, menghitung total konsumsi energi harian (kWh), serta menganalisis perangkat mana yang paling boros menggunakan algoritma *sorting* dan *searching*.

## 🚀 Fitur Utama
- [x] **Manajemen Data (CRUD)**: Menambah, menampilkan, mengubah, dan menghapus data perangkat.
- [x] **Analisis Statistik**: Menghitung total konsumsi (kWh) dan mengidentifikasi perangkat dengan daya tertinggi.
- [x] **Algoritma Sorting**: 
    - *Selection Sort* untuk mengurutkan daya perangkat dari yang terbesar.
    - *Insertion Sort* untuk mengurutkan nama perangkat secara alfabetis.
- [x] **Pencarian Cepat**: Implementasi *Binary Search* untuk mencari data berdasarkan nama ruangan.
- [x] **Validasi Input**: Memastikan data yang dimasukkan valid (tidak kosong & angka positif).

## 🛠️ Cara Menjalankan

1. **Pastikan Go sudah terinstall** di komputer kamu. Cek dengan perintah:
```bash
   go version
```
2. Clone/Download repository ini ke komputer kamu.
3. Buka Terminal atau Command Prompt, lalu masuk ke direktori proyek:
```bash
   cd tubes_powerlog
```
4. Jalankan aplikasi:
```bash
   go run powerlog.go
```

📂 Struktur Data
Aplikasi ini dibangun menggunakan struct untuk menjaga keteraturan data:
```bash
type Perangkat struct {
	namaPerangkat string
	ruangan       string
	dayaW         float64
	waktu         float64
}
```

⚙️ Kompatibilitas
Program ini mendukung fitur clear screen otomatis yang bisa berjalan di:

- ✅ Windows (cls)
- ✅ Linux/macOS (clear)

Dibuat untuk memenuhi tugas mata kuliah Algoritma & Pemrograman 2.
