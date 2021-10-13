package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strconv"

	"pborebuild/pinjol/entity"
	"pborebuild/pinjol/repository"
	"pborebuild/pinjol/schema"
)

func main() {
	Screen()
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Screen() {
	ClearScreen()
	fmt.Println("===========================================================")
	fmt.Println("               PT.ANGIN RIBUT SENTOSA PERKASA")
	fmt.Println("===========================================================")
	fmt.Println("Solusi pinjaman online terpercaya dan aman untuk kita semua")
	fmt.Println("===========================================================")
	fmt.Printf("\n")

	InputStepOne()
}

func InputStepOne() {
	input1 := bufio.NewScanner(os.Stdin)
	input2 := bufio.NewScanner(os.Stdin)
	input3 := bufio.NewScanner(os.Stdin)

	fmt.Printf("Apakah anda ingin melakukan pinjaman ? [yes/no]: ")
	input1.Scan()
	text1 := input1.Text()

	if text1 != "yes" {
		os.Exit(0)
	}

	fmt.Printf("Apakah saat ini anda sudah berkerja ? [yes/no]: ")
	input2.Scan()
	text2 := input2.Text()

	if text2 != "yes" {
		os.Exit(0)
	}

	fmt.Printf("Berapakah income pendapatan anda selama 1 bulan ?: ")
	input3.Scan()
	text3, err := strconv.ParseUint(input3.Text(), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	if text3 <= 4500000 {
		fmt.Printf("\n")
		fmt.Println(`Pemberitahuan: Mohon maaf anda tidak bisa melakukan pinjaman, dikarenakan income pendapatan anda dibawah batas standar ketentuan syarat yang kami tentukan.`)
		os.Exit(0)
	}

	Choice(text3)
}

func Choice(data uint64) {
	ClearScreen()
	fmt.Println("===================================")
	fmt.Println("         FORM BIODATA DIRI")
	fmt.Println("===================================")
	fmt.Printf("\n")

	name := bufio.NewScanner(os.Stdin)
	alamat := bufio.NewScanner(os.Stdin)
	ktp := bufio.NewScanner(os.Stdin)
	kk := bufio.NewScanner(os.Stdin)
	pinjaman := bufio.NewScanner(os.Stdin)
	angsuran := bufio.NewScanner(os.Stdin)

	fmt.Printf("Masukan nama anda ?: ")
	name.Scan()
	nameOutput := name.Text()

	if reflect.TypeOf(nameOutput) != reflect.TypeOf("string") {
		os.Exit(1)
	}

	fmt.Printf("Masukan alamat anda ?: ")
	alamat.Scan()
	alamatOutput := alamat.Text()

	if reflect.TypeOf(alamatOutput) != reflect.TypeOf("string") {
		os.Exit(0)
	}

	fmt.Printf("Masukan No KTP anda ?: ")
	ktp.Scan()
	ktpOutput, err := strconv.ParseUint(ktp.Text(), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	if ktpOutput < 16 {
		fmt.Printf("\n")
		fmt.Println("No KTP tidak valid")
		os.Exit(0)
	}

	fmt.Printf("Masukan No KK anda ?: ")
	kk.Scan()
	kkOuput, err := strconv.ParseUint(kk.Text(), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	if kkOuput < 16 {
		fmt.Printf("\n")
		fmt.Println("No KK tidak valid")
		os.Exit(0)
	}

	fmt.Printf("Masukan jumlah pinjaman yang diajukan ?: ")
	pinjaman.Scan()
	pinjamanOuput, err := strconv.ParseUint(pinjaman.Text(), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	if pinjamanOuput < 50000 {
		fmt.Printf("\n")
		fmt.Println("Minimal pinjaman diatas Rp 50.000")
		os.Exit(0)
	}

	fmt.Printf("Masukan lama angsuran pegembalian pinjaman ?: ")

	angsuran.Scan()
	angsuranOutput, err := strconv.ParseUint(angsuran.Text(), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	if angsuranOutput > 12 {
		fmt.Printf("\n")
		fmt.Println("Maksimal pengembalian pinjaman dibawah 1 tahun")
		os.Exit(0)
	}

	var input schema.SchemaPinjol
	input.Name = nameOutput
	input.Alamat = alamatOutput
	input.Kk = kkOuput
	input.Pinjaman = pinjamanOuput
	input.Angsuran = angsuranOutput
	input.Salary = data
	input.Bunga = 2

	res := repository.NewPinjol(&input)
	Ouput(res)
}

func Ouput(data *entity.EntityPinjol) {
	ClearScreen()
	fmt.Println("======================================")
	fmt.Println("       DATA TRANSAKSI PINJAMAN")
	fmt.Println("======================================")
	fmt.Printf("\n")
	fmt.Printf("Nama: %s", data.Name)
	fmt.Printf("\n")
	fmt.Printf("Dana Pinjaman: %d", data.Pinjaman)
	fmt.Printf("\n")
	fmt.Printf("Limit Pinjaman: %d", data.LimitPinjaman)
	fmt.Printf("\n")
	fmt.Printf("Bunga Pinjaman: %v", data.Bunga)
	fmt.Printf("\n")
	fmt.Printf("Total Pengembalian Pinjaman: %d", data.Pinjaman*data.Bunga)
	fmt.Printf("\n")
	fmt.Println("======================================")
}
