package virtualbox

type VM struct {
	uuid string
	name string
}

func NewVM(uuid, name string) *VM {
	return &VM{
		uuid: uuid,
		name: name,
	}
}

func (v *VM) UUID() string {
	return v.uuid
}

func (v *VM) Name() string {
	return v.name
}
