# hackathon
    
**简要描述：** 

- 用户验证手机号接口

**请求URL：** 
- `https:47.99.58.131:8080//api/regist `
  
**请求方式：**
- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|phone |是  |string |手机号   |

 **返回示例**

``` 
{
    "Msg": "203013",
    "status": "sms_success"
}
```

- 用户注册接口

**请求URL：** 
- ` https:47.99.58.131:8080/api/regist_confirm `
  
**请求方式：**
- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|phone |是  |string |手机号   |
|password |是  |string | 密码    |
|nickname     |是  |string | 昵称    |

 **返回示例**

``` 
 {
    "Msg": "注册成功",
    "status": 0
}
```


    
**简要描述：** 

- 用户登陆接口

**请求URL：** 
- `https:47.99.58.131:8080/api/login `
  
**请求方式：**
- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|phone |是  |string |手机号   |
|password |是  |string | 密码    |

 **返回示例**

``` 
 {
    "nickname": "djfa;f",
    "phone": "12345678901",
    "statuse": "success",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyUGFzc3dvcmQiOiIxIiwiVXNlcklkIjozLCJleHAiOjE1ODcxNTk0NTIsImlhdCI6MTU4NjU1NDY1MiwiaXNzIjoiZ3VndWd1Z3UiLCJzdWIiOiJ1c2VyIHRva2VuIn0.BbaoYN8jObdLh7PHeU5E4wHPsRunCIYfujHbeC2J0V0",
    "userid": 3
}
```

    
**简要描述：** 

- 用户修改密码接口

**请求URL：** 
- `https:47.99.58.131:8080/change_password`
  
**请求方式：**
- PUT 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|newpassword |是  |string |新密码   |

 **返回示例**
``` 
{
    "Msg": "密码修改成功",
    "NewToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyUGFzc3dvcmQiOiIyMzQ1NTU1IiwiVXNlcklkIjoyLCJleHAiOjE1ODcxNjAwNjcsImlhdCI6MTU4NjU1NTI2NywiaXNzIjoiZ3VndWd1Z3UiLCJzdWIiOiJ1c2VyIHRva2VuIn0.MqtriSaBz8dprZQZoGie_eX5NoG4wUl0DQtSa6WWsqQ"
}
```


    
**简要描述：** 

- 用户修改签名，昵称接口

**请求URL：** 
- `https:47.99.58.131:8080/change_password`
  
**请求方式：**
- PUT 

**参数：** 至少有一个参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|nick_name |否  |string |新昵称   |
|signature |否  |string |新签名   |


 **返回示例**
 ```
{
    "Msg": "修改信息成功"
}
```




- 发表秘密

**请求URL：** 
- `https:47.99.58.131:8080/api/release_secret `
  
**请求方式：**
- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|topic_of_conversation |是  |string |话题   |
|title |是  |string |标题   |
|text |是  |string |故事内容   |


 **返回示例**

``` 
{
    "Msg": "故事发表成功"
}
```





- 点赞

**请求URL：** 
- `https:47.99.58.131:8080/api/praise_points `
  
**请求方式：**
- PUT 

**参数：** 

|参数名|必选|类型|说明|UT
|:----    |:---|:----- |-----   |
|article_id |是  |string |话题   |

 **返回示例**

``` 
{
	"status":0,
    "Msg": "success"
}
```




- 点赞

**请求URL：** 
- `https:47.99.58.131:8080/api/submit_comments `
  
**请求方式：**
- POST

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|user_id |是  |string |用户ID   |
|story_id |是  |string |秘密ID   |
|father_comment_id |是  |string |父评论id
|comment_text |是  |string |评论内容


 **返回示例**

``` 
{
    "Msg": "评论成功"
}
```




- 获得文章列表

**请求URL：** 
- `https:47.99.58.131:8080/api/getPost
  
**请求方式：**
- GET



 **返回示例**

``` 
 {
            "ID": 2,
            "CreatedAt": "2020-04-11T18:05:54+08:00",
            "UpdatedAt": "2020-04-11T18:08:57+08:00",
            "DeletedAt": null,
            "topic_of_conversation": "xx",
            "images_address": "",
            "AuthorID": 3,
            "title": "xxxx",
            "text": "xxxxxx",
            "PraisePoints": 0,
            "Collection": 0
        }
```
