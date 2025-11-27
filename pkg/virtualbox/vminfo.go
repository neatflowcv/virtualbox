package virtualbox

type VMInfo struct {
	uuid   string
	name   string
	status string
}

func NewVMInfo(uuid, name, status string) *VMInfo {
	return &VMInfo{
		uuid:   uuid,
		name:   name,
		status: status,
	}
}

func (i *VMInfo) UUID() string {
	return i.uuid
}

func (i *VMInfo) Name() string {
	return i.name
}

func (i *VMInfo) Status() string {
	return i.status
}
