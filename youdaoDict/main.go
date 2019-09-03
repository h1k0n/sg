package main

import (
        "encoding/xml"
        "fmt"
)

type Wordbook struct {
        Items []WordItem `xml:"wordbook"`
}
type WordItem struct {
        XMLName  xml.Name `xml:"item"`
        Word     string   `xml:"word"`
        Trans    CDATA    `xml:"trans"`
        Phonetic CDATA    `xml:"phonetic"`
        Tags     string   `xml:"tags"`
        Progress string   `xml:"progress"`
}
type CDATA string

func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
        return e.EncodeElement(struct {
                string `xml:",cdata"`
        }{string(c)}, start)
}

func main() {
        var a WordItem
        a.Trans = "aa"

        var wb Wordbook

        wb.Items = append(wb.Items, a)
        xmlOutPut, _ := xml.MarshalIndent(&wb, "", "  ")
        fmt.Println(string(xmlOutPut))

}
