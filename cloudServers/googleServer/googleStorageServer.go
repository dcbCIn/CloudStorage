package main

import (
	"github.com/dcbCIn/CloudStorage/cloudLib/google"
	"github.com/dcbCIn/CloudStorage/shared"
	"github.com/dcbCIn/MidCloud/distribution"
	"github.com/dcbCIn/MidCloud/lib"
	"github.com/dcbCIn/MidCloud/services/common"
)

func main() {
	lib.PrintlnInfo("Initializing server GoogleStorage")

	lp := dist.NewLookupProxy(shared.NAME_SERVER_IP, shared.NAME_SERVER_PORT)
	err := lp.Bind("googleCloudFunctions", common.ClientProxy{Ip: shared.GOOGLE_SERVER_IP, Port: shared.GOOGLE_SERVER_PORT, ObjectId: 2000})
	lib.FailOnError(err, "Error at lookup")
	err = lp.Close()
	lib.FailOnError(err, "Error at closing lookup")

	// escuta na porta tcp configurada
	var inv dist.InvokerImpl
	inv.Register(2000, &google.GoogleFunctions{})

	err = inv.Invoke(shared.GOOGLE_SERVER_PORT, shared.CONNECTIONS)
	lib.FailOnError(err, "Error calling invoker.")

	lib.PrintlnInfo("GoogleStorage server finished")
}
