package hashTable

import "testing"

// TestHashTable 测试哈希表功能
func TestHashTable(t *testing.T) {
	// 初始化哈希表
	hashTable := InitHashTable()

	// 有效插入测试
	t.Log("开始测试有效插入--key1,key2")
	hashTable.Insert("key1")
	hashTable.Insert("key2")
	hashTable.Print()

	// 正常查找测试
	t.Log("开始测试正常查找--key1,key2")
	if hashTable.Search("key1") == nil {
		t.Errorf("期望找到key1，但是未找到")
	}
	if hashTable.Search("key2") == nil {
		t.Errorf("期望找到key2，但是未找到")
	}

	// 查找不存在的键
	t.Log("开始测试查找不存在的键--key3")
	if hashTable.Search("key3") != nil {
		t.Errorf("不应找到不存在的键key3")
	}

	// 删除测试
	t.Log("开始测试删除--key1")
	if hashTable.Delete("key1") == nil {
		t.Errorf("期望删除key1，但删除失败")
	}
	hashTable.Print()
	if hashTable.Search("key1") != nil {
		t.Errorf("期望找不到已删除的key1，但仍然找到")
	}

	// 再次删除同一个键，应该返回nil
	t.Log("再次删除同一个键，应该返回nil--key3")
	if hashTable.Delete("key3") != nil {
		t.Errorf("不应删除已不存在的key3")
	}
	hashTable.Print()

	// 删除并查找另一个键
	t.Log("开始测试删除--key1")
	if hashTable.Delete("key1") == nil {
		t.Errorf("期望删除key1，但删除失败")
	}
	hashTable.Print()
	if hashTable.Search("key1") != nil {
		t.Errorf("期望找不到已删除的key1，但仍然找到")
	}

	// 测试边界情况：插入空键
	t.Log("测试边界情况：插入空键")
	hashTable.Insert("")
	t.Log("开始测试查找空键")
	if hashTable.Search("") == nil {
		t.Errorf("期望找到空键，但未找到")
	}

	// 检验哈希表的大小
	t.Log("开始测试哈希表大小")
	if hashTable.size != 0 {
		t.Errorf("期望哈希表大小为0，但实际为%v", hashTable.size)
	}

	// 继续插入以验证边界情况
	t.Log("继续插入以验证边界情况--key8,key9")
	hashTable.Insert("key8")
	hashTable.Insert("key9")
	if hashTable.size != 2 {
		t.Errorf("期望哈希表大小为2，但实际为%v", hashTable.size)
	}
}
