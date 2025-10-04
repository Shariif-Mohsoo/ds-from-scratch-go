package patterns

// import (
// 	"bytes"
// 	"encoding/base64"
// 	"fmt"
// 	"image"
// 	_ "image/gif"
// 	_ "image/jpeg"
// 	_ "image/png"
// 	"log"
// 	"os"
// 	"strings"

// 	"github.com/chai2010/webp"
// 	"github.com/google/uuid"
// 	// "modernc.org/libc/uuid/uuid"
// )

// func base64ToRawImage(base64Img string) image.Image {
// 	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64Img))

// 	img, _, err := image.Decode(reader)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return img

// }

// func encodeToWebp(img image.Image) bytes.Buffer {
// 	var buf bytes.Buffer
// 	if err := webp.Encode(&buf, img, &webp.Options{Lossless: true}); err != nil {
// 		log.Fatal(err)
// 	}
// 	return buf
// }

// func saveToDisk(imgBuf bytes.Buffer) string {
// 	filename := fmt.Sprintf("%v.webp", uuid.New().String())
// 	os.WriteFile(filename, imgBuf.Bytes(), 0644)
// 	return filename
// }

// func makeWork(base64Images ...string) <-chan string {
// 	out := make(chan string)

// 	go func() {
// 		for _, encodingImg := range base64Images {
// 			out <- encodingImg
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func pipeline[I any, O any](quit <-chan struct{}, input <-chan I, process func(I) O) <-chan O {
// 	out := make(chan O)
// 	go func() {
// 		defer close(out)
// 		for in := range input {
// 			select{
// 			case out <- process(in):
// 			case <- quit:
// 				return
// 			}
// 		}
// 	}()
// 	return out
// }

// func PipelineCancalation() {
// 	//load data into pipeline
// 	base64Images := makeWork(img1, img2, img3)

// 	quit := make(chan struct {})
// 	var signal struct{}

// 	//stages of pipeline...
// 	// decode base64 into image format
// 	rawImages := pipeline(quit,base64Images, base64ToRawImage)
// 	// encode as webp
// 	webImages := pipeline(quit,rawImages, encodeToWebp)

// 	quit <- signal
// 	//  save images to disk
// 	filenames := pipeline(quit,webImages, saveToDisk)
// 	for name := range filenames {
// 		fmt.Println(name)
// 	}
// }

// const img1 = "iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAIAAACRXR/mAAAAS0lEQVR4nO3OsQEAEADAMPz/Mw9YMjE0F2Tu8aP1OnBXS9QStUQtUUvUErVELVFL1BK1RC1RS9QStUQtUUvUErVELVFL1BK1RC1xAEGqAWOFuDKrAAAAAElFTkSuQmCC"
// const img2 = "iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAIAAACRXR/mAAAAS0lEQVR4nO3OsQEAEADAMPz/Mw9YOjEkF2SOPT60XgfutAqtQqvQKrQKrUKr0Cq0Cq1Cq9AqtAqtQqvQKrQKrUKr0Cq0Cq1Cq9AqDkCrAWMf8NpdAAAAAElFTkSuQmCC"
// const img3 = "iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAIAAACRXR/mAAAATUlEQVR4nO3OMQHAIBAAsQf/nlsDLJlguCjImvnmPft24KyWqCVqiVqilqglaolaopaoJWqJWqKWqCVqiVqilqglaolaopaoJWqJWuIHP6wBY/cJXlsAAAAASUVORK5CYII="
