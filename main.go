package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	tickTime := 60 * time.Second
	if len(os.Args) > 1 {
		tmp, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("cannot parse int")
		}
		tickTime = time.Duration(tmp) * time.Second
	}
	ctx := context.Background()
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	initPos, err := GetMousePosition()
	currentPos := initPos
	if err != nil {
		log.Fatalf("cannot get mouse position: %v", err)
	}

	tick := time.Tick(tickTime)
	fmt.Printf(".")

	for {
		select {
		case <-tick:
			fmt.Printf(".")
			var pressWin bool
			if currentPos == initPos {
				currentPos.X += 200
			} else {
				pressWin = true
				currentPos.X -= 200
			}
			err = MoveMouse(currentPos)
			if err != nil {
				log.Fatalf("cannot move mouse position: %v", err)
			}

			if pressWin {
				err = PressLeftWindows()
				if err != nil {
					log.Fatalf("cannot press left windows: %v", err)
				}
				<-time.After(500 * time.Millisecond)
				err = PressLeftWindows()
				if err != nil {
					log.Fatalf("cannot press left windows: %v", err)
				}
			}

		case <-ctx.Done():
			fmt.Println("done")
			return
		}
	}

}
