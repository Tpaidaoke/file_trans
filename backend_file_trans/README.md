# 📦 file_trans 后端服务

一个基于 **Go 语言** 的文件传输后端服务，支持文本与文件传输，具备 **发送 / 接收 / 过期管理 / 记录查询** 等完整功能，适用于临时文件分享、内部文件中转等场景。

---

## ✨ 项目概述

`file_trans` 是一个基于 **Gin + MySQL + Redis + MinIO** 构建的文件传输后端系统，采用 **MVC 分层架构**，通过取件码机制实现安全的文件访问，并支持过期控制。

**核心能力：**

- 📄 文本 & 文件传输  
- 📦 多文件自动压缩（ZIP）  
- 🔐 6 位取件码机制  
- ⏱️ 过期时间管理  
- ☁️ MinIO 对象存储  
- 🚀 高性能 & 可扩展  

---

## 🧰 技术栈

| 分类 | 技术 |
|----|----|
| 开发语言 | Go 1.23.4 |
| Web 框架 | Gin v1.11.0 |
| ORM | GORM v1.31.1 |
| 数据库 | MySQL |
| 缓存 | Redis |
| 文件存储 | MinIO |
| 配置管理 | YAML |
| 日志 | slog（Go 1.21+） |
| 开发工具 | GoLand / air |

---

## 📁 项目结构

```text
backend_file_trans/
├── main.go
├── config.yaml
├── air.toml
├── go.mod
├── go.sum
├── conf/
├── controller/
├── service/
├── repository/
├── model/
├── database/
├── router/
├── utils/
├── logger/
├── test/
├── logs/
└── README.md
```

---

## 🚀 核心功能

### 文件发送
- 支持文本 / 文件上传  
- 多文件自动 ZIP 压缩  
- 生成 6 位取件码  
- 支持过期时间设置  

### 文件接收
- 取件码校验  
- 生成预签名下载链接  
- 文件状态更新  

### 记录查询
- 当天发送记录  
- 当天接收记录  

---

## ▶️ 启动方式

```bash
cd backend_file_trans
go mod tidy
go run main.go
```

或使用热重载：

```bash
air
```

---

## 🌟 项目特点

- MVC 分层架构清晰
- MySQL + Redis + MinIO 多存储协同
- 取件码 + 过期机制保障安全
- 结构化日志，易于排错
- 高扩展性设计

---

## 🔮 后续扩展

- JWT 鉴权
- 文件分片上传
- 管理后台
- 文件加密
- 通知系统
