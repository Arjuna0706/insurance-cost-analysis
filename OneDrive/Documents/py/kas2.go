package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const Maks = 100
const KasBulanan = 50000

type Pembayaran struct {
	Bulan   string
	Nominal int
	Tanggal string
	Lunas   bool
}

type Mahasiswa struct {
	NIM          string
	Nama         string
	Kelas        string
	Pembayaran   [12]Pembayaran
	TotalBayar   int
	TotalTunggak int
}

var reader = bufio.NewReader(os.Stdin)
var daftarBulan = [12]string{
	"Januari", "Februari", "Maret", "April", "Mei", "Juni",
	"Juli", "Agustus", "September", "Oktober", "November", "Desember",
}

func bacaInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func hanyaAngka(s string) bool {
	if s == "" {
		return false
	}

	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}

	return true
}

func inputNIM() string {
	for {
		nim := bacaInput("Masukkan NIM: ")
		if hanyaAngka(nim) {
			return nim
		}
		fmt.Println("NIM hanya boleh berisi angka!")
	}
}

func inputNIMDenganPrompt(prompt string) string {
	for {
		nim := bacaInput(prompt)
		if hanyaAngka(nim) {
			return nim
		}
		fmt.Println("NIM hanya boleh berisi angka!")
	}
}

func hanyaHurufDanSpasi(s string) bool {
	if strings.TrimSpace(s) == "" {
		return false
	}

	for _, ch := range s {
		if ch != ' ' &&
			(ch < 'A' || ch > 'Z') &&
			(ch < 'a' || ch > 'z') {
			return false
		}
	}

	return true
}

func inputNama() string {
	for {
		nama := bacaInput("Masukkan Nama Lengkap: ")
		if hanyaHurufDanSpasi(nama) {
			return nama
		}
		fmt.Println("Nama hanya boleh berisi huruf dan spasi!")
	}
}

func bacaInt(prompt string) int {
	var angka int

	for {
		input := bacaInput(prompt)
		jumlahTerbaca, err := fmt.Sscanf(input, "%d", &angka)
		if err == nil && jumlahTerbaca == 1 {
			return angka
		}
		fmt.Println("Input harus berupa angka. Silakan coba lagi.")
	}
}

func tampilkanMenu() {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println(" Aplikasi Informasi Kas Mahasiswa")
	fmt.Println(" SIKAS")
	fmt.Println("========================================")
	fmt.Println("1. Tambah data mahasiswa")
	fmt.Println("2. Ubah data mahasiswa")
	fmt.Println("3. Hapus data mahasiswa")
	fmt.Println("4. Catat pembayaran kas")
	fmt.Println("5. Cari mahasiswa belum bayar")
	fmt.Println("6. Urutkan data mahasiswa")
	fmt.Println("7. Tampilkan statistik kas")
	fmt.Println("8. Tampilkan seluruh data mahasiswa")
	fmt.Println("9. Keluar")
	fmt.Println("========================================")
}

func inisialisasiPembayaran() [12]Pembayaran {
	var pembayaran [12]Pembayaran

	for i := 0; i < 12; i++ {
		pembayaran[i] = Pembayaran{
			Bulan:   daftarBulan[i],
			Nominal: 0,
			Tanggal: "-",
			Lunas:   false,
		}
	}

	return pembayaran
}

func cariIndexNIM(data [Maks]Mahasiswa, jumlah int, nim string) int {
	for i := 0; i < jumlah; i++ {
		if data[i].NIM == nim {
			return i
		}
	}
	return -1
}

func hitungUlangTotal(mhs *Mahasiswa) {
	totalBayar := 0
	totalTunggak := 0

	for i := 0; i < 12; i++ {
		totalBayar += mhs.Pembayaran[i].Nominal
		if mhs.Pembayaran[i].Nominal >= KasBulanan {
			mhs.Pembayaran[i].Lunas = true
		} else {
			mhs.Pembayaran[i].Lunas = false
			totalTunggak += KasBulanan - mhs.Pembayaran[i].Nominal
		}
	}

	mhs.TotalBayar = totalBayar
	mhs.TotalTunggak = totalTunggak
}

