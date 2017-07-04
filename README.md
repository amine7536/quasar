# Quasar [![Build Status](https://img.shields.io/travis/amine7536/quasar/master.svg?style=flat-square)](https://travis-ci.org/amine7536/quasar)


Most astronomers think a quasar is a black hole with matter falling into it. Quasar collects BGP events from Peers and sends them to multiple outputs (Stdout, Logstash ...)


## Getting Started

You need a working `Golang` developpement environnement with `glide` for dependency management  
Please see https://golang.org and https://glide.sh

### Install dependencies

```bash
[abenseddik@macpro] git clone https://github.com/amine7536/quasar.git
[abenseddik@macpro] cd quasar
[abenseddik@macpro] glide install
```

This will install all the project dependencies in the `vendor` folder.  
As of **Golang 1.6** `vendor` folder is automatically added the the `$GOPATH` during build process.

### Build

#### Manually

```bash
[abenseddik@macpro] go build -o build/quasar
```

#### Using the `Makefile`

- Use the `make` target `quasar` :  

```bash
[abenseddik@macpro] make quasar

mkdir -p build
go env
GOARCH="amd64"
GOBIN=""
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="/usr/share/gocode:/home/abenseddik/gocode"
GORACE=""
GOROOT="/usr/lib/golang"
GOTOOLDIR="/usr/lib/golang/pkg/tool/linux_amd64"
GO15VENDOREXPERIMENT="1"
CC="gcc"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0"
CXX="g++"
CGO_ENABLED="1"
go build -ldflags="-w"  -o build/quasar
```

The built binary is in the `build` folder.

```bash
[abenseddik@macpro] ll build/quasar
-rwxr-xr-x  1 amine  staff    12M Apr 21 15:38 quasar
```

Golang produces a static binary with all the dependencies and the Golang runtime embedded.

#### Build the RPM

- Use the `make` target `rpm` to build the rpm package :

```bash
[abenseddik@macpro] make clean-all
[abenseddik@macpro] make rpm
...
go build -ldflags="-w"  -o build/quasar
mkdir -p tmp/
rm -rf tmp/quasar
mkdir -p tmp/quasar/
cp build/quasar tmp/quasar/quasar
cp quasar.json tmp/quasar/quasar.json
cd tmp && tar czf quasar.tar.gz quasar/
chmod +x deploy/buildrpm.sh
cp deploy/buildrpm.sh tmp/buildrpm.sh
cd tmp && ./buildrpm.sh ../deploy/quasar.spec.centos `../build/quasar version`
...
cp tmp/rpm/RPMS/x86_64/quasar-*.rpm build/
```

The resulting RPM is in the build folder :

```bash
[abenseddik@macpro] ls build
total 30256
-rwxr-xr-x  1 amine  staff    12M Apr 21 15:38 quasar
-rw-r--r--  1 amine  staff   3.3M Apr 21 15:38 quasar-0.3.1-1.el7.x86_64.rpm
```

#### RPM Info

```bash
[abenseddik@macpro] rpm -qp --info build/quasar-0.3.1-1.el7.x86_64.rpm
Name        : quasar
Version     : 0.3.1
Release     : 1.el7
Architecture: x86_64
Install Date: (not installed)
Group       : default
Size        : 12071752
License     : MIT License
Signature   : (none)
Source RPM  : quasar-0.3.1-1.el7.src.rpm
Build Date  : Fri 21 Apr 2017 03:38:11 PM CEST
Build Host  : buildvm.centos73-2.golang.lab
Relocations : (not relocatable)
Packager    : Amine Benseddik <amine.benseddik@gmail.com>
URL         : https://github.com/amine7536/quasar
Summary     : Collects BGP events from Peers and sends them to Logstash
Description :
Most astronomers think a quasar is a black hole with matter falling into it.
Quasar collects BGP events from Peers and sends them to multiple outputs (Stdout, Logstash ...)
```

#### RPM Files

```bash
[abenseddik@macpro] rpm -qlp build/quasar-0.3.1-1.el7.x86_64.rpm
/etc/quasar/quasar.json
/etc/sysconfig/quasar
/usr/bin/quasar
/usr/lib/systemd/system/quasar.service
```


## Configuration

### Sample configuration file

```
{
  "routerid": "10.2.2.2",
  "asn": 65001,
  "api": false,
  "logs": {
    "level": "debug",
    "file": "quasar.log",
    "format": "text"
  },
  "neighbors":[{
      "address": "10.2.2.3",
      "asn": 65000
  }],
  "outputs": {
    "logstash": {
      "host": "10.2.2.4",
      "port": "3000"
    }
  }
}
```

### Usage

```
/usr/bin/quasar -c quasar-dev.json
```

### Installation

Simply copy `build/quasar` on your linux host or use the rpm package :

```bash
[abenseddik@golang.lab] sudo rpm -Uvh https://github.com/amine7536/quasar/releases/download/v0.3.1/quasar-0.3.1-1.el7.x86_64.rpm
Retrieving https://github.com/amine7536/quasar/releases/download/v0.3.1/quasar-0.3.1-1.el7.x86_64.rpm
Preparing...                          ################################# [100%]
Updating / installing...
   1:quasar-0.3.1-1.el7               ################################# [100%]
```

Edit `/etc/quasar/quasar.json` and adjust the configuration as needed, by default most of the configuration fields are empty then start `quasar` :

```bash
[abenseddik@golang.lab] sudo systemctl start quasar
```

Check the status :

```bash
[abenseddik@golang.lab] sudo systemctl status quasar
● quasar.service - Quasar BGP Collector
   Loaded: loaded (/usr/lib/systemd/system/quasar.service; disabled; vendor preset: disabled)
   Active: active (running) since Tue 2017-05-02 12:18:44 UTC; 1s ago
     Docs: https://github.com/amine7536/quasar
 Main PID: 540 (quasar)
   CGroup: /system.slice/quasar.service
           └─540 /usr/bin/quasar -c /etc/quasar/quasar.json

May 02 12:18:44 buildvm.centos73-2.golang.lab systemd[1]: Started Quasar BGP Collector.
May 02 12:18:44 buildvm.centos73-2.golang.lab systemd[1]: Starting Quasar BGP Collector...

```

