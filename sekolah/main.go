package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"pborebuild/sekolah/entity"
	"pborebuild/sekolah/repository"
	"pborebuild/sekolah/schema"
)

func main() {
	Screen()
	Input()
}

func Screen() {
	// clear screen console
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println("=============================================")
	fmt.Println("            RA.Al-Karomah Bogor")
	fmt.Println("=============================================")
	fmt.Println("")
	fmt.Println("Silahkan pilih bayaran yang ingin dibayar !!!")
	fmt.Println("1). Pengembangan")
	fmt.Println("2). Spp")
	fmt.Println("3). Kegiatan")
	fmt.Println("4). Raport")
	fmt.Println("---------------------------------------------")
	fmt.Printf("\n")
}

func Input() {

	input := schema.SchemaBiayaPendidikan{
		UangPengembangan: 3000000,
		UangSpp:          1200000,
		UangKegiatan:     3500000,
		UangRaport:       500000,
	}

	data := repository.NewBiayaPendidikan(&input)

	pegembangan := bufio.NewScanner(os.Stdin)
	spp := bufio.NewScanner(os.Stdin)
	kegiatan := bufio.NewScanner(os.Stdin)
	raport := bufio.NewScanner(os.Stdin)

	fmt.Printf("[Pembayaran] - pilih 1 jika ya, pilih 2 jika tidak: ")
	pegembangan.Scan()

	fmt.Printf("[Spp] - pilih 1 jika ya, pilih 2 jika tidak: ")
	spp.Scan()

	fmt.Printf("[Kegiatan] - pilih 1 jika ya, pilih 2 jika tidak: ")
	kegiatan.Scan()

	fmt.Printf("[Raport] - pilih 1 jika ya, pilih 2 jika tidak: ")
	raport.Scan()

	pegembanganOutput, _ := strconv.ParseUint(pegembangan.Text(), 8, 64)
	sppOuput, _ := strconv.ParseUint(spp.Text(), 8, 64)
	kegiatanOutput, _ := strconv.ParseUint(kegiatan.Text(), 8, 64)
	raportOuput, _ := strconv.ParseUint(raport.Text(), 8, 64)

	fmt.Printf("\n")

	Choice(pegembanganOutput, sppOuput, kegiatanOutput, raportOuput, data)
}

