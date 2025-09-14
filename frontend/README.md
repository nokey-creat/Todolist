# Todo List 前端项目

这是一个基于React和Ant Design构建的Todo List应用前端。该应用提供了任务管理的完整功能，包括用户注册、登录、创建任务、编辑任务、标记完成状态以及删除任务等功能。

## 功能特点

- 用户认证：支持用户注册和登录
- 任务管理：创建、编辑、删除和查看任务
- 任务状态：可以标记任务为已完成或未完成
- 截止日期：支持设置任务的截止日期
- 响应式设计：适配各种屏幕尺寸

## 技术栈

- React: 用于构建用户界面的JavaScript库
- React Router: 用于页面路由
- Ant Design: UI组件库
- Axios: 用于发送HTTP请求
- Dayjs: 日期处理库

## 如何运行

1. 确保已安装Node.js (推荐v14或更高版本)
2. 安装依赖：
   ```
   npm install
   ```
3. 启动开发服务器：
   ```
   npm start
   ```
4. 打开浏览器访问 [http://localhost:3000](http://localhost:3000)

## 后端API

该前端应用依赖于Golang实现的后端API，主要包括以下接口：

### 认证相关
- POST /api/auth/register: 用户注册
- POST /api/auth/login: 用户登录

### 任务相关
- GET /api/tasks: 获取所有任务
- POST /api/tasks: 创建新任务
- GET /api/tasks/:id: 获取单个任务详情
- PATCH /api/tasks/:id: 更新任务
- DELETE /api/tasks/:id: 删除任务
- PATCH /api/tasks/:id/completed: 切换任务完成状态

## 构建生产版本

```
npm run build
```

这将在 `build` 文件夹中生成应用的生产版本。
