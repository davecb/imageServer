package imageServer

import (

	//"github.com/nfnt/resize"
	"image"
	//"image/jpeg"
	//"log"
	//"bytes"
	//"image/png"
	"time"
	"fmt"
)

// T is a debugging tool shared by the server components
var sample  image.Image

// Image strictly resizes an image.
func resizeImage(contents, key string, width, height, quality uint, grayScale bool, name, imgType string) string {
	defer T.Begin("<contents>", key, width, height, quality, grayScale, name, imgType)()

	return ""

	//buf := new(bytes.Buffer)
	//
	//initial := time.Now()
	//m := resize.Resize(width, height, contents, resize.NearestNeighbor)   // FIXME
	//switch {
	//case imgType == "jpg":
	//	opt := jpeg.Options{Quality: int(quality)}
	//	err := jpeg.Encode(buf, m, &opt)
	//	if err != nil {
	//		log.Fatalf("jpg write failure %v\n", err)
	//	}
	//case imgType == "png":
	//	err := png.Encode(buf, m)
	//	if err != nil {
	//		log.Fatalf("png write failure %v\n", err)
	//	}
	//	//ico
	//	// jpg
	//	// pdf
	//	// png
	//
	//default:
	//	log.Fatal("not a jpg") // FIXME
	//}
	//resizeTime := time.Since(initial)
	//reportPerformance(initial, resizeTime, 0,0, buf.Len(), 200, key)
	//return buf.String()
}

// reportPerformance in standard format
func reportPerformance(initial time.Time, latency, xferTime,
		thinkTime time.Duration, length int, rc int, key string) {

	fmt.Printf("%s %f %f %f %d %s %d RESIZE\n",
	initial.Format("2006-01-02 15:04:05.000"),
	latency.Seconds(), xferTime.Seconds(), thinkTime.Seconds(),
	length, key, rc)
}
