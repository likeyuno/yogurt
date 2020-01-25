package main

import (
	"github.com/kallydev/yogurt/service/message"
	"log"
)

func runQQBot() error {
	var err error
	message.QQBotInstance, err = message.NewQQBot()
	if err != nil {
		return err
	}
	return message.QQBotInstance.Run()
}

func runTGBot() error {
	var err error
	message.TGBotInstance, err = message.NewTGBot(message.Token)
	if err != nil {
		return err
	}
	return message.TGBotInstance.Run()
}

func main() {
	var errCh = make(chan error)
	go func() {
		errCh <- runQQBot()
	}()
	go func() {
		errCh <- runTGBot()
	}()
	select {
	case err := <-errCh:
		log.Fatalln(err)
	}
}
