package utils

var(
	ERR_CODE = map[int]string{
		//通用错误
		10001: "GET_REQUEST_BODY_ERROR",	//获取request的body错误
		10002: "DATA_UNMARSHAL_ERROR",		//json反序列化错误
		10003: "SQL_ERROR",					//数据库错误
		10004: "PARAM_ERROR",				//参数错误
		10005: "REQUEST_IS_NIL",			//request body为空

		//用户错误
		20001: "USER_PASSWORD_ERROR",		//密码错误
		20002: "USER_SESSION_ERROR",		//获取用户session错误
		20003: "USER_ALREADY_LOGIN",		//用户已经登录
		20004: "USER_NOT_LOGIN",			//用户未登录

		//信用卡错误
		30001: "USER_ALREADY_HAS_CARD",		//用户已经拥有该银行信用卡
	}
)