# Macouflage
Macouflage is a MAC address anonymization tool similar in functionality to
[GNU Mac Changer](http://directory.fsf.org/wiki/GNU_MAC_Changer). The main
difference is that Macouflage supports additional modes of operation such as
generating spoofed MAC addresses from a list of popular network devices.

# Data sources

Macouflage ships with a JSON database derived from IEEE's OUI 
(Organizationally Unique Identifier)/MA-L (MAC Address Block Large) [registry](http://standards.ieee.org/develop/regauth/oui/oui.txt).

The popular mode currently uses the data gathered  by Etienne Perot's
[macchiato](https://github.com/EtiennePerot/macchiato) project. 

The data is merged from these data sources by the supplementary [ouiner](https://github.com/mckinney-subgraph/ouiner) 
project.

# Usage

```
$ macouflage
NAME:
   macouflage - macouflage is a MAC address anonymization tool
USAGE:
   macouflage -i/--interface <device> [-b/--bia] command [command options] [arguments...]
VERSION:
   0.1
AUTHOR(S):
   David McKinney <mckinney@subgraph.com> 
COMMANDS:
   show				Print the MAC address and exit
   ending			Don't change the vendor bytes (generate last three bytes: XX:XX:XX:??:??:??)
   another			Set random vendor MAC of the same kind
   any				Set random vendor MAC of any kind
   permanent			Reset to original, permanent hardware MAC
   random			Set fully random MAC
   popular			Set a MAC from the popular vendors list
   list (popular)		Print known vendors
   search			Search vendor names
   mac				Set the MAC XX:XX:XX:XX:XX:XX
   help, h			Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   -i, --interface 	Target device (required)
   -b, --bia		Pretend to be a burned-in-address
   --help, -h		show help
   --version, -v	print the version
```

# Examples

## List vendors
```
$ macouflage list
#	VendorPrefix	Vendor
1	00:00:00	XEROX CORPORATION
2	00:00:01	XEROX CORPORATION
3	00:00:02	XEROX CORPORATION
4	00:00:03	XEROX CORPORATION
etc.
```

## List popular vendors
```
$ macouflage list popular
#	VendorPrefix	Vendor
1	00:00:48	SEIKO EPSON CORPORATION
2	00:01:29	DFI Inc.
3	00:03:0D	Uniwill Computer Corp.
4	00:08:54	Netronix, Inc.
etc.
```

## Popular mode
First take down the interface:
```
$ sudo ip link set eth0 down
```
Then change the MAC:
```
$ sudo macouflage -i eth0 popular                                                                                                                                                                                  -- INSERT --
Current MAC:   (Some MAC) (Some vendor))
Permanent MAC: (Some MAC) (Some vendor)
New MAC:	bc:5f:f4:79:32:a6 (ASRock Incorporation)
```


