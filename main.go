package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type NameRecord struct {
	PlatformID string `xml:"platformID,attr"`
	PlatEncID  string `xml:"platEncID,attr"`
	LangID     string `xml:"langID,attr"`
	NameID     string `xml:"nameID,attr"`
	// Unicode    string `xml:"unicode,attr"`
	Text string `xml:",chardata"`
}

type NameTable struct {
	XMLName     xml.Name     `xml:"name"`
	NameRecords []NameRecord `xml:"namerecord"`
}

func xmlToCSV(reader io.Reader, writer io.Writer) error {
	xmlData, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("無法讀取 stdin: %v", err)
	}

	var nameTable NameTable
	if err := xml.Unmarshal(xmlData, &nameTable); err != nil {
		return fmt.Errorf("無法解析 XML: %v", err)
	}

	csvWriter := csv.NewWriter(writer)
	defer csvWriter.Flush()

	header := []string{
		"platformID", "platEncID", "langID",
		// "unicode",
		"nameID",
		"text",
	}
	if err := csvWriter.Write(header); err != nil {
		return fmt.Errorf("無法寫入 CSV 標頭: %v", err)
	}

	for _, record := range nameTable.NameRecords {
		text := strings.ReplaceAll(strings.TrimSpace(record.Text), `"`, `""`)
		row := []string{
			record.PlatformID,
			record.PlatEncID,
			record.LangID,
			record.NameID,
			// record.Unicode,
			text,
		}
		if err := csvWriter.Write(row); err != nil {
			return fmt.Errorf("無法寫入 CSV 記錄: %v", err)
		}
	}

	return nil
}

func main() {
	if err := xmlToCSV(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "錯誤: %v\n", err)
		os.Exit(1)
	}
}
