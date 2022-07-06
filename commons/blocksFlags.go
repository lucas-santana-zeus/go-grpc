package commons

import "flag"

var (
	PORTService = flag.String("port blocks service", ":8080", "Blocks service server port")

	PORTClient = flag.String("port blocks client", ":8000", "Blocks client server port")
)
