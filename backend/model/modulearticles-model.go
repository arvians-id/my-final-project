package model

type GetModuleArticlesResponse struct {
	Id       int    `json:"id"`
	ModuleId int    `json:"module_id"`
	Content  string `json:"content"`
}
type CreateModuleArticlesRequest struct {
	Id       int    `json:"id"`
	ModuleId int    `json:"module_id"`
	Content  string `json:"content"`
}
type UpdateModuleArticlesRequest struct {
	Id       int    `json:"id"`
	ModuleId int    `json:"module_id"`
	Content  string `json:"content"`
}
