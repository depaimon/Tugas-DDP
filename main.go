// Package main digunakan untuk mendefinisikan fungsi utama.
package main 

// mengimpor fusionalitas yang diperlukan.
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

// struct user input dibuat untuk menyimpan input pengguna 
type UserInput struct {
	Nama string
}

// struct QuestionAnswer dibuat untuk menyimpan pertanyaan dan jawaban
type QuestionAnswer struct {
	Pertanyaan string
	Opsi       []string
	Jawaban    string
}

// main function.
func main() {
	// membuat variabel untuk menyimpan imput pengguna dan mengonfigurasi kebenaran,kesalahan dengan type data integer.
	var input UserInput
	var correct, incorrect int

	// meminta nama pengguna.
	fmt.Println() // baris kosong untuk pemisahan.
	fmt.Print("Masukan nama Anda: ")
	fmt.Scanln(&input.Nama)
	fmt.Println()

	// kumpulan pertanyaan dengan slice.
	questions := []QuestionAnswer{
		{"Berapakah hasil dari 1+1+?", []string{"A. Jendela", "B. 3", "C. 2", "D. Semua salah"}, "C"},
		{"Hewan apa yang kaki nya 2?", []string{"A. Gajah", "B. Anjing", "C. Pohon", "D. Ayam"}, "D"},
		{"Berapakah hasil dari 8 x 6?", []string{"A. 42", "B. 48", "C. 56", "D. 64"}, "B"},
		{"Ayah Nobel mempunyai 5 anak yang bernama Lana, Lini, Lunu, Lene dan siapakah nama anak yang terakhir?", []string{"A. Lono", "B. Lono Kopling","C. Nobel","D. Nobel Kopling" },"C"},
		{"Ada apa diujung langit?", []string{"A. Angkasa", "B. Awan", "C. Planet", "D. Huruf t"}, "D"},
	}

	// konsep iterasi menggunakan range pada slice 
	for _, qa := range questions {
		fmt.Println(qa.Pertanyaan)
		for _, option := range qa.Opsi {
			fmt.Println(option)
		}

		//membaca input dari pengguna melalui termianl atau konsol
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Jawaban (A/B/C/D): ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		// konfigurasi kebenaran dan kesalahan, jika benar maka +1 atau sebaliknya
		correct += map[bool]int{strings.EqualFold(answer, qa.Jawaban): 1, true: 0}[true]
		incorrect += map[bool]int{strings.EqualFold(answer, qa.Jawaban): 0, true: 1}[true]

		// baris kosong untuk pemisahan(<br>).
		fmt.Println()
	}

	// menampilkan skor akhir.
	fmt.Println("Statistic Kuis")
	fmt.Println("Nama        : ", input.Nama)
	fmt.Println("Skor        : ", correct)
	fmt.Println("Jawaban Benar  : ", correct)
	fmt.Println("Jawaban Salah   : ", incorrect)
}