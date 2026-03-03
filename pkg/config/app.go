package config
import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/gorilla/mux"
	// "github.com/HannaShiferaw/go-bookstore/pkg/routes"
)
var DB *gorm.DB

func Connect(){
	d, err := gorm.Open("mysql", "hanna:hannapass@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic("err")
	}	
	db=d
}
func GetDB() *gorm.DB{
	return db
}