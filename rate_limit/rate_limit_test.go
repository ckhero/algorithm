/**
 *@Description
 *@ClassName rate_limit_test
 *@Date 2021/5/19 下午4:28
 *@Author ckhero
 */

package rate_limit

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewLeakyBucket(t *testing.T) {
	leakyBucket := NewLeakyBucket()
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i ++ {
		wg.Add(1)
		go func(num int) {
			leakyBucket.GetTicket()
			fmt.Printf("[%s] 第 %d 次;\n", time.Now().String(), num)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func TestNewTicketBucket(t *testing.T) {
	ticketBucket := NewTicketBucket()
	wg := sync.WaitGroup{}
	time.Sleep(5 * time.Second)
	for i := 0; i < 10; i ++ {
		wg.Add(1)
		go func(num int) {
			ticketBucket.GetTicket()
			fmt.Printf("[%s] 第 %d 次;\n", time.Now().String(), num)
			wg.Done()
		}(i)
	}
	wg.Wait()
}