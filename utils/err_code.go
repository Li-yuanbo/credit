package utils

var(
	ERR_CODE = map[int]string{
		//通用错误
		10001: "GET_REQUEST_BODY_ERROR",	//获取request的body错误
		10002: "DATA_UNMARSHAL_ERROR",		//json反序列化错误
		10003: "SQL_ERROR",					//数据库获取错误
		10004: "PARAM_ERROR",				//参数错误
		10005: "REQUEST_IS_NIL",			//request body为空

		//用户错误
		20001: "USER_PASSWORD_ERROR",		//密码错误
		20002: "USER_SESSION_ERROR",		//获取用户session错误
		20003: "USER_ALREADY_LOGIN",		//用户已经登录
		20004: "USER_NOT_LOGIN",			//用户未登录
		20005: "OTHER_USER_NOT_FOUND",		//收款人信息错误
		20006: "OPERATE_USER_NOT_FOUND",	//打款人信息错误
		20007: "USER_NAME_ERROR",			//用户名错误
		20008: "USER_POWER_ERROR",			//用户admin权限校验错误

		//信用卡错误
		30001: "USER_ALREADY_HAS_CARD",		//用户已经拥有该银行信用卡
		30002: "USER_NOT_HAVE_CARD",	//收款人信用卡号错误
		30003: "CARD_BALANCE_TOO_LESS",		//余额不足
	}
)