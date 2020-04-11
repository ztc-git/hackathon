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


