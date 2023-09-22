# lorca + umi base demo project

[中文说明/Chinese Documentation](./README_CN.md)

## version
```
os: win10
umi 3.5.26
node v16.20.2
go1.21.1
chrome 117.**
```
nodejs always has unexpected version problem. please note.

## Getting Started

Install dependencies,

```bash
$ yarn
```

Start the dev server, bind 8000 port.

```bash
$ yarn start
```

Start the go dev server, use shell, show localhost:8000

```bash
go run .
```

or vscode Shortcut key `F5` when open any go file


# By hand create this project
```bash
# create umi link to https://umijs.org/docs/getting-started
yarn create @umijs/umi-app
yarn

# go init
go mod init jieshao.loract-umi-demo
go get github.com/zserge/lorca

# config go lorca ... link to
# https://github.com/zserge/lorca/tree/master/examples
# https://github.com/cnbattle/lorca-vue-demo
# https://github.com/erkkah/lorca-ts-react-starter


go mod vendor
```

# build
```shell
# build resource
yarn build
# build binary program
yarn go:build:win
```