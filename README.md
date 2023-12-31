[English](./README.md) | [简体中文](./README.zh-CN.md)

## guardnode

The session guard node, the request goes here first, and then the reverse proxy

### Parameter Description

| Field  | Type             | Description                                                                                 |
|--------|------------------|---------------------------------------------------------------------------------------------|
| l      | Address          | Listening address, the default value is 0.0.0.0:18080                                     |
| f      | Address          | Forward to parent website eg:https://google.com                                             |
| update | bool             | Update the latest version from GitHub, it will verify the upgrade package signature, sha512 |


### Instructions for use

Identify the `o-host` field in the header as the original host field.

### build

` go build -trimpath -ldflags="-w -s" `

### upgrade

`./guardnode -update`

After v0.0.6, the program will automatically update the latest version from GitHub, and verify the sha512 and gpg signature of the file. The gpg signature id is `189BE79683369DA3`

### releases

`goreleaser release --skip-publish --skip-validate --clean`


### docker
`docker run --name session-guard-node --restart unless-stopped -p 18080:18080   chenjia404/guardnode`

### Linux 

last https://github.com/chenjia404/guardnode/releases
```
wget https://github.com/chenjia404/guardnode/releases/download/v0.0.2/guardnode_0.0.2_linux_amd64.tar.gz 
tar -vxzf guardnode_0.0.2_linux_amd64.tar
chmod +x guardnode
./guardnode -l 0.0.0.0:18080
```

session proxy running successfully