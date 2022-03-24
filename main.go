package main

import (
	"fmt"
	"time"

	"github.com/rbozburun/eye-notifier-go/notifier"
)

func main() {
	// 20 min = 1200 secs
	i := -1
	count := 1200
	for true {
		i = i + 1

		fmt.Println(i)
		for i == count {
			go notifier.SendNotification()
			i = 0
		}
		time.Sleep(1 * time.Second)

	}

}
