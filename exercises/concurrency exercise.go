package main

func StartDecipher(senderChan chan string, decipherer func(encrypted string) string) chan string {
	ch := make(chan string, 5)
	go func() {
		for st := range senderChan {
			deciphered := decipherer(st)
			ch <- deciphered
		}
		close(ch)
	}()
	return ch
}
