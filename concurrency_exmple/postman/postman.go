package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func postman(ctx context.Context, wg *sync.WaitGroup, transferPoint chan<- string, id int, mail string) {
  defer wg.Done()

 for {
  select {
    case  <-ctx.Done():
      fmt.Println("Я почтальон", id, "и я закончил нести письмо")
      return
    default:
      fmt.Println("Я почтальон", id, "и я начал нести письмо", mail)
      time.Sleep(1 * time.Second)
      fmt.Println("Я почтальон", id, "и я донёс письмо", mail)

      transferPoint <- mail
      fmt.Println("Я почтальон", id, "и я передал письмо", mail)
  }
 }
}

func postmanToMail(id int) string {
  ptm := map[int]string {
    1: "Hello world!",
    2: "Lorem ipsum",
    3: "Never gonna give you up",
  }

  mail, ok := ptm[id]
  if !ok {
    return "Fortune"
  }

  return mail
}

func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
  mailTransferPoint := make(chan string)
  wg := &sync.WaitGroup{}

  for i := 1; i <= postmanCount; i++ {
    wg.Add(1)
    go postman(ctx, wg, mailTransferPoint, i, postmanToMail(i))
  }

  go func() {
    wg.Wait()
    close(mailTransferPoint)
  }()

  return mailTransferPoint
}
