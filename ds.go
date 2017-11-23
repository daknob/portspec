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
}
