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

***

## Set Configuration File & Location
It will accept a full path and the file name can be any valid filename
### Long Format
`hedynip --config=hedynip.json`

### Short Format
`hedynip -c hedynip.json`

***

## List all Tunnels
### Long Format
`hedynip all --username=<ACCOUNT_USERNAME> --password=<ACCOUNT_PASSWORD>`
### Short Format
`hedynip a -u <ACCOUNT_USERNAME> - p <ACCOUNT_PASSWORD>`

***

## List single tunnel
### Long Format
`hedynip tunnel --id=<TUNNEL_ID> --username=<ACCOUNT_USERNAME> --password=<ACCOUNT_PASSWORD> or <TUNNEL_UPDATE_KEY>`
### Short Format
`hedynip t -i <TUNNEL_ID> -u <ACCOUNT_USERNAME> -p <ACCOUNT_PASSWORD> or <TUNNEL_UPDATE_KEY>`

***

## Update IPv4
if --ip or -m are not provided hedynip will get the IPv4 automatically.
### Long Format
`hedynip update --id=<TUNNEL_ID> --ip=<CLIENT IPV4>  --username=<ACCOUNT_USERNAME> --password=<ACCOUNT_PASSWORD> or <TUNNEL_UPDATE_KEY>`
### Short Format
`hedynip update -i <TUNNEL_ID> -m <CLIENT IPV4>  -u <ACCOUNT_USERNAME> -p <ACCOUNT_PASSWORD> or <TUNNEL_UPDATE_KEY>`

***

## Get local IPv4
### Long Format
`hedynip getmyip`
### Short Format
`hedynip g`

***

## Show Current Configuration 
it will read hedynip.json or the configuration file provided by -c, --config
### Long Format
`hedynip showconfig`
### Short Format
`hedynip s`

***

## Config file format
You can add as many tunnel_id tags you want
```
{
    "name": "hedynip",
    "version": "1",
    "credentials": {
        "username": "",
        "password": ""
    },
    "tunnel_id": {
        "updatekey": ""
    }
}
```
Example:
```
{
    "name": "hedynip",
    "version": "1",
    "credentials": {
        "username": "myusername",
        "password": "mypassword"
    },
    "000000": {
        "updatekey": "xD2YdAnNnhlMwYf3qX7n3Y"
    },
    "111111": {
        "updatekey": "J0QjHjjd8sbtXaseiq3PyY"
    }    
}
```

***

## 3rd Party Libraries Used
* kingpin | https://github.com/alecthomas/kingpin
* tablewriter | https://github.com/olekukonko/tablewriter
* config | https://github.com/gookit/config

***

eSGR is not affiliated with or endorsed by Hurricane Electric (https://he.net)