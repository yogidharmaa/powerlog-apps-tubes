		if tambah.namaPerangkat == dataPerangkat[i].namaPerangkat {
			if tambah.ruangan == dataPerangkat[i].ruangan {
				fmt.Print("Nama perangkat sudah ada diruangan yang sama")
				pause()
				return
			}
		}
	}