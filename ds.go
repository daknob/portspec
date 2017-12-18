package main

/*
config contains the format of the PortSpec Configuration File
*/
type config struct {
	Interval      int              `yaml:"interval"`
	IPv6          bool             `yaml:"ipv6"`
	ParallelScans int              `yaml:"parallelscans"`
	Hosts         map[string][]int `yaml:"hosts"`
	ScanPorts     []int            `yaml:"scanports"`
	SendEmail     bool             `yaml:"sendemail"`
	SMTPServer    string           `yaml:"smtpserver"`
	SMTPPort      int              `yaml:"smtpport"`
	SMTPUsername  string           `yaml:"smtpusername"`
	SMTPPassword  string           `yaml:"smtppassword"`
	AlertEmail    []string         `yaml:"alertemail"`
	FromEmail     string           `yaml:"fromemail"`
}
