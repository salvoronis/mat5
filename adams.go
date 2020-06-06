package main

type point struct{
	x,y float32
}

func adams(f func(float32,float32) float32, ans []point, rightX, h float32) []point{
	n := 2
	h2 := h/2
	for ans[n].x < rightX - h2 {
		k1 := f(ans[n].x, ans[n].y)*h
		k2 := f(ans[n-1].x, ans[n].y)*h
		k3 := f(ans[n-2].x, ans[n].y)*h
		delt := (23*k1-16*k2+5*k3)/12
		newpoint := point{x:ans[n].x+h, y:ans[n].y+delt}
		ans = append(ans, newpoint)
		n++
	}
	return ans
}

func eiler(f func(float32,float32) float32, startPoint point, rightX, h float32) []point{
	ans := []point{startPoint}
	n := 0
	for ans[n].x < rightX - h/2 {
		newy := ans[n].y + h*f(ans[n].x,ans[n].y)
		newx := ans[n].x + h
		point := point{x:newx,y:newy}
		ans = append(ans, point)
		n++
	}
	return ans
}