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
> - 回复评论
> - 获取帖子评论
> - 获取评论的回复

#### 创建圈子

在创建圈子的时候就创建 圈子管理员，默认为token对应的用户，并将圈子的状态设置为"申请中"，

审核通过可以用其他接口完成(比如admin)，审核成功前圈子管理员权限为无权限(-1)，成功后修改圈子管理员权限为正常权限，若审核失败则删除圈子和圈子管理员

```json
url: /community/create_community
method: POST
header:{
    token:str
}
data:
{
    "community_name": str,
    "community_description": str,
    "background": url
}
return:
{
    code: int
    msg: str
    data: {
    	"community_id": int
	}
}
```

#### 圈子的信息更新

管理员可以进行圈子信息的更新

可以修改的字段是圈子名字和圈子简介

data中需要带community_id

```json
// 用户注册并返回登陆后的token
url: /community/update_profile
method: POST
header:{
    token:str
}
data:
{
    community_id: int,
    //圈子基本信息里的任何字段
}
return:
{
    code: int
    msg: str
    data: {}
}
```

#### 获取圈子内的帖子 【不需要登陆】

默认page为1，第一页

```json
url: /activity/get_activities
method: GET
queryparam:{
    community_id: str, // 就不需要带引号的那种 /?community_id=1&page=1
    page: str
}
return:
{
    code: int
    msg: str
    data: {
    	total_cnt: int,
    	total_page: int,
    	page: int, // 当前
    	cnt: int,
    	activities:{
    		{
    			activity_id: str,
    			activity_info: str,
    			collect_num: int,
    			comment_num: int,
    			read_num: int,
    			tags:{
                    {
                        tag_name,
                        tag_type
                    },
					...
				},
				media_url: url,
				media_type: str, //picture or video
				nick_name: str, // 创建人的昵称
                user_id: str, // 创建人id 
                avatar: url, // 创建人头像
				user_type: str, // manager or normal
				created_at: str
			},
			...
		}
	}
}
```

#### 收藏圈子和取消收藏

```js
url: /community/collect  | /community/uncollect
method: POST
header:{
    token:str
}
data:
{
    community_id: int
}
return:
{
    code: int
    msg: str
    data: {}
}
```

#### 查看收藏的圈子列表

默认第一页page=1

```json
url: /community/get_collect
method: GET
header:{
    token:str
}
queryparam:{
    page: str
}
data:
{
   
}
return:
{
    code: int
    msg: str
    data: {
    	total_cnt: int,
    	total_page: int,
    	page: int, // 当前
    	cnt: int,
    	communities:{
    		{
    			community_id: int,
    			community_creator: str, // user_id
    			nick_name: str,
    			avatar: url,
    			community_name: str,
    			community_description: str,
    			backgroud: url,
			},
			...
		}
	}
}
```



#### 创建帖子

```json
url: /activity/add_activity
method: POST
header:{
    token:str
}
data:
{
    community_id: int,
    activity_info: str,
    media_id: str,
    media_type: str // pircture or video
}
return:
{
    code: int
    msg: str
    data: {
    	activity_id: str
	}
}
```



#### 查看帖子详情【不需要登陆】

？ 所有字段都需要吗

```json
url: /activity/get_activity_info
method: GET
queryparam:{
    activity_id: str
}
data:
{
}
return:
{
    code: int
    msg: str
    data: {
    	activity_info: str,
    	collect_num: int,
    	comment_num: int,
    	read_num: int,
        tags:{
    		{
    			tag_name,
    			tag_type
			},
			...
		},
    	media_url: url,
    	media_type: str, // picture or video
    	nick_name: str, // 创建人的昵称
    	user_id: str, // 创建人id 
    	user_type: str, // manager or normal
    	avatar: url, // 创建人头像
    	created_at: str,
	}
}
```



#### 收藏和取消收藏 帖子

```json
url: /activity/collect | /activity/uncollect
method: POST
header:{
    token: str
}
data:
{
    activity_id: str
}
return:
{
    code: int
    msg: str
    data: {}
}
```



#### 查看收藏的帖子列表

默认第一页

```json
url: /activity/get_collect
method: GET
header:{
    token: str
}
queryparam:{
    page: str
}
data:
{
    
}
return:
{
    code: int
    msg: str
    data: {
    	total_cnt: int,
    	total_page: int,
    	page: int, // 当前
    	cnt: int,
    	activities:{
    		{
    			activity_id: str,
    			activity_info: str,
    			collect_num: int,
    			comment_num: int,
    			read_num: int,
    			tags:{
                    {
                        tag_name: str,
                        tag_type: int
                    },
					...
				},
				media_url: url,
				media_type: str, //picture or video
				nick_name: str, // 创建人的昵称
                user_id: str, // 创建人id 
				user_type: str, // manager or normal
                avatar: url, // 创建人头像
				created_at: str
			},
			...
		}
	}
}
```



#### 获取帖子点赞

? 要不要登录呢

```json
url: /activity/get_activity_like
method: GET
header:{
    token: str
}
queryparam:{
    activity_id: str
}
data:
{
    
}
return:
{
    code: int
    msg: str
    data: {
    	total_cnt: int,
    	total_page: int,
    	page: int, // 当前
    	cnt: int,
    	likes:{
    		{
    			nick_name: str, // 创建人的昵称
                user_id: str // 创建人id 
                avatar: url // 创建人头像
			},
			...
		}
	}
}
```



#### 点赞和取消点赞 帖子

```json
url: /activity/like_activity | /activity/dislike_activity
method: POST
header:{
    token: str
}
data:
{
    activity_id: str
}
return:
{
    code: int
    msg: str
    data: {}
}
```



#### 评论帖子

```json
url: /activity/comment
method: POST
header:{
    token: str
}
data:
{
    item_id: str, // activity_id
    content: str
}
return:
{
    code: int
    msg: str
    data: {
    	"comment_id": int
	}
}
```



#### 回复评论

```json
url: /activity/reply
method: POST
header:{
    token: str
}
data:
{
    item_id: str, // 放comment_id
    content: str
}
return:
{
    code: int
    msg: str
    data: {
    	"comment_id": int
	}
}
```





#### 获取帖子评论 【不需要登陆】

```json
url: /activity/get_activity_comment
method: GET
queryparam:{
    activity_id: str
    page: str
}
data:
{
    
}
return:
{
    code: int
    msg: str
    data: {
    	total_cnt: int,
    	total_page: int,
    	page: int, // 当前
    	cnt: int,
    	comments:{
    		{
    			comment_id: int,
    			content: str,
    			nick_name: str, // 创建人的昵称
                user_id: str // 创建人id 
                avatar: url // 创建人头像
    			reply_num: int,
    			created_at: str
			},
			...
		}
	}
}
```



#### 获取评论的回复 【不需要登陆】

```json
url: /activity/get_comment_reply
method: GET
queryparam:{
    comment_id: str
    page: str
}
data:
{
    
}
return:
{
    code: int
    msg: str
    data: {
    	total_cnt: int,
    	total_page: int,
    	page: int, // 当前
    	cnt: int,
    	comments:{
    		{
    			comment_id: int,
    			content: str,
    			nick_name: str, // 创建人的昵称
                user_id: str // 创建人id 
                avatar: url // 创建人头像
    			reply_num: int, //字段会返回一直是0 可以忽略
    			created_at: str
			},
			...
		}
	}
}
```





