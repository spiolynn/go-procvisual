package models

import "time"

type Commit_recorder struct {
	Change_num        int    `gorm:"column:change_num;PRIMARY_KEY"`
	Prj_sn  string `gorm:"column:prj_sn"`
	Prj_name  string `gorm:"column:prj_name"`
	Coderepo_name string    `gorm:"column:coderepo_name"`
	Branch_name string    `gorm:"column:branch_name"`
	Commit_msg string    `gorm:"column:commit_msg"`
	Update_lines int    `gorm:"column:update_lines"`
	Author string    `gorm:"column:author"`
	Commit_time *time.Time `gorm:"column:commit_time"`
	Change_id string    `gorm:"column:change_id"`
	Config_sys string    `gorm:"column:config_sys"`
}

type Result struct {
	Date  string
	Total int64
}

type Result1 struct {
	author string
	lineC  int
	commitC int
	commitT *time.Time
}

// 代码提交数 request
type CommitInfo struct {
	PrjSn     		string `json:"prjSn" valid:"Required;MaxSize(50)"`
	PrjName   		string `json:"prjName" valid:"MaxSize(50)"`
	StartTime 		string `json:"startTime" valid:"Required;MaxSize(50)"`
	EndTime        	string `json:"endTime" valid:"Required;MaxSize(50)"`
	CodeRepoName    string `json:"codeRepoName" valid:"Required;MaxSize(50)"`
	BranchName   	string `json:"branchName" valid:"Required;MaxSize(50)"`
}


// 代码贡献率 request
type SharedInfo struct {
	PrjSn     		string `json:"prjSn" valid:"Required;MaxSize(50)"`
	PrjName   		string `json:"prjName" valid:"MaxSize(50)"`
	CodeRepoName    string `json:"codeRepoName" valid:"Required;MaxSize(50)"`
	BranchName   	string `json:"branchName" valid:"Required;MaxSize(50)"`
	PageSize  		int		`json:"pageSize" valid:"Required;MaxSize(50)"`
	PageNum   		int 	`json:"pagenum" valid:"Required;MaxSize(50)"`
}
