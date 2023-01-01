package main

import "fmt" 

func swap (x *int, y *int){
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func stringSwap(x *string, y *string){
	var temp string
	temp = *x
	*x = *y
	*y = temp
}

func sort(begin *[]int, duration *[]int, name *[]string, n int){
	for i:=0 ; i<n-1;i++{
		for j:=0 ; j< n - i - 1; j++{
			if((*duration)[j] > (*duration)[j+1]){
				swap(&(*begin)[j], &(*begin)[j + 1])
				swap(&(*duration)[j], &(*duration)[j + 1])
				stringSwap(&(*name)[j], &(*name)[j + 1])
			}
		}
	}
}

func intro(begin []int, duration []int, name []string, n int)([]int, []int, []string){
	var temp int
	var t string
	for i:=0 ; i < n ; i++ {
		fmt.Println("Process name:")
		fmt.Scanln(&t)
		name = append(name, t)
		fmt.Println("Process begin:")
		fmt.Scanf("%d", &temp)
		begin = append(begin, temp)
		fmt.Println("Process duration:")
		_, err3 := fmt.Scanf("%d", &temp)
		duration = append(duration, temp)
		_ = err3
	}
	return begin, duration, name
}

func sjf(begin []int, duration []int, name []string, n int, timer int, seconds int, wait float64)(int, float64){
	var flag bool
	for {
		flag = true
		for i:=0 ; i < n && flag ; i++ {
			flag = false
			if(duration[i] != 0){
				flag = true
				break
			}
		}
		if (!flag) {
			break
		}
		for i:=0 ; i<n ; i++ {
			if(timer >= begin[i] && duration[i] != 0){
				for c:=0 ; c<duration[i] ; c++{
					fmt.Printf("%s\t", name[i])
					fmt.Printf("%d\n", seconds)
					seconds++
				}
				timer = timer + duration[i]
				wait = wait + float64(timer) - float64(begin[i]) - float64(duration[i])
				duration[i] = 0
				flag = false
			}
		}
		if(flag){
			timer++;
			fmt.Printf("no\t")
			fmt.Printf("%d\n", seconds)
			seconds++
		}
	}
	return timer, wait
}

func main(){
	var n int
	wait := 0.0
	timer := 0
	seconds := 1
	var begin = make([]int,n) 
	var duration = make([]int,n) 
	var name = make([]string,n) 
	fmt.Println("How many processes do you have?")
	_, err1 := fmt.Scanf("%d", &n)
	_ = err1

	begin, duration, name = intro(begin, duration, name, n)
	sort(&begin, &duration, &name, n)
	timer, wait = sjf(begin, duration, name, n, timer, seconds, wait)

	fmt.Println("Total time:", timer)
	fmt.Println("Average waiting time:", wait/float64(n))
}