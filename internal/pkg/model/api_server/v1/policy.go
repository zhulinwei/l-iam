package v1

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/ory/ladon"
	"gorm.io/gorm"
)

// Policy 后续完善的过程中需要考虑类型（如系统级别或者自定义级别）和作用域（全局或者项目级别）
type Policy struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Name      string         `json:"name" gorm:"name"`
	Policy    *AuthzPolicy   `json:"policy" gorm:"policy"`
	CreatedAt time.Time      `json:"created_at,omitempty" gorm:"created_at"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"deleted_at"`
}

type AuthzPolicy struct {
	ladon.DefaultPolicy
}

func (a *AuthzPolicy) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, &a)
}

func (a AuthzPolicy) Value() (value driver.Value, err error) {
	bytes, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return string(bytes), nil
}

func (Policy) TableName() string {
	return "policy"
}

type PolicyQueryUri struct {
	PolicyId uint `uri:"policy_id"`
}

type PolicyQueryOptions struct {
	Name string `form:"name"`
	Page int    `form:"page" binding:"required,number,gt=0"`
	Size int    `form:"size" binding:"required,number,gt=0"`
}

type PolicyList struct {
	Total int64     `json:"total"`
	List  []*Policy `json:"list"`
}

func (p *Policy) Validate() error {
	// todo 实现自定义校验
	return nil
}
