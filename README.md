# Qizhan 风格企业官网（Ant Design Pro + Kratos + MySQL）

该项目基于你的要求搭建：

- 前端：Ant Design Pro 风格（React + TypeScript + Ant Design）
- 后端：基于 `go-kratos/kratos-layout` 思路组织的服务结构
- 数据库：MySQL（包含初始化 SQL）

## 项目结构

- `frontend/` 前端项目
- `backend/` Kratos 后端服务
- `deploy/mysql/init.sql` MySQL 初始化脚本

## 快速启动

### 1) MySQL

```bash
mysql -uroot -p < deploy/mysql/init.sql
```

### 2) 启动后端

```bash
cd backend
cp configs/config.example.yaml configs/config.yaml
go mod tidy
go run ./cmd/server
```

默认监听：`http://localhost:8000`

接口：`GET /api/about`

### 3) 启动前端

```bash
cd frontend
npm install
npm run dev
```

默认地址：`http://localhost:5173`

## 说明

前端页面参考了 qizhan.ca/about 的信息组织方式：

- 公司介绍
- 业务范围
- 服务优势
- 发展里程碑
- 联系方式

并通过后端接口动态读取内容，可扩展为 CMS 配置化管理。
