package main

import (
	"fmt"
)

func main() {

	var personNames []string
	var personBMIs []float64
	var personFatRate []float64
	var personSuggestions []string

	for i := 0; i < 3; i++ {
		type person struct {
			name        string
			weight      float64
			tall        float64
			age         int
			sex         string
			bmi         float64
			fatRate     float64
			suggestions string
		}
		//输入顾客信息
		var p person
		fmt.Printf("输入第%d位顾客的信息：\n", i+1)
		fmt.Print("姓名：")
		fmt.Scanln(&p.name)
		personNames = append(personNames, p.name)

		fmt.Print("体重（千克）：")
		fmt.Scanln(&p.weight)

		fmt.Print("身高（米）：")
		fmt.Scanln(&p.tall)

		p.bmi = p.weight / (p.tall * p.tall)
		personBMIs = append(personBMIs, p.bmi)

		fmt.Print("年龄：")
		fmt.Scanln(&p.age)

		var sexWeight int
		fmt.Print("性别（男/女）：")
		fmt.Scanln(&p.sex)

		if p.sex == "男" {
			sexWeight = 1
		} else {
			sexWeight = 0
		}
		p.fatRate = (1.2*p.bmi + 0.23*float64(p.age) - 5.4 - 10.8*float64(sexWeight)) / 100
		personFatRate = append(personFatRate, p.fatRate)

		if p.sex == "男" {
			//编写男性的体脂率与体脂状态表
			if p.age >= 18 && p.age <= 39 {
				if p.fatRate <= 0.1 {
					p.suggestions = "目前是：偏瘦，要多吃多锻炼，增强体质。"
				} else if p.fatRate > 0.1 && p.fatRate <= 0.16 {
					p.suggestions = "目前是：标准，太棒了，要保持哦。"
				} else if p.fatRate > 0.16 && p.fatRate <= 0.21 {
					p.suggestions = "目前是：偏胖，吃完饭多散散步，消化消化。"
				} else if p.fatRate > 0.21 && p.fatRate <= 0.26 {
					p.suggestions = "目前是：肥胖，少吃点，多运动运动。"
				} else {
					p.suggestions = "目前是：非常肥胖，健身游泳跑步，即刻开始。"
				}
			} else if p.age >= 40 && p.age <= 59 {
				if p.fatRate <= 0.11 {
					p.suggestions = "目前是：偏瘦，要多吃多锻炼，增强体质。"
				} else if p.fatRate > 0.11 && p.fatRate <= 0.17 {
					p.suggestions = "目前是：标准，太棒了，要保持哦。"
				} else if p.fatRate > 0.17 && p.fatRate <= 0.22 {
					p.suggestions = "目前是：偏胖，吃完饭多散散步，消化消化。"
				} else if p.fatRate > 0.22 && p.fatRate <= 0.27 {
					p.suggestions = "目前是：肥胖，少吃点，多运动运动。"
				} else {
					p.suggestions = "目前是：非常肥胖，健身游泳跑步，即刻开始。"
				}
			} else if p.age >= 60 {
				if p.fatRate <= 0.13 {
					p.suggestions = "目前是：偏瘦，要多吃多锻炼，增强体质。"
				} else if p.fatRate > 0.13 && p.fatRate <= 0.19 {
					p.suggestions = "目前是：标准，太棒了，要保持哦。"
				} else if p.fatRate > 0.19 && p.fatRate <= 0.24 {
					p.suggestions = "目前是：偏胖，吃完饭多散散步，消化消化。"
				} else if p.fatRate > 0.24 && p.fatRate <= 0.29 {
					p.suggestions = "目前是：肥胖，少吃点，多运动运动。"
				} else {
					p.suggestions = "目前是：非常肥胖，健身游泳跑步，即刻开始。"
				}
			} else {
				p.suggestions = "我们不看未成年人体脂率，变化太大，无法判断。"
			}
		} else {
			//编写女性的体脂率与体脂状态表
			if p.age >= 18 && p.age <= 39 {
				if p.fatRate <= 0.2 {
					p.suggestions = "目前是：偏瘦，要多吃多锻炼，增强体质。"
				} else if p.fatRate > 0.2 && p.fatRate <= 0.27 {
					p.suggestions = "目前是：标准，太棒了，要保持哦。"
				} else if p.fatRate > 0.27 && p.fatRate <= 0.34 {
					p.suggestions = "目前是：偏胖，吃完饭多散散步，消化消化。"
				} else if p.fatRate > 0.34 && p.fatRate <= 0.39 {
					p.suggestions = "目前是：肥胖，少吃点，多运动运动。"
				} else {
					p.suggestions = "目前是：非常肥胖，健身游泳跑步，即刻开始。"
				}
			} else if p.age >= 40 && p.age <= 59 {
				if p.fatRate <= 0.21 {
					p.suggestions = "目前是：偏瘦，要多吃多锻炼，增强体质。"
				} else if p.fatRate > 0.21 && p.fatRate <= 0.28 {
					p.suggestions = "目前是：标准，太棒了，要保持哦。"
				} else if p.fatRate > 0.28 && p.fatRate <= 0.35 {
					p.suggestions = "目前是：偏胖，吃完饭多散散步，消化消化。"
				} else if p.fatRate > 0.35 && p.fatRate <= 0.40 {
					p.suggestions = "目前是：肥胖，少吃点，多运动运动。"
				} else {
					p.suggestions = "目前是：非常肥胖，健身游泳跑步，即刻开始。"
				}
			} else if p.age >= 60 {
				if p.fatRate <= 0.22 {
					p.suggestions = "目前是：偏瘦，要多吃多锻炼，增强体质。"
				} else if p.fatRate > 0.22 && p.fatRate <= 0.29 {
					p.suggestions = "目前是：标准，太棒了，要保持哦。"
				} else if p.fatRate > 0.29 && p.fatRate <= 0.36 {
					p.suggestions = "目前是：偏胖，吃完饭多散散步，消化消化。"
				} else if p.fatRate > 0.36 && p.fatRate <= 0.41 {
					p.suggestions = "目前是：肥胖，少吃点，多运动运动。"
				} else {
					p.suggestions = "目前是：非常肥胖，健身游泳跑步，即刻开始。"
				}
			} else {
				p.suggestions = "我们不看未成年人体脂率，变化太大，无法判断。"
			}
		}
		personSuggestions = append(personSuggestions, p.suggestions)

	}

	personNum := 0
	//输出顾客测量结果
	for j := 0; j < 3; j++ {
		fmt.Printf("顾客%s的BMI指数是%.2f，体脂率为%.2f，%s\n", personNames[j], personBMIs[j], personFatRate[j], personSuggestions[j])
		personNum += 1
	}
	//输出总人数和平均BMI信息
	averageBMI := (personBMIs[0] + personBMIs[1] + personBMIs[2]) / 3
	fmt.Printf("顾客一共%d人，平均BMI为%.2f。\n", personNum, averageBMI)
}
