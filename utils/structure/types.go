package structure

type CreateJobs struct {
	Type    string `json:"image_type"`
	Payload ImageRequest
}
type ImageRequest struct {
	S3URI         string `json:"uri"`
	TargetType    string `json:"target_type"`
	PriorityIndex string `json:"prioriy_index"`
}
type UploadFile struct {
	FileName string `json:"file_name"`
}
type UploadFileResponse struct {
	FileName   string `json:"file_name"`
	PresignUrl string `json:"presign_url"`
}

type CreateUser struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
	Password string `json:"password"`
}
type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserAccessToken struct {
	TokenId string `json:"token_id"`
	Token   string `json:"token"`
}
