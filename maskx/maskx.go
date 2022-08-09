package maskx

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/panyaxbo/libs/logx"
	"github.com/spf13/viper"
)

var defaultSensitiveData = []string{
	"name",
	"surName",
	"firstName",
	"firstNameTh",
	"firstNameEn",
	"lastName",
	"lastNameTh",
	"lastNameEn",
	"citizenId",
	"passport",
	"passportNo",
	"identification",
	"national",
	"nationality",
	"card",
	"cardNo",
	"phone",
	"phoneNo",
	"number",
	"username",
	"password",
	"email",
	"address",
	"accountNo",
	"dateOfBirth",
	"dob",
	"dateOfBirthTh",
	"dateOfBirthEn",
}

// var (
// 	gcmUat = NewAES([]byte(viper.GetString("crypto.ekyc.uat.aeskey")), []byte(viper.GetString("crypto.ekyc.uat.aesnonce")))
// )

type Mask interface {
	Json(b []byte, env string) (*string, error)
}

type mask struct {
	sensitiveField []string
	gcmUat         *AES
	gcmPrd         *AES
}

func Init(fields ...[]string) Mask {
	var f = defaultSensitiveData

	initConfig()
	if len(fields) > 0 {
		f = append(f, fields[0]...)
	}
	uat := NewAES([]byte("thisisaeskey16bt"), []byte("thisisnoncee"))
	prd := NewAES([]byte("kbcptgpm5OegmJN7"), []byte("nlOqfSqvdTj8"))
	return &mask{
		sensitiveField: f,
		gcmUat:         uat,
		gcmPrd:         prd,
	}
}

func initConfig() {
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		logx.Panic(err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	runtime.GOMAXPROCS(1)

	logx.Init(viper.GetString("log.level"), viper.GetString("log.env"))
}

func (m mask) Json(b []byte, env string) (*string, error) {
	var storage []interface{}
	p := make(chan bool, 10)
	var wg sync.WaitGroup
	wg.Add(1)
	p <- true
	err := m.walkThrough(b, &storage, p, &wg)
	if err != nil {
		return nil, err
	}
	wg.Wait()
	return m.masking(b, storage, env)
}

//walkThrough will recursive until no more array or map
func (m mask) walkThrough(b []byte, storage *[]interface{}, p chan bool, wg *sync.WaitGroup) error {
	defer func() {
		<-p
		wg.Done()
	}()
	var gson interface{}
	err := json.Unmarshal(b, &gson)
	if err != nil {
		return err
	}
	switch t := gson.(type) {
	case map[string]interface{}:
		for k, v := range t {
			switch v := v.(type) {
			case string:
				m.sensitive(k, v, storage)
			case float64:
				m.sensitive(k, v, storage)
			case int32:
				m.sensitive(k, v, storage)
			case []interface{}:
				for _, val := range v {
					err = m.next(val, storage, p, wg)
					if err != nil {
						return err
					}
				}
			case map[string]interface{}:
				err = m.next(v, storage, p, wg)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (m mask) next(v interface{}, storage *[]interface{}, p chan bool, wg *sync.WaitGroup) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	wg.Add(1)
	p <- true
	go m.walkThrough(b, storage, p, wg)
	return nil
}

func (m mask) sensitive(k string, v interface{}, storage *[]interface{}) {
	ok := strings.Contains(strings.ToLower(fmt.Sprint(m.sensitiveField)), strings.ToLower(k))
	if ok {
		//append sensitive value to storage
		*storage = append(*storage, v)
	}
}

func (m mask) masking(j []byte, d []interface{}, env string) (*string, error) {
	body := string(j)
	if len(d) == 0 {
		return &body, nil
	}
	for _, val := range d {
		if strings.EqualFold(env, "prd") || strings.EqualFold(env, "prod") || strings.EqualFold(env, "production") {
			body = strings.ReplaceAll(body, typeCasting(val.(interface{})), randomMask(m.gcmPrd, typeCasting(val.(interface{}))))
		} else {
			body = strings.ReplaceAll(body, typeCasting(val.(interface{})), randomMask(m.gcmUat, typeCasting(val.(interface{}))))
		}
	}
	return &body, nil
}

func randomMask(gcm *AES, c string) string {
	if len(c) == 0 {
		return c
	}
	// var r = []rune(c)
	// var cl = len(r)
	// var size = initMaskSize(cl)
	// var count int
	// raffle := make(map[int]int, size)
	// for i := 0; i < cl; i++ {
	// 	count += 1 //avoid random forever
	// 	if len(raffle) == size || count == 10 {
	// 		//break if mask enough
	// 		break
	// 	}
	// 	v := randPos(cl)
	// 	if _, ok := raffle[v]; ok {
	// 		i -= 1
	// 		continue
	// 	}
	// 	//case not mask yet
	// 	if len(r)-1 >= v {
	// 		r[v] = '*'
	// 		raffle[v] = v
	// 	}
	// }
	// return string(r)
	return gcm.Encrypt(c)

}

func randPos(max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	ra := rand.New(source)
	return ra.Intn(max)
}

func initMaskSize(l int) int {
	if l == 1 {
		return l
	}
	return l / 2
}

func typeCasting(d interface{}) string {
	switch c := d.(type) {
	case string:
		return c
	case int64:
		return fmt.Sprint(c)
	case float64:
		return strconv.FormatFloat(c, 'f', -1, 64)
	default:
		return fmt.Sprint(d)
	}
}
