package pixelset

type ValueDomain interface {
	Contains(p Value) bool
	Enumerate() ValueList
}

type Value struct {
	Region [4]uint32
	Size   [2]uint32
}

//

type ValueList []Value

func (vl ValueList) toValueMap() valueMap {
	vm := valueMap{}

	for _, v := range vl {
		vm[v] = struct{}{}
	}

	return vm
}

//

type valueMap map[Value]struct{}

func (pm valueMap) List() ValueList {
	var vl ValueList

	for item := range pm {
		vl = append(vl, item)
	}

	return vl
}
