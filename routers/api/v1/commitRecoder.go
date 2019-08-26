package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go-procvisual/models"
	"go-procvisual/pkgs/app"
	"go-procvisual/pkgs/e"
	"go-procvisual/pkgs/logging"
	"go-procvisual/pkgs/util"
	"go-procvisual/routers/api"
	"go-procvisual/service"
	"net/http"
)


func GetcodeCommitInfoByRepo(c *gin.Context) {

	logging.Logs.Infof("GetcodeCommitInfoByRepo in : %v", c.Request.RequestURI)
	// 1 返回数据准备
	appG := app.Gin{c}
	var data = make(map[string]interface{})

	// 2 json 数据绑定
	var _DCommitInfo models.CommitInfo
	if err := c.ShouldBindJSON(&_DCommitInfo); err != nil {
		logging.Logs.Errorf(" Bind fail  %v ", err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, data)
		return
	}
	logging.Logs.Infof("in: %v", _DCommitInfo)

	// 3 数据校验
	valid := validation.Validation{}
	ok, _ := valid.Valid(&_DCommitInfo)

	if !ok {
		// 验证失败
		validErrorMsg := api.GetErrorInfo(valid.Errors)
		data["msg"] = validErrorMsg
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, data)
	} else {
		// 验证通过
		_DCommitInfoMap := util.Struct_to_Map(_DCommitInfo)
		jsonStr, err := service.CodeCommitInfoByRepo(_DCommitInfoMap)
		if err != nil {
			appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, data)
		} else {
			data["status"] = "00"
			data["commitList"] = jsonStr
			data["prjSn"] = _DCommitInfo.PrjSn
			data["prjName"] = _DCommitInfo.PrjName
			data["startTime"] = _DCommitInfo.StartTime
			data["endTime"] = _DCommitInfo.EndTime
			data["codeRepoName"] = _DCommitInfo.CodeRepoName
			data["branchName"] = _DCommitInfo.BranchName
			logging.Logs.Infof("out: %v", data)
			appG.Response(http.StatusOK, e.SUCCESS, data)
		}
	}

}


/*
@feature: 查看项目下代码工程代码共享率
 */
func GetcodeSharedByRepo(c *gin.Context) {

	logging.Logs.Infof("GetcodeSharedByRepo in : %v", c.Request.RequestURI)
	// 1 返回数据准备
	appG := app.Gin{c}
	var data = make(map[string]interface{})

	// 2 json 数据绑定
	var _DCommitInfo models.CommitInfo
	if err := c.ShouldBindJSON(&_DCommitInfo); err != nil {
		logging.Logs.Errorf(" Bind fail  %v ", err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, data)
		return
	}
	logging.Logs.Infof("in: %v", _DCommitInfo)

	// 3 数据校验
	valid := validation.Validation{}
	ok, _ := valid.Valid(&_DCommitInfo)

	if !ok {
		// 验证失败
		validErrorMsg := api.GetErrorInfo(valid.Errors)
		data["msg"] = validErrorMsg
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, data)
	} else {
		// 验证通过
		_DCommitInfoMap := util.Struct_to_Map(_DCommitInfo)
		jsonStr, err := service.CodeCommitInfoByRepo(_DCommitInfoMap)
		if err != nil {
			appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, data)
		} else {
			data["status"] = "00"
			data["commitList"] = jsonStr
			data["prjSn"] = _DCommitInfo.PrjSn
			data["prjName"] = _DCommitInfo.PrjName
			data["startTime"] = _DCommitInfo.StartTime
			data["endTime"] = _DCommitInfo.EndTime
			data["codeRepoName"] = _DCommitInfo.CodeRepoName
			data["branchName"] = _DCommitInfo.BranchName
			logging.Logs.Infof("out: %v", data)
			appG.Response(http.StatusOK, e.SUCCESS, data)
		}
	}

}




