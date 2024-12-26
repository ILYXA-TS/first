package main

// func main() {
// 	// ss()
// 	// ss1()
// 	// ss2()
// 	atomicevalue()
// }

// func ss() {

// 	timer := time.Now()

// 	var count int
// 	var wg sync.WaitGroup
// 	var mu sync.Mutex

// 	for q := 0; q < 1000; q++ {
// 		wg.Add(1)

// 		go func() {

// 			defer wg.Done()
// 			mu.Lock()
// 			count++
// 			mu.Unlock()
// 		}()

// 	}
// 	wg.Wait()

// 	fmt.Println(count)
// 	fmt.Println(time.Now().Sub(timer).Seconds())
// }

// func ss1() {

// 	timer := time.Now()

// 	var count int64
// 	var wg sync.WaitGroup

// 	for q := 0; q < 1000; q++ {
// 		wg.Add(1)

// 		go func() {

// 			defer wg.Done()

// 			atomic.AddInt64(&count, 1)

// 		}()

// 	}
// 	wg.Wait()

// 	fmt.Println(count)
// 	fmt.Println(time.Now().Sub(timer).Seconds())
// }

// // Если нужно чтобы только одна горутина сделала действие
// func ss2() {

// 	var (
// 		counter int64
// 		wg      sync.WaitGroup
// 	)

// 	wg.Add(100)

// 	for q := 0; q < 100; q++ {
// 		go func(q int) {
// 			defer wg.Done()
// 			if !atomic.CompareAndSwapInt64(&counter, 0, 1) {
// 				return
// 			}
// 			fmt.Println(q)
// 		}(q)
// 	}
// 	wg.Wait()
// 	fmt.Println(counter)
// }

// func atomicevalue() {
// 	var value atomic.Value

// 	value.Load()                             // показывает переменую
// 	value.Store(1)                           // устоновка значения
// 	value.Swap(3)                            // устоновка значение и возвращает старое
// 	fmt.Println(value.CompareAndSwap(3, 44)) // замена одно на другое и возвращает bool
// 	fmt.Println(value.CompareAndSwap(2, 44)) // false
// 	fmt.Println(value.Load())

// }
