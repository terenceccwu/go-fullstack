package app_plugin

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

type CallbackFunc func()

type Application struct {
	startCallbacks    []CallbackFunc
	shutdownCallbacks []CallbackFunc
}

func New() *Application {
	return &Application{}
}

func (a *Application) OnStart(cb CallbackFunc) {
	a.startCallbacks = append(a.startCallbacks, cb)
}

func (a *Application) OnShutdown(cb CallbackFunc) {
	a.shutdownCallbacks = append(a.shutdownCallbacks, cb)
}

func (a *Application) Shutdown() {
	a.invokeCallbacksReverse(a.shutdownCallbacks)
}

func (a *Application) Start() {
	a.invokeCallbacks(a.startCallbacks)

	// clean up after use
	a.startCallbacks = nil
}

func (a *Application) WaitForShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("Gracefully shutting down...")
	a.Shutdown()
	log.Println("Bye.")
	os.Exit(0)
}

func (a *Application) invokeCallbacks(cbs []CallbackFunc) {
	for i := 0; i < len(cbs); i++ {
		cb := cbs[i]
		cb()
	}
}

func (a *Application) invokeCallbacksReverse(cbs []CallbackFunc) {
	// call sequence is the reverse order of registration
	for i := len(cbs) - 1; i >= 0; i-- {
		cb := cbs[i]
		cb()
	}
}
