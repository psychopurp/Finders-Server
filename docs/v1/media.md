# 媒体类通用接口

### API 前缀: /api/v1/

---

### 1. Media 相关接口

###### 接口功能

> - 图片上传
> - 视频上传

#### 图片上传

```js
url: /media/upload_image
method: POST
header:{
    token:str
}
data:
{
     //相关上传信息
}
return:
{
    code: int
    msg: str
    data: str  //image_url 文件地址

}
```

#### 视频上传

```js
url: /media/upload_video
method: POST
header:{
    token:str
}
data:
{
    //相关上传信息
}
return:
{
    code: int
    msg: str
    data: str  //video_url 文件地址
}
```
