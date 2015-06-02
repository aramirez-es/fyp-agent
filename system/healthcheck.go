package system

import (
    "github.com/cloudfoundry/gosigar"
    "fmt"
    "os"
)

func getUsedRam() uint64 {
    mem := sigar.Mem{}
    mem.Get()

    return mem.Used
}

func getUsedHdd() uint64 {
    fslist := sigar.FileSystemList{}
    fslist.Get()

    var usedHdd uint64 = 0

    for _, fs := range fslist.List {
        dir_name := fs.DirName
        usage := sigar.FileSystemUsage{}

        usage.Get(dir_name)
        usedHdd += usage.Used
    }

    return usedHdd * 1024
}

func getCpuLoad() float64 {
    concreteSigar := sigar.ConcreteSigar{}

    uptime := sigar.Uptime{}
    uptime.Get()

    avg, _ := concreteSigar.GetLoadAverage()

    return avg.One
}

func SendHealthCheck() {
    machineUuid := getSystemUUID()
    cpuLoad := getCpuLoad()
    usedRam := getUsedRam()
    usedHdd := getUsedHdd()

    bodyInJson := fmt.Sprintf(
        `{"id":"%s", "cpu_load":"%f", "used_ram":"%d", "used_hdd":"%d"}`, machineUuid, cpuLoad, usedRam, usedHdd)
    body := []byte(bodyInJson)

    resp, err := performRequest("POST", fmt.Sprintf("https://api.pfc.aramirez.es/systems/%s/healths", machineUuid), body)

    if err != nil {
        fmt.Println(os.Stderr, "Error sending health check.")
        fmt.Println(os.Stderr, err)
        os.Exit(-1)
    } else {
        fmt.Println(os.Stdout, fmt.Sprintf("Healthcheck sent for uuid \"%s\" with info '%s'", machineUuid, bodyInJson))
        defer resp.Body.Close()
    }
}