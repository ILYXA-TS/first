package main

// func main() {
// 	workerpool()
// }

// func ss() {

// 	ctx := context.Background()

// 	peremena := context.WithValue(ctx, "name", "ilya")
// 	fmt.Println(peremena.Value("name"))
// 	wihtca, ca := context.WithCancel(ctx)
// 	fmt.Println(wihtca.Err())
// 	ca()
// 	fmt.Println(wihtca.Err())

// 	withdeadline, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
// 	fmt.Println(withdeadline.Deadline())
// 	fmt.Println(withdeadline.Err())
// 	fmt.Println(<-withdeadline.Done())

// 	withtimeout, cancel := context.WithTimeout(ctx, time.Second*1)
// 	defer cancel()
// 	fmt.Println(<-withtimeout.Done())
// }

// func workerpool() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	ctx, cancel = context.WithTimeout(context.Background(), time.Microsecond*20)
// 	defer cancel()

// 	cputoprcessor, cpuprocessor := make(chan int, 5), make(chan int, 5)

// 	var wg sync.WaitGroup

// 	for i := 0; i <= runtime.NumCPU(); i++ {

// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			worker(ctx, cputoprcessor, cpuprocessor)
// 		}()
// 	}

// 	go func() {
// 		for q := 0; q < 1000; q++ {

// 			cputoprcessor <- q

// 		}
// 		close(cputoprcessor)
// 	}()

// 	go func() {
// 		wg.Wait()
// 		close(cpuprocessor)
// 	}()

// 	var conter int

// 	for result := range cpuprocessor {
// 		conter++
// 		fmt.Println(result)
// 	}
// 	fmt.Println(conter)
// }

// func worker(ctx context.Context, cputoprocessor <-chan int, cpuproc chan<- int) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return
// 		case val, ok := <-cputoprocessor:
// 			if !ok {
// 				return
// 			}
// 			time.Sleep(time.Millisecond)
// 			cpuproc <- val * val
// 		}
// 	}
// }
