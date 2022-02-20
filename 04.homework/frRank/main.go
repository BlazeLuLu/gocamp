package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type RandItem struct {
	Name    string
	FatRate float64
}

type FatRateRank struct {
	items []RandItem
}

//func (r *FatRateRank) inputRecord(name string, fatRate float64) {
//	r.items = append(r.items, RandItem{
//		Name:    name,
//		FatRate: fatRate,
//	})
//}

func (r *FatRateRank) inputRecord(i int, channel chan RandItem) {
	name := "customer" + strconv.Itoa(i+1)
	rand.Seed(time.Now().UnixNano())
	randomFatRate, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", 0+rand.Float64()*(0.4)), 64)
	randItem := RandItem{Name: name, FatRate: randomFatRate}
	channel <- randItem

	//item := <-channel
	//r.items = append(r.items, item)
	//return r.items
}

func (r *FatRateRank) getRank(name string) (rank int, fatRate float64) {
	sort.Slice(r.items, func(i, j int) bool {
		return r.items[i].FatRate < r.items[j].FatRate
	})
	frs := map[float64]struct{}{}
	for _, item := range r.items {
		frs[item.FatRate] = struct{}{}
		if item.Name == name {
			fatRate = item.FatRate
		}
	}
	rankArr := make([]float64, 0, len(frs))
	for k := range frs {
		rankArr = append(rankArr, k)
	}
	sort.Float64s(rankArr)
	for i, frItem := range rankArr {
		if frItem == fatRate {
			rank = i + 1
			break
		}
	}
	fmt.Println("姓名：", name, "排名：", rank, "体脂：", fatRate)
	return
}

func main() {
	r := &FatRateRank{}
	itemChannel := make(chan RandItem)
	for i := 0; i < 1000; i++ {
		go r.inputRecord(i, itemChannel)
		//name := "customer" + strconv.Itoa(i+1)
		//rand.Seed(time.Now().UnixNano())
		////randomFatRate := 0 + rand.Float64()*(0.4)
		//randomFatRate, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", 0+rand.Float64()*(0.4)), 64)
		//fmt.Println(name, randomFatRate)
		//go r.inputRecord(name, randomFatRate)
		r.items = append(r.items, <-itemChannel)
	}

	//fmt.Println(r.items)
	for j := 0; j < 1000; j++ {
		name := "customer" + strconv.Itoa(j+1)
		r.getRank(name)
	}
}
