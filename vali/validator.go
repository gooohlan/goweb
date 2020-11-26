package vali

import (
	"github.com/gin-gonic/gin/binding"
	zh "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
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
		// 添加label标签,实现错误信息全汉化
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("label")
		})
		// 验证器注册翻译器
		zh_translations.RegisterDefaultTranslations(v, trans)
		v.RegisterValidation("checkMobile", checkMobile)
		v.RegisterTranslation("checkMobile", trans, func(ut ut.Translator) error {
			return ut.Add("checkMobile", "{0}长度不等于11位或{1}格式错误!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("checkMobile", fe.Field(), fe.Field())

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
	return strings.Join(errList, " | ")
}

func checkMobile(fl validator.FieldLevel) bool {
	return false
}