package main

import (
	"context"
	"fmt"
	"runtime"
	"sync/atomic"

	//"runtime"
	"sync"
	//"io/ioutil"
	"math"
	//"os"
	"strings"
	//"strings"
	"time"
)

var a, b, c, d int
var k int = 45

func gg(hi int) (x, y int) {
	x = hi - 5
	y = 5
	a = 45
	return hi, a
}

func ffor() {
	var sum int = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func while() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
}

func iff(x float64) float64 {
	if x < 0 {
		fmt.Println("Дебил")
		return math.Sqrt(-x)
	} else if x == 10 {
		return 10
	} else {
		x += x
	}
	return x
}

func sswitch(x int) {
	switch x {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("Больше двух")
	}
}

func Newton(x float64) float64 {
	var z float64 = 1
	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func ttime() {
	x := time.Now()
	fmt.Println(x)
	fmt.Println(x.Format("2006-01-02"))
	fmt.Println(x.Format("2006-01-02 15:04:05"))
}

func ddefer() {
	defer fmt.Println("world")
	fmt.Println("hello")
	//hello
	//world
}

var x int = 10
var p *int = &x

func pointer() {
	foo := 23
	pointerFOO := &foo
	println(pointerFOO)  // адрес
	println(*pointerFOO) // значение по адресу
	*pointerFOO = 10     // изменение foo по адресу
}

// * разыменование
// & Адрес
func changepointer(x *int) {
	change := *x
	change = change * change
	*x = change
} //меняет переданное значение

type Struct struct {
	Name string
	Age  int
	//v := Struct{
	//	Name: "Pete",
	//	Age:  20,
	//}
	//vv := Struct{"Vlad", 20}
	// p := &v
	// (*p).Age = 20
	// p.Age = 20
}

func arrays() {
	var a [2]int
	a[0] = 1
	a[1] = 2
	println(a[1]) // 2
	b := [3]int{1, 2, 3}
	fmt.Println(b)

}
func slice() {
	arr := [4]int{1, 2, 3, 4}
	var slice []int = arr[1:4]
	fmt.Println(slice) // 1, 2, 3
	s := []int{}
	s = append(s, 1)        //{1}
	s = append(s, 2)        //{1, 2	}
	println(len(s), cap(s)) // длинна и вместимость
}

func rrange() {
	var gg = []int{1, 2, 3, 4, 5}
	for i, v := range gg {
		fmt.Println(i, v)
	}
	for _, v := range gg {
		fmt.Println(v)
	}
}

var pow = []int{1, 2, 3, 4, 5}
var sslice []int = pow[1:2]

func mmap() {
	var m = map[int]int{
		1: 1,
		2: 2,
	}
	mm := make(map[string]int)
	mm["hi"] = 1
	fmt.Print(mm["hi"])
	//value, ok :=mm["hi"] если в "hi" есть value, ok == true
	delete(mm, "hi")
	for i, v := range m {
		fmt.Println(i, v)
	}
}

func twoSum(nums []int, target int) []int {
	for j := 0; j < len(nums); j++ {
		for i := j + 1; i < len(nums); i++ {
			if nums[j]+nums[i] == target {
				return []int{j, i}
			}
		}
	}
	return nil
}

func isPalindrome(x int) bool {
	y := fmt.Sprintf("%d", x)
	s := make([]string, len(y))
	j := len(s) - 1
	for i, ch := range y {
		s[i] = string(ch)
	}
	for i := 0; i < len(s)/2; i++ {
		g := s[i]
		s[i] = s[j]
		s[j] = g
		j--
		fmt.Println(s)
	}
	if strings.Join(s, "") == y {
		print(1)
		return true
	} else {
		print(0)
		return false
	}
}

type Inter interface {
	Write()
	Read()
}

type Person struct {
	Name string
	Age  int
}

func (i *Person) Write(name string, age int) {
	i.Name = name
	i.Age = age
}

func (i Person) Read() {
	fmt.Println(i.Name)
	fmt.Println(i.Age)
}

func emptyInter() {
	var i interface{} = "hi"
	s, ok := i.(string) //передает ok true и s=i если тип стринг
	fmt.Println(s, ok)
}

func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// ошибки через Error() string
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	var z float64 = 1
	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

/* для корректного вывода код в main
result, err := Sqrt(2)
	if err == nil {
		fmt.Println(result)
	}
x, err := Sqrt(-2)
	if err == nil{fmt.Println(x)
	}else{fmt.Println(err)}
*/

type List[T any] struct {
	next *List[T]
	val  T
	gg   string
}

func Find[T any](head *List[T], target string) (T, bool) {
	for n := head; n != nil; n = n.next {
		if n.gg == target {
			return n.val, true
		}
	}
	var zero T
	return zero, false
}

/* в main
node3 := List[int]{val: 3, gg: "завершение"}
    node2 := List[int]{val: 2, next: &node3, gg: "победа"}
    node1 := List[int]{val: 1, next: &node2, gg: "начало"}

    // Ищем элемент с gg = "победа"
    if val, found := Find(&node1, "победа"); found {
        fmt.Println("Найдено значение:", val)
    } else {
        fmt.Println("Значение не найдено")
    }
*/

/*
	работа с фАЙЛАМИ

	func main() {
		//запись
		data := []byte("Text")
		er := ioutil.WriteFile("text1.txt", data, 0600)
		if er != nil {
			fmt.Println("Не создается файл")
		}
		//чтение
		filedata, err := ioutil.ReadFile("text.txt")
		if err == nil {
			fmt.Println(string(filedata))
		} else {
			fmt.Printf("ошибка:%v\n", err)
		}
		// удаление
		os.Remove("text1.txt")
	}
*/

var (
	data  int
	mutex sync.RWMutex
	wg    sync.WaitGroup
)

/*/ Чтение данных
readData := func() {
	mutex.RLock()         // Блокировка для чтения
	defer mutex.RUnlock() // Разблокировка
	defer wg.Done()       // Сообщаем, что горутина завершилась
	fmt.Println("Read:", data)
}

// Запись данных
writeData := func(val int) {
	mutex.Lock()         // Блокировка для записи
	defer mutex.Unlock() // Разблокировка
	defer wg.Done()      // Сообщаем, что горутина завершилась
	data = val
	fmt.Println("Write:", data)
}

// Запускаем горутины
for i := 1; i <= 3; i++ {
	wg.Add(1)
	go readData()

	wg.Add(1)
	go writeData(i)
}

wg.Wait()

func say(ch chan int) {
	for i := 0; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

var f4 int = 123

*/

// func CHrange(wg *sync.WaitGroup) {
// 	buffchan := make(chan int)
// 	nums := []int{1, 2, 3, 4, 5}
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for _, num := range nums {

// 			buffchan <- num
// 		}
// 		close(buffchan)
// 	}()
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for {
// 			v, ok := <-buffchan
// 			fmt.Println(v, ok)
// 			if !ok {
// 				break
// 			}

// 		}
// 	}()
// 	cchan := make(chan int)
// 	go func() {
// 		for _, v := range nums {
// 			cchan <- v
// 		}
// 		close(cchan)
// 	}()
// 	//time.Sleep(5 * time.Second)
// 	for v := range cchan {
// 		fmt.Println(v)
// 	}

// }

// func main() {
// 	var wg sync.WaitGroup
// 	CHrange(&wg)
// 	wg.Wait()
// }

func WorkerPool() {
	ctx, cancel := context.WithCancel(context.Background())
	//ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*30)
	defer cancel()
	wg := &sync.WaitGroup{}
	numToProc := make(chan int, 5)
	ProcNum := make(chan int, 5)
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, numToProc, ProcNum)
		}()
	}
	go func() {
		for i := 0; i < 500; i++ {
			numToProc <- i
		}
		close(numToProc)
	}()

	go func() {
		wg.Wait()
		close(ProcNum)
	}()
	var counter int
	for resValue := range ProcNum {
		counter++
		fmt.Println(resValue)
	}
	fmt.Println(counter)

}

