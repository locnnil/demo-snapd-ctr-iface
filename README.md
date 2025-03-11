# Demo sanpd REST API

## Setup

Follow the steps:

- Build this snap;
- install it then
- Connect [snapd-control interface](https://snapcraft.io/docs/snapd-control-interface)

```bash

snapcraft -v

sudo snap install *.snap --dangerous

sudo snap connect sdct:snapd-control
```

## Using Curl

- run the command as root user:

```bash
sudo sdct.curl -sS --unix-socket /run/snapd.socket http://localhost/v2/snaps/system/conf -X PUT -d '{"system": { "kernel": { "dangerous-cmdline-append": "buz=bazz foo=bar" } } }'
```

Example result:

```console
ubuntu@rtuc:~$ sudo sdct.curl -sS --unix-socket /run/snapd.socket http://localhost/v2/snaps/system/conf -X PUT -d '{"system": { "kernel": { "dangerous-cmdline-append": "buz=bazz foo=bar" } } }'
{"type":"async","status-code":202,"status":"Accepted","result":null,"change":"18"}ubuntu@rtuc:~$
ubuntu@rtuc:~$
```

Retrieving the set value:

```console
ubuntu@rtuc:~$ sudo sdct.curl -sS --unix-socket /run/snapd.socket http://localhost/v2/snaps/system/conf -X GET -d 'system.kernel.dangerous-cmdline-append'
{"type":"sync","status-code":200,"status":"OK","result":{"cloud":{"name":"multipass"},"seed":{"loaded":true},"system":{"hostname":"rtuc","kernel":{"dangerous-cmdline-append":"buz=bazz foo=bar"},"network":{"netplan":{"network":{"ethernets":{"default":{"dhcp-identifier":"mac","dhcp4":true,"match":{"macaddress":"52:54:00:9c:4d:c8"}}},"version":2}}},"timezone":"Etc/UTC"}}}ubuntu@rtuc:~$
ubuntu@rtuc:~$
```

## Using the GO app

- run the snap app command as root user:

```bash
sudo sdct
```

Example result:

```console
ubuntu@rtuc:~$ sudo sdct
HTTP RESPONSE:
HTTP/1.1 202 Accepted
Content-Length: 82
Content-Type: application/json
Date: Tue, 11 Mar 2025 15:13:57 GMT

{"type":"async","status-code":202,"status":"Accepted","result":null,"change":"40"}ubuntu@rtuc:~$
ubuntu@rtuc:~$
```
