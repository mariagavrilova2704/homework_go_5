package ioflow

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
)

// TODO: ReadAllText должен прочитать все данные из io.Reader и вернуть строку.
func ReadAllText(r io.Reader) (string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// TODO: CountBytes должен посчитать количество байт, доступных через Reader, и вернуть ошибку чтения без потери.
func CountBytes(r io.Reader) (int, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	return len(data), nil
}

// TODO: WriteString должен записать text в Writer и вернуть количество записанных байт и ошибку.
func WriteString(w io.Writer, text string) (int, error) {
	n, err := fmt.Fprint(w, text)
	if err != nil {
		return n, err
	}
	return n, nil
}

// TODO: WriteLines должен записать каждую строку с переводом строки. При ошибке остановиться и вернуть её.
func WriteLines(w io.Writer, lines []string) error {
	for _, line := range lines {
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return nil
}

// TODO: CopyAll должен скопировать все данные из src в dst и вернуть число байт.
func CopyAll(dst io.Writer, src io.Reader) (int64, error) {
	data, err := io.ReadAll(src)
	if err != nil {
		return 0, err
	}

	return io.Copy(dst, strings.NewReader(string(data)))

	// return io.Copy(dst,src)
}

// TODO: CopyUpper должен прочитать src, привести текст к верхнему регистру и записать в dst.
func CopyUpper(dst io.Writer, src io.Reader) (int64, error) {
	data, err := io.ReadAll(src)
	if err != nil {
		return 0, err
	}

	return io.Copy(dst, strings.NewReader(strings.ToUpper(string(data))))
}

// TODO: ScanLines должен вернуть все строки из Reader без символов перевода строки.
func ScanLines(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// TODO: CountNonEmptyLines должен посчитать строки, которые после TrimSpace не пустые.
func CountNonEmptyLines(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	counter := 0
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) != "" {
			counter++
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return counter, nil
}

// TODO: FirstLine должен вернуть первую строку. Если строк нет, вернуть пустую строку без ошибки.
func FirstLine(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		return scanner.Text(), nil
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", nil
}

// TODO: ReadCSVLike должен прочитать строки, разделить каждую по запятой и обрезать пробелы у ячеек.
func ReadCSVLike(r io.Reader) ([][]string, error) {
	scanner := bufio.NewScanner(r)
	cells := [][]string{}
	for scanner.Scan() {
		lines := scanner.Text()
		parts := strings.Split(lines, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		cells = append(cells, parts)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return cells, nil
}

// TODO: LimitRead должен прочитать не больше limit байт и вернуть строку. Отрицательный limit считать нулём.
func LimitRead(r io.Reader, limit int) (string, error) {
	if limit < 0 {
		limit = 0
	}
	buf := make([]byte, limit)
	n, err := io.ReadFull(r, buf)

	if err == io.EOF || err == io.ErrUnexpectedEOF {
		err = nil
	}

	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

// TODO: RepeatToWriter должен записать text count раз подряд. count <= 0 ничего не пишет.
func RepeatToWriter(w io.Writer, text string, count int) error {
	if count <= 0 {
		return nil
	}

	for i := 0; i < count; i++ {
		_, err := io.WriteString(w, text)

		if err != nil {
			return err
		}
	}
	return nil
}

// TODO: BufferReport должен вернуть строки, записанные через Writer в формате "<index>:<line>\n".
func BufferReport(lines []string) string {
	var buf bytes.Buffer

	for i, line := range lines {
		result := fmt.Sprintf("%d:%s\n", i, line)
		fmt.Fprint(&buf, result)
	}

	final := buf.String()
	return final
}

// TODO: ReadAndTrim должен прочитать весь Reader и вернуть строку без пробелов по краям.
func ReadAndTrim(r io.Reader) (string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	str := strings.TrimSpace(string(data))
	return str, nil
}

// TODO: WriteKeyValues должен записать пары map в отсортированном по ключу порядке: "key=value\n".
func WriteKeyValues(w io.Writer, values map[string]string) error {

	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, key := range keys {
		_, err := fmt.Fprintf(w, "%s=%s\n", key, values[key])
		if err != nil {
			return err
		}
	}
	return nil
}

func Example() string {
	var out strings.Builder
	text, _ := ReadAllText(strings.NewReader("hello"))
	out.WriteString("read: " + text + "\n")
	count, _ := CountNonEmptyLines(strings.NewReader("a\n\n b \n"))
	out.WriteString("non-empty: ")
	out.WriteString(strings.TrimSpace(BufferReport([]string{string(rune('0' + count))})))
	return out.String()
}