func tambahDataMahasiswa(data *[Maks]Mahasiswa, jumlah *int) {
	if *jumlah >= Maks {
		fmt.Println("Data mahasiswa sudah penuh.")
		return
	}

	nim := inputNIM()
	if cariIndexNIM(*data, *jumlah, nim) != -1 {
		fmt.Println("NIM sudah terdaftar.")
		return
	}

	nama := inputNama()
	kelas := bacaInput("Masukkan Kelas: ")
	if kelas == "" {
		fmt.Println("Kelas tidak boleh kosong.")
		return
	}

	data[*jumlah] = Mahasiswa{
		NIM:          nim,
		Nama:         nama,
		Kelas:        kelas,
		Pembayaran:   inisialisasiPembayaran(),
		TotalBayar:   0,
		TotalTunggak: 12 * KasBulanan,
	}

	(*jumlah)++
	fmt.Println("Data mahasiswa berhasil ditambahkan.")
}

func ubahDataMahasiswa(data *[Maks]Mahasiswa, jumlah int) {
	nim := inputNIMDenganPrompt("Masukkan NIM mahasiswa yang akan diubah: ")
	index := cariIndexNIM(*data, jumlah, nim)

	if index == -1 {
		fmt.Println("Mahasiswa dengan NIM tersebut tidak ditemukan.")
		return
	}

	fmt.Println("Data lama:")
	tampilkanRingkasMahasiswa(data[index])

	nama := inputNama()
	kelas := bacaInput("Masukkan Kelas baru: ")
	if kelas == "" {
		fmt.Println("Kelas tidak boleh kosong.")
		return
	}

	data[index].Nama = nama
	data[index].Kelas = kelas
	fmt.Println("Data mahasiswa berhasil diubah.")
}

func hapusDataMahasiswa(data *[Maks]Mahasiswa, jumlah *int) {
	nim := inputNIMDenganPrompt("Masukkan NIM mahasiswa yang akan dihapus: ")
	index := cariIndexNIM(*data, *jumlah, nim)

	if index == -1 {
		fmt.Println("Mahasiswa dengan NIM tersebut tidak ditemukan.")
		return
	}

	for i := index; i < *jumlah-1; i++ {
		data[i] = data[i+1]
	}

	data[*jumlah-1] = Mahasiswa{}
	(*jumlah)--
	fmt.Println("Data mahasiswa berhasil dihapus.")
}

func pilihBulan() int {
	fmt.Println("Daftar Bulan:")
	for i := 0; i < 12; i++ {
		fmt.Printf("%2d. %s\n", i+1, daftarBulan[i])
	}

	for {
		pilihan := bacaInt("Pilih bulan (1-12): ")
		if pilihan >= 1 && pilihan <= 12 {
			return pilihan - 1
		}
		fmt.Println("Pilihan bulan harus antara 1 sampai 12.")
	}
}

func bacaNominal() int {
	for {
		nominal := bacaInt("Masukkan Nominal Pembayaran: ")
		if nominal >= 0 {
			return nominal
		}
		fmt.Println("Nominal tidak boleh negatif.")
	}
}

func catatPembayaran(data *[Maks]Mahasiswa, jumlah int) {
	nim := inputNIMDenganPrompt("Masukkan NIM mahasiswa: ")
	index := cariIndexNIM(*data, jumlah, nim)

	if index == -1 {
		fmt.Println("Mahasiswa dengan NIM tersebut tidak ditemukan.")
		return
	}

	bulan := pilihBulan()
	nominal := bacaNominal()
	tanggal := bacaInput("Masukkan Tanggal Pembayaran: ")
	if tanggal == "" {
		tanggal = "-"
	}

	data[index].Pembayaran[bulan].Nominal = nominal
	data[index].Pembayaran[bulan].Tanggal = tanggal
	data[index].Pembayaran[bulan].Lunas = nominal >= KasBulanan
	hitungUlangTotal(&data[index])

	fmt.Println("Pembayaran berhasil dicatat.")
	if data[index].Pembayaran[bulan].Lunas {
		fmt.Println("Status bulan", daftarBulan[bulan]+": Lunas")
	} else {
		fmt.Println("Status bulan", daftarBulan[bulan]+": Belum lunas")
	}
}

