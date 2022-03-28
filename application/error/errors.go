package error

import (
	"git.bitkinetic.com/go/uitls/errors"
)

//go:generate gengoerror
const (
	// ErrForbidden 没有访问权限
	// Forbidden
	ErrForbidden errors.ErrorCode = 403
	// ErrSrcNotFound 资源不存在
	// source not found
	ErrSrcNotFound errors.ErrorCode = 404
	// ErrInternalServerError 服务内部错误
	// Internal Server Error
	ErrInternalServerError errors.ErrorCode = 500
)

// 基础错误码	1019000-1019099
const (
	// ErrParamsUnpacking 参数无法解码
	// Parameter failed to unpack
	ErrParamsUnpacking errors.ErrorCode = iota + 1019000
	// ErrAuthFail 鉴权失败
	// Authentication failed
	ErrAuthFail
	// ErrNotSupportAreaCode 不支持的国家区号
	// Unsupported area code
	ErrNotSupportAreaCode
	// ErrInvalidPhoneAndZoneNum 请填写正确格式的手机号
	// Invalid format phone
	ErrInvalidPhoneAndZoneNum
	// ErrPassInvalid 密码不符合格式要求
	// Invalid format password
	ErrPassInvalid
	// ErrUnknownActionType 未知的操作类型
	// Unknown operation type
	ErrUnknownActionType
	// ErrSendSMSLimit 您的账号今日发送验证码次数已达上限，请明日再尝试
	// Your account has reached the limit for sending authentication code today. Please try again tomorrow.
	ErrSendSMSLimit
	// ErrSendSMSFail 验证码发送失败
	// Authentication code failed to send
	ErrSendSMSFail
	// ErrVerifyCodeIncorrect 验证码错误或已失效
	// Authentication code incorrect or invalid
	ErrVerifyCodeIncorrect
	// ErrUnknownError 未知的异常，请稍后再试
	// unknown error
	ErrUnknownError
	// ErrParameterValidateFailed 参数校验失败
	// Parameters failed to validate
	ErrParameterValidateFailed
	// ErrSQLTransactionFail 事务启动失败
	// Failed to start transaction
	ErrSQLTransactionFail
	// ErrQueryFailed 查询失败
	// Failed to get data
	ErrQueryFailed
	// ErrOperateFailed 操作失败
	// Operation failed
	ErrOperateFailed
	// ErrRegFailed 注册失败
	// Registration failed
	ErrRegFailed
	// ErrPhoneUnRegistered 账号不存在，请注册账号
	// Incorrect username
	ErrPhoneUnRegistered
	// ErrPassErrLoginLimit 您过于频繁的尝试登录，建议10分钟后再次尝试。
	// You have been restricted from logging in 10 minutes.
	ErrPassErrLoginLimit
	// ErrPasswordIncorrect 账号密码错误
	// Incorrect username or password
	ErrPasswordIncorrect
	// ErrModPassFailed 修改密码失败，请稍后再试
	// Failed to reset password
	ErrModPassFailed
	// ErrUserNotFund 用户不存在
	// Unknown user
	ErrUserNotFund
	// ErrDbErr 数据操作失败
	// database error
	ErrDbErr
	// ErrInvalidDvToken 无效的推送账号
	// invalid device token
	ErrInvalidDvToken
	// ErrMssageTieleTooLong 公告标题长度超过128个字符的限制
	// notice title too long
	ErrMssageTieleTooLong
	// ErrMssageContentTooLong 公告内容长度超过5000个字符的限制
	// notice content too long
	ErrMssageContentTooLong
	// ErrNameIsEmpty 昵称不能为空
	// Nickname should be filled
	ErrNameIsEmpty
	// ErrNameFormat 用户昵称不符合规则
	// Invalid nickname
	ErrNameFormat
	// ErrFailToUpdateUserInfo 个人信息更新失败，请稍后再试
	// Failed to modify your info
	ErrFailToUpdateUserInfo
	// ErrAddressFormat 地址不符合规则
	// Invalid address
	ErrAddressFormat
	// ErrNotEnoughScore 积分不够提现
	// No Enough score to withdraw
	ErrNotEnoughScore
	// ErrFailedToWithdraw 积分提现失败
	// Failed to withdraw
	ErrFailedToWithdraw
	// ErrInvalidWithdrawScore 提现额度不符合规则
	// Invalid withdraw score
	ErrInvalidWithdrawScore
	// ErrFormatContact 联系方式不符合规则
	// Invalid contact information
	ErrFormatContact
	// ErrTransationStartFail 事务启动失败
	// Transaction start failed
	ErrTransationStartFail
	// ErrInvalidNoticeID 无效的消息编号
	ErrInvalidNoticeID
	// ErrIncorrectOldPass 旧密码不正确
	// ErrIncorrect Old Password
	ErrIncorrectOldPass
	// ErrPhoneRegistered 手机号已注册
	// Registered phone
	ErrPhoneRegistered
	// ErrModPhoneFailed 修改手机号失败
	// Failed to modify your phone
	ErrModPhoneFailed
	// ErrNoticeNotFund 不存在的公告详情
	// notice not fund
	ErrNoticeNotFund
	// ErrInvalidRoleID 无效的角色ID
	// invalid role id
	ErrInvalidRoleID
	// ErrUpdateGroupPolicyFail 更新组规则失败
	// update group policy fail
	ErrUpdateGroupPolicyFail
	// ErrUpdateActPolicyFail 更新组规则失败
	// update power policy fail
	ErrUpdateActPolicyFail
	// ErrInvalidGroupUID 请提供有效的 uid
	// invalid group uid
	ErrInvalidGroupUID
	// ErrInvalidTarget 无效的目标人群参数
	// invalid target params
	ErrInvalidTarget
	// ErrUserForbiddenLogin 当前账户已被封禁，请联系客服（gurux.defi@gmail.com）进行解决。
	// Your account has been blocked, please contact customer service (gurux.defi@gmail.com) to solve it.
	ErrUserForbiddenLogin
	// ErrPhoneNotChange 新的手机号不能和当前手机号一致，请更换其它手机号
	// The new phone number cannot be the same as the current phone number
	ErrPhoneNotChange
	// ErrNewPhoneHasRegisted 新输入的手机号已被注册，请更换其它手机号
	// The newly entered mobile phone number has been registered
	ErrNewPhoneHasRegisted
	// JWT token 相关错误
	// ErrParseClaims claims 解析失败
	// The custom Claims parse failed
	ErrParseClaims
	// ErrExpiredToken token 已过期
	// The token has expired
	ErrExpiredToken
	// ErrErrorToken 错误的 token
	// Error Token
	ErrErrorToken
	// ErrUnknownToken 未知 token
	// Unknow Token
	ErrUnknownToken

)

