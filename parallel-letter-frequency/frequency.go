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
	wg := new(sync.WaitGroup)
	for _, str := range strList {
		wg.Add(1)
		go func(input string) {
			for _, r := range input {
				a.Lock()
				m[r]++
				a.Unlock()
			}
			wg.Done()
		}(str)
	}
	wg.Wait()
	return m
}

/*最佳实现1
func ConcurrentFrequency(s []string) FreqMap {
	var wg sync.WaitGroup
	counters := make([]FreqMap, len(s))

	for i, sentence := range s {
		wg.Add(1)
		go func(counters []FreqMap, i int, sentence string) {
			counters[i] = Frequency(sentence)
			wg.Done()
		}(counters, i, sentence)
	}

	wg.Wait()

	result := FreqMap{}
	for _, counter := range counters {
		for key, value := range counter {
			result[key] += value
		}
	}

	return result
}
*/
/*最佳实现2
func ConcurrentFrequency(ss []string) FreqMap {
	fm := FreqMap{}
	cc := make(chan FreqMap)

	for _, s := range ss {
		go func(s string) {
			cc <- Frequency(s)
		}(s)
	}

	for i := 0; i < len(ss); i++ {
		fm.merge(<-cc)
	}

	return fm
}

func (f FreqMap) merge(f2 FreqMap) FreqMap {
	for k, v := range f2 {
		lock.RLock()
		_, ok := f[k]
		lock.RUnlock()

		lock.Lock()
		if !ok {
			f[k] = 0
		}
		f[k] += v
		lock.Unlock()

	}

	return f
}
*/
