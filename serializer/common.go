package serializer

// 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// TokenData
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

// DataList 带有总数的Data结构
type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

// 带总数的返回
func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: "success",
	}
}
