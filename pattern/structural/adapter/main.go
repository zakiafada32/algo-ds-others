package main

func main() {
	client := &client{}

	mac := &mac{}
	client.insertSquareUsbInComputer(mac)

	windowsMachine := &windows{}
	windowsMachineAdapter := &windowAdapter{
		windowMachine: windowsMachine,
	}
	client.insertSquareUsbInComputer(windowsMachineAdapter)

}
