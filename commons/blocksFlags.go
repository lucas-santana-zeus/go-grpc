package commons

import "flag"

var (
	PORTService = flag.String("port blocks service", ":8080", "Blocks service server port")

	PORTClient = flag.String("port blocks client", ":8080", "Blocks client server port")

	PORTGinAPI = flag.String("port gin api", ":8000", "Gin APi server port")
)
