package vo

type Douyin_user_response struct {
	Status_code int    `json:"status_code"`
	Status_msg  string `json:"status_msg"`
	Commonuser  `json:"user"`
}
type Douyin_video_response struct {
	Status_code int     `json:"status_code"`
	Status_msg  string  `json:"status_msg"`
	Video_list  []Video `json:"video_list"`
}
type Douyin_comment_response struct {
	Status_code int     `json:"status_code"`
	Status_msg  string  `json:"status_msg"`
	Comment     Comment `json:"comment"`
}
type Douyin_comment_list_response struct {
	Status_code  int       `json:"status_code"`
	Status_msg   string    `json:"status_msg"`
	Comment_list []Comment `json:"comment_list"`
}
type Douyin_follow_list_response struct {
	Status_code int          `json:"status_code"`
	Status_msg  string       `json:"status_msg"`
	User_list   []Commonuser `json:"user_list"`
}
type Douyin_chat_response struct {
	Status_code  int       `json:"status_code"`
	Status_msg   string    `json:"status_msg"`
	Message_list []Message `json:"message_list"`
}
