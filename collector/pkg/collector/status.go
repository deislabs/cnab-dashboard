package collector

type Status struct {
	Failed  int `json:"failed"`
	Pending int `json:"pending"`
}

func GetStatus(opts Options) (Status, error) {
	return Status{
		Failed: 0, Pending: 1,
	}, nil
}
