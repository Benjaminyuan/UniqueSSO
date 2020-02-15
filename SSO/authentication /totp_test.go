package authentication
import(
	"fmt"
	"testing"
)

func TestGenerateTOTP(t *testing.T){
	for i := 0;i < 20;i++ {
		fmt.Println(GenerateTOTP(i))
	}
	
}