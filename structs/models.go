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

type Agent struct {
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
	ID              int    `json:"ticket_id"`
	Subject         string `json:"subject"`
	Description     string `json:"description"`
	Category        int    `json:"category"`
	SubCategory     int    `json:"sub_category"`
	Priority        string `json:"priority"`
	SLA             int    `json:"sla"`
	StaffID         int    `json:"staff_id"`
	AgentID         int    `json:"agent_id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DueAt           string `json:"due_at"`
	AssetID         string `json:"asset_id"`
	RelatedTicketID int    `json:"related_ticket_id"`
	Tag             string `json:"tag"`
	Site            string `json:"site"`
	AttachmentID    int    `json:"attachment"`
	Status          string `json:"status"`
}

type Asset struct {
	AssetID       string `json:"asset_id"`
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
	SlaName        string `json:"sla_name"`
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

type Categories struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type SubCategories struct {
	SubCategoryID   int    `json:"sub_category_id"`
	SubCategoryName string `json:"sub_category_name"`
	CategoryID      int    `json:"category_id"`
}
