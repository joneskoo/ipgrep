ipgrep
======

IPv4/IPv6 CIDR netmask aware grep utility. ipgrep can match both
log files by CIDR mask or e.g. firewall CIDR masks by IP address.

Usage
--------

    $ ipgrep 4.3.2.1 sample_files/ipv4.txt
    kuudes 4.3.2.1m255.255.255.0
    seven 4.3.2.1m31

Installing
------------

    $ go install github.com/joneskoo/ipgrep


