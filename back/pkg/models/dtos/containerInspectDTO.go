package dtos

type ContainerInspectDTO struct {
	ID        string   `json:"id"`
  Name      string   `json:"name"`
  Status    string   `json:"status"`
  Image     string   `json:"image"`
  Ports     []string `json:"ports"`
  StartedAt string   `json:"started_at"`
  Mounts    []string `json:"mounts"`
}
