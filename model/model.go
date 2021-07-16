package model

type model struct {
	Id           string `json:"id"` //士兵id
	Rarity       string //士兵稀有度
	CombatPoints string //战力
	Cvc          string //版本
	UnlockArena  string //当前解锁阶段

}
