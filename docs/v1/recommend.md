# 首页推荐

### API 前缀: /api/v1/

---

### 假的首页推荐

```json
url: /recommend/main
method: GET
data:
{
}
return:
{
    code: int
    msg: str
    data: {
        cnt: int,
    	cards: [
    		{
    			card_id: int,
    			item_id: str,
    			item_type: int, // 用户:1, 帖子:2，动态:3
			}
    	]
    }
}
```



#### 用户卡片

```json
url: /recommend/user_info
method: GET
queryparam:{
    userId: str,
}
data:
{
}
return:
{
    code: int
    msg: str
    data: {
        user_id: str,
    	avatar: str,
    	nick_name: str,
    	signature: str,
    	shared_communities： [
    		{
    			community_id: int,
    			community_name: str
			}
    	]
    }
}
```



#### 帖子卡片

```json
url: /recommend/activity_info
method: GET
data:
{
}
return:
{
    code: int
    msg: str
    data: {
        activity_title: str,
    	activity_info: str,
    	nick_name: str,
    	user_id: str,
    	avatar: str,
    	community_id: int,
    	community_name: str,
    	medias: [
    		{
    			media_url: str,
    			media_type: int,
			}
    	]
    }
}
```



#### 动态卡片

```json
url: /recommend/moment_info
method: GET
data:
{
}
return:
{
    code: int
    msg: str
    data: {
    	nick_name: str,
    	avatar: str,
        user_id: str,
    	moment_id: int,
    	moment_info: str,
    	medias: [
    		{
    			media_url: str,
    			media_type: int,
			}
    	],
		created_at: str,
    }
}
```

