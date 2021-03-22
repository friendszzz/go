package main
import "fmt"

func main() {
		var money int = 100000
		i := 0
		for {
			if money > 50000 {
			money = money-(money % 5)
			i++
		} else {
				break
	} 
	fmt.Printf("%v\n",i)
	for {
		if money <= 50000 && money >0 {
				money = money - 1000
				i++
		}else {
			break
		}
		fmt.Printf("%v\n",i)	
	}

}
fmt.Printf("%v,%v",money,i)
}