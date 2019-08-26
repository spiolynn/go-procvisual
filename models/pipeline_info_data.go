package models

import "time"

type Pipeline_info struct {
	Pipeline_id           				string `gorm:"column:pipeline_id"`
	Pipeline_name                       string `gorm:"column:pipeline_name"`
	Pipeline_full_path                  string `gorm:"column:pipeline_full_path"`
	Prj_id                              string `gorm:"column:prj_id"`
	Prj_name                            string `gorm:"column:prj_name"`
	Prj_model                           string `gorm:"column:prj_model"`
	Prj_model_id                        string `gorm:"column:prj_model_id"`
	Prj_type                            string `gorm:"column:prj_type"`
	Source_manage_way                   string `gorm:"column:source_manage_way"`
	Git_code_repo                       string `gorm:"column:git_code_repo"`
	Code_repo_ssh                       string `gorm:"column:code_repo_ssh"`
	Git_code_branch                     string `gorm:"column:git_code_branch"`
	Git_credential                      string `gorm:"column:git_credential"`
	Archive_files_path                  string `gorm:"column:archive_files_path"`
	java_version                        string `gorm:"column:java_version"`
	Trigger_way                         string `gorm:"column:trigger_way"`
	Trigger_gerrit_server               string `gorm:"column:trigger_gerrit_server"`
	Trigger_code_repo                   string `gorm:"column:trigger_code_repo"`
	Trigger_code_branch                 string `gorm:"column:trigger_code_branch"`
	Trigger_after_commit_gerrit         string `gorm:"column:trigger_after_commit_gerrit"`
	Trigger_after_commit_git            string `gorm:"column:trigger_after_commit_git"`
	Timing_calendar                     string `gorm:"column:timing_calendar"`
	Create_persion                      string `gorm:"column:create_persion"`
	Version                             string `gorm:"column:version"`
	Status                              string `gorm:"column:status"`
	Remark1                             string `gorm:"column:remark1"`
	Id 									int	   `gorm:"column:id";PRIMARY_KEY`
	C_name                              string `gorm:"column:c_name"`
	Release_flag                        string `gorm:"column:release_flag"`
	Create_time 						*time.Time `gorm:"column:create_time"`
	About_prj                           string `gorm:"column:about_prj"`
	Create_type                         string `gorm:"column:create_type"`
	File_name                           string `gorm:"column:file_name"`
	Cc_component                        string `gorm:"column:cc_component"`
	Cc_config_sys                       string `gorm:"column:cc_config_sys"`
	Cc_name                             string `gorm:"column:cc_name"`
	Cc_passwd                           string `gorm:"column:cc_passwd"`
	Cc_stream                           string `gorm:"column:cc_stream"`
	Cc_center_ip                        string `gorm:"column:cc_center_ip"`
	Project_type                        string `gorm:"column:project_type"`
}


type Piple_build_info struct{
	Pipeline_id				string		`gorm:"column:pipeline_id"`
	Pipeline_build_id       string      `gorm:"column:pipeline_build_id"`
	Jenkins_build_number	int  		`gorm:"column:jenkins_build_number"`
	Start_time				*time.Time	`gorm:"column:start_time"`
	Exe_total_time			int			`gorm:"column:exe_total_time"`
	End_time				*time.Time  `gorm:"column:end_time"`
	Exe_status				int     	`gorm:"column:exe_status"`
	Trigger_way             string      `gorm:"column:trigger_way"`
	Build_person            string      `gorm:"column:build_person"`
	Remark2                 string      `gorm:"column:remark2"`
	Remark3                 string      `gorm:"column:remark3"`
	Ed						int         `gorm:"column:id";PRIMARY_KEY`
	About_branch            string      `gorm:"column:about_branch"`
	Prj_id                  string      `gorm:"column:prj_id"`
}


type PipelineHistory struct {
	prjSn string `json:"prjSn" valid:"Required;MaxSize(50)"`
	PipelineList []string `json:"pipelineList" valid:"Required;MaxSize(50)"`
}




