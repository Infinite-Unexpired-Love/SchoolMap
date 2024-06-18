# 简介

校园地图项目是一款功能完善的后端服务器，现已上线并通过全面测试。该项目采用Go语言、Redis和MySQL数据库技术，确保普通用户能够快速获取数据，并允许管理员灵活地动态修改数据。项目代码结构分明，提供详尽的API文档，支持快速开发前端应用，包括小程序和APP等各种终端，助力服务快速上线。

# 目录

- [简介](#简介)
- [项目环境](#项目环境)
- [快速启动](#快速启动)
- [API文档](#API文档)
  - [登录](#登录)
  - [获取图形验证码](#获取图形验证码)
  - [地标](#地标)
    - [获取所有地标](#获取所有地标)
    - [新建地标](#新建地标)
    - [更新地标](#更新地标)
    - [删除地标](#删除地标)
  - [筛选](#筛选)
    - [获取所有筛选](#获取所有筛选)
    - [新建筛选](#新建筛选)
    - [删除筛选](#删除筛选)
  - [通知](#通知)
    - [获取所有通知](#获取所有通知)
    - [新建通知](#新建通知)
    - [更新通知](#更新通知)
    - [删除通知](#删除通知)
  - [用户反馈](#用户反馈)
    - [获取所有反馈](#获取所有反馈)
    - [新建反馈](#新建反馈)
    - [删除反馈](#删除反馈)

# 项目环境

go_version:1.19.6

redis_version:5.0.14.1

mysql_version:8.0.36

# 快速启动

1. 安装
2. 修改/config.yaml文件中的内容
3. 修改/utils/token.go文件下`var jwtKey = []byte("your_secret_key")`为你自己的key
4. 修改import路径中的`TGU-MAP`为你自己的项目根目录
5. 在项目根目录下运行`go mod tidy`后运行`go run main.go`


# API文档

## POST 登录

POST /login

> Body 请求参数

```json
{
  "mobile": "string",
  "password": "string",
  "captchaId": "string",
  "captcha": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» mobile|body|string| 是 |none|
|» password|body|string| 是 |none|
|» captchaId|body|string| 是 |none|
|» captcha|body|string| 是 |与验证码ID相对应的验证码，验证成功后即删|

> 返回示例

> 200 Response

```json
{
  "token": "string",
  "username": "string",
  "mobile": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» token|string|true|none||none|
|» username|string|true|none||none|
|» mobile|string|true|none||none|

## GET 获取图形验证码

GET /ca

> 返回示例

> 200 Response

```json
{
  "captchaId": "string",
  "picPath": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» captchaId|string|true|none||none|
|» picPath|string|true|none||base64编码|

# 地标

## GET 获取所有地标

GET /li/

返回树形结构数据的json序列化字符串

> 返回示例

> 200 Response

```json
{
  "msg": "string",
  "data": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» msg|string|true|none||none|
|» data|string|true|none||none|

## POST 新建地标

POST /li/item/{id}

在Id=id的父节点下插入节点，id为0的根节点已自动创建
返回所有marker树形结构数据的json序列化字符串
权限接口

> Body 请求参数

```json
{
  "title": "string",
  "desc": "string",
  "contact": "string",
  "latitude": 0,
  "longitude": 0,
  "iconName": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|
|Authorization|header|string| 是 |none|
|body|body|object| 否 |none|
|» title|body|string| 是 |none|
|» desc|body|string| 是 |none|
|» contact|body|string| 否 |none|
|» latitude|body|number| 否 |none|
|» longitude|body|number| 否 |none|
|» iconName|body|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "msg": "string",
  "data": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» msg|string|true|none||none|
|» data|string|true|none||none|

## PUT 更新地标

PUT /li/item/{id}

返回所有marker树形结构数据的json序列化字符串
权限接口

> Body 请求参数

```json
{
  "title": "string",
  "desc": "string",
  "contact": "string",
  "latitude": 0,
  "longitude": 0,
  "iconName": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|
|Authorization|header|string| 是 |none|
|body|body|object| 否 |none|
|» title|body|string| 是 |none|
|» desc|body|string| 是 |none|
|» contact|body|string| 否 |none|
|» latitude|body|number| 否 |none|
|» longitude|body|number| 否 |none|
|» iconName|body|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "msg": "string",
  "data": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» msg|string|true|none||none|
|» data|string|true|none||none|

## DELETE 删除地标

DELETE /li/item/{id}

返回所有marker树形结构数据的json序列化字符串
权限接口

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|
|Authorization|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "msg": "string",
  "data": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» msg|string|true|none||none|
|» data|string|true|none||none|

# 筛选（分类）

## GET 获取所有筛选

GET /al/

返回树形结构数据的json序列化字符串

> 返回示例

> 200 Response

```json
{
  "msg": "string",
  "data": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» msg|string|true|none||none|
|» data|string|true|none||none|

## POST 新建筛选

POST /al/item

返回所有alias树形结构数据的json序列化字符串
权限接口

> Body 请求参数

```json
{
  "title": "string",
  "markers": [
    {
      "id": 0
    }
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 是 |none|
|body|body|object| 否 |none|
|» title|body|string| 是 |none|
|» markers|body|[object]| 是 |none|
|»» id|body|number| 是 |none|

> 返回示例

> 200 Response

```json
{
  "msg": "string",
  "data": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» msg|string|true|none||none|
|» data|string|true|none||none|

## DELETE 删除筛选

DELETE /al/item/{id}

返回所有alias树形结构数据的json序列化字符串
权限接口

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|
|Authorization|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "msg": "string",
  "data": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» msg|string|true|none||none|
|» data|string|true|none||none|

# 通知

## GET 获取所有通知

GET /no/

返回所有数据的json序列化字符串

> 返回示例

> 200 Response

```json
{
  "msg": "string",
  "data": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» msg|string|true|none||none|
|» data|string|true|none||none|

## POST 新建通知

POST /no/item

返回所有notice树形结构数据的json序列化字符串
权限接口

> Body 请求参数

```json
{
  "title": "string",
  "publishTime": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 是 |none|
|body|body|object| 否 |none|
|» title|body|string| 是 |none|
|» publishTime|body|string| 是 |格式为“2024-5-24”|

> 返回示例

> 200 Response

```json
{
  "msg": "string",
  "data": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» msg|string|true|none||none|
|» data|string|true|none||none|

## PUT 更新通知

PUT /no/item/{id}

返回所有notice数据的json序列化字符串
权限接口

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|
|Authorization|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "msg": "string",
  "data": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» msg|string|true|none||none|
|» data|string|true|none||none|

## DELETE 删除通知

DELETE /no/iten/{id}

返回所有notice树形结构数据的json序列化字符串
权限接口

> Body 请求参数

```json
{}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|
|Authorization|header|string| 是 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{
  "msg": "string",
  "data": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» msg|string|true|none||none|
|» data|string|true|none||none|

# 用户反馈

## GET 获取所有反馈

GET /fe/

返回所有feedback的json序列化字符串
权限接口

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "msg": "string",
  "data": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» msg|string|true|none||none|
|» data|string|true|none||none|

## POST 新建反馈

POST /fe/item

用户反馈是给管理员看的，自然不需要返回所有数据给用户

> Body 请求参数

```json
{
  "title": "string",
  "category": 0,
  "contact": "string",
  "publishTime": "string",
  "detail": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» title|body|string| 是 |none|
|» category|body|integer| 是 |0为信息维护，1为反馈建议|
|» contact|body|string| 否 |none|
|» publishTime|body|string| 是 |格式：“2024-5-23”，默认为今天|
|» detail|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» msg|string|true|none||none|

## DELETE 删除feedback

DELETE /fe/item/{id}

返回所有feedback的json序列化字符串
权限接口

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|
|Authorization|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "msg": "string",
  "data": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» msg|string|true|none||none|
|» data|string|true|none||none|

