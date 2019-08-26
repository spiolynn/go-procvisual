package service

import (
	"go-procvisual/models"
	"go-procvisual/pkgs/util"
)


func CodeCommitInfoByRepo(commitInfoMap map[string]interface{}) (string, error) {

	// 返回
	data := ""

	// adminmap -> struct
	var commitInfo models.CommitInfo
	util.Map_to_Struct(commitInfoMap, &commitInfo)

	// call model
	_commitInfolist, err := models.GetCommitNumbyTime(commitInfo.PrjSn,
		commitInfo.CodeRepoName,commitInfo.BranchName,commitInfo.StartTime,commitInfo.EndTime)

	if err != nil {
		return data, err
	} else {
		json_str, _ := util.Struct_to_Json(_commitInfolist)
		return json_str, nil

	}

}