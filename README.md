ipgrep
======

IPv4/IPv6 CIDR netmask aware grep utility. ipgrep can match both
log files by CIDR mask or e.g. firewall CIDR masks by IP address.

Usage (with [go installed and GOPATH set][GOPATH]):

    $ go get github.com/joneskoo/ipgrep
    $ $GOPATH/bin/ipgrep 1.2.3.0/24 sample_files/ipv4.txt
    eka 1.2.3.4
    toka 1.2.3.4/24
    neljäs 1.2.3.4/8
    viides 1.2.3.4/255.255.0.0
    phoo 1.2.3.6/64
    xyz 1.2.3.4

[GOPATH]: https://golang.org/doc/code.html#GOPATH "How to Write Go Code - The GOPATH environment variable"
