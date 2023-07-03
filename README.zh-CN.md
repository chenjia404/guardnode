[English](./README.md) | [简体中文](./README.zh-CN.md) 
## guardnode

session 守卫节点，请求先到这里，然后再反向代理

### 参数说明

| 字段        | 类型          | 说明                             |
|-----------|-------------|--------------------------------|
| l         | ip          | 监听的地址，默认值 127.0.0.1:18080      |
| f      | Address          | 转发到上级网站 eg:https://google.com  |
| update    | bool        | 从GitHub更新最新版，会验证升级包签名、sha512   |


### 使用说明

识别 header 里面 `o-host` 字段作为原始host字段。


### 更新

`./guardnode -update`

程序会自动从GitHub更新最新版版本，会验证文件的sha512和gpg签名，gpg签名id为 `189BE79683369DA3`

### 打包

`goreleaser release --skip-publish --skip-validate --clean`


### docker
docker run --name session-guard-node --restart unless-stopped -p 18080:18080   chenjia404/guardnode
