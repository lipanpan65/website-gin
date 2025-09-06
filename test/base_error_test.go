package test

import (
	"testing"
	"website-gin/utils/errors"
)

func TestBaseErrorMethods(t *testing.T) {
	// 测试业务错误
	businessErr := errors.NewBusinessError("1001", "用户不存在")
	
	// 测试基本方法
	if businessErr.GetCode() != "1001" {
		t.Errorf("Expected code '1001', got '%s'", businessErr.GetCode())
	}
	
	if businessErr.Error() != "用户不存在" {
		t.Errorf("Expected message '用户不存在', got '%s'", businessErr.Error())
	}
	
	if businessErr.GetType() != errors.TypeBusiness {
		t.Errorf("Expected type 'business', got '%s'", businessErr.GetType())
	}
	
	// 测试类型判断方法
	if !businessErr.IsBusinessError() {
		t.Error("Expected IsBusinessError() to return true")
	}
	
	if businessErr.IsTechnicalError() {
		t.Error("Expected IsTechnicalError() to return false")
	}
	
	// 测试String方法
	expected := "[business] 1001: 用户不存在"
	if businessErr.String() != expected {
		t.Errorf("Expected String() to return '%s', got '%s'", expected, businessErr.String())
	}
	
	// 测试技术错误
	techErr := errors.NewTechnicalError("5000", "数据库连接失败")
	
	if techErr.GetCode() != "5000" {
		t.Errorf("Expected code '5000', got '%s'", techErr.GetCode())
	}
	
	if !techErr.IsTechnicalError() {
		t.Error("Expected IsTechnicalError() to return true")
	}
	
	if techErr.IsBusinessError() {
		t.Error("Expected IsBusinessError() to return false")
	}
	
	// 测试技术错误的String方法
	expectedTech := "[technical] 5000: 数据库连接失败"
	if techErr.String() != expectedTech {
		t.Errorf("Expected String() to return '%s', got '%s'", expectedTech, techErr.String())
	}
}