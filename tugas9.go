package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	defer recovery()

	var password string
	fmt.Print("...\nSilahkan masukan password: ") //masukan password
	fmt.Scanln(&password)

	if valid, err := validasi(password); valid {
		fmt.Println("...\nSelamat Anda Berhasil Login ke Akun: [Andika Andriana]\n...")
	} else {
		panic(err.Error())
		fmt.Println("Erorr!!!")
	}
}

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("...\nTerjadi Error!!\n...", r)
	} else {
		fmt.Println("Aplikasi Berjalan Dengan Sempurna\n...")
	}
}

func validasi(input string) (bool, error) {
	if strings.TrimSpace(input) != "admin" {
		return false, errors.New("\nPassword Yang Anda Masukan Salah!\n...")
	}
	return true, nil
}
