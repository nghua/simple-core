package terms

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestAdd(t *testing.T) {
	ok, err := Add(0, "hua", nil)
	if err != nil {
		t.Errorf("新增Term出错：%v", err)
	}

	if !ok {
		t.Error("新增Term出错")
	}
}

func TestGet(t *testing.T) {
	term, err := Get(1, 0, "name")
	if err != nil {
		t.Errorf("获取term出错：%v", err)
	}
	t.Logf("term信息：%#v", term)
}

func TestAlter(t *testing.T) {
	ok, err := Alter(1, "hahaha", nil)
	if err != nil {
		t.Errorf("更改Term出错：%v", err)
	}

	if !ok {
		t.Error("更改Term出错")
	}
}

func TestGetList(t *testing.T) {
	list, err := GetList(0, 0, 1, "name")
	if err != nil {
		t.Errorf("获取term出错：%v", err)
	}
	t.Logf("term信息：%#v", list[0])
}

func TestGetNonNullList(t *testing.T) {
	list, err := GetNonNullList(0, 0, 1, "name")
	if err != nil {
		t.Errorf("获取term出错：%v", err)
	}
	t.Logf("term信息：%#v", list[0])
}

func TestDelete(t *testing.T) {
	ok, err := Delete(1)
	if err != nil {
		t.Errorf("删除Term出错：%v", err)
	}

	if !ok {
		t.Error("删除Term出错")
	}
}
