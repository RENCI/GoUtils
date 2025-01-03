package Misc

import (
	"encoding/json"
	"log"
	"time"
)

func TimeTrack(start time.Time, name string, data map[any]any, logger *log.Logger) {
	elapsed := time.Since(start)
	datajson, _ := json.Marshal(data)
	logger.Printf("%s took %s", name, elapsed)
	logger.Printf("Data: %s", string(datajson))
}
