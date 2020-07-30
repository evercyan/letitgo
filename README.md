# letitgo

[![Build Status](https://www.travis-ci.org/evercyan/letitgo.svg?branch=master)](https://www.travis-ci.org/evercyan/letitgo)
[![codecov](https://codecov.io/gh/evercyan/leetcli/branch/master/graph/badge.svg?token=RbJTUtAlvl)](https://codecov.io/gh/evercyan/letitgo)

> some utils that use frequently

more info in *_test.go

```
.
|____crypto
| |____encode.go        # 常用的一些 encode 操作, 如 Base64Encode, JsonEncode..
| |____aes.go           # AES 加解密
| |____snowflake.go     # 雪花 id
| |____jwt.go           # golang jwt(github.com/dgrijalva/jwt-go)
|____util
| |____util.go          # 工具集合, Md5, Guid, Rand...
| |____is.go            # Is 集合, IsInt, IsEmail...
| |____to.go            # To 集合, ToString, ToCamelCase...
|____redis
| |____redis.go         # 基于 redigo 封装
|____json
| |____json.go          # 动态取 json 字符串数据
|____request
| |____http.go          # 简单的 http 请求
|
```
