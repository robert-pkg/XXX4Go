import time
import json
import requests
import hashlib

APP_Name = "XXXApp"
APP_Secret = "62F8EA91-5180-7C55-2713-B970943696D6"

# 计算md5码
def GetMD5(src):
    m2 = hashlib.md5()
    m2.update(src.encode('utf-8'))
    return m2.hexdigest()

def callInterface(device_type, data, userID, token, serviceName, method):
    ts = int(time.time())
    strData = json.dumps(data)
    strSign = APP_Name + str(ts) + strData + APP_Secret
    strSign = GetMD5(strSign)

    headers = {
        'app_name': APP_Name,
        'app_ver':"1.0.0",
        'device_type': device_type,
        'sign': strSign,
        'ts': str(ts),
        'token': token,
        'uid': str(userID)
        }

    url = "http://127.0.0.1:9999/api/%s/%s" % (serviceName, method)
    try:
        r = requests.post(url=url, headers=headers, data=strData)
        if r.status_code != 200:
            print("http状态码错误")
            return False, ""

        print(r.text)
        return True, r.text
    except Exception as e:
        print("shttp调用失败. \n异常信息：【%s】" % (repr(e)))
        return False, ""     


def sendVcode(mobile):    
    data = {
        "mobile": mobile,
    } 

    isSuccess,respText = callInterface("web", data, 0, "", "XXXLoginServer", "SendVerifyCode")
    if not isSuccess:
        return False, ""



    print("SendVerifyCode resp:%s" % respText)
    respObj = json.loads(respText)    

    code = 0
    if 'code' in respObj.keys():
        code = respObj["code"]

    if code != 0:
        print("发送验证码失败")
        return False, ""  

    vcode = ""
    if 'vcode' in respObj.keys():
        vcode = respObj["vcode"]
    return True, vcode


# 根据手机号、验证码，获取 userID、token
def login(mobile, vCode):
    data = {
        "mobile": mobile,
        "device_type": "web",   # 设备类型. web, android, ios
        "verify_code": vCode,
    }

    isSuccess,respText = callInterface("web", data, 0, "", "XXXLoginServer", "Login")
    if not isSuccess:
        return False, 0, "", 0  

    #print("login resp:%s" % respText)
    respObj = json.loads(respText)    

    code = 0
    if 'code' in respObj.keys():
        code = respObj["code"]

    if code != 0:
        print("验证不通过")
        return False, 0, "", 0    

    respData = None
    if 'data' in respObj.keys():
        respData = respObj["data"]
    if respData is None: 
        return False, 0, "", 0    

    
    userID = -1
    if 'user_id' in respData.keys():
        userID = respData["user_id"]

    token = ""
    if 'token' in respData.keys():
        token = respData["token"]

    expireTS = 0
    if 'expire_ts' in respData.keys():   
        expireTS = respData["expire_ts"]

    if userID <= 0:
        return False, 0, "", 0 

    return True, userID, token, expireTS

def sayHello(userID, token, msg):
    data = {
        "message": msg,
    }

    isSuccess,respText = callInterface("web", data, userID, token, "XXXUserServer", "SayHello")
     # respObj = json.loads(respText)    
    return isSuccess,respText
    


def main():
    mobile = "10000000000"
    success,vcode = sendVcode(mobile) 
    if not success:
        return

    success,userID,token,expireTS = login(mobile, vcode)
    if not success:
        return
    else:
        print("登录成功. 返回: userID:[%d], token:[%s], expireTS:[%d]" % (userID, token, expireTS))

    for i in range(1,5):
        msg = " world %d" % i
        isSuccess, respText = sayHello(userID, token, msg)    
        if not isSuccess:
            print("fail")
            break

        print("sayHello resp:%s" % respText)

if __name__ == '__main__':
    main()