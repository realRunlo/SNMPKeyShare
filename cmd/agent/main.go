package main

import (
    "os"
    "flag"
    "fmt"
    "github.com/realRunlo/SNMPKeyShare/pkg/agent/parser"  
    log "github.com/sirupsen/logrus"
    km "github.com/realRunlo/SNMPKeyShare/pkg/agent/keyManagement" 
)

// 126 - 33 + 1 = 94 - range of visable ASCII chars
const FM_SIZE int = 94
// Agent server configurations
var configs parser.AgentConf
// FM Matrix
var fm_matrix [][]int32 

func customUsage(){
    fmt.Println("usage: agent [-h] [-f <path>]")
    fmt.Println("These are common commands used:")
    flag.PrintDefaults()
}

// Init function (runs before main)
func init() {

    // Text format
    log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
    // Output to stdout instead of the default stderr
    log.SetOutput(os.Stdout)

    // Only log the warning severity or above.
    log.SetLevel(log.InfoLevel)
    
    // Set flags usage output
    flag.Usage = customUsage

    // Default agent configurations
    configs.Port = 9595
    configs.Master_key = "abcdefghijklmn"
    configs.Update_interval = 10
    configs.Expire_interval = 10

    fm_matrix = km.Generate_random_matrix(FM_SIZE)
    log.Info("FM matrix generated")
}


func main() {
    // CLI parser - flags parser
    config_file := flag.String("F","default", "Agent configuration file\n	Example:\n"+parser.Config_example())
    debug_mode := flag.Bool("D",false,"Enable debug mode")
    flag.Parse()
    
    if *debug_mode {
	 // Enable debug mode
	 log.SetLevel(log.DebugLevel)	
    }

    if *config_file == "default"{
	log.Info("Using default agent configurations")
    }else{
	// Parse config file
	configs = parser.Parse_config_file(*config_file)
    }

    var z = km.Generate_matrix(fm_matrix,configs.Master_key)
    var updated_z = km.Update_matrix(z)
    km.Keygen(fm_matrix,updated_z,1)
}
