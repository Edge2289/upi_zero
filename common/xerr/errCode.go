package xerr

// OK 成功返回
const OK uint32 = 200

/** (前三位代表业务，后三位代表具体功能) **/

// 全局错误码
const SERVER_COMMON_ERROR uint32 = 100001 // 服务器错误
const REUQEST_PARAM_ERROR uint32 = 100002 // 参数错误
const TOKEN_EXPIRE_ERROR uint32 = 100003 // token失效
const TOKEN_GENERATE_ERROR uint32 = 100004 // 生成token失败
const NOT_ROUTE uint32 = 100005 // 生成token失败

const DB_ERROR = 100010 // db链接异常
const DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100011 // 更新数据为0

// 其他业务方
