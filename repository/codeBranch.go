package repository

import (
	"fmt"

	"github.com/sharelin-linpeng/easyRun/db"
)

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

type codeBranchService struct {
}

// 查询列表
func (codeBranchService) QueryList() []CodeBranch {
	sqlStr := "SELECT id, branch_name, git_url,branch,dir,commond,repo_local FROM code_branch"

	var codeBranch []CodeBranch
	if err := db.DB.Select(&codeBranch, sqlStr); err != nil {
		fmt.Printf("codeBranchService.QueryList, err:%v\n", err)
	}
	return codeBranch
}

// 根据ID查询
func (codeBranchService) QueryById(id string) *CodeBranch {
	sqlStr := "SELECT id, branch_name, git_url,branch,dir,commond,repo_local FROM code_branch where id = ?"

	var codeBranch CodeBranch
	if err := db.DB.Get(&codeBranch, sqlStr, id); err != nil {
		fmt.Printf("codeBranchService.QueryById(%s), err:%v\n", id, err)
		return nil
	}
	return &codeBranch
}

// 根据ID查询
func (codeBranchService) Add(codeBranch CodeBranch) {
	sqlStr := "INSERT INTO code_branch(id, branch_name, git_url,branch,dir,commond,repo_local) VALUE( ?, ?, ?, ?, ?, ?, ?)"
	if _, err := db.DB.Exec(sqlStr, codeBranch.Id, codeBranch.BranchName, codeBranch.GitUrl, codeBranch.Branch, codeBranch.Dir, codeBranch.Commond, codeBranch.RepoLocal); err != nil {
		fmt.Printf("codeBranchService.Add(), err:%v\n", err)
	}
}

// 更新
func (codeBranchService) Update(codeBranch CodeBranch) {
	sqlStr := "UPDATE  code_branch SET branch_name = ?, git_url = ?, branch = ?, dir = ?, commond=?, repo_local=?  WHERE id = ?"
	if _, err := db.DB.Exec(sqlStr, codeBranch.BranchName, codeBranch.GitUrl, codeBranch.Branch, codeBranch.Dir, codeBranch.Commond, codeBranch.RepoLocal, codeBranch.Id); err != nil {
		fmt.Printf("codeBranchService.Update(), err:%v\n", err)
	}
}

// 删除
func (codeBranchService) Delete(id string) {
	sqlStr := "DELETE  FROM code_branch  WHERE id = ?"
	if _, err := db.DB.Exec(sqlStr, id); err != nil {
		fmt.Printf("codeBranchService.Delete(%s), err:%v\n", id, err)
	}
}

var CodeBranchService = &codeBranchService{}
