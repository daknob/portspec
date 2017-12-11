# PortSpec

PortSpec is a tool to ensure all hosts in your network or under your control
have open ports that are up to the specification you set.

By setting a list of hosts and the ports you expect to be open, PortSpec will
run every predefined amount of time and scan these hosts for open TCP ports.
After all open ports are found, PortSpec will compare this list with the one
you defined, and it will alert you if something is not in spec.

Currently it is being under active development, and has very minimal
functionality, fit for a small particular purpose that it was developed for,
but hopefully some day it can work and have more features that will address
more needs.
