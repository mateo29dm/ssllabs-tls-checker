package main

type HostResponse struct {
	Host          string     `json:"host"`
	Status        string     `json:"status"`
	StatusMessage string     `json:"statusMessage"`
	Endpoints     []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	IPAddress     string          `json:"ipAddress"`
	StatusMessage string          `json:"statusMessage"`
	Grade         string          `json:"grade"`
	Details       EndpointDetails `json:"details"`
}

type EndpointDetails struct {
	Protocols  []Protocol `json:"protocols"`
	Heartbleed bool       `json:"heartbleed"`
	Poodle     bool       `json:"poodle"`
	Freak      bool       `json:"freak"`
	Logjam     bool       `json:"logjam"`
}

type Protocol struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Q       *int   `json:"q"`
}
