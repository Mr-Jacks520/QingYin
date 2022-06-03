# 项目结构

```shell
├── api
├── config
├── core
├── global
├── initialize
├── middleware
├── model
│   ├── request
│   └── response
├── resource
├── router
├── service
└── utils
```

| 文件夹       | 说明                    | 描述                        |
| ------------ | ----------------------- | --------------------------- |
| `api`     | 接口/controller层             | 相当于controller层这里参考的下面那个项目的架构,比较复杂 |
| `config`     | 配置包                  | config.yaml对应的配置结构体 |
| `core`       | 核心文件                | 核心组件(zap, viper, server)的初始化 |
| `global`     | 全局对象                | 全局对象 |
| `initialize` | 初始化 | router,redis,gorm,validator, timer的初始化 |
| `middleware` | 中间件层 | 用于存放 `gin` 中间件代码 |
| `model`      | 模型层                  | 模型对应数据表              |
| `--request`  | 入参结构体              | 接收前端发送到后端的数据。  |
| `--response` | 出参结构体              | 返回给前端的数据结构体      |
| `resource`   | 静态资源文件夹          | 负责存放静态文件                |
| `router`     | 路由层                  | 路由层 |
| `service`    | service层               | 存放业务逻辑问题 |
| `utils`      | 工具包                  | 工具函数封装            |

## 项目架构参考

[相关项目](https://github.com/flipped-aurora/gin-vue-admin.git)

## 5.31-6.1测试访问接口结果

![image-20220601002408274.png](https://s2.loli.net/2022/06/02/IB84bySNMqTWX3P.png)
=======
![image-20220601002408274](https://s2.loli.net/2022/06/01/ld5gOEVLxCpUZsR.png)

**并未实现业务逻辑，只是进行了伪数据测试**

## 6.1-6.2数据库结构搭建

![image-20220602012937404.png](https://s2.loli.net/2022/06/02/SUCMiGgj7YnIhPy.png)

![image-20220602013202537.png](https://s2.loli.net/2022/06/02/r7FaAnNHQoUiqhS.png)

## 6.2-6.3用户业务逻辑实现

![register.png](https://s2.loli.net/2022/06/03/13siChgURbXFw9v.png)

![login.png](https://s2.loli.net/2022/06/03/WHEznYVQ2ob8gSl.png)

![userinfo.png](https://s2.loli.net/2022/06/03/YDArcEuLqzP6M7d.png)

![debug.png](https://s2.loli.net/2022/06/03/ALIwj9O4cRXbDsZ.png)
