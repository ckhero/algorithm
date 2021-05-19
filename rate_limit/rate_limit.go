/**
 *@Description
 *@ClassName rate_limit
 *@Date 2021/5/19 下午4:17
 *@Author ckhero
 */

package rate_limit

import (
	"time"
)

type RateLimit interface {
	GetTicket() struct{}
}

// 漏桶
type LeakyBucket struct {
	ticker *time.Ticker
}

func NewLeakyBucket() RateLimit {
	leakyBucket := &LeakyBucket{
		ticker: time.NewTicker(time.Second),
	}
	return leakyBucket
}

func (l *LeakyBucket) GetTicket() struct{} {
	<- l.ticker.C
	return struct{}{}
}

// 令牌桶

type TicketBucket struct {
	ticker *time.Ticker
	bucketCh chan struct{}
}

func NewTicketBucket() RateLimit {
	ticketBucket := &TicketBucket{
		ticker: time.NewTicker(time.Second),
		bucketCh: make(chan struct{}, 10),
	}
	go ticketBucket.GenTicket()
	return ticketBucket
}

func(t *TicketBucket) GetTicket() struct{} {
	return <- t.bucketCh
}

func(t *TicketBucket) GenTicket() struct{} {
	for {
		<- t.ticker.C
		t.bucketCh <- struct{}{}
	}
}
