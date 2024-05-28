# go frame

> 本地环境要安装 gf 环境，如果没有，可以根据[官网](https://goframe.org/pages/viewpage.action?pageId=1114399)介绍进行安装。

- clone 项目
  ```sh
  $ git clone -b myapp https://github.com/clin211/goframe-practice.git
  ```
- 安装依赖
  ```sh
  $ go mod tidy
  ```

- 启动项目
  ```sh
  $ gf run main.go
  ```

## 项目架构介绍
```tree
.
├── api                 // 对外提供服务的输入/输出数据结构定义。考虑到版本管理需要，往往以api/xxx/v1...存在
│   └── hello
├── hack                // 存放项目开发工具、脚本等内容。例如，CLI工具的配置，各种shell/bat脚本等文件
│   ├── config.yaml
│   ├── hack-cli.mk
│   └── hack.mk
├── internal            // 业务逻辑存放目录
│   ├── cmd             // 命令行管理目录
│   ├── consts          // 项目所有常量定义
│   ├── controller      // 接收/解析用户输入参数的入口/接口层
│   ├── dao             // 数据访问对象，这是一层抽象对象，用于和底层数据库交互，仅包含最基础的 CURD 方法
│   ├── logic           // 业务逻辑封装管理，特定的业务逻辑实现和封装。往往是项目中最复杂的部分
│   ├── model           // 数据结构管理模块，管理数据实体对象，以及输入与输出数据结构定义
│   │   ├── do          // 用于dao数据操作中业务模型与实例模型转换，由工具维护，用户不能修改
│   │   └── entity      // 数据模型是模型与数据集合的一对一关系，由工具维护，用户不能修改
│   ├── packed
│   └── service         // 用于业务模块解耦的接口定义层。具体的接口实现在 logic 中进行注入
├── manifest            // 包含程序编译、部署、运行、配置的文件
│   ├── config          // 配置文件存放目录
│   ├── deploy          // 部署相关的文件。默认提供了Kubernetes集群化部署的Yaml模板，通过 kustomize 管理
│   ├── docker          // Docker镜像相关依赖文件，脚本文件等等
│   ├── i18n            // 国际化
│   └── protobuf        // GRPC协议时使用的protobuf协议定义文件，协议文件编译后生成go文件到api目录
├── resource            // 静态资源文件
│   ├── public
│   └── template
├── utility
├── go.mod              // 使用 Go Module 包管理的依赖描述文件
├── go.sum
├── Makefile
├── README.MD           // 描述文件
└── main.go             // 程序入口文件
```