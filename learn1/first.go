package main

// func main() {
// arr := []int{1, 2, 3, 4, 5}
// fmt.Println(mm(arr...))
// fmt.Println(arr)

// es := make(map[int64]User)
// es[1] = User{
// 	id:   1,
// 	name: "jonh1",
// }
// es[2] = User{
// 	id:   2,
// 	name: "jonh2",
// }
// es[3] = User{
// 	id:   3,
// 	name: "jonh3",
// }
// es[4] = User{
// 	id:   4,
// 	name: "jonh4",
// }
// es[5] = User{
// 	id:   5,
// 	name: "jonh5",
// }
// fmt.Println(*mm(4, es))
// ca()
// }

// func mm(id int64, mapa map[int64]User) *User {
// 	if es, ok := mapa[id]; ok {
// 		return &es
// 	}
// 	return nil
// }

// func mm(ar ...int) []int {
// 	arr := []int{}
// 	arr = append(arr, ar...)
// 	return arr

// }

// type User struct {
// 	name string
// 	id   int64
// }

//	defer func() {
//		panic := recover()
//		fmt.Println(panic)
//	}()
//
// panic("gg")
// func main() {
// 	ss()
// }

// func with() {
// 	var wg sync.WaitGroup

// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func() {
// 			fmt.Println(i + 1)
// 			defer wg.Done()
// 		}()

// 	}
// 	wg.Wait()
// 	fmt.Println("exit")
// }

// func with1() {
// 	t := time.Now()

// 	var counter int
// 	var wg sync.WaitGroup
// 	var mu sync.Mutex

// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			time.Sleep(time.Nanosecond)
// 			mu.Lock()
// 			counter++
// 			mu.Unlock()
// 			defer wg.Done()
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println(counter)
// 	fmt.Println(time.Now().Sub(t).Seconds())
// }

// func ca() {
// 	can := make(chan int, 3)
// 	go func() {
// 		can <- 1
// 		can <- 2
// 		can <- 3
// 		defer close(can)
// 	}()

// 	for q := range can {
// 		fmt.Println(q)
// 	}

// }

// func slice1() {
// 	slice := []byte("Hello world")
// 	es := (*byte)(unsafe.SliceData(slice))
// 	str := unsafe.String(es, len(slice))
// 	fmt.Println(str)
// 	slice[0] = 'q'
// 	fmt.Println(str)
// }

func slice2() {
	// slice := make([]int, 3, 6)
	// slice[5] = 10 panic
	// slice = slice[:6]
	// slice[4] = 10
	// fmt.Println(slice)

	// slice2 := []int{1, 2}
	// slice3 := []int{1, 2}
	// fmt.Println(reflect.DeepEqual(slice2, slice3))

}

// func ss() {
// 	buff := make(chan int)
// 	go func() {
// 		buff <- 1
// 		buff <- 2
// 		buff <- 3
// 		buff <- 4
// 		buff <- 5
// 		time.Sleep(time.Second * 2)
// 		close(buff)
// 	}()
// 	go func() {
// 		buff <- 1
// 		buff <- 2
// 		buff <- 3
// 		buff <- 4
// 		buff <- 5
// 	}()

// 	go func() {
// 		time.Sleep(time.Second * 4)
// 		fmt.Println("gg")
// 	}()

// 	for q := range buff {
// 		fmt.Println(q)
// 	}
// 	time.Sleep(time.Second * 6)
// }

// func ss() {
// 	es := make(chan int, 5)
// 	go func() {
// 		es <- 1
// 		es <- 2
// 		time.Sleep(time.Second * 5)
// 		es <- 3
// 		es <- 4
// 		es <- 5
// 		defer close(es)
// 	}()

// 	for q := range es {
// 		fmt.Println(q)
// 		time.Sleep(time.Second)

// 	}
// }
