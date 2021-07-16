package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type model struct {
	Id           string `json:"id"` //士兵id
	Rarity       string //士兵稀有度
	CombatPoints string //战力
	Cvc          string //版本
	UnlockArena  string //当前解锁阶段

}

func getMap() (mm map[string]model) {
	bytes, _ := ioutil.ReadFile("../../conf/new.json")
	m := make(map[string]model)
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		fmt.Println(err)
		return
	}
	return m
}

// FindSoldierByRaUnCv 1）输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
func FindSoldierByRaUnCv(rarity, unlockArena, cvc string) (s string) {
	m1 := getMap()
	m2 := make(map[string]string)
	for s, k := range m1 {
		if k.Rarity == rarity && k.UnlockArena <= unlockArena && k.Cvc == cvc {
			jsons1, _ := json.Marshal(k)
			m2[s] = string(jsons1)
		}
	}
	jsons2, _ := json.Marshal(m2)
	return string(jsons2)
}

// FindSoldierRaById 2）输入士兵id获取稀有度
func FindSoldierRaById(id string) (s string) {
	m1 := getMap()
	com := m1[id].Rarity
	return com
}

// FindSoldierCoById 3）输入士兵id获取战力
func FindSoldierCoById(id string) (s string) {
	m1 := getMap()
	com := m1[id].CombatPoints
	return com
}

// FindSoldierByCv 4）输入cvc获取所有合法的士兵
func FindSoldierByCv(cvc string) (s string) {
	m1 := getMap()
	m2 := make(map[string]string)
	for s, k := range m1 {
		if k.Cvc == cvc {
			jsons, _ := json.Marshal(k)
			m2[s] = string(jsons)
		}
	}
	jsons2, _ := json.Marshal(m2)
	return string(jsons2)
}

// FindSoldierByUn 5）获取每个阶段解锁相应士兵的json数据
func FindSoldierByUn() string {
	//将new.json转为map
	m1 := getMap()
	//key为解锁阶段，value为一个切片，切片里放符合解锁阶段的士兵信息
	m2 := map[string][]string{}
	s := "UnlockArena: "
	for _, k := range m1 {
		//将士兵信息转为json字符串
		jsons, _ := json.Marshal(k)
		if _, ok := m2[s+k.UnlockArena]; ok {
			//解锁阶段已经存储的话，就在切片后面添加士兵数据
			m2[s+k.UnlockArena] = append(m2[s+k.UnlockArena], string(jsons))
		} else {
			//解锁阶段未储存就新建切片储存
			sli := make([]string, 0)
			sli = append(sli, string(jsons))
			m2[s+k.UnlockArena] = sli
		}
	}
	//返回每个阶段解锁的士兵数据
	jsons2, _ := json.Marshal(m2)
	return string(jsons2)

}
