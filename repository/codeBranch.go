package repository

import (
	"fmt"

	"github.com/sharelin-linpeng/easyRun/common/aes"
	"github.com/sharelin-linpeng/easyRun/common/db"
	"github.com/sharelin-linpeng/easyRun/entity"
)

type codeBranchService struct {
}

// 查询列表
func (codeBranchService) QueryList() ([]entity.CodeBranch, error) {
	sqlStr := "SELECT id, branch_name, git_url,branch,dir,commond,repo_local,user,auth,message,status FROM codebranch"

	var codeBranch []entity.CodeBranch
	if err := db.DB.Select(&codeBranch, sqlStr); err != nil {
		fmt.Printf("codeBranchService.QueryList, err:%v\n", err)
		return nil, err
	}
	return codeBranch, nil
}

// 根据ID查询
func (codeBranchService) QueryById(id string) (*entity.CodeBranch, error) {
	sqlStr := "SELECT id, branch_name, git_url,branch,dir,commond,repo_local,user,auth,message,status FROM codebranch where id = ?"

	var codeBranch entity.CodeBranch
	if err := db.DB.Get(&codeBranch, sqlStr, id); err != nil {
		fmt.Printf("codeBranchService.QueryById(%s), err:%v\n", id, err)
		return nil, err
	}
	return &codeBranch, nil
}

// 根据ID查询
func (codeBranchService) Add(codeBranch entity.CodeBranch) error {
	codeBranch.Auth, _ = aes.EnPwdCode([]byte(codeBranch.Auth))
	codeBranch.User, _ = aes.EnPwdCode([]byte(codeBranch.User))
	codeBranch.Message = ""
	codeBranch.Status = int(entity.CODE_BRANCH_NOT_INIT)
	sqlStr := "INSERT INTO codebranch( branch_name, git_url,branch,dir,commond,repo_local,user,auth,message,status) VALUE( ?, ?, ?, ?, ?, ?, ?, ?,?,?)"
	if _, err := db.DB.Exec(sqlStr, codeBranch.BranchName, codeBranch.GitUrl, codeBranch.Branch, codeBranch.Dir, codeBranch.Commond, codeBranch.RepoLocal, codeBranch.User, codeBranch.Auth, codeBranch.Message, codeBranch.Status); err != nil {
		fmt.Printf("codeBranchService.Add(), err:%v\n", err)
		return err
	}
	return nil
}

// 更新
func (codeBranchService) Update(codeBranch entity.CodeBranch) error {
	sqlStr := "UPDATE  codebranch SET branch_name = ?, git_url = ?, branch = ?, dir = ?, commond=?, repo_local=?, user=?, auth=?  WHERE id = ?"
	codeBranch.Auth, _ = aes.EnPwdCode([]byte(codeBranch.Auth))
	codeBranch.User, _ = aes.EnPwdCode([]byte(codeBranch.User))
	if _, err := db.DB.Exec(sqlStr, codeBranch.BranchName, codeBranch.GitUrl, codeBranch.Branch, codeBranch.Dir, codeBranch.Commond, codeBranch.RepoLocal, codeBranch.User, codeBranch.Auth, codeBranch.Id); err != nil {
		fmt.Printf("codeBranchService.Update(), err:%v\n", err)
		return err
	}
	return nil
}

func (codeBranchService) UpdateStatus(id string, massage string, status int) error {
	sqlStr := "UPDATE  codebranch SET message=?, status=?  WHERE id = ?"
	if _, err := db.DB.Exec(sqlStr, massage, status, id); err != nil {
		fmt.Printf("UpdateStatus.Update(), err:%v\n", err)
		return err
	}
	return nil
}

// 删除
func (codeBranchService) Delete(id string) error {
	sqlStr := "DELETE  FROM codebranch  WHERE id = ?"
	if _, err := db.DB.Exec(sqlStr, id); err != nil {
		fmt.Printf("codeBranchService.Delete(%s), err:%v\n", id, err)
		return err
	}
	return nil
}

var CodeBranchService = &codeBranchService{}
