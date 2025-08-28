package branchDto

type BranchRequest struct {
	ID               uint64 `json:"id" gorm:"primaryKey"`
	BranchCode       string `json:"branch_code"`
	BranchLetterCode string `json:"branch_letter_code"`
	BranchName       string `json:"branch_name"`
	BranchLocation   string `json:"branch_location"`
	BranchRadius     int32  `json:"branch_radius"`
	IsProject        int    `json:"is_project"`
	IsSub            int    `json:"is_sub"`
	BranchParentCode string `json:"branch_parent_code"`
	Meal             int    `json:"meal"`
	IsActive         int    `json:"is_active"`
}
