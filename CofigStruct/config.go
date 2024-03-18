package configstruct

type Appliance struct {
	Domain      string `json:"domain"`
	ApplianceID string `json:"applianceID"`
	Status      string `json:"status"`
}

// SetStatus updates the status of the appliance.
func (a *Appliance) SetStatus(newStatus string) {
	a.Status = newStatus
}

// GetStatus returns the current status of the appliance.
func (a *Appliance) GetStatus() string {
	return a.Status
}

// Device that can control one or more appliances -> connected appliances
type Device struct {
	DeviceID              string   `json:"deviceID"` // Identifying device 
	ConnectedApplianceIDs []string `json:"connectedAppliances"` // ConnectedApplianceIDs saves the IDs of appliances connected to each device.
	ConnectedAppliances   []*Appliance
}

// ConfigurationStructure holds the configuration needed for the application to work with devices and appliances.
type ConfigurationStructure struct {
	PrivateKey string       `json:"privateKey"`
	PublicKey  string       `json:"publicKey"`
	HMACKey    string       `json:"hmacKey"`
	AESKey     string       `json:"aesKey"`
	Appliances []*Appliance `json:"appliances"`
	Devices    []*Device    `json:"devices"`
}
