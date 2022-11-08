package main

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/golang/freetype"
	"github.com/nfnt/resize"
)

// GetImage 根据路径获取一个图片
func GetImage(imagePath string) *image.Image {
	// 读取image文件
	imageFile, _ := ioutil.ReadFile(imagePath)
	// bytes.buffer 是一个缓冲byte类型的缓冲器存放着都是byte
	imageBytes := bytes.NewBuffer(imageFile)
	// 解码
	im, _, err := image.Decode(imageBytes)
	if err != nil {
		log.Printf("解码图像失败%#v", err)
		return nil
	}
	return &im
}

// equalScaleImageFromWidth 根据宽度，等比例缩放图片, 返回新的像素值（等比例处理后的）
// maxWidth 图片最大宽度 如果超过此值就按照宽度 targetWidth 等比例缩放图片
func equalScaleImageFromWidth(img image.Image, maxWidth int, targetWidth int) *image.Image {
	// 原图像界限范围
	bounds := img.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	log.Printf("原图宽=%d,高=%d", dx, dy)
	m := img
	if dx > maxWidth {
		// 按照宽度缩放
		m = resize.Resize(uint(targetWidth), 0, img, resize.Lanczos3)
	}
	return &m
}

// SaveNewImage 保存一个新的图片文件
func SaveNewImage(targetPath string, newImg *image.Image) {
	f, err := os.Create(targetPath)
	if err != nil {
		log.Printf("创建文件失败%#v", err)
	}
	defer f.Close()
	// 获取文件后缀判断编码方法
	suffix := path.Ext(targetPath) // strings.Replace(path.Ext(targetPath), ".", "", -1)
	if suffix == ".jpg" || suffix == ".jpge" {
		err = jpeg.Encode(f, *newImg, nil)
	} else if suffix == "png" {
		err = png.Encode(f, *newImg)
	} else if suffix == "gif" {
		err = gif.Encode(f, *newImg, nil)
	} else {
		log.Printf("非可处理图片格式")
	}
	if err != nil {
		log.Printf("保存一个新的图片失败%#v", err)
	}
	log.Printf("保存一个新的图片成功！")
}

// ImageToChars 图像转成字符展示
// RGBA(R,G,B,A) 像素值
// R：红色值 G：绿色值 B：蓝色值 A：Alpha透明度
func ImageToChars(img *image.Image, inputChars string) string {
	sourceImage := *img
	bounds := sourceImage.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	imgChars := "**  "
	if inputChars != "" {
		imgChars = inputChars
	}
	resultString := ""
	intSliceRGB := []int{}
	maxIntRGB := 0
	minIntRGB := 255 * 3
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			colorRgb := sourceImage.At(j, i)
			// 获取 uint32 像素值 >> 8 转换为 255 值
			r, g, b, _ := colorRgb.RGBA()
			// sumRGB 越大越趋近于白色，越小越趋近于黑色
			sumRGB := int(uint8(r>>8)) + int(uint8(g>>8)) + int(uint8(b>>8))
			// 找到最大值和最小值，方便将像素值划分不同区间段
			if maxIntRGB < sumRGB {
				maxIntRGB = sumRGB
			}
			if minIntRGB > sumRGB {
				minIntRGB = sumRGB
			}
			intSliceRGB = append(intSliceRGB, sumRGB)
		}
	}
	for index, val := range intSliceRGB {
		// partLen为区间跨度 我们按照传入字符串元素个数平均分配像素值，将像素值分成几个区间
		//  +1 防止下标溢出
		// 例如 像素值 0~500 用两个字符替换 0~250 251~500 两个区间分别与其对应即可
		partLen := (maxIntRGB-minIntRGB)/(len(imgChars)) + 1
		// 根据像素值取不同的字符替换像素
		str := string(imgChars[(val-minIntRGB)/partLen])
		resultString += str
		// 判断换行
		if (index+1)%dx == 0 {
			resultString += "\n"
		}
	}
	return resultString
}

