package mapper

import (
	"sync"
)

var mutex = &sync.Mutex{}

func mapperThread(Input []string, MapperOutput map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < len(Input); i++ {
		mutex.Lock()
		MapperOutput[Input[i]] += 1
		mutex.Unlock()
	}
}

//Map the string arrays to threads and return a map that
//combines
func Mapper(SplittedInput [][]string) map[string]int {
	MapperOutput := make(map[string]int)
	var i int
	var wg sync.WaitGroup

	for i = 0; i < len(SplittedInput); i++ {
		wg.Add(1)
		mapperThread(SplittedInput[i], MapperOutput, &wg)
	}

	wg.Wait()

	return MapperOutput
}
