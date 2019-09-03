package main

import (
        "bufio"
        "encoding/xml"
        "fmt"
        "io"
        "io/ioutil"
        "log"
        "os"
        "strings"
)

type Wordbook struct {
        XMLName xml.Name `xml:"wordbook"`
        Items   []WordItem
}
type WordItem struct {
        XMLName  xml.Name `xml:"item"`
        Word     string   `xml:"word"`
        Trans    CDATA    `xml:"trans"`
        Phonetic CDATA    `xml:"phonetic"`
        Tags     string   `xml:"tags"`
        Progress int      `xml:"progress"`
}
type CDATA string

func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
        return e.EncodeElement(struct {
                string `xml:",cdata"`
        }{string(c)}, start)
}

func main() {
        var wb Wordbook

        f, err := os.Open("import.txt")
        if err != nil {
                panic(err)
        }
        defer f.Close()

        rd := bufio.NewReader(f)
        for {
                line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

                if err != nil || io.EOF == err {
                        break
                }
                line = strings.TrimSpace(line)
                if resSlice := strings.Split(line, "|"); len(resSlice) >= 2 {

                        var a WordItem
                        a.Word = resSlice[0]
                        a.Trans = CDATA(resSlice[1])
                        a.Tags = "rustlangtrans"
                        a.Progress = -1
                        wb.Items = append(wb.Items, a)
                }

        }
        xmlOutPut, _ := xml.MarshalIndent(&wb, "", "  ")
        fmt.Println(string(xmlOutPut))
        err = ioutil.WriteFile("output.xml", xmlOutPut, 0644)
        if err != nil {
                log.Fatal(err)
        }

}