// app错误码	1019200-1019399
const (
	// ErrNewsOff 文章已下架
	ErrNewsOff errors.ErrorCode = iota + 1019200
)

// admin错误码	1019400-1019599
const (
	// ErrTitleNotBlank 标题不能为空
	ErrTitleNotBlank errors.ErrorCode = iota + 1019400
	// ErrContentNotBlank 内容不能为空
	ErrContentNotBlank
	// ErrSourceNotBlank 来源不能为空
	ErrSourceNotBlank
	// ErrFakeCountFailed 初始化阅读量不正确
	ErrFakeCountFailed
	// ErrTitleLengthLimited 标题长度超过限制
	ErrTitleLengthLimited
	// ErrContentLengthLimited 内容长度超过限制
	ErrContentLengthLimited
	// ErrSourceLengthLimited 来源长度超过限制
	ErrSourceLengthLimited
	// ErrRewardPointsNotBlank 悬赏积分不正确
	ErrRewardPointsNotBlank
	// ErrPromoteEndTimeFailed 推广结束时间不能小于当前时间
	ErrPromoteEndTimeFailed
	// ErrTargetCntFailed 预期流量数不正确
	ErrTargetCntFailed
	// ErrFakeCntFailed 初始化流量不正确
	ErrFakeCntFailed
	// ErrAdUIDNotFound 广告主不存在
	ErrAdUIDNotFound
	// ErrAvailableBalanceLimited 可用积分不足
	ErrAvailableBalanceLimited
	// ErrInvalidAdUID 无效的广告主id
	ErrInvalidAdUID
	// ErrFailedWithdrawReject 提现申请拒绝失败
	ErrFailedWithdrawReject
	// ErrFailedWithdrawPass 提现申请通过失败
	ErrFailedWithdrawPass
	// ErrWithdrawRejectReasonIsEmpty 提现申请拒绝原因不能为空
	ErrWithdrawRejectReasonIsEmpty
	// ErrWithdrawRejectReasonInvalid 提现申请拒绝原因长度不合法
	ErrWithdrawRejectReasonInvalid
	// ErrWithdrawRecordHandleDone 提现申请已处理不能重复处理
	ErrWithdrawRecordHandleDone
	// ErrNewsNotFound 文章不存在
	ErrNewsNotFound
	// ErrUserScoreChargeFailed 用户积分充值失败
	ErrUserScoreChargeFailed
	// ErrModUserStatus 修改用户启用禁用状态失败
	ErrModUserStatus
	// ErrNotFlowUser 当前用户不是流量主用户
	ErrNotFlowUser
	// ErrNotFlowUserWithdraw 非流量主用户不能提现
	ErrNotFlowUserWithdraw
	// ErrNewsScoreWasSettled 不允许变更已结算广告文的状态
	ErrNewsScoreWasSettled
)
