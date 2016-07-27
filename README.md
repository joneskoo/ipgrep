ipgrep
======

IPv4 netmask aware grep for firewall scripts etc.

Usage:

    $ go install github.com/joneskoo/ipgrep
    $ $GOPATH/bin/ipgrep 1.2.3.0/24 sample_files/ipv4.txt
    eka 1.2.3.4
    toka 1.2.3.4/24
    neljäs 1.2.3.4/8
    viides 1.2.3.4/255.255.0.0
    phoo 1.2.3.6/64
    xyz 1.2.3.4
