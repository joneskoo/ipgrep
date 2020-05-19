ipgrep
======

IPv4/IPv6 CIDR netmask aware file pattern searcher.
Searches lines from log files by IP address or CIDR mask.
Can also match an IP address search term against firewall
configuration with netmasks.

    $ ipgrep 4.3.2.1 example.conf
    allow all from 4.3.2.1m255.255.255.0
    deny from 4.3.2.1m31

Supported formats:

 * 1.2.3.4 (plain IPv4)
 * 1.2.3.4/24 (IPv4 CIDR)
 * 1.2.3.4m24 (IPv4 m notation)
 * 1.2.3.4m255.255.255.0 (IPv4 m mask)
 * 2001:db8::1 (plain IPv6)
 * 2001:db8::1/64 (IPv6 CIDR)

Installing
------------

    $ go get github.com/joneskoo/ipgrep/...
