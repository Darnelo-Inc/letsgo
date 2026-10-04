package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func miner(
  ctx context.Context,
  wg *sync.WaitGroup,
  transferPoint chan<- int,
  id int,
  power int,
) {
  defer wg.Done()

  for{
    select {
    case <-ctx.Done():
      fmt.Println("Я шахтёр", id, "и я завершил работу")
      return
    case <-time.After(1 * time.Second):
      fmt.Println("Я шахтёр", id, "и я добыл", power, "угля")
    }

    select {
    case <-ctx.Done():
      fmt.Println("Я шахтёр", id, "и я завершил работу")
      return
    case transferPoint <- power:
      fmt.Println("Я шахтёр", id, "и я передал", power, "угля")
    }

    // select {
    // case <-ctx.Done():
    //   fmt.Println("Я шахтёр", id, "и я начал завершил работу")
    //   return
    // default:
    //   fmt.Println("Я шахтёр", id, "и я начал добывать уголь")
    //   time.Sleep(1 * time.Second)
    //   fmt.Println("Я шахтёр", id, "и я добыл", power, "угля")

    //   transferPoint <- power
    //   fmt.Println("Я шахтёр", id, "и я передал", power, "угля")
    // }
  }
}

func MinerPool(ctx context.Context, minerCount int) <-chan int {
  coalTransferPoint := make(chan int)
  wg := &sync.WaitGroup{}

  for i := 1; i <= minerCount; i++ {
    wg.Add(1)
    go miner(ctx, wg, coalTransferPoint, i, i * 10)
  }

  go func() {
    wg.Wait()
    close(coalTransferPoint)
  }()

  return coalTransferPoint
}
