package main

import "fmt"

func main() {

	var css []string
	fmt.Printf("%#v\n", css)
	fmt.Printf("%T\n", css)

	css = append(css, "go")
	css = append(css, "go")
	css = append(css, "go")
	css = append(css, "go")
	fmt.Println(css)

	//初始化
	//1.从数组直接创建
	//2.使用string{}
	//3.make
	css1 := [5]string{"aaa", "bbb", "ccc", "ddd", "eee"}
	css2 := css1[0:2] //左开右闭[)
	fmt.Println(css2)

	//make初始化长度为5，申请了10个容量，所以空间内可以赋值
	//如果直接定义切片没有开辟空间，赋值会报错：
	//css3 := []string
	//css3[0] = "aaa" //报错，应该用append追加
	css3 := make([]string, 5, 10)
	css3[0] = "aaa"
	fmt.Println(css3)

	var css4 []string
	css4 = append(css4, "aaa")
	fmt.Println(css4)

	//访问切片数据
	fmt.Println(css1[:])
	fmt.Println(css1[1:])
	fmt.Println(css1[:4])
	fmt.Println(css1[1:4])

	//追加
	cosSlice := []string{"aaa", "bbb", "ccc", "ddd", "eee"}
	fmt.Println(len(cosSlice), cap(cosSlice))
	cosSlice = append(cosSlice, "fff")
	fmt.Println(len(cosSlice), cap(cosSlice))

	cosSlice2 := []string{"a", "b", "c", "d", "e"}
	cosSlice3 := append(cosSlice, cosSlice2...) //第二个参数是string。用...把数组切片打散
	fmt.Println(cosSlice3, len(cosSlice3), cap(cosSlice3))

	//删除元素
	cosSlice4 := []string{"aaa", "bbb", "ccc", "ddd", "eee"}
	//删除 ccc
	cosSlice5 := append(cosSlice4[:2], cosSlice4[3:]...)
	fmt.Println(cosSlice5, len(cosSlice5), cap(cosSlice5))

	//复制slice
	cosSliceCopy := cosSlice
	cosSliceCopy2 := cosSlice[:]
	fmt.Println(cosSliceCopy, cosSliceCopy2, len(cosSliceCopy), cap(cosSliceCopy), len(cosSliceCopy), cap(cosSliceCopy))

	var cosSliceCopy3 []string    //空间为0，就算copy也无法赋值，
	copy(cosSliceCopy3, cosSlice) //所以cosSliceCopy3为空[]
	fmt.Println(cosSlice, cosSliceCopy3)
	var cosSliceCopy4 = make([]string, 5, 10)
	copy(cosSliceCopy4, cosSlice) //原切片有6个元素，但是cosSliceCopy4初始长度只有5，所以这里只复制到5个元素
	fmt.Println(cosSlice, cosSliceCopy4, len(cosSliceCopy4), cap(cosSliceCopy4))
	//修改原数据
	//用[:]复制的切片会跟着修改cosSliceCopy4
	//用copy方法复制的切片不受影响cosSliceCopy2
	cosSlice[0] = "000"
	fmt.Println(cosSlice, cosSliceCopy4, cosSliceCopy2)

	// slice 底层实现
	// 函数参数传递的时候是 值传递 还是 引用传递
	// 是值传递，整个slice结构体复制了一份，但是结构体里是使用指针指向存放数据的地址
	// 所以复制时，是值传递，但是内部的数据是引用同一份
	// 只有在复制的切片发生扩容时，复制的切片里指向数据的指针才会重新指向扩容后新数据的地址
	arrSlice := []string{"1", "2", "3"}
	printSlice(arrSlice)
	fmt.Println(arrSlice)

	// sls = append(sls, "aaa")
	//为什么append必须要有返回值,而不是直接在传入的切片上修改数据就好了
	//因为是值传递，如果append没有发生扩容，那么sls里的指针不变，没有问题
	//如果append后发生了扩容，那么sls里的指针会被修改，指向扩容后新的数据地址，在append外部的sls里的指针指向的还是原来的地址

}

/*
切片底层实现：

	struct {
		pointer, //指针指向实际存放数据的地址
		len,	 //长度
		cap		 //容量
	}

当切片容量不够时会翻倍扩容，如果扩容后大于1024，改为每次扩容25%
*/
func printSlice(arrslc []string) {
	//传递的是值，但是切片里的数据是使用指针进行引用的，所以直接改会改到切片里指针指向的数据
	arrslc[0] = "000"
	//append追加后发生扩容，切片内部的指针指向了新的数据地址，所以传进来的原切片数据不受影响
	fmt.Println("append前，容量：", cap(arrslc))
	arrslc = append(arrslc, "4")
	fmt.Println("append后，容量：", cap(arrslc))
	for _, v := range arrslc {
		fmt.Println(v)
	}
}
