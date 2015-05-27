package main

import (
    "github.com/aramirez-es/agent/system"
    "time"
)

func main() {
    system.SendSystemInformation()

    sendSystemInformationInterval := time.NewTicker(24 * time.Hour)
    healthCheckInterval := time.NewTicker(30 * time.Second)

    for {
        select {
            case <- sendSystemInformationInterval.C:
                system.SendSystemInformation()
            case <- healthCheckInterval.C:
                system.SendHealthCheck()
        }
    }
}