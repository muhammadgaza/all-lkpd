package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(response http.ResponseWriter, request *http.Response)  {
	
	temp, err := template.ParseFiles("views/home.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)
}

func soal1(response http.ResponseWriter, requset *http.Request)  {
	var pabp, mtk, dpk float64

	fmt.Print("Masukkan nilai PABP: ")
	fmt.Scan(&pabp)

	fmt.Print("Masukkan nilai MTK: ")
	fmt.Scan(&mtk)

	fmt.Print("Masukkan nilai DPK: ")
	fmt.Scan(&dpk)

	rataRata := (pabp + mtk + dpk) / 3.0
	grade := ""

	if rataRata >= 80 && rataRata <= 100 {
		grade = "A"
	} else if rataRata >= 75 {
		grade = "B"
	} else if rataRata >= 65 {
		grade = "C"
	} else if rataRata >= 45 {
		grade = "D"
	} else {
		grade = "E"
	}

	fmt.Printf("Rata-rata nilai: %.2f\n", rataRata)
	fmt.Printf("Grade: %s\n", grade)

	temp, err := template.ParseFiles("views/soal1.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)


}

func Soal2(response http.ResponseWriter, requset *http.Request)  {
	var data string
	fmt.Print("Masukkan angka dengan format gddmmyyyynn: ")
	fmt.Scanln(&data)

	hasilPenguraian := parseAngka(data)

	fmt.Printf("Informasi dari angka %s adalah:\n", data)
	for key, value := range hasilPenguraian {
		fmt.Printf("%s: %s\n", key, value)
	}

	temp, err := template.ParseFiles("views/soal2.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(response, nil)

}

func parseAngka(data string) map[string]string {
	golongan := string(data[0])
	tanggal := data[1:3]
	bulan := data[3:5]
	tahun := data[5:9]
	kodeIdentitas := data[9:11]

	bulanMap := map[string]string{
		"01": "JANUARI",
		"02": "FEBRUARI",
		"03": "MARET",
		"04": "APRIL",
		"05": "MEI",
		"06": "JUNI",
		"07": "JULI",
		"08": "AGUSTUS",
		"09": "SEPTEMBER",
		"10": "OKTOBER",
		"11": "NOVEMBER",
		"12": "DESEMBER",
	}

	tanggalLahir := tanggal + " " 
	bulanLahir := bulanMap[bulan] + " " 
	tahunLahir := tahun 

	informasi := map[string]string{
		"Golongan":      golongan,
		"Tanggal Lahir": tanggalLahir,
		"bulan Lahir": bulanLahir,
		"Tahun Lahir": tahunLahir,
		"Kode Identitas": kodeIdentitas,
	}

	return informasi

	
}

