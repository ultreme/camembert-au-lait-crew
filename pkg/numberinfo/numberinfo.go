package calcnumberinfo

type CALCNumberInfo struct {
	number float64
}

func New(number float64) *CALCNumberInfo {
	return &CALCNumberInfo{number: number}
}

func (n *CALCNumberInfo) All() map[string]interface{} {
	ret := map[string]interface{}{}
	ret["number"] = n.number
	return ret
}
