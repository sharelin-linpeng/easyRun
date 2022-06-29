package repository

import (
	"fmt"

	"github.com/sharelin-linpeng/easyRun/common/db"
	"github.com/sharelin-linpeng/easyRun/entity"
)

type codeBranchService struct {
}

// 查询列表
func (codeBranchService) QueryList() []entity.CodeBranch {
	sqlStr := "SELECT id, branch_name, git_url,branch,dir,commond,repo_local FROM code_branch"

	var codeBranch []entity.CodeBranch
	if err := db.DB.Select(&codeBranch, sqlStr); err != nil {
		fmt.Printf("codeBranchService.QueryList, err:%v\n", err)
	}
	return codeBranch
}

// 根据ID查询
func (codeBranchService) QueryById(id string) *entity.CodeBranch {
	sqlStr := "SELECT id, branch_name, git_url,branch,dir,commond,repo_local FROM code_branch where id = ?"

	var codeBranch entity.CodeBranch
	if err := db.DB.Get(&codeBranch, sqlStr, id); err != nil {
		fmt.Printf("codeBranchService.QueryById(%s), err:%v\n", id, err)
		return nil
	}
	return &codeBranch
}

// 根据ID查询
func (codeBranchService) Add(codeBranch entity.CodeBranch) {
	sqlStr := "INSERT INTO code_branch(id, branch_name, git_url,branch,dir,commond,repo_local) VALUE( ?, ?, ?, ?, ?, ?, ?)"
	if _, err := db.DB.Exec(sqlStr, codeBranch.Id, codeBranch.BranchName, codeBranch.GitUrl, codeBranch.Branch, codeBranch.Dir, codeBranch.Commond, codeBranch.RepoLocal); err != nil {
		fmt.Printf("codeBranchService.Add(), err:%v\n", err)
	}
}

// 更新
func (codeBranchService) Update(codeBranch entity.CodeBranch) {
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
