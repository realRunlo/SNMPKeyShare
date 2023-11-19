package main

import (
    "os"
     log "github.com/sirupsen/logrus"
     km "github.com/realRunlo/SNMPKeyShare/pkg/agent/keyManagement" 
)

func init() {
  // Log as JSON instead of the default ASCII formatter.
 // log.SetFormatter(&log.JSONFormatter{})
    log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
  // Output to stdout instead of the default stderr
  // Can be any io.Writer, see below for File example
  log.SetOutput(os.Stdout)

  // Only log the warning severity or above.
  log.SetLevel(log.TraceLevel)
}

// 126 - 33 + 1 = 94 - range of visable ASCII chars
const FM_SIZE int = 94

func main() {
    var master_key string = "abcdefghijklmn"
    var fm_matrix [][]int32 = km.Generate_random_matrix(FM_SIZE)
    var z = km.Generate_matrix(fm_matrix,master_key)
    var updated_z = km.Update_matrix(z)
    km.Keygen(fm_matrix,updated_z,1)
}