func worker(ctx context.Context, toProc <-chan int, procnum chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case value, ok := <-toProc:
			if !ok {
				return
			}
			time.Sleep(time.Millisecond)
			procnum <- value * value
		}
	}
}

// channel as promise
func someRequest(n int) <-chan string {
	newchan := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		newchan <- fmt.Sprint("ваш номер %d", n)
	}()
	return newchan
}

func ChanAsPromise() {
	firstChan := someRequest(1)
	secChan := someRequest(2)

	//какая-то работа
	//time.Sleep(5 * time.Second)
	fmt.Println(<-firstChan, <-secChan)
}

// channel as mutex
func chanAsMutex() {
	var counter int
	mutexChan := make(chan struct{}, 1)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutexChan <- struct{}{}
			counter++
			<-mutexChan
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

func atom() {
	var counter int64
	atomic.AddInt64(&counter, 5)            //изменение переменной
	fmt.Println(atomic.LoadInt64(&counter)) //загрузка переменной
	atomic.StoreInt64(&counter, 10)         //новое значение переменной
	fmt.Println(atomic.LoadInt64(&counter))
	fmt.Println(atomic.SwapInt64(&counter, 99)) //новое значение переменной + возврат старого значения
	fmt.Println(counter)
	if !atomic.CompareAndSwapInt64(&counter, 5, 10) {
		fmt.Println("тут")
		return
	}
}

// generic дженерики !_!
func showSum() {
	floats := []float64{1.0, 2.0, 3.0}
	ints := []int{1, 2, 3, 4}
	fmt.Println(independantSum(floats))
	fmt.Println(independantSum(ints))
}

func independantSum[V int | float64](numbers []V) V {
	var sum V
	for _, n := range numbers {
		sum += n
	}
	return sum
}

//Пакеты (не шмали :)
//go get "golang.org/x/sync/..."

func romanToInt(s string) int {
	runes := []rune(s)
	var strings []string
	var sum int
	for _, v := range runes {
		strings = append(strings, string(v))
	}
	strings = append(strings, "Z")
	for i := 0; i < len(strings)-1; i++ {
		switch {
		case strings[i] == "I" && strings[i+1] == "V":
			sum += 4
			i += 1
		case strings[i] == "I" && strings[i+1] == "X":
			sum += 9
			i += 1
		case strings[i] == "I" && strings[i+1] != "X" && strings[i+1] != "V":
			sum += 1
		case strings[i] == "X" && strings[i+1] == "L":
			sum += 40
			i += 1
		case strings[i] == "X" && strings[i+1] == "C":
			sum += 90
			i += 1
		case strings[i] == "X" && strings[i+1] != "L" && strings[i+1] != "C":
			sum += 10
		case strings[i] == "C" && strings[i+1] == "D":
			sum += 400
			i += 1
		case strings[i] == "C" && strings[i+1] == "M":
			sum += 900
			i += 1
		case strings[i] == "C" && strings[i+1] != "D" && strings[i+1] != "M":
			sum += 100
		case strings[i] == "V":
			sum += 5
		case strings[i] == "L":
			sum += 50
		case strings[i] == "D":
			sum += 500
		case strings[i] == "M":
			sum += 1000
		}
	}
	return sum
}

func longestCommonPrefix(strs []string) string {
	var a []string
	first := strs[0]
	var final string
	for _, v := range first {
		a = append(a, string(v))
	}
	for _, v1 := range strs {
		for i, v2 := range v1 {
			if len(a) > len(v1) {
				a = a[:len(v1)]
			}
			a = append(a, "Z")
			if string(v2) != a[i] {
				a = a[:i]
				break
			}
		}
	}
	//a{a, b}
	//a
	for _, v := range a {
		if v == "Z" {
			break
		}
		final = final + v
	}
	return final
}
func main() {
	x := []string{"ab", "a"}
	fmt.Println(longestCommonPrefix(x))
	y := []string{"ab", "a"}
	y = y[:0]
	fmt.Println(y)
	print("gggggggg")
}
