

### 项目概述

这是一个使用 Go 语言和 Gin 框架构建的 ToDoList 应用程序。它包括用户注册、登录、任务管理等功能，并且使用 JWT 进行身份验证。

---

## 项目目录结构

```plaintext
ToDoList/
├── api/
│   ├── tasks.go
│   └── user.go
├── conf/
│   ├── conf.go
│   └── config.ini
├── middleware/
│   └── jwt.go
├── model/
│   ├── init.go
│   ├── migrate.go
│   ├── tasks.go
│   └── user.go
├── pkg/
│   └── utils/
│       └── utils.go
├── routes/
│   └── routers.go
├── serializer/
│   ├── common.go
│   ├── task.go
│   └── user.go
├── service/
│   ├── task.go
│   └── user.go
└── main.go
```


---

## 项目依赖

以下是项目中使用的依赖包及其版本信息：

```go
module ToDoList

go 1.23.3

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/sessions v1.0.2
	github.com/gin-gonic/gin v1.10.0
	github.com/jinzhu/gorm v1.9.16
	github.com/sirupsen/logrus v1.9.3
	golang.org/x/crypto v0.32.0
	gopkg.in/ini.v1 v1.67.0
)

require (
	github.com/bytedance/sonic v1.11.6 // indirect
	github.com/bytedance/sonic/loader v0.1.1 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.20.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/gorilla/context v1.1.2 // indirect
	github.com/gorilla/securecookie v1.1.2 // indirect
	github.com/gorilla/sessions v1.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.8.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
```


---

## 配置文件

### `conf/config.ini`

```ini
[service]
AppMode = "debug"
HttpPort = ":8080"

[mysql]
Db = "todolist"
DbHost = "localhost"
DbPort = "3306"
DbUser = "root"
DbPassWord = "password"
DbName = "todolist"
```


### `conf/conf.go`

```go
package conf

import (
	"ToDoList/model"
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	AppMode  string
	HttpPort string

	RedisAddr  string
	RedisPw    string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径")
	}
	LoadServer(file)
	LoadMysql(file)
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	model.Database(path)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
```


---

## API 接口说明

### 用户模块

#### 注册用户

- **URL**: `/user/register`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
      "user_name": "string",
      "password": "string"
  }
  ```

- **Response**:
    - **Success**:
      ```json
      {
          "status": 200,
          "msg": "注册成功"
      }
      ```

    - **Error**:
      ```json
      {
          "status": 400,
          "msg": "用户名已存在"
      }
      ```


#### 登录用户

- **URL**: `/user/login`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
      "user_name": "string",
      "password": "string"
  }
  ```

- **Response**:
    - **Success**:
      ```json
      {
          "status": 200,
          "data": {
              "user": {
                  "id": 1,
                  "user_name": "string"
              },
              "token": "string"
          },
          "msg": "登录成功"
      }
      ```

    - **Error**:
      ```json
      {
          "status": 400,
          "msg": "用户不存在"
      }
      ```


### 任务模块

#### 创建任务

- **URL**: `/tasks`
- **Method**: `POST`
- **Headers**:
    - `Authorization: Bearer <token>`
- **Request Body**:
  ```json
  {
      "title": "string",
      "content": "string",
      "status": 0
  }
  ```

- **Response**:
    - **Success**:
      ```json
      {
          "status": 200,
          "msg": "任务创建成功"
      }
      ```

    - **Error**:
      ```json
      {
          "status": 400,
          "msg": "任务创建失败"
      }
      ```


#### 获取任务详情

- **URL**: `/tasks/{id}`
- **Method**: `GET`
- **Headers**:
    - `Authorization: Bearer <token>`
- **Response**:
    - **Success**:
      ```json
      {
          "status": 200,
          "msg": "任务查询成功",
          "data": {
              "id": 1,
              "title": "string",
              "content": "string",
              "status": 0,
              "start_time": 1672531199,
              "end_time": 0
          }
      }
      ```

    - **Error**:
      ```json
      {
          "status": 400,
          "msg": "任务查询失败"
      }
      ```


#### 获取任务列表

- **URL**: `/tasks`
- **Method**: `GET`
- **Headers**:
    - `Authorization: Bearer <token>`
- **Query Parameters**:
    - `page_num`: int (默认 1)
    - `page_size`: int (默认 15)
- **Response**:
    - **Success**:
      ```json
      {
          "status": 200,
          "msg": "success",
          "data": {
              "item": [
                  {
                      "id": 1,
                      "title": "string",
                      "content": "string",
                      "status": 0,
                      "start_time": 1672531199,
                      "end_time": 0
                  }
              ],
              "total": 1
          }
      }
      ```

    - **Error**:
      ```json
      {
          "status": 400,
          "msg": "任务查询失败"
      }
      ```


#### 更新任务

- **URL**: `/tasks/{id}`
- **Method**: `PUT`
- **Headers**:
    - `Authorization: Bearer <token>`
- **Request Body**:
  ```json
  {
      "title": "string",
      "content": "string",
      "status": 0
  }
  ```

- **Response**:
    - **Success**:
      ```json
      {
          "status": 200,
          "msg": "任务更新成功"
      }
      ```

    - **Error**:
      ```json
      {
          "status": 400,
          "msg": "任务更新失败"
      }
      ```


#### 删除任务

- **URL**: `/tasks/{id}`
- **Method**: `DELETE`
- **Headers**:
    - `Authorization: Bearer <token>`
- **Response**:
    - **Success**:
      ```json
      {
          "status": 200,
          "msg": "任务删除成功"
      }
      ```

    - **Error**:
      ```json
      {
          "status": 400,
          "msg": "任务删除失败"
      }
      ```


---

## 中间件

### JWT 中间件 (`middleware/jwt.go`)

用于解析和验证请求中的 JWT Token，确保用户身份合法。

---

## 主函数 (`main.go`)

```go
package main

import (
	"ToDoList/conf"
	"ToDoList/routes"
)

func main() {
	// 1. 加载配置文件
	conf.Init()
	// 2. 初始化路由
	r := routes.NewRouter()
	// 3. 启动服务
	r.Run(conf.HttpPort)
}
```


---

## 数据模型 (`model/user.go`, `model/tasks.go`)

定义了应用程序中使用的数据模型，包括用户和任务。

---

## 序列化器 (`serializer/common.go`, `serializer/task.go`, `serializer/user.go`)

用于将数据模型转换为 JSON 格式，以便在 API 响应中返回。

---

## 服务层 (`service/user.go`, `service/task.go`)

包含业务逻辑处理，如用户注册、登录、任务创建、更新等。

---

## 路由配置 (`routes/routers.go`)

定义了应用程序的路由规则，将 HTTP 请求映射到相应的处理函数。

---

