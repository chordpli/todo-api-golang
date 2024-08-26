package response

import (
	"encoding/json"
	"net/http"
)

// Response 구조체는 HTTP 응답을 구조화합니다.
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ResponseJSON은 HTTP 응답을 JSON 형식으로 반환합니다.
func ResponseJSON(w http.ResponseWriter, httpCode, errCode int, msg string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
}

// BindAndValid는 요청의 바인딩과 유효성 검사를 수행합니다.
func BindAndValid(r *http.Request, form interface{}) (int, int) {
	// 요청 본문을 바인딩
	err := json.NewDecoder(r.Body).Decode(form)
	if err != nil {
		return http.StatusBadRequest, 400
	}

	// 데이터 유효성 검사
	validate := validator.New()
	err = validate.Struct(form)
	if err != nil {
		return http.StatusBadRequest, 400
	}

	return http.StatusOK, 200
}
