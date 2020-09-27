![Author](https://img.shields.io/badge/Author-4everlynn-red)
![Golang](https://img.shields.io/badge/100%25-Golang-olive)
# Journal

> A small but beautiful, simple and performance-free job report command line job (CLI) (based on code version management warehouse)

# Features

- No need for Git environment, just a project folder managed by Git
- Automatically extract today's report according to the module
- Support closing part of warehouse file extraction
- External configuration file support

# Getting Start

## .journal.yaml

```yaml
version: '1'
maintainer: Edward
git:
  boneco-admin:
    name: 发货宝管理界面
    # path to your 
    path: /Users/disware/Documents/Java/boneco-backend
    disable: true
  boneco-portal:
    name: 发货宝管理入口界面
    path: /Users/disware/Vue/hitotek/portal
  boneco-api:
    name: 发货宝接口
    path: /Users/disware/Documents/Java/boneco-backend
  journal:
    name: Journal
    path: /Users/disware/Documents/Go/src/diswares.com.journal
```

## Usage

```bash
./journal stylish
```

### Output


```text
Config maintained by Edward is loaded successfully
20200927 日报
Journal
  1、first commit
  2、Added the option of disable report generation in the warehouse
发货宝管理入口界面
  1、add ci
  2、add ci
  3、CI 每次都运行(临时)
  4、cleanup image
  5、仓库地址修改
发货宝接口
  1、commit all
  2、commit all
  3、commit all
  4、version、depot、express and region etc.
  5、impl depot & depot-express api
  6、update query-core's version
  7、version of multipart is change, introduce query-core
  8、commit current code
  9、仓库下关联商品、仓库下关联快递，获取仓库与商品、快递的关联关系
  10、更换鉴权网关
  11、commit register client api
  12、client api for auth to register
  13、change genius version
  14、region changed
  15、update region api path
  16、add shipper api
  17、add global ex.h & shipper api recreate
  18、commit sys code
```
