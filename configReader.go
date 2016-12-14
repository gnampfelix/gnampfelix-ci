package main

import (
    "encoding/json"
    "io/ioutil"
    "bytes"
)

type ConfigFile struct {
    Port int
    Keyfile string
    Certificate string
    PreventHTTPS bool
    CiRoot string
}

func (c *ConfigFile)UnmarshalJSON(data []byte) error {
    var (
        resultMap map[string]interface{}
        ok bool
        settingString string
    )

    err := json.Unmarshal(data, &resultMap)
    if err != nil {
        return err
    }
    c.PreventHTTPS, ok = resultMap["PreventHTTPS"].(bool)
    if !ok || !c.PreventHTTPS {
        settingString = "HTTPS"
    } else {
        settingString = "HTTP"
    }
    Info.Printf("Using %s-Mode.\n", settingString)

    tmpPort, ok := resultMap["Port"].(float64)
    if !ok {
        if c.PreventHTTPS {
            c.Port = 80
        } else {
            c.Port = 443
        }
    } else {
        c.Port = int(tmpPort)
    }
    Info.Printf("Using port %d.\n", c.Port)

    c.Keyfile, ok = resultMap["Keyfile"].(string)
    if !ok {
        c.Keyfile = "key.ppk"
        Warning.Printf("No keyfile specified. Using default name.\n")
    }
    Info.Printf("Using %s as keyfile.\n", c.Keyfile)

    c.Certificate, ok = resultMap["Certificate"].(string)
    if !ok {
        c.Certificate = "certificate.pem"
        Warning.Printf("No certificate specified. Using default name.\n")
    }
    Info.Printf("Using %s as certificate.\n", c.Certificate)

    c.CiRoot, ok = resultMap["CiRoot"].(string)
    if !ok {
        c.CiRoot = "./"
    }
    Info.Printf("CI-Root folder is %s.\n", c.CiRoot)

    return nil
}

func readConfing() (ConfigFile, error) {
    var config ConfigFile
    var compactData bytes.Buffer
    data, err := ioutil.ReadFile("gnampfile")
    if err != nil {
        return ConfigFile{}, err
    }
    err = json.Compact(&compactData, data)
    if err != nil {
        return ConfigFile{}, err
    }
    err = json.Unmarshal(compactData.Bytes(), &config)

    return config, nil
}
