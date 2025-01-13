package WindowsSandbox

import (
	"encoding/xml"
)

type Sandbox struct {
	Configuration Configuration
}

func (s *Sandbox) ToXML() (string, error) {
	xmlBytes, err := xml.MarshalIndent(s.Configuration, "", "    ")
	if err != nil {
		return "", err
	}
	return string(xmlBytes), nil
}

func (s *Sandbox) FromXML(xmlString string) error {
	return xml.Unmarshal([]byte(xmlString), s)
}

type Configuration struct {
	VGPU                 State         `xml:"vGPU"`
	Networking           State         `xml:"Networking"`
	MappedFolders        MappedFolders `xml:"MappedFolders"`
	LogonCommand         LogonCommand  `xml:"LogonCommand"`
	AudioInput           State         `xml:"AudioInput"`
	VideoInput           State         `xml:"VideoInput"`
	ProtectedClient      State         `xml:"ProtectedClient"`
	PrinterRedirection   State         `xml:"PrinterRedirection"`
	ClipboardRedirection State         `xml:"ClipboardRedirection"`
	MemoryInMB           int           `xml:"MemoryInMB"`
}

type MappedFolders struct {
	MappedFolders []MappedFolder `xml:"MappedFolder"`
}

type MappedFolder struct {
	HostFolder    string `xml:"HostFolder"`
	SandboxFolder string `xml:"SandboxFolder"`
	ReadOnly      bool   `xml:"ReadOnly"`
}

type LogonCommand struct {
	Command string `xml:"Command"`
}

type State string

const (
	ENABLED  State = "Enable"
	DISABLED State = "Disable"
	DEFAULT  State = ""
)
