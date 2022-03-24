package notifier

import (
	"fmt"

	"github.com/gen2brain/beeep"
)

// Write notifier here
func SendNotification() {

	err := beeep.Notify("Eye Health Reminder", "Please look at a different place!", "warning.png")
	if err != nil {
		fmt.Println("Notification error!")
		panic(err)
	}

	bp := beeep.Beep(440.0, 800)
	if bp != nil {
		fmt.Println("Beep error!")
		panic(bp)
	}

}
