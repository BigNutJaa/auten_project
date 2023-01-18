package products

type Request struct {
	Name   string
	Detail string
	Qty    int32
	Token  string
}

type FitterUpdateProducts struct {
	Name      string
	Detail    string
	QtyUpdate int32
	Id        int32
	Token     string
}

type UpdateResponseProducts struct {
	Name   string
	Detail string
	Qty    int32
	Id     int32
}
