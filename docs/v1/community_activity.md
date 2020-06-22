# 圈子接口

### API 前缀: /api/v1/

---

### 1. Community 和 Activity 相关接口

###### 接口功能

> - 创建圈子
> - 圈子的信息更新
> - 获取圈子内的帖子 【不需要登陆】
> - 收藏和取消收藏 圈子
> - 查看收藏的圈子列表
>
> ---
>
> - 创建帖子
> - 查看帖子详情
> - 收藏和取消收藏 帖子
> - 查看收藏的帖子列表
> - 获取帖子点赞
> - 点赞和取消点赞 帖子
> - 评论帖子
> - 获取帖子评论

#### 用户登陆和注册

```js
// 使用用户名+密码或者手机号+验证码等方式登陆
// 如果手机号不存在则注册
url: /user/login
method: POST
data:
{
    userName | phone : str
    password | 验证码 : str  //加密后的密码串
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
