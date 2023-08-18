package vo

type Douyin_user_login_response struct {
	Status_code int    `json:"status_code"`
	Status_msg  string `json:"status_msg"`
	User_id     uint   `json:"user_id"`
	Token       string `json:"token"`
}
type Douyin_user_register_response struct {
	Status_code int    `json:"status_code"`
	Status_msg  string `json:"status_msg"`
	User_id     uint   `json:"user_id"`
	Token       string `json:"token"`
}