func Choice(pengembanganChoice uint64, sppChoice uint64, kegiatanChoice uint64, raportChoice uint64, data *entity.EntityBiayaPendidikan) {
	if pengembanganChoice == 1 && sppChoice == 1 && kegiatanChoice == 1 && raportChoice == 1 {

		fmt.Println("Terimakasih Telah Melunasi Semua Pembayaran")

	} else if pengembanganChoice == 1 && sppChoice == 1 && kegiatanChoice == 2 && raportChoice == 1 {

		totalPrice := data.UangPengembangan + data.UangSpp + data.UangRaport
		fmt.Printf("Uang Kegiatan Belum Dibayar Dibayar Rp.%d", data.UangKegiatan)
		fmt.Printf("\n")
		fmt.Printf("Total Yang Sudah Dibayarkan Rp.%d", totalPrice)
		fmt.Printf("\n")

	} else if pengembanganChoice == 1 && sppChoice == 1 && kegiatanChoice == 1 && raportChoice == 2 {

		totalPrice := data.UangPengembangan + data.UangSpp + data.UangKegiatan
		fmt.Printf("Uang Raport Belum Dibayar Dibayar Rp.%d", data.UangRaport)
		fmt.Printf("\n")
		fmt.Printf("Total Yang Sudah Dibayarkan Rp.%d", totalPrice)
		fmt.Printf("\n")

	} else if pengembanganChoice == 1 && sppChoice == 1 && kegiatanChoice == 2 && raportChoice == 2 {

		totalPrice := data.UangPengembangan + data.UangSpp
		fmt.Printf("Uang Kegiatan Belum Dibayar Dibayar Rp.%d", data.UangKegiatan)
		fmt.Printf("Uang Raport Belum Dibayar Rp.%d", data.UangRaport)
		fmt.Printf("Total Yang Sudah Dibayarkan Rp.%d", totalPrice)

	} else if pengembanganChoice == 2 && sppChoice == 1 && kegiatanChoice == 1 && raportChoice == 1 {

		totalPrice := data.UangSpp + data.UangKegiatan + data.UangRaport
		fmt.Printf("Uang Pengembangan Belum Dibayar Rp.%d", data.UangPengembangan)
		fmt.Printf("Total Yang Sudah Dibayarkan Rp.%d", totalPrice)

	} else if pengembanganChoice == 2 && sppChoice == 1 && kegiatanChoice == 2 && raportChoice == 1 {

		totalPrice := data.UangSpp + data.UangRaport
		fmt.Printf("Uang Pengembangan Belum Dibayar Rp.%d", data.UangPengembangan)
		fmt.Printf("Uang Kegiatan Belum Dibayar Rp.%d", data.UangKegiatan)
		fmt.Printf("Total Yang Sudah Dibayarkan Rp.%d", totalPrice)

	} else if pengembanganChoice == 2 && sppChoice == 1 && kegiatanChoice == 1 && raportChoice == 2 {

		totalPrice := data.UangSpp + data.UangKegiatan
		fmt.Printf("Uang Pengembangan Belum Dibayar Rp.%d", data.UangPengembangan)
		fmt.Printf("Uang Raport Belum Dibayar Rp.%d", data.UangRaport)
		fmt.Printf("Total Yang Sudah Dibayarkan Rp.%d", totalPrice)

	} else if pengembanganChoice == 2 && sppChoice == 1 && kegiatanChoice == 2 && raportChoice == 2 {

		totalPrice := data.UangSpp
		fmt.Printf("Uang Pengembangan Belum Dibayar Rp.%d", data.UangPengembangan)
		fmt.Printf("Uang Kegiatan Belum Dibayar Rp.%d", data.UangRaport)
		fmt.Printf("Uang Raport Belum Dibayar Rp.%d", data.UangRaport)
		fmt.Printf("Total Yang Sudah Dibayarkan Rp.%d", totalPrice)

	} else if pengembanganChoice == 1 && sppChoice == 2 && kegiatanChoice == 1 && raportChoice == 1 {

		totalPrice := data.UangPengembangan + data.UangKegiatan + data.UangRaport
		fmt.Printf("uang spp Belum Dibayar Rp.%d", data.UangSpp)
		fmt.Printf("Total Yang Sudah Dibayarkan Rp.%d", totalPrice)

	} else if pengembanganChoice == 1 && sppChoice == 2 && kegiatanChoice == 2 && raportChoice == 1 {

		totalPrice := data.UangPengembangan + data.UangRaport
		fmt.Printf("uang spp Belum Dibayar Rp.%d", data.UangSpp)
		fmt.Printf("Uang Kegiatan Belum Dibayar Rp.%d", data.UangKegiatan)
		fmt.Printf("Total Yang Sudah Dibayarkan Rp.%d", totalPrice)

	} else if pengembanganChoice == 1 && sppChoice == 2 && kegiatanChoice == 1 && raportChoice == 2 {

		totalPrice := data.UangPengembangan + data.UangKegiatan
		fmt.Printf("uang spp Belum Dibayar Rp.%d", data.UangSpp)
		fmt.Printf("Uang Raport Belum Dibayar Rp.%d", data.UangRaport)
		fmt.Printf("Total Yang Sudah Dibayarkan Rp.%d", totalPrice)

	} else if pengembanganChoice == 2 && sppChoice == 2 && kegiatanChoice == 1 && raportChoice == 2 {

		totalPrice := data.UangKegiatan
		fmt.Printf("Uang Pengembangan Belum Dibayar Rp.%d", data.UangPengembangan)
		fmt.Printf("uang spp Belum Dibayar Rp.%d", data.UangSpp)
		fmt.Printf("Uang Raport Belum Dibayar Rp.%d", data.UangRaport)
		fmt.Printf("Total Yang Sudah Dibayarkan Rp.%d", totalPrice)

	} else if pengembanganChoice == 2 && sppChoice == 2 && kegiatanChoice == 1 && raportChoice == 1 {

		totalPrice := data.UangKegiatan + data.UangRaport
		fmt.Printf("Uang Pengembangan Belum Dibayar Rp.%d", data.UangPengembangan)
		fmt.Printf("uang spp Belum Dibayar Rp.%d", data.UangSpp)
		fmt.Printf("Total Yang Sudah Dibayarkan Rp.%d", totalPrice)

	} else if pengembanganChoice == 2 && sppChoice == 2 && kegiatanChoice == 2 && raportChoice == 1 {

		totalPrice := data.UangRaport
		fmt.Printf("Uang Pengembangan Belum Dibayar Rp.%d", data.UangPengembangan)
		fmt.Printf("Uang spp Belum Dibayar Rp.%d", data.UangSpp)
		fmt.Printf("Uang Kegiatan Belum Dibayar Rp.%d", data.UangKegiatan)
		fmt.Printf("Total Yang Sudah Dibayarkan Rp.%d", totalPrice)

	} else if pengembanganChoice == 2 && sppChoice == 2 && kegiatanChoice == 2 && raportChoice == 2 {
		fmt.Printf("Harap Lunasi Segera Pembayaran !!!")
	} else {
		fmt.Println("Pembayaran Tidak Tersedia !!!")
	}
}
