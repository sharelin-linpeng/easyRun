package entity

type Application struct {
	Id          string `db:"id" json:"id"`
	AppName     string `db:"app_name" json:"appName"`
	AppWorkPath string `db:"app_workpath" json:"appWorkPath"`
	AppFile     string `db:"app_file" json:"appFile"`
}

type CodeBranch struct {
	// ID
	Id string `db:"id" json:"id"`
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
}

type Machine struct {
	// ID
	Id string `db:"id" json:"id"`
	// 机器ID
	Ip string `db:"ip" json:"ip"`
	// 登录名称
	LoginName string `db:"login_name" json:"loginName"`
	// 登录密码
	Password string `db:"password" json:"password"`
	// 机器环境
	Env string `db:"env" json:"env"`
}

type PublishInfo struct {
	// ID
	Id string `db:"id" json:"id"`
	// 应用ID
	ApplicationId string `db:"application_id" json:"applicationId"`
	// 机器ID
	MachineId string `db:"machine_id" json:"machineId"`
	// 关联分支
	BranchId string `db:"branch_id" json:"branchId"`
	// 状态
	Status string `db:"status" json:"status"`
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
}
