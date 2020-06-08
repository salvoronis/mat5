package main

type point struct{
	x,y float32
}

func Adams(f func(float32,float32) float32, ans []point, rightX, h float32) []point{
	n := 3
	for ans[n].x < rightX - h/2 {
		k1 := f(ans[n].x, ans[n].y) - f(ans[n-1].x, ans[n-1].y)
		k2 := f(ans[n].x, ans[n].y) -2*f(ans[n-1].x, ans[n-1].y) + f(ans[n-2].x, ans[n-2].y)
		k3 := f(ans[n].x, ans[n].y) -3*f(ans[n-1].x, ans[n-1].y) +3*f(ans[n-2].x, ans[n-2].y) - f(ans[n-3].x, ans[n-3].y)
		newy := ans[n].y + h*f(ans[n].x, ans[n].y)+(6*h*h*k1 + 5*h*h*h*k2+4.5*h*h*h*h*k3)/12
		newpoint := point{x:ans[n].x+h, y:newy}
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