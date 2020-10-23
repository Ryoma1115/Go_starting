import "io"

type Reader interface {
	Read(p []byte) (n int, err error)
}

f, err := os.Open("foo.txt")
if err != nil {
	log.Fatal(err)
}

bs := make([]byte, 128)
n, err := f.Read(bs)