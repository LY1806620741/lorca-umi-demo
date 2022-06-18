# umi project

## Getting Started

Install dependencies,

```bash
$ yarn
```

Start the dev server,

```bash
$ yarn start
```


# 手动构建
```bash
# 创建umi https://umijs.org/zh-CN/docs/getting-started
yarn create @umijs/umi-app
yarn

# go 初始化
go mod init jieshao.loract-umi-demo
go get github.com/zserge/lorca

# 配置 go lorca ... 参考以下
# https://github.com/zserge/lorca/tree/master/examples
# https://github.com/cnbattle/lorca-vue-demo
# https://github.com/erkkah/lorca-ts-react-starter


go mod vendor
```

# 开发
```shell
# umi 运行
yarn start

# go 运行 vscode F5
go run .
```

# 构建
```shell
yarn build
yarn go:build
```