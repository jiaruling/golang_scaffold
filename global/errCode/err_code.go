package errCode

type ErrInfo = string

const (
	SuccessEn        ErrInfo = "success"
	SuccessZh        ErrInfo = "成功"
	CreatedSuccessEn ErrInfo = "created success"
	CreatedSuccessZh ErrInfo = "创建成功"
	NotFoundEn       ErrInfo = "not found"
	NotFoundZh       ErrInfo = "访问资源丢失"
	NoForbiddenEn    ErrInfo = "no forbidden"
	NoForbiddenZh      ErrInfo = "没有访问权限"
)
