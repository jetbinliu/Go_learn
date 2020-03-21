package main

import "fmt"

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (left int) {
	//1.依次拿到每个人的名字
	//2.拿到一个人名，根据金币的规则分金币
	//2.1 每个人的金币数应该保存到distribution中
	//2.2 还要记录剩余的金币数
	fmt.Println("coins int:", coins)
	usersMap := make(map[string]int, 20)
	for _, user := range users {
		//fmt.Println(user)
		usersMap[user] = 0
		for _, str := range user {
			switch str {
			case 'e', 'E':
				usersMap[user] = usersMap[user] + 1
				coins = coins - 1
			case 'i', 'I':
				usersMap[user] = usersMap[user] + 2
				coins = coins - 2
			case 'o', 'O':
				usersMap[user] = usersMap[user] + 3
				coins = coins - 3
			case 'u', 'U':
				usersMap[user] = usersMap[user] + 4
				coins = coins - 4
			}

		}
	}

	fmt.Println(usersMap)
	//fmt.Println("coins end:", coins)
	return coins
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	//dispatchCoin()

}
