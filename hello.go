package main

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

func newUUID() string {
    out, err := exec.Command("uuidgen").Output()
    if err != nil {
        return ""
    }
    return strings.TrimSpace(string(out))
}

func main() {
    var hostname, _ = os.Hostname()
    var ip = getExternalIp()
    var uuid = newUUID()

//    var i, _ = net.Interfaces()
//    for _, b := range i {
//        fmt.Println(b.HardwareAddr)
//    }

    body := []byte(fmt.Sprintf(`{"id":"%s", "hostname":"%s", "ip":"%s"}`, uuid, hostname, ip))
    fmt.Println(body)

    pem, _ := ioutil.ReadFile("ca.cert.pem")
    pool := x509.NewCertPool()
    pool.AppendCertsFromPEM([]byte(pem))

    tr := &http.Transport{
        TLSClientConfig:    &tls.Config{RootCAs: pool},
        DisableCompression: true,
    }
    client := &http.Client{Transport: tr}
//    resp, err := client.Post("https://api.pfc.aramirez.es/systems", body, nil)
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