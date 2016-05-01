package dba

type materialText struct {
	Mid int32                 `json:"mid"`
	Mcontent string           `json:"mcontent"`
}

func GetText(materialId int32) (*materialText, error){
	arr := []*materialText{}
	err := readTable("zx_m_texts", &arr)

	if err != nil {
		return nil, err
	}

	for _, item := range arr {
		if item.Mid == materialId {
			return item, nil
		}
	}

	arr = nil

	return nil, ErrNotFound
}
