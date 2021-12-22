package utils

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func WaitSignal() {
	log.Println("start")
	var endWaiter sync.WaitGroup
	endWaiter.Add(1)
	signalChannel := make(chan os.Signal, 1)
	// signal.Notifyを使ってシグナルを待つ
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		log.Printf("SIGNAL %d received, then shutting down...\n", <-signalChannel)
		endWaiter.Done()
	}()
	// シグナルが来ればここのwaitが解除される
	endWaiter.Wait()
	log.Println("end")
}