// WriteStringToFile 通过 io.WriteString 写入文件
func WriteStringToFile(fileName string, writeInfo string) {
	f, err := os.Create(fileName)
	if err != nil {
		log.Printf("创建文件失败:%s，日志%#v", fileName, err)
		return
	}
	log.Printf("成功创建文件:%s", fileName)
	defer f.Close()
	// 将文件写进去
	if _, err = io.WriteString(f, writeInfo); err != nil {
		log.Printf("WriteStringToFile 写入文件失败:%+v", err)
		return
	}
	log.Printf("WriteStringToFile 写入文件成功")
}

func crateImageFromString(fileName string, imgString string, width int, height int) *image.Image {
	imgFile, err := os.Create(fileName)
	if err != nil {
		log.Printf("创建文件失败:%s，日志%#v", fileName, err)
		return nil
	}
	defer imgFile.Close()
	// 创建位图,坐标x,y,长宽x,y
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	// 画背景,这里可根据喜好画出背景颜色
	bounds := img.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			// 设置某个点的颜色 z这里设置白色
			img.Set(x, y, color.RGBA{uint8(255), uint8(255), 255, 255})
		}
	}
	// 读字体数据
	fontBytes, err := ioutil.ReadFile("ht.ttf")
	if err != nil {
		log.Println(err)
		return nil
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return nil
	}
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(200)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.Black)
	// 设置字体显示位置
	// int(c.PointToFixed(200)>>8)
	log.Printf("**********>%d", int(c.PointToFixed(200)>>6))
	pt := freetype.Pt((width-len([]rune(imgString))*200)/2, int(c.PointToFixed(200)>>8))
	_, err = c.DrawString(imgString, pt)
	if err != nil {
		log.Printf("crateImageFromString 字符串画图失败:%#v", err)
		return nil
	}
	// 保存图像到文件
	err = png.Encode(imgFile, img)
	if err != nil {
		log.Fatal(err)
	}
	return GetImage(fileName)
}

// CrateImageFromString 根据字符串创建图片
func CrateImageFromString(fileName string, imgString string, width int, height int, fontSize float64) *image.Image {
	// 读字体数据
	fontBytes, err := ioutil.ReadFile("ht.ttf")
	if err != nil {
		log.Printf("CrateImageFromString 读字体信息失败:%#v", err)
		return nil
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Printf("CrateImageFromString 解析字体失败:%#v", err)
		return nil
	}

	// 设置图片范围和底色
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), image.White, image.ZP, draw.Src)

	c := freetype.NewContext()
	c.SetDPI(100) // 设置像素分辨率 每英寸点数 越大像素越好
	c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.Black)

	// 设置字体显示位置
	// >>6 把Int26_6转回int
	pt := freetype.Pt(10, int(c.PointToFixed(fontSize)>>6))
	for _, s := range strings.Split(imgString, "\n") {
		_, err = c.DrawString(s, pt)
		if err != nil {
			log.Printf("CrateImageFromString 字符串画图失败:%#v", err)
			return nil
		}
		pt.Y += c.PointToFixed(fontSize)
	}

	// 创建图片文件
	imgFile, err := os.Create(fileName)
	if err != nil {
		log.Printf("创建文件失败:%s，日志%#v", fileName, err)
		return nil
	}
	defer imgFile.Close()

	// 保存图像到文件
	err = png.Encode(imgFile, img)
	if err != nil {
		log.Fatal(err)
	}
	return GetImage(fileName)
}

func main() {
	WriteStringToFile("james.txt", ImageToChars(equalScaleImageFromWidth(*(GetImage("1.jpg")), 200, 200), "#$*@    "))
	WriteStringToFile("smallSuperMan.txt", ImageToChars(equalScaleImageFromWidth(*(CrateImageFromString("1.png", "小超人\n BUG \n不会飞", 500, 500, 50)), 150, 150), "|.    "))
}
