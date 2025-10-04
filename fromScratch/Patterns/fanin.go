package patterns

// import (
// 	"bytes"
// 	"encoding/base64"
// 	"fmt"
// 	"image"
// 	"sync"

// 	_ "image/gif"

// 	_ "image/jpeg"

// 	_ "image/png"
// 	"log"
// 	"os"
// 	"strings"

// 	"github.com/chai2010/webp"
// 	"github.com/google/uuid"
// )

// // "modernc.org/libc/uuid/uuid"

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
// 			select {
// 			case out <- process(in):
// 			case <-quit:
// 				return
// 			}
// 		}
// 	}()
// 	return out
// }

// // will take multiple channels and merge them into single channel.
// func fanIn[T any](channels ...<-chan T) <-chan T {
// 	var wg sync.WaitGroup

// 	out := make(chan T)

// 	wg.Add(len(channels))

// 	for _, ch := range channels {
// 		go func(in <-chan T) {
// 			for i := range in {
// 				out <- i
// 			}
// 			wg.Done()
// 		}(ch)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()

// 	return out

// }

// func Fan_in() {
// 	//load data into pipeline
// 	base64Images := makeWork(img1, img2, img3)

// 	//stages of pipeline...
// 	// decode base64 into image format

// 	rawImages1 := pipeline(base64Images, base64ToRawImage)
// 	rawImages2 := pipeline(base64Images, base64ToRawImage)
// 	rawImages3 := pipeline(base64Images, base64ToRawImage)

// 	rawImages := fanIn(rawImages1, rawImages2, rawImages3)

// 	// encode as webp
// 	webImages1 := pipeline(rawImages, encodeToWebp)
// 	webImages2 := pipeline(rawImages, encodeToWebp)
// 	webImages3 := pipeline(rawImages, encodeToWebp)
// 	webImages := fanIn(webImages1, webImages2, webImages3)

// 	//  save images to disk
// 	filenames1 := pipeline(webImages, saveToDisk)
// 	filenames2 := pipeline(webImages, saveToDisk)
// 	filenames3 := pipeline(webImages, saveToDisk)
// 	filenames := fanIn(filenames1, filenames2, filenames3)
// 	for name := range filenames {
// 		fmt.Println(name)
// 	}
// }

// const img1 = "iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAIAAACRXR/mAAAAS0lEQVR4nO3OsQEAEADAMPz/Mw9YMjE0F2Tu8aP1OnBXS9QStUQtUUvUErVELVFL1BK1RC1RS9QStUQtUUvUErVELVFL1BK1RC1xAEGqAWOFuDKrAAAAAElFTkSuQmCC"
// const img2 = "iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAIAAACRXR/mAAAAS0lEQVR4nO3OsQEAEADAMPz/Mw9YOjEkF2SOPT60XgfutAqtQqvQKrQKrUKr0Cq0Cq1Cq9AqtAqtQqvQKrQKrUKr0Cq0Cq1Cq9AqDkCrAWMf8NpdAAAAAElFTkSuQmCC"
// const img3 = "iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAIAAACRXR/mAAAATUlEQVR4nO3OMQHAIBAAsQf/nlsDLJlguCjImvnmPft24KyWqCVqiVqilqglaolaopaoJWqJWqKWqCVqiVqilqglaolaopaoJWqJWuIHP6wBY/cJXlsAAAAASUVORK5CYII="
