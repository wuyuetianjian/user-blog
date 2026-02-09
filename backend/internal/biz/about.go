package biz

type AboutInfo struct {
	HeroTitle       string      `json:"heroTitle"`
	HeroSubtitle    string      `json:"heroSubtitle"`
	CompanyOverview string      `json:"companyOverview"`
	Businesses      []string    `json:"businesses"`
	Advantages      []string    `json:"advantages"`
	Milestones      []Milestone `json:"milestones"`
	Contact         Contact     `json:"contact"`
}

type Milestone struct {
	Year    string `json:"year"`
	Content string `json:"content"`
}

type Contact struct {
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type AboutRepo interface {
	GetAbout() (*AboutInfo, error)
}

type AboutUsecase struct {
	repo AboutRepo
}

func NewAboutUsecase(repo AboutRepo) *AboutUsecase {
	return &AboutUsecase{repo: repo}
}

func (u *AboutUsecase) Get() (*AboutInfo, error) {
	return u.repo.GetAbout()
}
