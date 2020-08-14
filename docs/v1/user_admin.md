# 用户和管理员接口

### API 前缀: /api/v1/

---

### 1. User 相关接口

###### 接口功能

> - 用户的注册和登陆 【不需要登陆(鉴权)】
> - 个人信息更新
> - 关注和取消关注用户
> - 查看用户的关注和粉丝列表 【不需要登陆】
> - 添加和移除黑名单
> - 查看黑名单列表

#### 用户登陆和注册

```js
// 使用用户名+密码或者手机号+验证码等方式登陆
// 如果手机号不存在则注册
url: /user/login
method: POST
data:
{
    userName | phone : str
    password | code : str  //加密后的密码串
}
return:
{
    code: int
    msg: str
    data: {
        token: str
    }
}
```

#### 用户信息更新

```js
// 用户注册并返回登陆后的token
url: /user/update_profile
method: POST
header:{
    token:str
}
data:
{
    //用户基本信息和详细信息里的任何字段
}
return:
{
    code: int
    msg: str
    data: {}
}
```

#### 关注和取消关注用户

```js
url: /user/follow  | /user/unfollow
method: POST
header:{
    token:str
}
data:
{
    userId: str
}
return:
{
    code: int
    msg: str
    data: {}
}
```

#### 添加和移除黑名单

```js
url: /user/add_denylist  | /user/remove_denylist
method: POST
header:{
    token:str
}
data:
{
    userId: str
}
return:
{
    code: int
    msg: str
    data: {}
}
```

#### 查看黑名单列表

```js
url: /user/get_denylist
method: GET
header:{
    token:str
}
queryparam:{}
return:
{
    code: int
    msg: str
    data: [
        {
            userId: str,
            avatar: str,
            nickName: str,
            introduction: str
        },
        {
            //...
        }
    ]
}
```

#### 查看用户的关注和粉丝列表 【不需要登陆】

```js
url: /user/get_fans  |  /user/get_follow
method: GET
queryparam:{
    userId: str
}
return:
{
    code: int
    msg: str
    data: [
        {
            userId: str,
            avatar: str,
            nickName: str,
            introduction: str
        },
        {
            //...
        }
    ]
}
```

---

### 2. Admin 相关接口

###### 接口功能

> - 后台系统管理员的注册，登陆
> - 个人信息更新

#### 管理员登陆和注册

```js
// 使用用户名+密码或者手机号+验证码等方式登陆
// 如果手机号不存在则注册
url: /admin/login
method: POST
data:
{
    userName | phone : str
    password | code : str  //加密后的密码串
}
return:
{
    code: int
    msg: str
    data: {
        token: str
    }
}
```

#### 管理员信息更新

```js
url: /admin/update_profile
method: POST
header:{
    token:str
}
data:
{
    //管理员基本信息信息里的任何字段
}
return:
{
    code: int
    msg: str
    data: {}
}
```
