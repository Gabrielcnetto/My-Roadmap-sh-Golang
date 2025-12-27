package main

/*
ESCREVENDO USANDO O MAKE CONVENCIONAL PARA UM MAP

	func main() {
		c := make(map[string]string)
		wg := sync.WaitGroup{}
		var lock sync.Mutex
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(n int) {
				k, v := strconv.Itoa(n), strconv.Itoa(n)
				lock.Lock()
				c[k] = v
				lock.Unlock()
				wg.Done()
			}(i)
		}
		wg.Wait()
		fmt.Println(c)
	}
*/
/*
 GERALMENTE QUANDO TEM GOROUNTINE, E DPS QUE SALVA PRECISA LER E QUER MANTER A VELOCIADED DO GOROUNTINE, É MAIS INTERESSANTE USAR O SYNC.MAP, CONTANTO QUE NAO TENHA VARIAS ESCRITAS
 ACONTECENDO AO MESMO TEMPO, SE FOR ISSO, MELHORAR USAR O MAKE
func main() {
	var syncMap sync.Map

	// store a key-value pair
	syncMap.Store("blog", "VictoriaMetrics")

	// load a value by key "blog"
	value, ok := syncMap.Load("blog")
	fmt.Println(value, ok)

	// delete a key-value pair by key "blog"
	syncMap.Delete("blog")
	value, ok = syncMap.Load("blog")
	fmt.Println(value, ok)
}
*/
// Output:
// VictoriaMetrics true
// <nil> false
