# hedynip
Update the IPv4 for the Hurricane Electric IPv6 Tunnel service for users with non static IPv4

## OS & Arch Support
* Linux - amd64, 386, arm, arm64, mips, mipsle, ppc64, ppc64le, s390x
* Windows - 386, amd64
* OS X - 386, amd64
* Dragonfly - amd64
* FreeBSD - 386, amd64, arm
* NetBSD - 386, amd64, arm
* openBSD - 386, amd64
* Solaris - amd64

## Build
In most cases this is enough to build:
```
go get github.com/eservicesgreece/hedynip
```

## List all Tunnels
### Long Format
`hedynip all --username=<ACCOUNT_USERNAME> --password=<ACCOUNT_PASSWORD>`
### Short Format
`hedynip a -u <ACCOUNT_USERNAME> - p <ACCOUNT_PASSWORD>`

## List single tunnel
### Long Format
`hedynip tunnel --id=<TUNNEL_ID> --username=<ACCOUNT_USERNAME> --password=<ACCOUNT_PASSWORD> or <TUNNEL_UPDATE_KEY>`
### Short Format
`hedynip t -i <TUNNEL_ID> -u <ACCOUNT_USERNAME> -p <ACCOUNT_PASSWORD> or <TUNNEL_UPDATE_KEY>`

## Update IPv4
if --ip or -m are not provided hedynip will get the IPv4 automatically.
### Long Format
`hedynip update --id=<TUNNEL_ID> --ip=<CLIENT IPV4>  --username=<ACCOUNT_USERNAME> --password=<ACCOUNT_PASSWORD> or <TUNNEL_UPDATE_KEY>`
### Short Format
`hedynip update -i <TUNNEL_ID> -m <CLIENT IPV4>  -u <ACCOUNT_USERNAME> -p <ACCOUNT_PASSWORD> or <TUNNEL_UPDATE_KEY>`

## Get local IPv4
### Long Format
`hedynip getmyip`
### Short Format
`hedynip g`

***

## 3rd Party Libraries Used
* kingpin | https://github.com/alecthomas/kingpin
* tablewriter | https://github.com/olekukonko/tablewriter
