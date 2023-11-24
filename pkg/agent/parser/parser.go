package parser

import (
    "os"
    "encoding/json"
     log "github.com/sirupsen/logrus"
)

type AgentConf struct{
    Port int 	//udp port
    Master_key string // master key (required)
    Update_interval int  //matrix update interval (seconds)
    Expire_interval int //key expire interval 	(seconds)
}
func Config_example() string{

    return "	{\n"+
    "		\"port\": 9595,\n" +
    "		\"master_key\": \"abcdefghijklmn\",\n" +
    "		\"update_interval\": 10,\n" +
    "		\"expire_interval\": 10\n" +
    "	}\n" +
    "Default configurations are loaded if flag is not used"
}

func Parse_config_file(filename string) AgentConf{
    
    var confs AgentConf
    
    // Read the contents of the file
    file_content, read_err := os.ReadFile(filename)
    if read_err != nil {
	log.Fatal(read_err)
    }

    parse_err := json.Unmarshal(file_content, &confs)
    if parse_err != nil {
	log.Fatal(parse_err)
    }
    log.Info("Agent configs loaded")
    return confs
}
