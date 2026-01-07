<p align="center">
  <strong>📦 file_trans 前后端分离的临时文件 / 文本传输系统</strong><br/>
  安全 · 简洁 · 可扩展
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.23+-00ADD8?logo=go" />
  <img src="https://img.shields.io/badge/Vue-3-42b883?logo=vue.js" />
  <img src="https://img.shields.io/badge/Gin-Web_Framework-blue" />
  <img src="https://img.shields.io/badge/License-MIT-green" />
</p>

## ✨ 项目简介

**file_trans** 是一个用于 **临时文件 / 文本安全传输** 的 Web 应用，采用 **前后端分离架构**，通过 **6 位取件码 + 过期机制** 实现文件的安全分享。

适用于：

- 内部文件中转
- 临时文件共享
- Demo / 教学 / 项目模板

------

## 🧩 系统架构

```text
file_trans
├── backend_file_trans/   # Go 后端服务
└── frontend_file_trans/  # Vue 3 前端应用
```

------

## 🚀 核心功能

### 📤 文件发送

- 文本 / 文件上传
- 多文件自动 ZIP 压缩
- 自动生成 6 位取件码
- 支持文件过期时间

### 📥 文件接收

- 取件码校验
- 预签名下载链接
- 文件状态管理

### 🌗 前端体验

- 深色 / 浅色主题切换
- 系统主题自动适配
- 平滑动画与现代 UI

### 📈 上传进度

- 组合式函数封装上传逻辑
- 单文件 / 总进度同步
- 可无缝替换真实接口

------

## 🧰 技术栈

### 后端

| 分类     | 技术     |
| -------- | -------- |
| 语言     | Go 1.23+ |
| 框架     | Gin      |
| ORM      | GORM     |
| 数据库   | MySQL    |
| 缓存     | Redis    |
| 对象存储 | MinIO    |
| 架构     | MVC      |

### 前端

| 分类 | 技术            |
| ---- | --------------- |
| 框架 | Vue 3           |
| API  | Composition API |
| 路由 | Vue Router      |
| 样式 | CSS Variables   |
| 架构 | Composables     |

------

## ▶️ 快速开始

### 启动后端

```bash
cd backend_file_trans
go mod tidy
go run main.go
```

### 启动前端

```bash
cd frontend_file_trans
npm install
npm run dev
```

------

## 📁 项目结构（简要）

```text
backend_file_trans/
├── controller/
├── service/
├── repository/
├── model/
└── router/

frontend_file_trans/
├── composables/
├── components/
├── views/
└── router/
```

------

## 🔌 前后端对接

- 前端通过 `axios + FormData` 上传文件
- 后端返回取件码与下载信息
- 下载使用 MinIO 预签名 URL

------

## 🌟 项目亮点

- ✅ 前后端分离，结构清晰
- ✅ 取件码 + 过期机制
- ✅ MinIO 对象存储
- ✅ Vue 3 组合式 API
- ✅ 适合作为真实项目或模板

------

## 🔮 后续规划

- JWT 用户体系
- 分片 / 断点续传
- 后台管理面板
- 文件加密
- CI / CD

------

## 📄 License

MIT License

------

<p align="center">
  如果这个项目对你有帮助，欢迎 ⭐ Star
</p>

