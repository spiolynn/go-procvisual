package models



// get data by id
// 主键查找
func GetPipelineById(id int) (*Pipeline_info, error) {

	var pl Pipeline_info
	err := db.Where("id = ?", id).First(&pl).Error
	if err != nil {
		return nil, err
	}
	return &pl, nil
}



