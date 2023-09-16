package main

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
}

func (m MultiFunctionPrinter) Print(d Document) {

}

func (m MultiFunctionPrinter) Fax(d Document) {

}

func (m MultiFunctionPrinter) Scan(d Document) {

}

type OldFashionPrinter struct {
}

func (m OldFashionPrinter) Print(d Document) {
	// ok
}

// Deprecated ...
func (m OldFashionPrinter) Fax(d Document) {
	panic("operation not support")
}

// Deprecated ...
func (m OldFashionPrinter) Scan(d Document) {
	panic("operation not support")
}

// ISP -> try to break up interface into small portion

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct {
}

func (m MyPrinter) Print(d Document) {

}

type PhotoCopier struct {
}

func (m PhotoCopier) Print(d Document) {

}

func (m PhotoCopier) Scan(d Document) {

}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	doc := Document{}
	m.printer.Print(doc)
}

func main() {

}
