package cofigstruct

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    configFile := "config.json"
    fileContent, err := ioutil.ReadFile(configFile)
    if err != nil {
        log.Fatalf("Error reading configuration file: %s", err)
    }

    var config ConfigurationStructure
    if err := json.Unmarshal(fileContent, &config); err != nil {
        log.Fatalf("Error decoding configuration: %s", err)
    }


// Resolve connected appliance references in each device
    applianceMap := make(map[string]*Appliance)
    for , appliance := range config.Appliances {
        applianceMap[appliance.ApplianceID] = appliance
    }

    for , device := range config.Devices {
        for _, applianceID := range device.ConnectedApplianceIDs {
            if appliance, exists := applianceMap[applianceID]; exists {
                device.ConnectedAppliances = append(device.ConnectedAppliances, appliance)
            }
        }
    }

    fmt.Printf("Configuration loaded: %+v\n", config)
}




/*
func main() {
    configFile := "config.json"
    fileContent, err := ioutil.ReadFile(configFile)
    if err != nil {
        log.Fatalf("Error reading configuration file: %s", err)
    }

    var config ConfigurationStructure
    if err := json.Unmarshal(fileContent, &config); err != nil {
        log.Fatalf("Error decoding configuration: %s", err)
    }
*/
