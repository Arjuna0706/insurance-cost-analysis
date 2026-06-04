package main

//import "fmt"

// Mengurutkan nama dari A ke Z (Ascending)
func selectionSortBerdasarkanNama(data *[Maks]Mahasiswa, jumlah int) {
	for i := 0; i < jumlah-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlah; j++ {
			if data[j].Nama < data[minIdx].Nama {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
}

func insertionSortBerdasarkanTunggakan(data *[Maks]Mahasiswa, jumlah int) {
	for i := 1; i < jumlah; i++ {
		key := data[i]
		j := i - 1

		for j >= 0 && data[j].TotalTunggakan < key.TotalTunggakan {
			data[j+1] = data[j]
			j--
		}

		data[j+1] = key
	}
}