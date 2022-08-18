package common

type Response struct {
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
	Status     int         `json:"status"`
	TotalCount int         `json:"total_count"`
	Errorcode  string      `json:"error_code"`
}

func ResponseHandler(errcode, language string, totalcount int, data interface{}) Response {
	ed := ErrorMessage[errcode]
	msg := ed.(map[string]interface{})[language].(string)
	status := ed.(map[string]interface{})["status"].(float64)
	return Response{Msg: msg, Data: data, Status: int(status), TotalCount: totalcount, Errorcode: errcode}
}

// func init() { log.SetFlags(log.Lshortfile | log.LstdFlags) }
