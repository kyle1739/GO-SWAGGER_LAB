package response

// 成功回傳格式測試1
// swagger:response Successful
 type Successful struct {
 	Body struct {
 		// 成功回傳格式測試2
 		// Required: true
 		// in: body
		ID uint `json:"id"`
		// 成功回傳格式測試3
		Title string `json:"title"`
		// 成功回傳格式測試4
		Completed bool `json:"completed"`
 	}
 }

// 失败回傳格式測試1
// swagger:response ErrorWapper
type ErrorWapper struct {
	// 失败回傳格式測試2
	//
	// in: body
    Body ErrorMessage
}

type ErrorMessage struct {
    // 失败回傳格式測試3
    //
    // Required: true
	ID uint `json:"id"`
	// 失败回傳格式測試4
	Title string `json:"title"`
	// 失败回傳格式測試5
	//
	Completed bool `json:"completed"`
}