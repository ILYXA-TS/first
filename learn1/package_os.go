package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//test()
	test1()
	//inputtest()
	three()
}

func test() {

	newfile, err := os.Create("voy_voy_f.txt")
	if err != nil {
		log.Fatal("EBAN 4E ZA NAZVANIE")
	}
	n, err := newfile.Write([]byte("GGS DOP DOP ES ES"))
	if err != nil {
		log.Fatal("e")
	}
	fmt.Println(n)

	es, err := newfile.WriteString("eseseseseses")
	fmt.Println(es)

	oo, err := os.Open(newfile.Name())

	buff := make([]byte, 100)

	oo.Read(buff)
	fmt.Println(string(buff))
	oo.Close()

	file, err := os.OpenFile(newfile.Name(), os.O_RDWR|os.O_APPEND, 0666)

	defer func() {
		file.Close()
	}()
	// os.ReadFile and os.WriteFile минует все эти действия и возвращает срез байтов | ошибку во врaйте
}

func test1() {
	count, _ := fmt.Println("gg ez")

	count, _ = fmt.Fprintln(os.Stderr, "stdout")

	fmt.Println(count)

	file, _ := os.Create("ILYXAEPTA.txt")

	defer func() {
		file.Close()
	}()

	count, _ = fmt.Fprintln(file, "hello\nworld")
}

func inputtest() {
	var (
		text string
		//text1 string
	)
	// text2, _ := fmt.Scan(&text, &text1)
	// fmt.Println(text2)

	// count, _ := fmt.Fscan(os.Stdin, &text, &text1)
	// fmt.Println(count)

	file, _ := os.Open("ILYXAEPTA.txt")
	gg, _ := fmt.Fscanln(file, &text)
	fmt.Println(gg, text)
}

func three() {
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
}
