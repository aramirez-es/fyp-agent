package main

import (
    "github.com/aramirez-es/agent/system"
    "time"
)

func main() {
    system.SendSystemInformation()

    sendSystemInformationInterval := time.NewTicker(5 * time.Second)
    healthCheckInterval := time.NewTicker(10 * time.Second)

    for {
        select {
            case <- sendSystemInformationInterval.C:
                system.SendSystemInformation()
            case <- healthCheckInterval.C:
                system.SendHealthCheck()
        }
    }
}