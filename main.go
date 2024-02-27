package main

import (
	"fmt"
	"goprj/domain"
	"math/rand"
	"strconv"
	"time"
)

const (
	rohalyky            = 50
	rohalykyPerQuestion = 5
)

var mineRohalyky = 0

var id uint64 = 1

func main() {
	fmt.Println("Вітаю у Грі \"РОГАЛИКИ-ОБРИГАЛИКИ\"")
	time.Sleep(1 * time.Second)
	for {

		menu()
		var punct = ""
		fmt.Scan(&punct)

		switch punct {
		case "1":
			play()
		case "2":
			fmt.Println("На таблицю кончених не хватило бюджету, тому що він пішов на подарунки колишнім Віталія-ПідКаблукомСтоялія")
		case "3":
			return
		default:
			fmt.Println("♥️Гарна спроба взломати пентагон!♥️")
		}
	}

}

func play() domain.User {
	fmt.Println("Відповідай і наригай!")

	var totalrohalyky int = rohalyky

	now := time.Now()
	for rohalyky > 0 {

		x, y := rand.Intn(100), rand.Intn(100)
		res := x + y
		fmt.Printf("%v + %v =?\n", x, y)
		var ans string
		fmt.Scan(&ans)
		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Printf("error: %s\n", err)
		} else {
		}
		if ansInt == res {
			if totalrohalyky < rohalykyPerQuestion {
				mineRohalyky += totalrohalyky
				totalrohalyky -= totalrohalyky
			}

			if rohalyky >= rohalykyPerQuestion {
				mineRohalyky += rohalykyPerQuestion
				totalrohalyky -= rohalykyPerQuestion
			}

			fmt.Printf("Красавчик, ти наригав достатньо! В тебе %v обригаликів\n", mineRohalyky)
		} else {

			totalrohalyky += 2
			mineRohalyky -= 2

			if mineRohalyky < 1 {
				time.Sleep(10 * time.Second)
				totalrohalyky = 50
				mineRohalyky = 0
				fmt.Println("Ти програв потну каточку :С ")
				break
			}

			fmt.Printf("Яке чарування <3... Спробуй ще, І ЩЕ, І ЩЕЕЕ!  P.S. На кону іще доля %v-ох обригаликів!\n", mineRohalyky)

		}
	}

	then := time.Now()
	timeSpent := then.Sub(now)
	if mineRohalyky >= 50 {
		fmt.Printf("ОЙОЙОЙОЙО! ТОБІ БУЛО ПОТРІБНО %v років для того щоб зібрати всі рогалики!", timeSpent)

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
	return domain.User{}
}

func menu() {
	fmt.Println("Виберіть крінж, від якого ви хочете померти:")
	fmt.Println("1: Почати крінжувати!")
	fmt.Println("2. Хто крінжанув більше!")
	fmt.Println("3. Померти від крінжу!")
}
