# 自建图床

* 监视本地文件夹
* 上传新增文件到私有 minio
* 自动将地址复制到剪切板
* 需要的环境变量

```bazaar
    ssl: true/false use http or https ,default false
    endpoint: minio server url , without "http" or "https"
    accessKey: minio 
    secretKey: minio
    watch_folder: local folder to watch 
```
 * 支持平台osx (linux, window not tested)