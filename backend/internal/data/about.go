package data

import (
	"encoding/json"

	"qizhan/backend/internal/biz"

	"gorm.io/gorm"
)

type aboutModel struct {
	ID              uint   `gorm:"primaryKey"`
	HeroTitle       string `gorm:"column:hero_title"`
	HeroSubtitle    string `gorm:"column:hero_subtitle"`
	CompanyOverview string `gorm:"column:company_overview"`
	Businesses      string `gorm:"column:businesses_json"`
	Advantages      string `gorm:"column:advantages_json"`
	Milestones      string `gorm:"column:milestones_json"`
	Phone           string `gorm:"column:phone"`
	Email           string `gorm:"column:email"`
	Address         string `gorm:"column:address"`
}

func (aboutModel) TableName() string {
	return "about_contents"
}

type aboutRepo struct {
	db *gorm.DB
}

func NewAboutRepo(db *gorm.DB) biz.AboutRepo {
	return &aboutRepo{db: db}
}

func (r *aboutRepo) GetAbout() (*biz.AboutInfo, error) {
	var model aboutModel
	if err := r.db.First(&model).Error; err != nil {
		return defaultAbout(), nil
	}

	var businesses []string
	var advantages []string
	var milestones []biz.Milestone
	_ = json.Unmarshal([]byte(model.Businesses), &businesses)
	_ = json.Unmarshal([]byte(model.Advantages), &advantages)
	_ = json.Unmarshal([]byte(model.Milestones), &milestones)

	return &biz.AboutInfo{
		HeroTitle:       model.HeroTitle,
		HeroSubtitle:    model.HeroSubtitle,
		CompanyOverview: model.CompanyOverview,
		Businesses:      businesses,
		Advantages:      advantages,
		Milestones:      milestones,
		Contact: biz.Contact{
			Phone:   model.Phone,
			Email:   model.Email,
			Address: model.Address,
		},
	}, nil
}

func defaultAbout() *biz.AboutInfo {
	return &biz.AboutInfo{
		HeroTitle:       "启展跨境咨询服务",
		HeroSubtitle:    "为企业与家庭提供专业、合规、可落地的北美服务方案",
		CompanyOverview: "我们是一家专注于跨境咨询服务的团队，服务覆盖商业落地、移民规划、留学申请、税务协同与本地生活支持，强调长期陪跑与结果导向。",
		Businesses:      []string{"企业出海咨询", "加拿大移民服务", "留学与签证方案", "税务与财务协同", "本地安家支持"},
		Advantages:      []string{"一对一顾问机制", "透明流程与费用", "本地持牌合作网络", "多语种服务团队", "项目进度实时同步"},
		Milestones: []biz.Milestone{
			{Year: "2018", Content: "团队成立，聚焦加拿大市场。"},
			{Year: "2020", Content: "服务升级，建立商业与家庭双线服务体系。"},
			{Year: "2022", Content: "上线数字化客户管理流程。"},
			{Year: "2024", Content: "拓展跨境企业综合服务能力。"},
		},
		Contact: biz.Contact{Phone: "+1-604-123-4567", Email: "service@qizhan.ca", Address: "Vancouver, BC, Canada"},
	}
}
