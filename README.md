ipgrep
======

IPv4/IPv6 CIDR netmask aware file pattern searcher.
Searches lines from log files by IP address or CIDR mask.
Can also match an IP address search term against firewall
configuration with netmasks.

    $ ipgrep 4.3.2.1 example.conf
    allow all from 4.3.2.1/24
    deny from 4.3.2.1/31

[![Go](https://github.com/joneskoo/ipgrep/workflows/Go/badge.svg)](https://github.com/joneskoo/ipgrep/actions?query=workflow%3AGo)

Supported formats:

 * 1.2.3.4 (plain IPv4)
 * 1.2.3.4/24 (IPv4 CIDR)
 * 1.2.3.4m24 (IPv4 m notation)
 * 2001:db8::1 (plain IPv6)
 * 2001:db8::1/64 (IPv6 CIDR)

Installing
------------

    $ go get github.com/joneskoo/ipgrep/...
