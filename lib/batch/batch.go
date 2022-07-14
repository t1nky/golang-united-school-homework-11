package batch

import (
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	inputChan := make(chan int64)
	usersChan := make(chan user)
	for i := 0; i < int(pool); i++ {
		go worker(inputChan, usersChan)
	}
	go func() {
		for i := int64(0); i < n; i++ {
			inputChan <- i
		}
		close(inputChan)
	}()

	result := make([]user, n)
	i := int64(0)
	for user := range usersChan {
		result[i] = user
		i++
		if i == n {
			break
		}
	}
	return result
}

func worker(in chan int64, returnChan chan user) {
	for i := range in {
		returnChan <- getOne(i)
	}
}
