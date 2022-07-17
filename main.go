package main

import (
	"clisimulator/theme"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()
var indent = '-'

//func showField(field reflect.StructField) bool {
//switch field.Name {
//case "state", "sizeCache", "unknownFields":
//return false
//}
//return true
//}

//func buildProto(valMap map[string]*binding.String, valPtr reflect.Value) {
//log.Debug("enter buildProto ", len(valMap))
//value := valPtr.Elem()
//for name, val := range valMap {
//field := value.FieldByName(name)
//v, _ := (*val).Get()
//log.Debug("name ", name, "kind ", field.Kind(), "value ", v)
//switch field.Kind() {
//case reflect.String:
//if field.CanSet() {
//strVal, _ := (*val).Get()
//field.SetString(strVal)
//log.Info("field", name, "set value", strVal)
//} else {
//log.Error("field error", field)
//}

//case reflect.Pointer:
//if name != "Msg" {
//continue
//}

//log.Info("enter msg field", field.Elem().Kind())
//str, _ := (*val).Get()
//strValue := reflect.ValueOf(&str)
//strValue.Elem().SetString(str)

//field.Set(strValue)

//// elem := field.Elem()
//// log.Debug("pointer", elem.Kind())
//// switch elem.Kind() {
//// case reflect.String:
////         log.Debug("ptr string")
////
//// }
//}
//}
//}

//func CreateProtoWindow() fyne.CanvasObject {
//content := container.NewVBox()
//pbObj := pb.SayHello{}
//typeObj := reflect.TypeOf(pbObj)
//content.Add(widget.NewLabel(typeObj.Name()))

//valueMap := make(map[string]*binding.String)
//for i := 0; i < typeObj.NumField(); i++ {
//field := typeObj.Field(i)

//if !showField(field) {
//continue
//}

//bStr := binding.NewString()
//valueMap[field.Name] = &bStr
//entry := widget.NewEntryWithData(bStr)
//form := &widget.Form{
//Items: []*widget.FormItem{
//{Text: fmt.Sprintf("%-10s", field.Name), Widget: entry}},
//}
//content.Add(form)
//}
//content.Add(widget.NewButton("build proto", func() {
//buildProto(valueMap, reflect.ValueOf(&pbObj))
//log.Printf("build proto click %+v \n", pbObj)
//}),
//)
//return content
//}

func CreateProtoWindow() fyne.CanvasObject {
	content := container.NewVBox()
	//pbObj := pb.SayHello{}
	content.Add(widget.NewLabel("proto message"))

	//valueMap := make(map[string]*binding.String)

	//msgDefs := ParseFile("msg.proto.parse")
	msgDefs := ParseFile("test.proto.parse")

	// for debug
	for _, msg := range msgDefs {
		fmt.Print(msg)

		for _, field := range msg.Fields {
			fmt.Print("\t", field)
		}
	}

	for i := 0; i < 2; i++ {

		bStr := binding.NewString()
		entry := widget.NewEntryWithData(bStr)
		form := &widget.Form{
			Items: []*widget.FormItem{
				{Text: fmt.Sprintf("%-10s", "field1"), Widget: entry}},
		}
		content.Add(form)
	}

	content.Add(widget.NewButton("build proto", func() {
	}))

	return content
}

func init() {
	log.SetLevel(logrus.DebugLevel)
}

func main() {
	// log.Info("Hello, world!")
	a := app.New()
	a.Settings().SetTheme(&theme.MyTheme{})

	w := a.NewWindow("Hello World")

	w.SetContent(CreateProtoWindow())

	w.Resize(fyne.NewSize(400, 300))
	w.CenterOnScreen()
	w.ShowAndRun()
}
