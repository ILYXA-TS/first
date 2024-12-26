package main

// func main() {
// 	ss()
// }

// func ss() {
// 	es := make(chan int)

// 	go func() {
// 		time.Sleep(time.Second * 2)
// 		es <- 1
// 	}()

// 	select {
// 	case cl := <-es:
// 		fmt.Println("read", cl)
// 	case <-time.After(time.Second):
// 		fmt.Println("Time up")

// 	}

// }

// func ss() {
// 	es := make(chan int)
// 	timer := time.After(time.Microsecond * 100)
// 	go func() {

// 		defer close(es)

// 		for i := 0; i < 1000; i++ {

// 			select {
// 			case <-timer:
// 				return
// 			default:
// 				es <- i
// 			}
// 		}

// 	}()

// 	for q := range es {
// 		fmt.Println(q)
// 	}
// }

// func grafshadown() {
// 	es := make(chan os.Signal, 1)

// 	signal.Notify(es, syscall.SIGINT, syscall.SIGTERM)

// 	timer := time.After(time.Second * 10)

// 	select {
// 	case <-timer:
// 		fmt.Println("time up")
// 	case k := <-es:
// 		fmt.Println("\nSYSCALL", k)
// 		return
// 	}
// }
