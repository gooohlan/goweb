package vali

import (
	"github.com/gin-gonic/gin/binding"
	zh "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"regexp"
	"strings"
)

var (
	trans      ut.Translator
	validate *validator.Validate
)

func Init(){
	// en := en.New()
	zh := zh.New()
	uni := ut.New(zh, zh)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ = uni.GetTranslator("zh")
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("label")
		})
		// 验证器注册翻译器
		zh_translations.RegisterDefaultTranslations(v, trans)

		v.RegisterValidation("checkPhone", checkPhone)
		v.RegisterTranslation("checkPhone", trans, func(ut ut.Translator) error {
			return ut.Add("checkPhone", "{0}格式错误!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("checkPhone", fe.Field(), fe.Field())

			return t
		})
		v.RegisterValidation("duration", checkDuration)
		v.RegisterTranslation("duration", trans, func(ut ut.Translator) error {
			return ut.Add("duration", "{0}格式错误!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("duration", fe.Field(), fe.Field())

			return t
		})
	}
}


func Translate(errs validator.ValidationErrors) string {
	var errList []string
	for _, e := range errs {
		// can translate each error one at a time.
		errList = append(errList, e.Translate(trans))
	}
	return strings.Join(errList, "|")
}

// 电话号码
func checkPhone(fl validator.FieldLevel) bool {
	phone := fl.Field().Interface().(string)
	re := `^1[3456789]\d{9}$`
	r := regexp.MustCompile(re)
	return r.MatchString(phone)
}

func checkDuration(fl validator.FieldLevel) bool {
	duration := fl.Field().Interface().(int)
	if duration == 1 || duration == 3 || duration == 12 {
		return true
	}
	return false
}