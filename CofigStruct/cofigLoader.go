package configstruct

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    configFile := "config.json" //file that will be configured
    fileContent, err := ioutil.ReadFile(configFile) // read entire file 
    if err != nil {
        log.Fatalf("Error reading configuration file: %s", err)
    }

    var config ConfigurationStructure
    if err := json.Unmarshal(fileContent, &config); err != nil {
        log.Fatalf("Error decoding configuration: %s", err) // If there is an error reading the file it will be logged and exit

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
