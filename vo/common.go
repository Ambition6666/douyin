package vo

type User struct {
	UID    uint       `gorm:"primaryKey" json:"id"`
	Pwd    string     `gorm:"not NULL"`
	Info   Commonuser `json:"info" gorm:"foreignKey:info_id"`
	InfoID uint       `json:"-"`
}

type Commonuser struct {
	ID               uint   `gorm:"primaryKey" json:"id"`
	Name             string `json:"name" gorm:"type:varchar(255)"`
	Follow_count     int    `json:"follow_count"`
	Follower_count   int    `json:"follower_count"`
	Is_follow        bool   `json:"is_follow"`
	Avatar           string `json:"avatar"`
	Background_image string `json:"background_image"`
	Signature        string `json:"signature"`
	Total_favorited  string `json:"total_favorited"`
	Work_count       int    `json:"work_count"`
	Favorite_count   int    `json:"favorite_count"`
}

type Video struct {
	Create_time    int64      `json:"create_time"`
	ID             int64      `json:"id" gorm:"primaryKey"`
	Author         Commonuser `json:"author" gorm:"-"`
	AuthorID       uint
	Play_url       string `json:"play_url"`
	Cover_url      string `json:"cover_url"`
	Favorite_count int64  `json:"favorite_count"`
	Comment_count  int64  `json:"comment_count"`
	Is_favorite    bool   `json:"is_favorite" gorm:"-"`
	Title          string `json:"title"`
}
type Comment struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	Content     string `json:"content"`
	Create_date string `json:"create_date"`
	PublisherID uint
	VideoID     int64      `json:"video_id"`
	User        Commonuser `json:"user" gorm:"-"`
}
type Message struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	Content     string `json:"content"`
	Create_time int64  `json:"create_time"`
	FromID      uint   `json:"from_user_id"`
	ToID        uint   `json:"to_user_id"`
}
