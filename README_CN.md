# lorca + umi 基础示例 工程

## 开始

安装依赖,

```bash
$ yarn
```

开启dev服务,

```bash
$ yarn start
```

开启lorca go界面程序

```bash
go run .
```

或者在打开的go文件中使用vscode的`F5`快捷键


# 手动构建工程
```bash
# 创建umi 参考 https://umijs.org/zh-CN/docs/getting-started
yarn create @umijs/umi-app
yarn

# go 初始化
go mod init jieshao.loract-umi-demo
go get github.com/zserge/lorca

# 配置 go lorca ... 参考以下项目
# https://github.com/zserge/lorca/tree/master/examples
# https://github.com/cnbattle/lorca-vue-demo
# https://github.com/erkkah/lorca-ts-react-starter


go mod vendor
```

# 构建
```shell
# 构建资源文件
yarn build
# 构建二进制文件
yarn go:build:win
```