# API 文档

## 数据交换模式

在所有的 POST 请求中使用 JSON 格式作为请求体格式，所有服务端响应使用 JSON 格式。响应格式如下：

```json
{
    "success": true,
    "hint": "",
    "data": {} 
}
```

下列接口定义中，服务器响应格式均指 data 中的子文档格式。

如果遇到错误，则 `success` 值一定为 `false`，且 `hint` 字段中包含错误的具体内容。

### POST /user/register

新用户注册，全部字段都不可为空。

请求：

```json
{
    "username": "用户名",
    "password": "密码",
    "profile_name": "昵称",
    "profile_bio": "简介",
    "profile_blog": "博客",
    "profile_twitter_username": "推特",
    "profile_company": "公司",
    "profile_location": "地址"
}
```

响应：

```json
{
    "_id": "用户 ID"
}
```

### POST /user/update

修改用户个人信息。可以修改的字段有：密码、昵称、简介、博客、推特、公司、地址，请求体中包含哪些字段就修改哪些。

请求：

```json
{
    "_id": "用户 ID",
    "password": "密码",
    "profile_name": "昵称",
    "profile_bio": "简介",
    "profile_blog": "博客",
    "profile_twitter_username": "推特",
    "profile_company": "公司",
    "profile_location": "地址"
}
```

响应：无

### POST /use/session

登录请求，并创建一个session。

请求：

```json
{
    "username": "用户名",
    "password": "密码"
}
```