func cariMahasiswaBelumBayar(data *[Maks]Mahasiswa, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}

	fmt.Println("Metode Pencarian:")
	fmt.Println("1. Sequential Search berdasarkan bulan belum lunas")
	fmt.Println("2. Binary Search berdasarkan nama")

	for {
		pilihan := bacaInt("Pilih metode pencarian: ")
		switch pilihan {
		case 1:
			sequentialSearchBelumBayar(*data, jumlah)
			return
		case 2:
			selectionSortNama(data, jumlah)
			cariBelumBayarDenganBinarySearch(*data, jumlah)
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func sequentialSearchBelumBayar(data [Maks]Mahasiswa, jumlah int) {
	bulan := pilihBulan()
	ditemukan := false

	fmt.Println()
	fmt.Println("Mahasiswa yang belum lunas bulan", daftarBulan[bulan]+":")
	for i := 0; i < jumlah; i++ {
		if !data[i].Pembayaran[bulan].Lunas {
			sisa := KasBulanan - data[i].Pembayaran[bulan].Nominal
			if sisa < 0 {
				sisa = 0
			}

			fmt.Printf("- %s | %s | %s | Sisa: Rp%d\n",
				data[i].NIM, data[i].Nama, data[i].Kelas, sisa)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Tidak ada mahasiswa yang menunggak pada bulan tersebut.")
	}
}

func cariBelumBayarDenganBinarySearch(data [Maks]Mahasiswa, jumlah int) {
	nama := bacaInput("Masukkan Nama Lengkap yang dicari: ")
	index := binarySearchNama(data, jumlah, nama)

	if index == -1 {
		fmt.Println("Mahasiswa dengan nama tersebut tidak ditemukan.")
		return
	}

	fmt.Println("Data ditemukan:")
	tampilkanRingkasMahasiswa(data[index])
	tampilkanBulanBelumLunas(data[index])
}

func binarySearchNama(data [Maks]Mahasiswa, jumlah int, nama string) int {
	kiri := 0
	kanan := jumlah - 1
	target := strings.ToLower(nama)

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		namaTengah := strings.ToLower(data[tengah].Nama)

		if namaTengah == target {
			return tengah
		} else if namaTengah < target {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	return -1
}

func tampilkanBulanBelumLunas(mhs Mahasiswa) {
	ditemukan := false

	fmt.Println("Bulan yang belum lunas:")
	for i := 0; i < 12; i++ {
		if !mhs.Pembayaran[i].Lunas {
			sisa := KasBulanan - mhs.Pembayaran[i].Nominal
			if sisa < 0 {
				sisa = 0
			}

			fmt.Printf("- %s | Dibayar: Rp%d | Sisa: Rp%d\n",
				mhs.Pembayaran[i].Bulan, mhs.Pembayaran[i].Nominal, sisa)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Semua bulan sudah lunas.")
	}
}

func urutkanDataMahasiswa(data *[Maks]Mahasiswa, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}

	fmt.Println("Metode Pengurutan:")
	fmt.Println("1. Selection Sort berdasarkan Nama (ascending)")
	fmt.Println("2. Insertion Sort berdasarkan Total Tunggakan (descending)")

	for {
		pilihan := bacaInt("Pilih metode pengurutan: ")
		switch pilihan {
		case 1:
			selectionSortNama(data, jumlah)
			fmt.Println("Data berhasil diurutkan berdasarkan nama secara ascending.")
			return
		case 2:
			insertionSortTotalTunggak(data, jumlah)
			fmt.Println("Data berhasil diurutkan berdasarkan total tunggakan secara descending.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func selectionSortNama(data *[Maks]Mahasiswa, jumlah int) {
	for i := 0; i < jumlah-1; i++ {
		indexMin := i

		for j := i + 1; j < jumlah; j++ {
			namaJ := strings.ToLower(data[j].Nama)
			namaMin := strings.ToLower(data[indexMin].Nama)

			if namaJ < namaMin {
				indexMin = j
			}
		}

		if indexMin != i {
			data[i], data[indexMin] = data[indexMin], data[i]
		}
	}
}

func insertionSortTotalTunggak(data *[Maks]Mahasiswa, jumlah int) {
	for i := 1; i < jumlah; i++ {
		kunci := data[i]
		j := i - 1

		for j >= 0 && data[j].TotalTunggak < kunci.TotalTunggak {
			data[j+1] = data[j]
			j--
		}

		data[j+1] = kunci
	}
}

func tampilkanStatistikKas(data [Maks]Mahasiswa, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}

	totalSaldo := 0
	jumlahLunasSemua := 0
	jumlahMenunggak := 0

	for i := 0; i < jumlah; i++ {
		totalSaldo += data[i].TotalBayar

		if data[i].TotalTunggak == 0 {
			jumlahLunasSemua++
		} else {
			jumlahMenunggak++
		}
	}

	rataRata := float64(totalSaldo) / float64(jumlah)

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println(" Statistik Kas")
	fmt.Println("========================================")
	fmt.Println("Total saldo kas               : Rp", totalSaldo)
	fmt.Println("Mahasiswa lunas semua bulan   :", jumlahLunasSemua)
	fmt.Println("Mahasiswa masih menunggak     :", jumlahMenunggak)
	fmt.Printf("Rata-rata pembayaran          : Rp%.2f\n", rataRata)
	fmt.Println("========================================")
}

func tampilkanSeluruhDataMahasiswa(data [Maks]Mahasiswa, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}

	fmt.Println()
	fmt.Println("================================================================================================")
	fmt.Printf("%-4s %-15s %-30s %-10s %-15s %-15s\n",
		"No", "NIM", "Nama Lengkap", "Kelas", "Total Bayar", "Total Tunggak")
	fmt.Println("================================================================================================")

	for i := 0; i < jumlah; i++ {
		fmt.Printf("%-4d %-15s %-30s %-10s Rp%-13d Rp%-13d\n",
			i+1,
			data[i].NIM,
			data[i].Nama,
			data[i].Kelas,
			data[i].TotalBayar,
			data[i].TotalTunggak)
	}

	fmt.Println("================================================================================================")
}

func tampilkanRingkasMahasiswa(mhs Mahasiswa) {
	fmt.Println("NIM            :", mhs.NIM)
	fmt.Println("Nama Lengkap   :", mhs.Nama)
	fmt.Println("Kelas          :", mhs.Kelas)
	fmt.Println("Total Bayar    : Rp", mhs.TotalBayar)
	fmt.Println("Total Tunggak  : Rp", mhs.TotalTunggak)
}

func main() {
	var dataMahasiswa [Maks]Mahasiswa
	var jumlahMahasiswa int
	jalan := true

	for jalan {
		tampilkanMenu()
		pilihan := bacaInt("Pilih menu: ")

		switch pilihan {
		case 1:
			tambahDataMahasiswa(&dataMahasiswa, &jumlahMahasiswa)
		case 2:
			ubahDataMahasiswa(&dataMahasiswa, jumlahMahasiswa)
		case 3:
			hapusDataMahasiswa(&dataMahasiswa, &jumlahMahasiswa)
		case 4:
			catatPembayaran(&dataMahasiswa, jumlahMahasiswa)
		case 5:
			cariMahasiswaBelumBayar(&dataMahasiswa, jumlahMahasiswa)
		case 6:
			urutkanDataMahasiswa(&dataMahasiswa, jumlahMahasiswa)
		case 7:
			tampilkanStatistikKas(dataMahasiswa, jumlahMahasiswa)
		case 8:
			tampilkanSeluruhDataMahasiswa(dataMahasiswa, jumlahMahasiswa)
		case 9:
			fmt.Println("Terima kasih telah menggunakan SIKAS.")
			jalan = false
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih menu 1 sampai 9.")
		}
	}
}
