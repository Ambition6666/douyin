package vo

type Douyin_feed_response struct {
	Status_code int     `json:"status_code"`
	Status_msg  string  `json:"status_msg"`
	Video_list  []Video `json:"video_list"`
	Next_time   int64   `json:"next_time"`
}
type Publish_response struct {
	Status_code int    `json:"status_code"`
	Status_msg  string `json:"status_msg"`
}
type Douyin_userlike_response struct {
	Status_code int     `json:"status_code"`
	Status_msg  string  `json:"status_msg"`
	Video_list  []Video `json:"video_list"`
}
