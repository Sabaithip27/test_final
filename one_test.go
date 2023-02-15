package testfinal

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
	// "time"
)

type Sabaithip struct {
	gorm.Model
	ID    string `valid:"matches(^DP\\d{6}$)~Not fromat"`
	Name  string `valid:"length(0|50)~Not lengeth 50,required~ชื่อห้ามว่าง"`
	Url   string `gorm:"uniqueIndex" valid:"url"`
	Age   uint   `valid:"range(1|50)"`
	Text  string
	Email string `gorm:"uniqueIndex" valid:"email,required~Not Null"`
}

func Test_ID(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sa := Sabaithip{
		ID:    "DP1234",
		Name:  "saa",
		Url:   "http://www.youtube.com/",
		Age:   25,
		Text:  "sabaithip",
		Email: "sabaithip@gmail.com",
	}
	ok, err := govalidator.ValidateStruct(sa)

	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Not fromat"))
}

func Test_Name(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sa := Sabaithip{
		Name:  "",
		Url:   "http://www.youtube.com/",
		Age:   25,
		Text:  "sabaithip",
		Email: "sabaithip@gmail.com",
	}
	ok, err := govalidator.ValidateStruct(sa)

	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("ชื่อห้ามว่าง"))
}

func Test_Nameee(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sa := Sabaithip{
		Name:  "SaaaaaaaaaaaaaaaaaaaaaaaSaaaqqqqaaaaaaaaaaaaaaaaaaaaSaaaaaaaaaaaaaaaaaaaaaaa",
		Url:   "http://www.youtube.com/",
		Age:   20,
		Text:  "AAA",
		Email: "sabaithip@gmail.com",
	}

	ok, err := govalidator.ValidateStruct(sa)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Not lengeth 50"))

}

func Test_email(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sa := Sabaithip{
		Name:  "Sabaithip",
		Url:   "http://www.youtube.com/",
		Age:   25,
		Text:  "sabaithip",
		Email: "sabaithip",
	}

	ok, err := govalidator.ValidateStruct(sa)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Email: sabaithip does not validate as email"))

}

func Test_Url(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sa := Sabaithip{
		Name:  "Sabaithip",
		Url:   "youtube",
		Age:   25,
		Text:  "sabaithip",
		Email: "sabaithip@gmail.com",
	}

	ok, err := govalidator.ValidateStruct(sa)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Url: youtube does not validate as url"))

}
func Test_EmailNotNull(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	sa := Sabaithip{
		Name:  "Sabaithip",
		Url:   "http://www.youtube.com/",
		Age:   25,
		Text:  "sabaithip",
		Email: "",
	}
	ok, err := govalidator.ValidateStruct(sa)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Not Null"))

}

func Test_Num(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	sa := Sabaithip{
		Name:  "Sabaithip",
		Url:   "http://www.youtube.com/",
		Age:   100,
		Text:  "sabaithip",
		Email: "sabaithip@gmail.com",
	}
	ok, err := govalidator.ValidateStruct(sa)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Age: 100 does not validate as range(1|50)"))

}

func Test_All(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	sa := Sabaithip{
		Name:  "Sabaithip",
		Url:   "http://www.youtube.com/",
		Age:   25,
		Text:  "sabaithip",
		Email: "sabaithip@gmail.com",
	}
	ok, err := govalidator.ValidateStruct(sa)
	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(err).To(gomega.BeNil())
}
