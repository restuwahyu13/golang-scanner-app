package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"pborebuild/toko/repository"
	"pborebuild/toko/schema"
)

func main() {
	Screen()

	firstweek := bufio.NewScanner(os.Stdin)
	lastweek := bufio.NewScanner(os.Stdin)

	fmt.Printf("Silahkan input penjualan baju gamis minggu pertama: ")
	firstweek.Scan()
	firstweekOutput, _ := strconv.ParseUint(firstweek.Text(), 10, 64)

	fmt.Printf("Silahkan input penjualan baju gamis sampai minggu terakhir: ")
	lastweek.Scan()
	lastweekOutput, _ := strconv.ParseUint(lastweek.Text(), 10, 64)

	var input schema.SchemaToko
	input.FirstWeek = firstweekOutput
	input.LastWeek = lastweekOutput
	input.TotalPenjualan = input.FirstWeek + input.LastWeek
	input.Jual = 150000
	input.Modal = 80000
	input.Keuntungan = 70000

	Choice(&input)
}

func Screen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println("=============== PERHITUNGAN =================")
	fmt.Println("Penjualan Baju Gamis (Selama Ramadhan 1442 H)")
	fmt.Println("=============== GAMIS FASHION ===============")
	fmt.Printf("\n")
}

func Choice(input *schema.SchemaToko) {

	if input.TotalPenjualan == 100 {

		data := repository.NewToko(input)

		fmt.Printf("\n")
		fmt.Println("=================================================")
		fmt.Println("Keuntungan Anda Mencapai Target Penjualan Sebulan")
		fmt.Println("=================================================")
		fmt.Printf("\n")
		fmt.Printf("Modal: %d", data.Modal)
		fmt.Printf("\n")
		fmt.Printf("Keuntungan: %d", data.Keuntungan)
		fmt.Printf("\n")
		fmt.Printf("Total Penjualan: %d", data.TotalPenjualan)
		fmt.Printf("\n")
	} else if input.TotalPenjualan < 100 {

		data := repository.NewToko(input)

		fmt.Printf("\n")
		fmt.Println("=======================================================")
		fmt.Println("Keuntungan Anda Tidak Mencapai Target Penjualan Sebulan")
		fmt.Println("=======================================================")
		fmt.Printf("\n")
		fmt.Printf("Modal: %d", data.Modal)
		fmt.Printf("\n")
		fmt.Printf("Keuntungan: %d", data.Keuntungan)
		fmt.Printf("\n")
		fmt.Printf("Total Penjualan: %d", data.TotalPenjualan)
		fmt.Printf("\n")
	} else {

		data := repository.NewToko(input)

		fmt.Printf("\n")
		fmt.Println("=================================================")
		fmt.Println("Anda Melebihi Pencapaian Target Penjualan Sebulan")
		fmt.Println("=================================================")
		fmt.Printf("\n")
		fmt.Printf("Modal: %d", data.Modal)
		fmt.Printf("\n")
		fmt.Printf("Keuntungan: %d", data.Keuntungan)
		fmt.Printf("\n")
		fmt.Printf("Total Penjualan: %d", data.TotalPenjualan)
		fmt.Printf("\n")
	}
}
