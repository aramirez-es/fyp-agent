package system

import (
    "fmt"
    "os"
    "net/http"
    "io/ioutil"
    "strings"
    "os/exec"
    "crypto/tls"
    "crypto/x509"
    "bytes"
    "runtime"
    "github.com/cloudfoundry/gosigar"
)

func getExternalIp() string {
    response, err := http.Get("http://myexternalip.com/raw")
    if err != nil {
        return ""
    }

    defer response.Body.Close()
    contents, err := ioutil.ReadAll(response.Body)

    if err != nil {
        return ""
    }

    return strings.TrimSpace(string(contents))
}

func getSystemUUID() string {

//    addresses, _ := net.Interfaces()
//    for _, iface := range addresses {
//        fmt.Println(iface.HardwareAddr)
//    }

    out, err := exec.Command("uuidgen").Output()
    if err != nil {
        return ""
    }
    return strings.TrimSpace(string(out))
}

func getTotalHdd() uint64 {
    fslist := sigar.FileSystemList{}
    fslist.Get()

    var totalHdd uint64 = 0

    for _, fs := range fslist.List {
        dir_name := fs.DirName
        usage := sigar.FileSystemUsage{}

        usage.Get(dir_name)
        totalHdd += usage.Total
    }

    return totalHdd * 1024
}

func getTotalRam() uint64 {
    mem := sigar.Mem{}
    mem.Get()

    return mem.Total
}

func SendSystemInformation() {

    var hostname, _ = os.Hostname()
    var operatingSystem = runtime.GOOS
    var ip = getExternalIp()
    var uuid = getSystemUUID()
    var cpus = runtime.NumCPU();
    var ram = getTotalRam();
    var hdd = getTotalHdd();

    body := []byte(
        fmt.Sprintf(
            `{"id":"%s", "hostname":"%s", "ip":"%s", "cpus":"%d", "ram": "%d", "hdd": "%d", "os":"%s"}`, uuid, hostname, ip, cpus, ram, hdd, operatingSystem))

    pool := x509.NewCertPool()
    pool.AppendCertsFromPEM(caCertificate)

    tr := &http.Transport{
        TLSClientConfig:    &tls.Config{RootCAs: pool},
        DisableCompression: true,
    }
    client := &http.Client{Transport: tr}
    req, err := http.NewRequest("POST", "https://api.pfc.aramirez.es/systems", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    resp, err := client.Do(req)

    if err != nil {
        fmt.Println("Error")
        fmt.Println(err)
    } else {
        fmt.Println("Ok")
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)
        fmt.Println(body)
    }
}