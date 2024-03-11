package main

import (
	"fmt"
	"goprj/domain"
	"math/rand"
	"strconv"
	"time"
)

const (
	rohalyky            = 30
	rohalykyPerQuestion = 5
	mineRohalykyDefault = 0
)

var id uint64 = 1

func main() {
	var leaderboard []domain.User

	fmt.Println("Вітаю у Грі \"РОГАЛИКИ-ОБРИГАЛИКИ\"")
	time.Sleep(1 * time.Second)
	for {

		menu()
		var punct = ""
		fmt.Scan(&punct)

		switch punct {
		case "1":
			var u = play()
			leaderboard = append(leaderboard, u)
		case "2":
			fmt.Println("Список невдах, піднятих Інтернетом:")
			for _, user := range leaderboard {
				fmt.Printf("IP :%v \nПогоняло-Обригало: %s\nРоків витратив: %v\n",
					user.Id,
					user.Name,
					user.Time)
			}

		case "3":
			return
		default:
			fmt.Println("♥️Гарна спроба взломати пентагон!♥️")
		}
	}
}

func play() domain.User {
	fmt.Println("Відповідай і наригай!")
	now := time.Now()

	var totalrohalyky int = rohalyky
	var mineRohalykyInGame int = mineRohalykyDefault

	var res int
	var operation string = ""

	var x, y, z int
	var ans string

	for totalrohalyky > 0 {

		x, y = rand.Intn(100), rand.Intn(100)
		z = rand.Intn(3)

		switch z {
		case 0:
			res = x + y
			operation = "+"

		case 1:
			res = x - y
			operation = "-"

		case 2:
			if x < 11 && y < 11 || x < 31 && x > 0 && y < 11 && y > 0 || x < 11 && x > 0 && y < 31 && y > 0 {
				res = x * y
				operation = "*"

			} else {
				res = x + y
				operation = "+"

			}

		}
		fmt.Printf("%v %v %v =?\n", x, operation, y)

		fmt.Scan(&ans)
		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Printf("error: %s\n", err)
		} else {
		}
		if ansInt == res {
			if totalrohalyky < rohalykyPerQuestion {
				mineRohalykyInGame += totalrohalyky
				totalrohalyky -= totalrohalyky
			}

			if totalrohalyky >= rohalykyPerQuestion {
				mineRohalykyInGame += rohalykyPerQuestion
				totalrohalyky -= rohalykyPerQuestion
			}

			fmt.Printf("Красавчик, ти наригав достатньо! В тебе %v обригаликів\n", mineRohalykyInGame)
		} else {

			totalrohalyky += 2
			mineRohalykyInGame -= 2

			if mineRohalykyInGame > 1 {
				fmt.Printf("Яке чарування <3... Спробуй ще, І ЩЕ, І ЩЕЕЕ!  P.S. На кону іще доля %v-ох обригаликів!\n", mineRohalykyInGame)
			}
			if mineRohalykyInGame < 1 {
				fmt.Println("Ти програв потну каточку :С ")
				time.Sleep(9 * time.Second)
				break
			}

		}
		if mineRohalykyInGame >= 30 {
			timeSpent := time.Since(now)
			fmt.Printf("ОЙОЙОЙОЙО! ТОБІ БУЛО ПОТРІБНО %v років для того щоб зібрати всі рогалики!\n", timeSpent)

			fmt.Println("Введіть ваше погоняло: ")
			var name = ""
			fmt.Scan(&name)

			var user = domain.User{
				Id:   id,
				Name: name,
				Time: timeSpent,
			}
			id++

			return user
		}
	}
	return domain.User{}
}

func menu() {
	fmt.Println("Виберіть крінж, від якого ви хочете померти:")
	fmt.Println("1: Почати крінжувати, забираючи рогалики з під під'їзду!")
	fmt.Println("2. Хто крінжанув більше(таблиця кончених)!")
	fmt.Println("3. Померти від крінжу (не видержати тякої тяжкої долі і лівнути)!")
}
