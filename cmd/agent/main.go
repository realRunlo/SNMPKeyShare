package main

import (
    "os"
     "github.com/realRunlo/SNMPKeyShare/pkg/agent/parser"  
     log "github.com/sirupsen/logrus"
     km "github.com/realRunlo/SNMPKeyShare/pkg/agent/keyManagement" 
)

func init() {

    // Text format
    log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
    // Output to stdout instead of the default stderr
    log.SetOutput(os.Stdout)

    // Only log the warning severity or above.
    log.SetLevel(log.TraceLevel)
}

// 126 - 33 + 1 = 94 - range of visable ASCII chars
const FM_SIZE int = 94

func main() {
     // Generate fm matrix on another thread
    fm_channel := make(chan [][]int32)
    go func() {
	 fm_channel <- km.Generate_random_matrix(FM_SIZE)
	 log.Info("FM matrix generated")
    }()
    // Parse config file
    var confs parser.AgentConf
    confs = parser.Parse_config_file("../../etc/agentConf.json")

    var fm_matrix [][]int32 = <-fm_channel

    var z = km.Generate_matrix(fm_matrix,confs.Master_key)
    var updated_z = km.Update_matrix(z)
    km.Keygen(fm_matrix,updated_z,1)
}
