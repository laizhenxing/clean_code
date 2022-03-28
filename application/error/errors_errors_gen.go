// Package error const code comment error message
// 由 git.bitkinetic.com/go/uitls/errors/gengoerror 自动生成的代码，禁止编辑
package error

import (
	"git.bitkinetic.com/go/uitls/errors"
	gl "git.bitkinetic.com/go/uitls/language"
	"golang.org/x/text/language"
)

// Err 错误处理变量
var Err *errors.Errors

func init() {
	Err = errors.NewErrors()

	Err.SetLanguagePackage(language.Chinese, errors.LanguagePackage{
		ErrAdUIDNotFound:               "广告主不存在",
		ErrAddressFormat:               "地址不符合规则",
		ErrAuthFail:                    "鉴权失败",
		ErrAvailableBalanceLimited:     "可用积分不足",
		ErrContentLengthLimited:        "内容长度超过限制",
		ErrContentNotBlank:             "内容不能为空",
		ErrDbErr:                       "数据操作失败",
		ErrFailToUpdateUserInfo:        "个人信息更新失败，请稍后再试",
		ErrFailedToWithdraw:            "积分提现失败",
		ErrFailedWithdrawPass:          "提现申请通过失败",
		ErrFailedWithdrawReject:        "提现申请拒绝失败",
		ErrFakeCntFailed:               "初始化流量不正确",
		ErrFakeCountFailed:             "初始化阅读量不正确",
		ErrForbidden:                   "没有访问权限",
		ErrFormatContact:               "联系方式不符合规则",
		ErrIncorrectOldPass:            "旧密码不正确",
		ErrInternalServerError:         "服务内部错误",
		ErrInvalidAdUID:                "无效的广告主id",
		ErrInvalidDvToken:              "无效的推送账号",
		ErrInvalidGroupUID:             "请提供有效的 uid",
		ErrInvalidNoticeID:             "无效的消息编号",
		ErrInvalidPhoneAndZoneNum:      "请填写正确格式的手机号",
		ErrInvalidRoleID:               "无效的角色ID",
		ErrInvalidTarget:               "无效的目标人群参数",
		ErrInvalidWithdrawScore:        "提现额度不符合规则",
		ErrModPassFailed:               "修改密码失败，请稍后再试",
		ErrModPhoneFailed:              "修改手机号失败",
		ErrModUserStatus:               "修改用户启用禁用状态失败",
		ErrMssageContentTooLong:        "公告内容长度超过5000个字符的限制",
		ErrMssageTieleTooLong:          "公告标题长度超过128个字符的限制",
		ErrNameFormat:                  "用户昵称不符合规则",
		ErrNameIsEmpty:                 "昵称不能为空",
		ErrNewPhoneHasRegisted:         "新输入的手机号已被注册，请更换其它手机号",
		ErrNewsNotFound:                "文章不存在",
		ErrNewsOff:                     "文章已下架",
		ErrNewsScoreWasSettled:         "不允许变更已结算广告文的状态",
		ErrNotEnoughScore:              "积分不够提现",
		ErrNotFlowUser:                 "当前用户不是流量主用户",
		ErrNotFlowUserWithdraw:         "非流量主用户不能提现",
		ErrNotSupportAreaCode:          "不支持的国家区号",
		ErrNoticeNotFund:               "不存在的公告详情",
		ErrOperateFailed:               "操作失败",
		ErrParameterValidateFailed:     "参数校验失败",
		ErrParamsUnpacking:             "参数无法解码",
		ErrPassErrLoginLimit:           "您过于频繁的尝试登录，建议10分钟后再次尝试。",
		ErrPassInvalid:                 "密码不符合格式要求",
		ErrPasswordIncorrect:           "账号密码错误",
		ErrPhoneNotChange:              "新的手机号不能和当前手机号一致，请更换其它手机号",
		ErrPhoneRegistered:             "手机号已注册",
		ErrPhoneUnRegistered:           "账号不存在，请注册账号",
		ErrPromoteEndTimeFailed:        "推广结束时间不能小于当前时间",
		ErrQueryFailed:                 "查询失败",
		ErrRegFailed:                   "注册失败",
		ErrRewardPointsNotBlank:        "悬赏积分不正确",
		ErrSQLTransactionFail:          "事务启动失败",
		ErrSendSMSFail:                 "验证码发送失败",
		ErrSendSMSLimit:                "您的账号今日发送验证码次数已达上限，请明日再尝试",
		ErrSourceLengthLimited:         "来源长度超过限制",
		ErrSourceNotBlank:              "来源不能为空",
		ErrSrcNotFound:                 "资源不存在",
		ErrTargetCntFailed:             "预期流量数不正确",
		ErrTitleLengthLimited:          "标题长度超过限制",
		ErrTitleNotBlank:               "标题不能为空",
		ErrTransationStartFail:         "事务启动失败",
		ErrUnknownActionType:           "未知的操作类型",
		ErrUnknownError:                "未知的异常，请稍后再试",
		ErrUpdateActPolicyFail:         "更新组规则失败",
		ErrUpdateGroupPolicyFail:       "更新组规则失败",
		ErrUserForbiddenLogin:          "当前账户已被封禁，请联系客服（gurux.defi@gmail.com）进行解决。",
		ErrUserNotFund:                 "用户不存在",
		ErrUserScoreChargeFailed:       "用户积分充值失败",
		ErrVerifyCodeIncorrect:         "验证码错误或已失效",
		ErrWithdrawRecordHandleDone:    "提现申请已处理不能重复处理",
		ErrWithdrawRejectReasonInvalid: "提现申请拒绝原因长度不合法",
		ErrWithdrawRejectReasonIsEmpty: "提现申请拒绝原因不能为空",
	})

	Err.SetLanguagePackage(language.English, errors.LanguagePackage{
		ErrAddressFormat:           "Invalid address",
		ErrAuthFail:                "Authentication failed",
		ErrDbErr:                   "database error",
		ErrFailToUpdateUserInfo:    "Failed to modify your info",
		ErrFailedToWithdraw:        "Failed to withdraw",
		ErrForbidden:               "Forbidden",
		ErrFormatContact:           "Invalid contact information",
		ErrIncorrectOldPass:        "ErrIncorrect Old Password",
		ErrInternalServerError:     "Internal Server Error",
		ErrInvalidDvToken:          "invalid device token",
		ErrInvalidGroupUID:         "invalid group uid",
		ErrInvalidPhoneAndZoneNum:  "Invalid format phone",
		ErrInvalidRoleID:           "invalid role id",
		ErrInvalidTarget:           "invalid target params",
		ErrInvalidWithdrawScore:    "Invalid withdraw score",
		ErrModPassFailed:           "Failed to reset password",
		ErrModPhoneFailed:          "Failed to modify your phone",
		ErrMssageContentTooLong:    "notice content too long",
		ErrMssageTieleTooLong:      "notice title too long",
		ErrNameFormat:              "Invalid nickname",
		ErrNameIsEmpty:             "Nickname should be filled",
		ErrNewPhoneHasRegisted:     "The newly entered mobile phone number has been registered",
		ErrNotEnoughScore:          "No Enough score to withdraw",
		ErrNotSupportAreaCode:      "Unsupported area code",
		ErrNoticeNotFund:           "notice not fund",
		ErrOperateFailed:           "Operation failed",
		ErrParameterValidateFailed: "Parameters failed to validate",
		ErrParamsUnpacking:         "Parameter failed to unpack",
		ErrPassErrLoginLimit:       "You have been restricted from logging in 10 minutes.",
		ErrPassInvalid:             "Invalid format password",
		ErrPasswordIncorrect:       "Incorrect username or password",
		ErrPhoneNotChange:          "The new phone number cannot be the same as the current phone number",
		ErrPhoneRegistered:         "Registered phone",
		ErrPhoneUnRegistered:       "Incorrect username",
		ErrQueryFailed:             "Failed to get data",
		ErrRegFailed:               "Registration failed",
		ErrSQLTransactionFail:      "Failed to start transaction",
		ErrSendSMSFail:             "Authentication code failed to send",
		ErrSendSMSLimit:            "Your account has reached the limit for sending authentication code today. Please try again tomorrow.",
		ErrSrcNotFound:             "source not found",
		ErrTransationStartFail:     "Transaction start failed",
		ErrUnknownActionType:       "Unknown operation type",
		ErrUnknownError:            "unknown error",
		ErrUpdateActPolicyFail:     "update power policy fail",
		ErrUpdateGroupPolicyFail:   "update group policy fail",
		ErrUserForbiddenLogin:      "Your account has been blocked, please contact customer service (gurux.defi@gmail.com) to solve it.",
		ErrUserNotFund:             "Unknown user",
		ErrVerifyCodeIncorrect:     "Authentication code incorrect or invalid",
	})

	l, err := Err.GenLanguagePackageBytes()
	if err != nil {
		panic("导出错误信息到语言包失败" + err.Error())
	}
	err = gl.InitLanguageWithBytes(l)
	if err != nil {
		panic("自动导入语言包失败" + err.Error())
	}

}
