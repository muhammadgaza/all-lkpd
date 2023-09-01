package main

import (
	
	"net/http"

	"all-lkpd/controllers/tugascontroller"

)

func main()  {
	
	http.HandleFunc("/", tugascontroller.Index)
	http.HandleFunc("/", tugascontroller.Soal1)
	http.HandleFunc("/", tugascontroller.Soal2)
	http.HandleFunc("/", tugascontroller.Soal3)

	http.ListenAndServe(":8080", nil)
}