package models

// get data by Change_num
// 主键查找
func GetCommitRecoderByChangeNum(id int) (*Commit_recorder, error) {

	var cr Commit_recorder
	err := db.Where("change_num = ?", id).First(&cr).Error
	if err != nil {
		return nil, err
	}
	return &cr, nil
}

/*
//
SELECT date(commit_time),count(*) FROM `commit_recorder` t where t.prj_sn = 'PJ1901400' and t.coderepo_name = 'BK2C/pecf/pecf-prod-app' and t.branch_name = 'dev_core' and to_days(commit_time) >= to_days('2019-08-01') and to_days(commit_time) <= to_days('2019-08-10') group by date(commit_time)
 */
func GetCommitNumbyTime(prj_sn string,coderepo_name string,branch_name string,start_time string,end_time string) ([]*Result, error) {
	var results []*Result
	err := db.Table("commit_recorder").
		Select("date(commit_time) as date,count(*) as total").
		Where("prj_sn=? " +
		"and coderepo_name=? " +
		"and branch_name=? " +
		"and to_days(commit_time) >= to_days(?) and to_days(commit_time) <= to_days(?)",
		prj_sn,coderepo_name,branch_name,start_time,end_time).
		Group("date(commit_time)").Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}


/*
SELECT author,sum(update_lines) as line ,count(*),max(commit_time) FROM `commit_recorder`  WHERE (prj_sn='PJ1901400' and coderepo_name='BK2C/pecf/pecf-prod-app' and branch_name='dev_core' ) group by author order by line desc
 */

func GetCommiShared(prj_sn string,coderepo_name string,branch_name string) ([]*Result1, error) {
	var results []*Result1
	err := db.Table("commit_recorder").
		Select("author,sum(update_lines) as lineC,count(*) as commitC,max(commit_time) as commitT").
		Where("prj_sn=? " +
			"and coderepo_name=? " +
			"and branch_name=?",
			prj_sn,coderepo_name,branch_name).
		Group("author").Order("lineC desc").Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}


