# nextfreeip
Find the next IP address that is not listed in DNS when given a CIDR address


## Usage

```shell

C:\> nextfreeip 192.168.1.0/26

192.168.1.0      SKIPPED - Network Boundary
192.168.1.1      device1.example.com.
192.168.1.2      device2.example.com.
192.168.1.3      device3.example.com.
192.168.1.4      device4.example.com.
192.168.1.5      device5.example.com.
192.168.1.6 is not in DNS

```

## Notes
* The program stops checking after `x.y.z.255` because it assumes a `/24` netmask when unspecified.

## Download
* Binaries for Windows, macOS and Linux can be found on the [releases](https://github.com/jftuga/nextfreeip/releases) page.
