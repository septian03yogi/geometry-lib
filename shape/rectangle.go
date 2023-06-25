package shape

//Nama field dan method harus diawali huruf besar

type Rectangle struct {
	//membuat struct dengan field width dan length
	Width  float32
	Length float32
}

func (rectangle *Rectangle) Area() float32 {
	//method Area akan menghitung luas
	return rectangle.Width * rectangle.Length
}

func (rectangle *Rectangle) Perimater() float32 {
	//method perimeter akan menghitung keliling rectangle
	return 2 * (rectangle.Length + rectangle.Width)
}
