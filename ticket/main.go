package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"pborebuild/ticket/repository"
	"pborebuild/ticket/schema"
)

func main() {
	Screen()
	Input()
}

func Screen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println("===========================================")
	fmt.Println("           BELI TICKET DOT COM")
	fmt.Println("===========================================")
	fmt.Println("")
	fmt.Println("Silahkan pilih nama pesawat yang di inginkan !!!")
	fmt.Println("1). Garuda")
	fmt.Println("2). Lion")
	fmt.Println("3). Citylink")
	fmt.Println("==============================================")
	fmt.Printf("\n")
}

func Input() {
	maskapai := bufio.NewScanner(os.Stdin)
	jumlahTicket := bufio.NewScanner(os.Stdin)

	fmt.Printf("Silahkan Pilih Maskapai ( [1] atau [2] atau [3] ): ")
	maskapai.Scan()
	choiceMaskapai, _ := strconv.ParseUint(maskapai.Text(), 10, 64)

	fmt.Printf("Masukan Jumlah Ticket Yang Akan Di Beli: ")
	jumlahTicket.Scan()
	choiceJumlahTicket, _ := strconv.ParseUint(jumlahTicket.Text(), 10, 64)

	var input schema.SchemaTicket
	input.MaskapaiCount = choiceMaskapai
	input.TicketCount = choiceJumlahTicket

	fmt.Printf("\n")

	Choice(&input)
}

func Choice(input *schema.SchemaTicket) {
	switch input.MaskapaiCount {
	case 1:
		fmt.Println("=============================")
		fmt.Println("Pesanan Ticket Pesawat Anda:")
		fmt.Println("=============================")
		fmt.Printf("\n")

		Garuda(input)

		fmt.Printf("\n")
		fmt.Println("=============================")
		fmt.Println("=============================")
	case 2:
		fmt.Println("=============================")
		fmt.Println("Pesanan Ticket Pesawat Anda:")
		fmt.Println("=============================")
		fmt.Printf("\n")

		Lion(input)

		fmt.Printf("\n")
		fmt.Println("=============================")
		fmt.Println("=============================")
	case 3:
		fmt.Println("=============================")
		fmt.Println("Pesanan Ticket Pesawat Anda:")
		fmt.Println("=============================")
		fmt.Printf("\n")

		Citylink(input)

		fmt.Printf("\n")
		fmt.Println("=============================")
		fmt.Println("=============================")
	default:
		fmt.Println("Maskapai Pesawat Tidak Tersedia")
	}
}

func Garuda(input *schema.SchemaTicket) {

	if input.TicketCount >= 50 {
		input.Pesawat = "Garuda"
		input.Harga = 400000
		input.Discount = 200000

		data := repository.NewTicket(input)

		fmt.Printf("Nama Pesawat: %s", data.Pesawat)
		fmt.Printf("\n")
		fmt.Printf("Harga Ticket: %d", data.Harga)
		fmt.Printf("\n")
		fmt.Printf("Discount: %d", data.Discount)
		fmt.Printf("\n")
		fmt.Printf("Total: %d", data.Total)
		fmt.Printf("\n")
	} else {
		input.Pesawat = "Garuda"
		input.Harga = 400000

		data := repository.NewTicket(input)

		fmt.Printf("Nama Pesawat: %s", data.Pesawat)
		fmt.Printf("\n")
		fmt.Printf("Harga Ticket: %d", data.Harga)
		fmt.Printf("\n")
		fmt.Printf("Total: %d", data.Total)
		fmt.Printf("\n")
	}

}

func Lion(input *schema.SchemaTicket) {

	if input.TicketCount >= 50 {
		input.Pesawat = "Lion"
		input.Harga = 200000
		input.Discount = 100000

		data := repository.NewTicket(input)

		fmt.Printf("Nama Pesawat: %s", data.Pesawat)
		fmt.Printf("\n")
		fmt.Printf("Harga Ticket: %d", data.Harga)
		fmt.Printf("\n")
		fmt.Printf("Discount: %d", data.Discount)
		fmt.Printf("\n")
		fmt.Printf("Total: %d", data.Total)
		fmt.Printf("\n")
	} else {
		input.Pesawat = "Lion - Tidak ada Diskon"
		input.Harga = 200000

		data := repository.NewTicket(input)

		fmt.Printf("Nama Pesawat: %s", data.Pesawat)
		fmt.Printf("\n")
		fmt.Printf("Harga Ticket: %d", data.Harga)
		fmt.Printf("\n")
		fmt.Printf("Total: %d", data.Total)
		fmt.Printf("\n")
	}
}

func Citylink(input *schema.SchemaTicket) {

	if input.TicketCount >= 50 {
		input.Pesawat = "Citylink"
		input.Harga = 300000
		input.Discount = 150000

		data := repository.NewTicket(input)

		fmt.Printf("Nama Pesawat: %s", data.Pesawat)
		fmt.Printf("\n")
		fmt.Printf("Harga Ticket: %d", data.Harga)
		fmt.Printf("\n")
		fmt.Printf("Discount: %d", data.Discount)
		fmt.Printf("\n")
		fmt.Printf("Total: %d", data.Total)
		fmt.Printf("\n")
	} else {
		input.Pesawat = "Citylink"
		input.Harga = 300000

		data := repository.NewTicket(input)

		fmt.Printf("Nama Pesawat: %s", data.Pesawat)
		fmt.Printf("\n")
		fmt.Printf("Harga Ticket: %d", data.Harga)
		fmt.Printf("\n")
		fmt.Printf("Total: %d", data.Total)
		fmt.Printf("\n")
	}
}
