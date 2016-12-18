//  config handles the basic handling of reading the gnampfile.
//  A gnampfile is a config file in JSON. The simplest gnampfile is
//  the file containing an empty object:
//      {
//      }
//  A better example is the following:
//      {
//          "PreventHTTPS": false,
//          "Port": 8080,
//          "Certificate": "certificate.pem",
//          "Keyfile": "key.ppk",
//          "CiRoot": "./",
//      }
package config

import (
    "encoding/json"
    "io/ioutil"
    "bytes"
    "log"
    "os"
    "github.com/Thamtham/gnampfelix-ci/domain"
)

type ConfigFile struct {
    Port int
    Keyfile string
    Certificate string
    PreventHTTPS bool
    CiRoot string
}

var (
    errorLog *log.Logger
    warningLog *log.Logger
    infoLog *log.Logger
    currentConfig ConfigFile
)

func init() {
    infoLog = log.New(os.Stdout, "[INFO]\t\t", 0)
    warningLog = log.New(os.Stdout, "[WARNING]\t\t", 0)
    errorLog = log.New(os.Stdout, "[ERROR]\t\t", 0)
    currentConfig = ConfigFile{
        Port: 443,
        Keyfile: "key.ppk",
        Certificate: "certificate.pem",
        CiRoot: "./",
        PreventHTTPS: false,
    }
}

//  Set the Logger for the configuration file handling.
//  The default logger of this package just prints the output
//  without any prefix to Stdout.
func SetLogger(err, warn, inf *log.Logger) {
    errorLog = err
    warningLog = warn
    infoLog = inf
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
    infoLog.Printf("Using %s-Mode.\n", settingString)

    tmpPort, ok := resultMap["Port"].(float64)
    if !ok {
        if c.PreventHTTPS {
            c.Port = 80
        }
    } else {
        c.Port = int(tmpPort)
    }
    infoLog.Printf("Using port %d.\n", c.Port)

    c.Keyfile, ok = resultMap["Keyfile"].(string)
    if !ok {
        warningLog.Printf("No keyfile specified. Using default name.\n")
    }
    infoLog.Printf("Using %s as keyfile.\n", c.Keyfile)

    c.Certificate, ok = resultMap["Certificate"].(string)
    if !ok {
        warningLog.Printf("No certificate specified. Using default name.\n")
    }
    infoLog.Printf("Using %s as certificate.\n", c.Certificate)

    c.CiRoot, ok = resultMap["CiRoot"].(string)
    if !ok {
        c.CiRoot = "./"
    }
    infoLog.Printf("CI-Root folder is %s.\n", c.CiRoot)

    return nil
}

//  Tries to read the "gnampfile" in the current directory and to convert it into
//  a ConfigFile. If the file could not be read correctly, the last working file
//  is returned. If no last config exists, an empty config is returned.
func ReadFile() (ConfigFile, error) {
    var compactData bytes.Buffer
    data, err := ioutil.ReadFile("gnampfile")
    if err != nil {
        return currentConfig, err
    }
    err = json.Compact(&compactData, data)
    if err != nil {
        return currentConfig, err
    }
    err = json.Unmarshal(compactData.Bytes(), &currentConfig)
    if err != nil {
        return currentConfig, err
    }

    return currentConfig, nil
}

func GetConfig() ConfigFile {
    return currentConfig
}

//  Based on the current valid config (or the default config) of the main application,
//  ReafRepoConfig reads the configuration file of the given repo and converts
//  it into a domain.RepoConfig object.
func ReadRepoConfig(repoName string) (domain.RepoConfig, error) {
    var compactData bytes.Buffer
    var repoConfig domain.RepoConfig
    data, err := ioutil.ReadFile(currentConfig.CiRoot + repoName + ".json")
    if err != nil {
        return domain.RepoConfig{}, err
    }
    err = json.Compact(&compactData, data)
    if err != nil {
        return domain.RepoConfig{}, err
    }

    err = json.Unmarshal(compactData.Bytes(), &repoConfig)
    if err != nil {
        return domain.RepoConfig{}, err
    }
    return repoConfig, nil
}
