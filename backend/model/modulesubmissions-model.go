package model

type GetModuleSubmissionsResponse struct {
	Id       int    `json:"id"`
	ModuleId int    `json:"module_id"`
	File     string `json:"file"`
	Type     string `json:"type"`
	MaxSize  int    `json:"max_size"`
}

type CreateModuleSubmissionsRequest struct {
	Id       int    `json:"id"`
	ModuleId int    `json:"module_id"`
	File     string `json:"file"`
	Type     string `json:"type"`
	MaxSize  int    `json:"max_size"`
}

type UpdateModuleSubmissionsRequest struct {
	Id       int    `json:"id"`
	ModuleId int    `json:"module_id"`
	File     string `json:"file"`
	Type     string `json:"type"`
	MaxSize  int    `json:"max_size"`
}
