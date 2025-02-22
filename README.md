# Go Yggdrasil Server

使用 Go 语言 Gin + GORM 框架编写的 Minecraft 登录协议服务端。

## 功能

+ 实现了 Minecraft 登录服务器时的认证部分以及材质部分。支持注册。
+ 兼容 [authlib-injector](https://github.com/yushijinhun/authlib-injector) 。
+ 支持使用在线账号（正版账号）登录，起到透明代理的功能。

## 用途

用于服务器管理员调试和测试时使用小号登录而不必关闭在线验证 (online-mode)。


禁止违反 [EULA](https://account.mojang.com/documents/minecraft_eula) 的行为。

## 用法

下载或编译得到可执行文件并运行，将会自动生成所需的配置文件和数据库文件。

配置文件格式详见 `config_example.ini`，请重命名为 `config.ini` 并放在执行目录下。

启动成功后在启动器（请使用第三方启动器）外置登录选项上填写运行的 URL 的根路径，比如 `http://localhost:8080`。

注册地址在 `/profile/`。

## 编译命令

- 先在frontent目录下运行`npm install && npm run build`

- 之后将dist目录复制进根目录并改名assets

- 最后在根目录下运行go编译（以需要sqlite和mysql为例）

```shell
go build --tags 'nomsgpack sqlite mysql'
```
