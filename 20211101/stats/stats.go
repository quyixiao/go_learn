package main

import (
	"fmt"
)

func main() {
	users := []string{"烟灰", "祥哥", "zhangsan", "lisi", "zhangsan", "wangwu", "zhaoliu", "zhangsan"}
	counts := map[string]int{}
	for _, user := range users {
		//fmt.Println(user)
		_, ok := counts[user]
		if !ok {
			counts[user] = 1
		} else {
			counts[user] = counts[user] + 1
		}
	}
	fmt.Println(counts) // map[lisi:1 wangwu:1 zhangsan:3 zhaoliu:1 烟灰:1 祥哥:1]

	counts1 := map[string]int{}
	for _, user := range users {
		counts1[user]++
	}
	fmt.Println(counts1) // map[lisi:1 wangwu:1 zhangsan:3 zhaoliu:1 烟灰:1 祥哥:1]


	dream := "I have a dream\n \nFive score years ago, a great American, in whose symbolic shadow we stand today, signed the Emancipation Proclamation. This momentous decree came as a great beacon light of hope to millions of Negro slaves who had been seared in the flames of withering injustice. It came as a joyous daybreak to end the long night of bad captivity.\n \n \nBut one hundred years later, the Negro still is not free. One hundred years later, the life of the Negro is still sadly crippled by the manacles of segregation and the chains of discrimination. One hundred years later, the Negro lives on a lonely island of poverty in the midst of a vast ocean of material prosperity. One hundred years later, the Negro is still languished in the corners of American society and finds himself an exile in his own land. So we’ve come here today to dramatize a shameful condition.\n \n \nI am not unmindful that some of you have come here out of great trials and tribulations. Some of you have come fresh from narrow jail cells. Some of you have come from areas where your quest for freedom left you battered by the storms of persecution and staggered by the winds of police brutality. You have been the veterans of creative suffering. Continue to work with the faith that unearned suffering is redemptive.\n \nGo back to Mississippi, go back to Alabama, go back to South Carolina, go back to Georgia, go back to Louisiana, go back to the slums and ghettos of our northern cities, knowing that somehow this situation can and will be changed. Let us not wallow in the valley of despair.\nI say to you today, my friends, so even though we face the difficulties of today and tomorrow, I still have a dream. It is a dream deeply rooted in the American dream.\n \nI have a dream that one day this nation will rise up, live up to the true meaning of its creed: “We hold these truths to be self-evident; that all men are created equal.”\n \nI have a dream that one day on the red hills of Georgia the sons of former slaves and the sons of former slave-owners will be able to sit down together at the table of brotherhood.\n \nI have a dream that one day even the state of Mississippi, a state sweltering with the heat of injustice, sweltering with the heat of oppression, will be transformed into an oasis of freedom and justice.\n \nI have a dream that my four children will one day live in a nation where they will not be judged by the color if their skin but by the content of their character.\n \nI have a dream today.\n \nI have a dream that one day down in Alabama with its governor having his lips dripping with the words of interposition and nullification, one day right down in Alabama little black boys and black girls will be able to join hands with little white boys and white girls as sisters and brothers.\n \nI have a dream today.";

	dreamMap := map[rune]int {}
	for _ ,v :=range  dream{
		dreamMap[v]++
	}

	for ch ,cnt := range dreamMap{
		fmt.Printf("%c %d \n",ch,cnt)
	}


}
