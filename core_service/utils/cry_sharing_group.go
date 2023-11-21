package utils

type SharingGroup struct {
	Name string
	ID   string
}

func (o *SharingGroup) GetGroupName() string {
	return o.Name
}
