# Todolist

一个todo list功能的web应用，实现了前后端分离。
后端提供了注册和登录、增删查改待办事项、更改完成情况的api，前端由ai生成。

项目demo: https://todolist.noclavis.com


## 二编

原代码架构比较混乱，现在将项目重构为handler-service-repository架构。
（不过大部分业务逻辑都很简单，解耦太多可能变得更复杂）

- handler 负责处理请求，获取请求内容后，传入具体的业务逻辑的控制器
- service 实现具体业务逻辑，调用models层访问数据
- repository(models) 封装与数据库交互的操作，向外提供操作数据的接口

其他小的修改：
- 去除了数据库自动迁移
- 统一处理错误，例如init等函数的错误返回到main中处理. (记得处理所有错误！)
- 删去global包


### 现项目结构

```
├─common   
│  ├─config                //配置相关
│  ├─middleware       //中间件
│  └─utils                   //工具函数
├─handler                 //handler层
├─models                  //数据层
├─router                    //路由
└─service                   //业务逻辑层
```

### 还需要改进的地方
- 日志处理，错误打包
- 统一api，目前比较混乱
- 尝试加入缓存层
- 加入单元测试

  