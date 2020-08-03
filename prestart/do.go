package prestart

// Do is a func will be called at init, registering the drivers of program
func Do() {
	initLogDriver()
	initConfigDriver()
}
