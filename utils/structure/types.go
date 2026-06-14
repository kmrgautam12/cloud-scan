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
