package implement

import "log"

type DirectHandler interface {
	CheckDirectLink(info string, title string) (bool, error)
}

type BillManager interface {
	Count()
}

type Download struct {
	DirectHandler DirectHandler
}

type Handler struct {
	billManager BillManager
}

type BillManagerI struct {
}

func (b *BillManagerI) Count() {

}

func (d *Download) Check() (bool, error) {
	return d.DirectHandler.CheckDirectLink("download", "download title")
}

func (d *Handler) CheckDirectLink(info string, title string) (bool, error) {
	log.Print("origin info: ", info)
	log.Print("origin title: ", title)
	return false, nil
}

func main() {
}
