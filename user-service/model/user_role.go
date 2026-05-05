package model

type UserRole struct {
	ID        uint  `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint  `json:"user_id" gorm:"not null"`
	RoleID    uint  `json:"role_id" gorm:"not null"`
	User      User  `gorm:"foreignKey:UserID"`
	Role      Role  `gorm:"foreignKey:RoleID"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type tabler interface {
	TableName() string
}

func (UserRole) TableName() string {
	return "user_role"
}
