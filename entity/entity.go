package entity

type Application struct {
	Id          string `db:"id" json:"id" form:"id" binding:"required"`
	AppName     string `db:"app_name" json:"appName" form:"appName" binding:"required"`
	AppWorkPath string `db:"app_workpath" json:"appWorkPath" form:"appWorkPath" binding:"required"`
	AppFile     string `db:"app_file" json:"appFile" form:"appFile" binding:"required"`
}

type STATUS int

const (
	CODE_BRANCH_NOT_INIT STATUS = iota
	CODE_BRANCH_OK
	CODE_BRANCH_REFRESH
	CODE_BRANCH_BUILD
)

type CodeBranch struct {

	// ID
	Id int `db:"id" json:"id" form:"id"`
	// 分支名称
	BranchName string `db:"branch_name" json:"branchName" form:"branchName" binding:"required"`
	// git地址
	GitUrl string `db:"git_url" json:"gitUrl" form:"gitUrl" binding:"required"`
	// 分支名称
	Branch string `db:"branch" json:"branch" form:"branch" binding:"required"`

	// 本地仓库路径
	User string `db:"user" json:"user" form:"user"`

	// 本地仓库路径
	Auth string `db:"auth" json:"auth" form:"auth"`
	// 代码存储路径
	Dir string `db:"dir" json:"dir" form:"dir" binding:"required"`
	// 构建命令
	Commond string `db:"commond" json:"commond" form:"commond" binding:"required"`
	// 本地仓库路径
	RepoLocal string `db:"repo_local" json:"repoLocal" form:"repoLocal" binding:"required"`

	// 代码信息
	Message string `db:"message" json:"message"`

	// 代码状态
	Status int `db:"status" json:"status"`
}

type Machine struct {
	// ID
	Id string `db:"id" json:"id" form:"id" binding:"required"`
	// 机器ID
	Ip string `db:"ip" json:"ip" form:"ip" binding:"required"`
	// 登录名称
	LoginName string `db:"login_name" json:"loginName" form:"loginName" binding:"required"`
	// 登录密码
	Password string `db:"password" json:"password" form:"password" binding:"required"`
	// 机器环境
	Env string `db:"env" json:"env" form:"env" binding:"required"`
}

type PublishInfo struct {
	// ID
	Id string `db:"id" json:"id" form:"id" binding:"required"`
	// 应用ID
	ApplicationId string `db:"application_id" json:"applicationId" form:"applicationId" binding:"required"`
	// 机器ID
	MachineId string `db:"machine_id" json:"machineId" form:"machineId" binding:"required"`
	// 关联分支
	BranchId string `db:"branch_id" json:"branchId" form:"branchId" binding:"required"`
	// 状态
	Status string `db:"status" json:"status" form:"status" binding:"required"`
}

type PublishBindingInfo struct {
	// ID
	Id string `db:"id" json:"id"`

	ApplicationId string `db:"application_id" json:"applicationId"`
	AppName       string `db:"app_name" json:"appName"`
	AppWorkPath   string `db:"app_workpath" json:"appWorkPath"`
	AppFile       string `db:"app_file" json:"appFile"`

	// ID
	BranchId string `db:"branch_id" json:"branchId"`
	// 分支名称
	BranchName string `db:"branch_name" json:"branchName"`
	// git地址
	GitUrl string `db:"git_url" json:"gitUrl"`
	// 分支名称
	Branch string `db:"branch" json:"branch"`
	// 代码存储路径
	Dir string `db:"dir" json:"dir"`
	// 构建命令
	Commond string `db:"commond" json:"commond"`
	// 本地仓库路径
	RepoLocal string `db:"repo_local" json:"repoLocal"`

	// ID
	MachineId string `db:"machine_id" json:"machineId"`
	// 机器ID
	Ip string `db:"ip" json:"ip"`
	// 登录名称
	LoginName string `db:"login_name" json:"loginName"`
	// 登录密码
	Password string `db:"password" json:"password"`
	// 机器环境
	Env string `db:"env" json:"env"`
	// 状态
	Status string `db:"status" json:"status"`
}

const (
	// 任务容器
	CONTAINER = "CONTAINER"
	// 更新代码
	UPDATE_CODE = "UPDATE_CODE"
	// 构建代码
	BUILD_CODE = "BUILD_CODE"
	// 发布应用
	PUBLISH_APP = "PUBLISH_APP"
)

type Task struct {
	// ID
	Id string `db:"id" json:"id"`
	// 任务父ID
	ParentId string `db:"parent_id" json:"parentId"`
	// 执行顺序
	Sequence int `db:"sequence" json:"sequence"`
	// 任务类型
	Type string `db:"type" json:"type"`
}

type UpdateCodeTask struct {
	// ID
	Id string `db:"id" json:"id"`
	// 代码仓库编号
	CodeBranchId string `db:"code_branch_id" json:"codeBranchId"`
}

type BuildCodeTask struct {
	// ID
	Id string `db:"id" json:"id"`
	// 代码仓库编号
	CodeBranchId string `db:"code_branch_id" json:"codeBranchId"`
	// 构建路径
	BuildPath string `db:"build_path" json:"buildPath"`
}

type PublishAppTask struct {
	// ID
	Id string `db:"id" json:"id"`
	// 代码仓库编号
	CodeBranchId string `db:"code_branch_id" json:"codeBranchId"`
	// 应用Id
	AppId string `db:"app_id" json:"appId"`
	// 机器ID
	MachineId string `db:"machine_id" json:"machineId"`
}
