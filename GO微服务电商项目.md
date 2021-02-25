# 第1章 GO微服务介绍、Docker 及go-micro 入门

## 1.1 微服务基本介绍
### 微服务介绍
- 什么是微服务?
  - 首先他是一种架构模式
  - 	相比较单体架构，微服务架构更独立，能够单独更新发布
  - 	微服务里面的服务仅仅用于某一个特定的业务功能
- 为什么需要微服务?
  - 	逻辑清晰
  - 	快速迭代
  - 	多语言灵活组合
- 微服务中的DDD是什么?
  - 	领域驱动设计（Domain Driven Design，简称DDD ）
  - 	定律∶康威定律（ Conway's Law ）
  - DDD作用——真正决定软件复杂性的是设计方法
    - 		有助于指导我们确定系统边界
    - 		能够聚焦在系统核心元素上
    - 		帮助我们拆分系统
  - DDD 常用概念-领域
    - 		领域：领域是有范围界限的，也可以说是有边界的
    - 		核心域：核心域是业务系统的核心价值
    - 		通用子域：所有子域的消费者，提供着通用服务
    - 		支撑子域：专注于业务系统的某一重要的业务
  - DDD常用概念-界限上下文
    - 理解：语文中的语境的意思
    - 方式：领域+界限上下文
    - 目的：不在于如何划分边界，而在于如何控制边界
  - DDD常用概念-领域模型
    - 理解：领域模型是对我们软件系统中要解决问题的抽象表达。
    - 领域：反应的是我们业务上需要解决的问题
    - 模型：我们针对该问题提出的解决方案

### DDD域微服务四层架构

![image-20210222223247199](C:\Users\25407\AppData\Roaming\Typora\typora-user-images\image-20210222223247199.png)

### 微服务架构

![image-20210222223309617](C:\Users\25407\AppData\Roaming\Typora\typora-user-images\image-20210222223309617.png)

### 回到微服务的设计原则上

- 要领域驱动设计，而不是数据驱动设计，也不是界面驱动设计
- 要边界清晰的微服务，而不是泥球小单体
- 要职能清晰的分层，而不是什么都放的大箩筐
- 要做自己能hold住的微服务，而不是过度拆分的微服务

## 1.2 Docker快速入门和使用

### Docker介绍
-	为什么需要Docker ?
  -	软件更新发布及部署低效，过程繁琐且需要人工介入
  -	环境一致性难以保证，不同环境之间迁移成本太高
  -	构建容易分发简单
-	应用场景？
  -	构建运行环境
  -	微服务
  -	CI/ CD
-	Docker的重要概念
  -	客户端Client：可运行docker指令
  -	服务器进程（Docker Daemon ) ：管理镜像和容器
  -	镜像仓库：存储镜像的仓库
-	Docker环境安装
-	Docker基本操作
  -	Docker 仓库操作: pull , push
  -	Docker镜像管理: images , rmi， build
  -	Docker生命周期管理:run , start , stop, rm

## 1.3 go-micro基础: Grpc和ProtoBuf

### RPC和gRPC介绍

- RPC是什么?
  - RPC代指远程过程调用(Remote Procedure Call )
  - 包含了传输协议和编码（对象序列号)协议
  - 允许运行于一台计算机的程序调用另一台计算机的子程序
- 使用RPC有什么好处?
  - 简单、通用、安全、效率
- gRPC又是什么?
  - gRPC是一个高性能、开源、通用的RPC框架
  - 基于HTTP2.0协议标准设计开发
  - 支持多语言，默认采用Protocol Buffers数据序列化协议

### ProtoBuf及详细语法介绍

```protobuf
//版本号
syntax = "proto3";

option go_package = "./;go.micro.service.product";//目标地址(相对于当前);包名

//包名
package go.micro.service.product;
//服务
service Product{ //命名规范：开头大写，驼峰
  rpc AddProduct(ProductIn) returns (ProductOut) {}
}
//消息格式
message ProductIn{
  int64 id = 1;//不是赋值，是字段标识符
  string product_name = 2;
}

message ProductOut{
  int64 product_id = 1;
}

```

### 使用Proto生成go文件

```shell
#路径下放入proto.exe
go get github.com/golang/protobuf/protoc-gen-go

#命令
protoc -I=./ --go_out=./ *.proto
protoc -I=./ --micro_out=./ *.proto
```

## 1.4go-micro组件架构及通讯原理

### Micro简介

