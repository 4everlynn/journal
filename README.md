![Author](https://img.shields.io/badge/org-Fatewa_H.S.C-red)
![Golang](https://img.shields.io/badge/Golang-olive)
![Build](https://www.travis-ci.org/4everlynn/journal.svg?branch=dev)
# Journal

> Small but beautiful, simple and performance command line tool (CLI) for daily programming

![stage](resources/stage.gif)

# How to download ?
See [RELEASE](https://github.com/4everlynn/journal/releases)

# Features

- No need for Git environment, just a project folder managed by Git
- Automatically extract today's report according to the module
- Support disabling some modules
- External configuration file support

# TODO
- [x] Generate daily work reports (multiple projects), free your hands (cmd. journal stylish)
- [ ] File monitoring (via daemon), impl code submission habits statistics (cmd. journal daemon)
- [x] Work communication space，realize lan command line chat and file transfer (cmd. journal exchange -c(-s) -H localhost -p 123456 )

# Getting Start

## .journal.yaml

```yaml
# config version, syntax may be different for different versions in the future
version: '1'
maintainer: 4everlynn
git:
  project:
    name: Project
    # path to your git repo
    path: /home/4everlynn/dev/projcet
    # is disabled
    disable: true
  # another project here
  journal:
    name: Journal
    path: /home/4everlynn/dev/journal
```

## Usages

### stylish (reporting)

```bash
./journal stylish -d/-w/-m
```

### Output

```text
Config maintained by Edward is loaded successfully
20200927 Daily/Weekly/Monthly
Journal
  1、first commit
  2、Added the option of disable report generation in the warehouse
Project
  1、commit all
  2、version、depot、express and region etc.
  3、impl depot & depot-express api
  4、update query-core's version
  5、version of multipart is change, introduce query-core
  6、commit current code
  7、commit register client api
  8、client api for auth to register
  9、change genius version
  10、region changed
  11、update region api path
  12、add shipper api
  13、add global ex.h & shipper api recreate
  14、commit sys code
```

### exchange (tcp based exchanging)

![exchange](resources/exchange.gif)

#### Principle

![exchange](resources/exchange.jpg)

  User (let's call it A for the time being) initiates the Server side, and all other users are as Clients to establish a connection with A.     
   
  The responsibility of A is to maintain and monitor the connection of the Clients and when Client has messages,      
it will dispatch the message to other Clients, and at the same time.    

  If message generated, it is also distributed to all clients, so from a role perspective, 
A is like the homeowner, and all other clients are members.   


#### server slide
```bash
./journal exchange -s -n SERVER
```

#### Output

```bash
SERVER HOST AT 127.0.0.1:9651
```

#### client
```bash
./journal exchange -c -H 127.0.0.1 -n CLIENT
```
