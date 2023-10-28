package structs

type Staff struct {
	StaffID      int    `json:"staff_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	StaffEmail   string `json:"staff_email"`
	Username     string `json:"username"`
	PositionID   int    `json:"position_id"`
	DepartmentID int    `json:"department_id"`
}

type ITSMAgent struct {
	AgentID      int    `json:"agent_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	AgentEmail   string `json:"agent_email"`
	Username     string `json:"username"`
	RoleID       string `json:"role_id"`
	Unit         string `json:"unit"`
	SupervisorID int    `json:"supervisor_id"`
}

type Ticket struct {
	TicketID        int    `json:"ticket_id"`
	Subject         string `json:"subject"`
	Description     string `json:"description"`
	Category        string `json:"category"`
	SubCategory     string `json:"sub_category"`
	Priority        string `json:"priority"`
	SLA             string `json:"sla"`
	StaffID         string `json:"staff_id"`
	AgentID         string `json:"agent_id"`
	CreatedAt       string `json:"created_at"`
	DueAt           string `json:"due_at"`
	AssetID         string `json:"asset_id"`
	RelatedTicketID string `json:"related_ticket_id"`
	Tag             string `json:"tag"`
	Site            string `json:"site"`
	AttachmentID    string `json:"attachment"`
}

type Asset struct {
	AssetID       int    `json:"asset_id"`
	AssetType     string `json:"asset_type"`
	AssetName     string `json:"asset_name"`
	Description   string `json:"description"`
	Manufacturer  string `json:"manufacturer"`
	Model         string `json:"model"`
	SerialNumber  string `json:"serial_number"`
	PurchaseDate  string `json:"purchase_date"`
	PurchasePrice string `json:"purchase_price"`
	Site          string `json:"site"`
	Status        string `json:"status"`
}

type Sla struct {
	SlaID          int    `json:"sla_id"`
	SlaName        string `json:"sla_id"`
	PriorityID     int    `json:"priority_id"`
	SatisfactionID int    `json:"satisfaction_id"`
	PolicyID       int    `json:"policy_id"`
}

type Priority struct {
	PriorityID    int    `json:"priority_id"`
	Name          string `json:"priority_name"`
	FirstResponse int    `json:"first_response"`
	Colour        string `json:"red"`
}

type Satisfaction struct {
	SatisfactionID int    `json:"satisfaction_id"`
	Name           string `json:"satisfaction_name"`
	Emoji          string `json:"emoji"`
}

type Policies struct {
	PolicyID     int    `json:"policy_id"`
	PolicyName   string `json:"policy_name"`
	EmbeddedLink string `json:"policy_embed"`
	PolicyUrl    string `json:"policy_url"`
}

type Position struct {
	PositionID   int    `json:"position_id"`
	PositionName string `json:"position_name"`
	CadreName    string `json:"cadre_name"`
}

type Department struct {
	DepartmentID   int    `json:"department_id"`
	DepartmentName string `json:"department_name"`
	Emoji          string `json:"emoji"`
}

type Role struct {
	RoleID   int    `json:"role_id"`
	RoleName string `json:"role_name"`
}
