package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"runtime"
	"sync"
	"time"
)

// func main() {
// 	//simplereader()
// 	//rowssimple()
// 	//writersimple()
// 	//workerpool1()
// 	//Writesimple2()
// }

func simplereader() {

	emkost := make([]byte, 10)

	es := Nums{number: "1,2,3,4,5,6.,7=,8g"}

	count, err := es.Read(emkost)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Println(string(emkost), count)

}

type Nums struct {
	number string
}

func (r Nums) Read(p []byte) (int, error) {
	var count int

	for i := 0; i < len(r.number); i++ {
		if r.number[i] >= '0' && r.number[i] < '9' {
			p[count] = r.number[i]
			count++
		}
	}
	return count, io.EOF
}

func rowssimple() {
	rowsreader := Rows{text: "первая строка\nвторая строка\nтретья строка"}

	var (
		err   error
		count int
	)

	row := make([]byte, 100)

	for err != io.EOF {
		count, err = rowsreader.Read(row)
		fmt.Println(string(row), count)
	}

}

type Rows struct {
	text string
}

func (s *Rows) Read(es []byte) (int, error) {
	var q int
	for q = 0; q < len(s.text); q++ {
		if s.text[q] == '\n' {
			s.text = s.text[q+1:]
			break
		}

		es[q] = s.text[q]

		if q == len(s.text)-1 {
			s.text = ""
			return q + 1, io.EOF

		}
	}
	return q + 1, nil
}

type Writers struct {
	stroka []byte
}

func writersimple() {
	rowsreader := []byte{'1', '2', '3', '4', '-', '6', '7', '+', '9', 't', '1'}

	writer := Writers{stroka: make([]byte, 100)}

	count, err := writer.Write(rowsreader)
	if err != nil && err != io.EOF {
		log.Fatal("gg tebe")
	}

	fmt.Println(string(writer.stroka), count)
}

func (s *Writers) Write(e []byte) (int, error) {
	var count int

	for q := 0; q < len(e); q++ {
		if e[q] >= '0' && e[q] <= '9' {
			s.stroka[count] = e[q]
			count++
		}
	}
	return count, nil

}

func workerpool1() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	procesorto, procesorin := make(chan int, 5), make(chan int, 5)
	var wg sync.WaitGroup
	for q := 0; q <= runtime.NumCPU(); q++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker1(ctx, procesorto, procesorin)
		}()

	}

	go func() {
		for q := 0; q < 1000; q++ {
			procesorto <- q
		}
		close(procesorto)
	}()

	go func() {
		wg.Wait()
		defer close(procesorin)
	}()

	var counter int

	for q := range procesorin {
		counter++
		fmt.Println(q)
	}
	fmt.Println(counter)
}

func worker1(ctx context.Context, procesorto1 <-chan int, procesorin chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case es, ok := <-procesorto1:
			if !ok {
				return
			}
			time.Sleep(time.Millisecond)
			procesorin <- es * es
		}
	}

}
