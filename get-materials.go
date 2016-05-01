package dba

type material struct {
	Id int32                `json:"id"`
	Mtype string            `json:"mtype"`
	CreateDate string       `json:"create_date"`
	LastChangeDate string   `json:"last_change_date"`
	Mname string            `json:"mname"`
	Mdesc string            `json:"mdesc"`
	OwnerId int32           `json:"owner_id"`
}

func GetMaterials(userId int32, mtype string) (
	[]*material,
	error){
	arr := []*material{}
	
	err := readTable("zx_materials", &arr)

	if err != nil {
		return nil, err
	}

	result := []*material{}
	
	for _, item := range arr {
		if item.OwnerId == userId && item.Mtype == mtype {
			result = append(result, item)
		}
	}

	return result, nil
}
