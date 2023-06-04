package lochal

type Lochal struct {
	data []interface{}
}

func Create() *Lochal {
	return &Lochal{data: make([]interface{}, 0, 5)}
}

func (ls *Lochal) Set(any interface{}) error {
	return nil
}
