# Demo sanpd REST API

## Using Curl

First build this snap, install it then run the command as root user:
```bash

snapcraft -v

sudo snap install *.snap --dangerous

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
