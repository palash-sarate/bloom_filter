package bloomFilter

import (
	"bytes"
	"fmt"
	"hash"
	"html/template"
	"math/rand"
	"testing"
	"time"

	"github.com/spaolacci/murmur3"
)

type TestResult struct {
    Size          int
    NumHashers    int
    FalsePositives int
    FalseNegatives int
}

func generateDataset(size int) []string {
    rand.Seed(time.Now().UnixNano())
    dataset := make([]string, size)
    for i := range dataset {
        dataset[i] = fmt.Sprintf("item%d", rand.Intn(size*10))
    }
    return dataset
}

func testBloomFilter(size int, numHashers int, dataset []string) TestResult {
    hashers := make([]hash.Hash32, numHashers)
    for i := 0; i < numHashers; i++ {
        hashers[i] = murmur3.New32WithSeed(uint32(i))
    }
    bf := NewBloomFilter(size, numHashers, hashers)

    // Add half of the dataset to the Bloom filter
    for i := 0; i < len(dataset)/2; i++ {
        bf.Add([]byte(dataset[i]))
    }

    falsePositives := 0
    falseNegatives := 0

    // Check the entire dataset
    for i := 0; i < len(dataset); i++ {
        contains := bf.Contains([]byte(dataset[i]))
        if i < len(dataset)/2 && !contains {
            falseNegatives++
        } else if i >= len(dataset)/2 && contains {
            falsePositives++
        }
    }

    return TestResult{
        Size:          size,
        NumHashers:    numHashers,
        FalsePositives: falsePositives,
        FalseNegatives: falseNegatives,
    }
}

func generateHTMLReport(results []TestResult) string {
    const tpl = `
<!DOCTYPE html>
<html>
<head>
    <title>Bloom Filter Test Report</title>
</head>
<body>
    <h1>Bloom Filter Test Report</h1>
    <table border="1">
        <tr>
            <th>Bit Array Size</th>
            <th>Number of Hash Functions</th>
            <th>False Positives</th>
            <th>False Negatives</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.Size}}</td>
            <td>{{.NumHashers}}</td>
            <td>{{.FalsePositives}}</td>
            <td>{{.FalseNegatives}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>
`
    t := template.Must(template.New("report").Parse(tpl))
    var buf bytes.Buffer
    t.Execute(&buf, results)
    return buf.String()
}

func TestBloomFilter(t *testing.T) {
    dataset := generateDataset(10000)
    results := []TestResult{}

    // Test different configurations
    for _, size := range []int{1000, 2000, 5000} {
        for _, numHashers := range []int{3, 5, 7} {
            result := testBloomFilter(size, numHashers, dataset)
            results = append(results, result)
        }
    }

    // Generate HTML report
    report := generateHTMLReport(results)
    fmt.Println(report)
}