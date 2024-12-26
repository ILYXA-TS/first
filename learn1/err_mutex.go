package main

// func main() {
// 	ss()
// }

// func ss() {
// 	firstcahnal := mainachn(1)
// 	secondchan := mainachn(2)

// 	fmt.Println(<-firstcahnal, <-secondchan)

// 	var count int

// 	chanal := make(chan struct{}, 1)

// 	var wg sync.WaitGroup

// 	for q := 0; q < 1000; q++ {
// 		wg.Add(1)
// 		go func() {

// 			chanal <- struct{}{}

// 			count++

// 			defer wg.Done()
// 			<-chanal

// 		}()

// 	}
// 	wg.Wait()
// 	fmt.Println(count)

// }

// func mainachn(num int) <-chan string {

// 	chanall := make(chan string)

// 	go func() {
// 		chanall <- fmt.Sprintf("response number %d", num)
// 	}()
// 	return chanall
// }

// not with errgroup

// func ss() {

// 	var err error

// 	ctx, cancel := context.WithCancel(context.Background())

// 	var wg sync.WaitGroup
// 	wg.Add(3)
// 	go func() {
// 		time.Sleep(time.Second * 3)
// 		defer wg.Done()

// 		select {
// 		case <-ctx.Done():
// 			return
// 		default:
// 			fmt.Println("first started")
// 			time.Sleep(time.Second)
// 		}

// 	}()

// 	go func() {
// 		time.Sleep(time.Second)
// 		defer wg.Done()

// 		select {
// 		case <-ctx.Done():
// 			return
// 		default:
// 			fmt.Println("second started")
// 			err = fmt.Errorf("Any error")
// 			cancel()
// 		}

// 	}()

// 	go func() {
// 		time.Sleep(time.Second)
// 		defer wg.Done()

// 		select {
// 		case <-ctx.Done():
// 			return
// 		default:
// 			fmt.Println("three started")
// 			time.Sleep(time.Second)
// 		}

// 	}()
// 	wg.Wait()
// 	fmt.Println(err)
// }

// with errgroup

// func ss() {

// 	g, ctx := errgroup.WithContext(context.Background())

// 	g.Go(func() error {
// 		time.Sleep(time.Second)
// 		select {
// 		case <-ctx.Done():
// 			return nil
// 		default:
// 			fmt.Println("first started")
// 			time.Sleep(time.Second)
// 			return nil
// 		}
// 	})

// 	g.Go(func() error {
// 		fmt.Println("second started")
// 		return fmt.Errorf("error")
// 	})

// 	g.Go(func() error {
// 		select {
// 		case <-ctx.Done():
// 		default:
// 			fmt.Println("third started")
// 			time.Sleep(time.Second)

// 		}
// 		return nil
// 	})

// 	if err := g.Wait(); err != nil {
// 		fmt.Println(err)
// 	}
// }
