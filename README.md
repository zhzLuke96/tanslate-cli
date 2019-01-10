# tanslate-cli
 Cli for assisting translation with online translation tools.

# Usage
```
$ trans --help
```

```
trans-cli version: trans-cli/0.0.1
Usage: trans [-f filename] [-c configfile] [-l logres] [-o outputfile]
  -c string
        config file name (default "./conf.json")
  -f string
        file name (default "none")
  -l    show result on stdout
  -o string
        output file name (default "./translate.txt")
```

# conf.json
```json
{
    "default":"baidu",
    "baidu":{
        "url":"https://fanyi-api.baidu.com/api/trans/vip/translate",
        "appid":"xxxxx",
        "key":"xxxxx"
    },
    "bing":{
        "url":""
    }
}
```
