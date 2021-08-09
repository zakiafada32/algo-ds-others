package main

type windowAdapter struct {
	windowMachine *windows
}

func (w *windowAdapter) insertInSquarePort() {
	w.windowMachine.insertInCirclePort()
}