- Micro是什么?
  - 用来构建和管理分布式程序的系统
  - Runtime（运行时）：用来管理配置，认证，网络等
    - Micro其中 Runtime (运行时）介绍
      - 他是工具集，工具名称是“micro"
      - 官方docker 版本是docker pull micro/micro
      - 课程扩展版本是docker pull cap1573/cap-micro
    - Micro其中 Runtime (运行时）组成
      - api : api网关
      - broker :允许异步消息的消息代理
      - network :通过微网络服务构建多云网络
      - new :服务模板生成器
      - proxy :建立在Go Micro 上的透明服务代理
      - registry :一个服务资源管理器
      - store :简单的状态存储
      - web : Web 仪表板允许您浏览服务
  - Framework（程序开发框架）：用来方便编写微服务
    - Micro其中Framework(go-micro)介绍
      - 它是对分布式系统的高度抽象
      - 提供分布式系统开发的核心库
      - 可插拔的架构，按需使用
    - Micro其中 Framework(go-micro)组件
      - 注册（Registry )︰提供了服务发现机制
      - 选择器（ Selector ) :能够实现负载均衡
      - 传输( Transport )︰服务与服务之间通信接口
      - Broker:提供异步通信的消息发布/订阅接口
      - 编码( Codec ):消息传输到两端时进行编码与解码
      - Server (服务端）
      - Client (客户端)
  - Clients (多语言客户端）：支持多语言访问服务端

### Micro其中 Framework(go-micro)组件架构图

![image-20210224094857235](C:\Users\25407\AppData\Roaming\Typora\typora-user-images\image-20210224094857235.png)

### Micro其中 Framework(go-micro)通信图

![image-20210224095008925](C:\Users\25407\AppData\Roaming\Typora\typora-user-images\image-20210224095008925.png)

### 开发流程

- Proto编写
- 服务端编写
- Client端编写

```shell
go get github.com/micro/go-micro


```

# 第二章 微服务模块开发

## 2.1 Docker在开发中的应用

### docker 中使用micro &项目目录搭建

- 项目目录搭建

  - 使用micro new生成项目初始目录

  - 在项目初始目录上添加目录

  - docker pull micro/micro

  - sudo docker pull micro/micro

  - 慕课网创建项目(演示)imocc

  - docker run --rm -v $(pwd):$(pwd) -w $(pwd) micro/micro new user（模块名称)

  - 比如你的目录是

    ```
    C:\docker.image.data\redis\data
    ```

    则需要写成

    ```
    /c/docker.image.data/redis/data
    
    C:\Users\25407\Desktop\DDMO>docker run --rm -v /C/Users/25407/Desktop/DDMO:/user -w /user micro/micro new user
    ```

```text
├── micro.mu
├── main.go
├── generate.go
├── handler
│   └── user.go
├── proto
│   └── user.proto
├── Dockerfile
├── Makefile
├── README.md
├── .gitignore
└── go.mod
```



### go module私有化设置& gorm使用

#### go module

- Go module设置演示
  - Go module基本设置
  - Go module私有仓库设置
  - 慕课网ssh配置
- Go module私有仓库设置
  - 使用go env命令查看本机参数
  - GOPRVATE="*.imooc.com"
- Go module加速设置
  - Linux或macOs export GOPROXY="https://goproxy.io
  - Linux或macOS固化可以把命令写到.bashrc或.bash_profile 文件当中
  - Windows cmd命令行设置set GOPROXY = "https://goproxy.io"
- Go module加速设置
  - Linux或macOs export GOPROXY="https://goproxy.io
  - Linux或macOS固化可以把命令写到.bashrc或.bash_profile 文件当中
  - Windows cmd命令行设置set GOPROXY = "https://goproxy.io"
- Go module 设置慕课网git转发
  - go get内部使用git clone命令，默认只支持公有仓库
  - 替换https为 ssh请求（注意，慕课网ssh端口是80不是22)
  - git config --global url."ssh:/lgit@git.imooc.com:80/" ,jnsteadOf"https://git.imooc.com/"

#### GORM基本介绍及使用

- 什么是GORM ?
  - go语言中实现数据库访问的ORM（对象关系映射）库
  - 使用这个库，我们可以利用面向对象的方法
  - 方便的对数据库中的数据进行CRUD(增删改查)
- GORM依赖库说明
  - go get github.com/jinzhu/gorm
  - go get github.com/go-sql-driver/mysql
  - gorm.Open("mysql","root:123456@/test?charset=utf8&parseTime=True&loc=Local")

### 电商用户领域功能开发& Docker打包go程序

#### 用户功能梳理及开发

- 用户领域需求分析
  - 用户信息管理
  - 用户登陆，注册
  - 用户鉴权
- go mod init github.com/XXXYYYZZZLB/micro-user



```shell
C:\Users\25407\Desktop\DDMO\user>protoc -I=./ --micro_out=./proto/user/ ./proto/user/user.proto

C:\Users\25407\Desktop\DDMO\user>protoc -I=./ --go_out=./proto/user/ ./proto/user/user.proto

go get github.com/jinzhu/gorm


```

