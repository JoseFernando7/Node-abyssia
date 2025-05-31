package models

type PortBinding struct {
    PrivatePort uint16 `json:"privatePort"`
    PublicPort  uint16 `json:"publicPort"`
    Type        string `json:"type"`
}

type VolumeMount struct {
    Source      string `json:"source"`
    Destination string `json:"destination"`
}

type NetworkInfo struct {
    Name      string `json:"name"`
    IPAddress string `json:"ipAddress"`
}

type ContainerDetails struct {
    ID        string         `json:"id"`
    Name      string         `json:"name"`
    Image     string         `json:"image"`
    Status    string         `json:"status"`
    StartedAt string         `json:"startedAt"`
    Uptime    string         `json:"uptime"`
    Ports     []PortBinding  `json:"ports"`
    Volumes   []VolumeMount  `json:"volumes"`
    Networks  []NetworkInfo  `json:"networks"`
}
