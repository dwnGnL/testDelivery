package models

type ProjectReq struct {
	Title           *string `json:"title"`
	Description     *string `json:"description"`
	MediaID         *int64  `json:"media_id"`
	Type            *string `json:"type"`
	BusinessOwner   *string `json:"business_owner"`
	Stage           *string `json:"stage"`
	LegacyEntity    *string `json:"legacy_entity"`
	Cluster         *string `json:"cluster"`
	OwnerID         *string `json:"owner_id"`
	OwnerPhoto      *string `json:"owner_photo"`
	Region          *string `json:"region"`
	Status          *string `json:"status"`
	Priority        *int    `json:"priority"`
	PipelineManager *string `json:"pipeline_manager"`
	ProjectManager  *string `json:"project_manager"`
}

type MilestoneFilter struct {
	MilestoneID int   `json:"milestone_id"`
	ProjectID   int64 `json:"project_id"`
}

type ProjectFilter struct {
	Cluster *string `json:"cluster"`
	Type    *string `json:"type"`
	Stage   *string `json:"stage"`
}

type UserCreateReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"` //for easy test
}

type ProcessReq struct {
	ProcessID int64  `json:"process_id"`
	Name      string `json:"name"`
}
