# nextfreeip
Find the next IP address that is not listed in DNS


## Usage

```shell

C:\> nextfreeip 192.168.1.1

192.168.1.1      device1.example.com.
192.168.1.2      device2.example.com.
192.168.1.3      device3.example.com.
192.168.1.4      device4.example.com.
192.168.1.5      device5.example.com.
192.168.1.6 is not in DNS

```

## Notes
* The program stops searching after checking the `x.y.z.255` address.

## Download
* Binaries for Windows, macOS and Linux can be found on the [releases](https://github.com/jftuga/nextfreeip/releases) page.
