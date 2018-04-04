package letter

import (
	"sync"
)

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(strList []string) FreqMap {
	m := FreqMap{}
	a := new(sync.Mutex)
	forever := make(chan bool)
	for _, str := range strList {
		go func(input string) {
			for _, r := range input {
				a.Lock()
				m[r]++
				a.Unlock()
			}
			forever <- true
		}(str)
	}
	<-forever
	<-forever
	<-forever
	return m
}
