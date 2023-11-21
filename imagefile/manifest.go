package imagefile

func (o *Manifest) HasSharingGroups() bool {
	return len(o.Groups.Groups) > 0
}

func (o *Manifest) GetSharingGroups() []string {
	resv := []string{}
	for _, item := range o.Groups.Groups {
		resv = append(resv, item.Certificate)
	}
	return resv
}
