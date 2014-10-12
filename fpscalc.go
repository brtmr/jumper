package main

const FPS_ARRAY_LENGTH = 20

type FpsCalc struct {
	cur        int
	fps        int
	frametimes [FPS_ARRAY_LENGTH]int
}

func (f *FpsCalc) Put(time int) {
	f.frametimes[f.cur] = time
	f.cur = (f.cur + 1) % FPS_ARRAY_LENGTH
	if f.cur == FPS_ARRAY_LENGTH-1 {
		total := 0
		for _, time := range f.frametimes {
			total += time
		}
		if total != 0 {
			f.fps = (1000 * FPS_ARRAY_LENGTH) / total
		}
	}
}

func (f FpsCalc) FPS() int {
	return f.fps
}

func Init_FpsCalc() FpsCalc {
	var arr [FPS_ARRAY_LENGTH]int
	return FpsCalc{0, 0, arr}
}
