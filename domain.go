package main

type Campaign struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Brand   string `json:"brand"`
	Country string `json:"country"`
	GEID    string `json:"geid"`
	//Conditions     []Condition `json:"conditions,omitempty"`
	CreatedBy     string `json:"created_by"`
	CreatedAt     string `json:"created_at,omitempty"`
	LastUpdatedAt string `json:"last_updated_at,omitempty"`
	LastUpdatedBy string `json:"last_updated_by,omitempty"`
	//Metadata       interface{} `json:"metadata,omitempty"`
	//WorkflowIds    []string    `json:"workflow_ids,omitempty"`
	IsActive       bool   `json:"is_active,omitempty"`
	ReasonInactive string `json:"reason_inactive,omitempty"`
}

type Condition struct {
	Type     string      `json:"type"`
	Operator string      `json:"operator"`
	Targets  []string    `json:"targets"`
	Value    interface{} `json:"value"`
}

func NewCampaign(id, status, name, code, brand, country, geid, createdBy string) Campaign {
	return Campaign{
		ID:      id,
		Status:  status,
		Name:    name,
		Code:    code,
		Brand:   brand,
		Country: country,
		GEID:    geid,
		//Conditions:     []Condition{},
		CreatedBy:     createdBy,
		CreatedAt:     "",
		LastUpdatedAt: "",
		LastUpdatedBy: "",
		//Metadata:       map[string]interface{}{},
		//WorkflowIds:    []string{},
		IsActive:       true,
		ReasonInactive: "",
	}
}
