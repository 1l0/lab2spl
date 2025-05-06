package construct

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"

	"github.com/1l0/lab2spl/model"
)

func Lab2Spl(path string, fps int) (string, error) {
	lab, err := loadLab(path)
	if err != nil {
		return "", err
	}
	log.Printf("lab loaded: %d", len(lab))
	spl := "DFSP\n"
	for _, el := range lab {
		frame := fmt.Sprintf("%.6f", math.Floor(el.Start*float64(fps)))
		value := fmt.Sprintf("%.6f", float64(el.Index))
		spl += fmt.Sprintf("%s %s\n", frame, value)
	}
	return spl, nil
}

func BOMAwareCSVReader(reader io.Reader) *csv.Reader {
	var transformer = unicode.BOMOverride(encoding.Nop.NewDecoder())
	return csv.NewReader(transform.NewReader(reader, transformer))
}

func loadLab(path string) (model.Lab, error) {
	readfile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer readfile.Close()
	reader := BOMAwareCSVReader(readfile)
	reader.Comma = ' '
	reader.TrimLeadingSpace = true
	recs, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	lab := make([]model.LabElement, 0, len(recs))
	for i, rec := range recs {
		if len(rec) != 3 {
			return nil, fmt.Errorf("lab file broken at :%d: %v", i, rec)
		}
		start, err := strconv.ParseFloat(rec[0], 64)
		if err != nil {
			return nil, err
		}
		end, err := strconv.ParseFloat(rec[1], 64)
		if err != nil {
			return nil, err
		}
		idx, ok := model.PhonemeIndex[rec[2]]
		if !ok {
			return nil, fmt.Errorf("unknown phoneme: %s", rec[2])
		}
		lab = append(lab, model.LabElement{
			Start: start * 0.0000001,
			End:   end * 0.0000001,
			Index: idx,
		})
	}
	log.Printf("lab elements loaded: %d\n", len(lab))
	return lab, nil
}
