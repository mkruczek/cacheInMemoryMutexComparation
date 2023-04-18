package channel

type Cache struct {
	data   map[string]int
	setCh  chan cacheObject
	delCh  chan string
	readCh chan string
	getCh  chan int
}

type cacheObject struct {
	key   string
	value int
}

func NewCache() *Cache {
	c := &Cache{
		data:   make(map[string]int),
		setCh:  make(chan cacheObject, 100),
		delCh:  make(chan string, 100),
		readCh: make(chan string, 100),
	}

	go c.loop()

	return c
}

func (c *Cache) loop() {
	for {
		select {
		case co := <-c.setCh:
			c.data[co.key] = co.value
		case k := <-c.delCh:
			delete(c.data, k)
		case k := <-c.readCh:
			c.getCh <- c.data[k]
		}
	}
}

func (c *Cache) Set(key string, val int) {
	co := cacheObject{
		key:   key,
		value: val,
	}

	c.setCh <- co
}

func (c *Cache) Get(key string) (int, bool) {
	c.readCh <- key
	result := <-c.getCh

	//naive check for zero value... in general, this is a bad idea ;)
	if result == 0 {
		return 0, false
	}

	return result, true
}

func (c *Cache) Delete(key string) {
	c.delCh <- key
}
