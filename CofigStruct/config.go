package configstruct

type Appliance struct {
	Domain      string `json:"domain"`
	ApplianceID string `json:"applianceID"`
	Status      string `json:"status"`
}

func (a *Appliance) SetStatus(newStatus string) {
	a.Status = newStatus
}

func (a *Appliance) GetStatus() string {
	return a.Status
}

type Device struct {
	DeviceID              string   `json:"deviceID"`
	ConnectedApplianceIDs []string `json:"connectedAppliances"`
	ConnectedAppliances   []*Appliance
}

type ConfigurationStructure struct {
	PrivateKey string       `json:"privateKey"`
	PublicKey  string       `json:"publicKey"`
	HMACKey    string       `json:"hmacKey"`
	AESKey     string       `json:"aesKey"`
	Appliances []*Appliance `json:"appliances"`
	Devices    []*Device    `json:"devices"`
}
