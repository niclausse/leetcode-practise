package golang

import "sync"

func printLetterAndNumber() {
	letter, number := make(chan bool), make(chan bool)

	go func() {
		i := 1

		for {
			select {
			case <-number:
				print(i)
				i++
				print(i)
				i++
				letter <- true
			}
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(w *sync.WaitGroup) {
		i := 'A'

		for {
			select {
			case <-letter:
				if i > 'Z' {
					w.Done()
					return
				}

				print(string(i))
				i++
				print(string(i))
				i++
				number <- true
			}
		}
	}(&wg)

	number <- true
	wg.Wait()
}
