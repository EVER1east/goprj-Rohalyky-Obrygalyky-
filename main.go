package main

import (
	"encoding/json"
	"fmt"
	"goprj/domain"
	"math/rand"
	"os"
	"sort"
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

	fmt.Println("Вітаю у Грі \"РОГАЛИКИ-ОБРИГАЛИКИ\"")
	time.Sleep(1 * time.Second)

	leaderboard := showRate()
	for _, u := range leaderboard {
		if u.Id > id {
			id = u.Id
		}

	}
	id++

	for {

		menu()
		var punct = ""
		fmt.Scan(&punct)

		switch punct {
		case "1":
			var u = play()
			var leaderboard = showRate()
			leaderboard = append(leaderboard, u)
			sortAndSave(leaderboard)
		case "2":
			var leaderboard = showRate()

			if leaderboard == nil {

			} else {
				fmt.Println("Список невдах, піднятих Інтернетом:")
				for _, user := range leaderboard {
					fmt.Printf("IP :%v | Погоняло-Обригало: %s | Років витратив: %v\n",
						user.Id,
						user.Name,
						user.Time)
				}
			}

		case "3":
			return
		case "7":
			clearRate()
		default:
			fmt.Println("♥️Гарна спроба взломати пентагон!♥️")
		}
	}
}

func play() domain.User {
	//Тут можна добавити по приколу інші варіації мотивації
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
				time.Sleep(8 * time.Second)
				break
			}

		}
		if mineRohalykyInGame >= 30 {
			timeSpent := time.Since(now).Round(time.Second)
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
	fmt.Println("7. Позбавитися конкурентів (не витримати тякої конкуренції...)")
}

func sortAndSave(leaderboard []domain.User) {
	sort.Slice(leaderboard, func(i, j int) bool {
		return leaderboard[i].Time < leaderboard[j].Time
	})

	file, err := os.OpenFile("leaderboard.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("Сталась помилка створення файлу :С %s\n", err)
		return
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Printf("Трабли с закриттям ротику файла Е_Е: %s\n", err)
		}
	}(file)

	var encoder = json.NewEncoder(file)
	err = encoder.Encode(leaderboard)
	if err != nil {
		fmt.Printf("От чорт. ПомилОЧКА в записі результатів у таблицю лідерів: %s\n", err)
		return
	}
}

func showRate() []domain.User {
	var info, err = os.Stat("leaderboard.json")
	if err != nil {
		fmt.Printf("        Поки що немає рекордів |:3_3:|\n")
		return nil
	}

	var leaderboard []domain.User
	if info.Size() != 0 {
		var file, err = os.Open("leaderboard.json")
		if err != nil {
			fmt.Printf("Сталась помилка ойойойо (РОЗМІР ВАЖЛИВИЙ!): %s\n", err)
			return nil
		}

		defer func(file *os.File) {
			err = file.Close()
			if err != nil {
				fmt.Printf("Трабли с файлом Е_Е: %s\n", err)
			}
		}(file)

		var decoder = json.NewDecoder(file)
		err = decoder.Decode(&leaderboard)
		if err != nil {
			fmt.Printf("Помилка декодування: %s\n", err)
			return nil
		}
	}
	return leaderboard
}

func clearRate() {
	var err = os.Remove("leaderboard.json")
	if err != nil {
		fmt.Printf("Помилка при видаленні даних вашої карти, прав на квартиру і паролю до OnlyFans аккаунту")
		return
	}
	fmt.Println("Ви успішно знищили всіх конкурентів")
}
