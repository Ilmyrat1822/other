package main

import (
	"fmt"
	"strings"
	"unicode"
)

type InheritanceResult int

const (
	None InheritanceResult = iota
	Mother
	Father
	Both
	Neutral
)

func (r InheritanceResult) String() string {
	switch r {
	case None:
		return "None"
	case Mother:
		return "Mother"
	case Father:
		return "Father"
	case Both:
		return "Both"
	case Neutral:
		return "Neutral"
	default:
		return "Unknown"
	}
}

type Horse struct {
	Name string
	DNA  map[string]string
}

type CellColor int

const (
	Transparent CellColor = iota
	Red
)

type AnalysisResult struct {
	Marker          string
	ChildValue      string
	FatherValue     string
	MotherValue     string
	Result          InheritanceResult
	FatherCellColor CellColor
	MotherCellColor CellColor
}

type DNAAnalyzer struct {
	NeutralRow int
}

func NewDNAAnalyzer() *DNAAnalyzer {
	return &DNAAnalyzer{NeutralRow: 1}
}

func (a *DNAAnalyzer) ExtractAlleles(dna string) map[string]bool {
	alleles := make(map[string]bool)

	if dna == "" {
		return alleles
	}

	cleaned := strings.NewReplacer(
		"(", "",
		")", "",
		"*", "",
		"?", "",
		"a.", "",
		"a", "",
		"e.", "",
		"e", "",
		" ", "",
	).Replace(dna)

	parts := strings.FieldsFunc(cleaned, func(r rune) bool {
		return r == '/' || r == ';' || r == ','
	})

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}

		for _, c := range trimmed {
			if unicode.IsLetter(c) || unicode.IsDigit(c) {
				alleles[strings.ToLower(string(c))] = true
			}
		}
	}

	return alleles
}

func (a *DNAAnalyzer) HasCommonAllele(dna1, dna2 string) bool {
	if dna1 == "" || dna2 == "" {
		return false
	}

	alleles1 := a.ExtractAlleles(dna1)
	alleles2 := a.ExtractAlleles(dna2)

	for allele := range alleles1 {
		if alleles2[allele] {
			return true
		}
	}

	return false
}

func (a *DNAAnalyzer) CanInheritFromBothParents(childAlleles, motherAlleles, fatherAlleles map[string]bool) bool {
	var childList []string
	for allele := range childAlleles {
		childList = append(childList, allele)
	}

	if len(childList) == 1 {
		return motherAlleles[childList[0]] && fatherAlleles[childList[0]]
	}

	if len(childList) >= 2 {
		allele1 := childList[0]
		allele2 := childList[1]

		return (motherAlleles[allele1] && fatherAlleles[allele2]) ||
			(motherAlleles[allele2] && fatherAlleles[allele1])
	}

	return false
}

// CanInheritFromParent checks if child can inherit from a specific parent
func (a *DNAAnalyzer) CanInheritFromParent(childAlleles, parentAlleles map[string]bool) bool {
	for allele := range childAlleles {
		if parentAlleles[allele] {
			return true
		}
	}
	return false
}

// AnalyzeInheritance determines inheritance pattern for a marker
func (a *DNAAnalyzer) AnalyzeInheritance(child, mother, father string) InheritanceResult {
	childAlleles := a.ExtractAlleles(child)
	motherAlleles := a.ExtractAlleles(mother)
	fatherAlleles := a.ExtractAlleles(father)

	if len(childAlleles) == 0 || len(motherAlleles) == 0 || len(fatherAlleles) == 0 {
		return None
	}

	canInheritFromBoth := a.CanInheritFromBothParents(childAlleles, motherAlleles, fatherAlleles)
	canInheritFromMother := a.CanInheritFromParent(childAlleles, motherAlleles)
	canInheritFromFather := a.CanInheritFromParent(childAlleles, fatherAlleles)

	if canInheritFromBoth {
		return Both
	} else if canInheritFromMother && !canInheritFromFather {
		return Mother
	} else if !canInheritFromMother && canInheritFromFather {
		return Father
	} else if canInheritFromMother && canInheritFromFather {
		return Neutral
	}

	return None
}

// AnalyzeThreeHorses analyzes child, father, and mother
func (a *DNAAnalyzer) AnalyzeThreeHorses(child, father, mother *Horse, markers []string) []AnalysisResult {
	results := make([]AnalysisResult, 0, len(markers))

	for _, marker := range markers {
		childValue := child.DNA[marker]
		fatherValue := father.DNA[marker]
		motherValue := mother.DNA[marker]

		childAlleles := a.ExtractAlleles(childValue)
		motherAlleles := a.ExtractAlleles(motherValue)
		fatherAlleles := a.ExtractAlleles(fatherValue)

		result := AnalysisResult{
			Marker:          marker,
			ChildValue:      childValue,
			FatherValue:     fatherValue,
			MotherValue:     motherValue,
			Result:          None,
			FatherCellColor: Transparent,
			MotherCellColor: Transparent,
		}

		if len(childAlleles) == 0 || len(motherAlleles) == 0 || len(fatherAlleles) == 0 {
			results = append(results, result)
			continue
		}

		canInheritFromBoth := a.CanInheritFromBothParents(childAlleles, motherAlleles, fatherAlleles)

		if canInheritFromBoth {
			result.Result = Both
		} else {
			canInheritFromMother := a.CanInheritFromParent(childAlleles, motherAlleles)
			canInheritFromFather := a.CanInheritFromParent(childAlleles, fatherAlleles)

			if canInheritFromMother && !canInheritFromFather {
				result.Result = Mother
				result.FatherCellColor = Red // Mark father red
			} else if !canInheritFromMother && canInheritFromFather {
				result.Result = Father
				result.MotherCellColor = Red // Mark mother red
			} else if canInheritFromMother && canInheritFromFather {
				result.Result = Neutral
				if a.NeutralRow == 1 {
					result.FatherCellColor = Red // Mother wins, mark father red
				} else if a.NeutralRow == 2 {
					result.MotherCellColor = Red // Father wins, mark mother red
				}
			} else {
				result.Result = None
				result.FatherCellColor = Red // Mark both red
				result.MotherCellColor = Red
			}
		}

		results = append(results, result)
	}

	return results
}

// AnalyzeTwoHorses analyzes child and single parent
func (a *DNAAnalyzer) AnalyzeTwoHorses(child, parent *Horse, markers []string) []AnalysisResult {
	results := make([]AnalysisResult, 0, len(markers))

	for _, marker := range markers {
		childValue := child.DNA[marker]
		parentValue := parent.DNA[marker]

		result := AnalysisResult{
			Marker:          marker,
			ChildValue:      childValue,
			FatherValue:     parentValue,
			Result:          None,
			FatherCellColor: Transparent,
		}

		if !a.HasCommonAllele(childValue, parentValue) {
			result.FatherCellColor = Red
		}

		results = append(results, result)
	}

	return results
}

// PrintResults prints the analysis results in a table format
func PrintResults(results []AnalysisResult, threeHorses bool) {
	fmt.Println("\n" + strings.Repeat("=", 100))
	if threeHorses {
		fmt.Printf("%-15s %-20s %-20s %-20s %-15s\n", "Marker", "Child", "Father", "Mother", "Result")
	} else {
		fmt.Printf("%-15s %-20s %-20s %-15s\n", "Marker", "Child", "Parent", "Match")
	}
	fmt.Println(strings.Repeat("=", 100))

	for _, r := range results {
		if threeHorses {
			fatherMark := ""
			motherMark := ""
			if r.FatherCellColor == Red {
				fatherMark = " [RED]"
			}
			if r.MotherCellColor == Red {
				motherMark = " [RED]"
			}

			fmt.Printf("%-15s %-20s %-20s %-20s %-15s\n",
				r.Marker,
				r.ChildValue,
				r.FatherValue+fatherMark,
				r.MotherValue+motherMark,
				r.Result.String())
		} else {
			parentMark := ""
			if r.FatherCellColor == Red {
				parentMark = " [NO MATCH]"
			} else {
				parentMark = " [MATCH]"
			}

			fmt.Printf("%-15s %-20s %-20s %-15s\n",
				r.Marker,
				r.ChildValue,
				r.FatherValue+parentMark,
				parentMark)
		}
	}
	fmt.Println(strings.Repeat("=", 100))
}

func main() {
	// Define markers (equivalent to columns in your WPF app)
	markers := []string{
		"VHL20", "HTG4", "AHT4", "HMS7", "HTG6", "AHT5", "HMS6",
		"ASB23", "ASB2", "HTG10", "HTG7", "HMS3", "HMS2",
		"ASB17", "LEX3", "HMS1", "CA425",
	}

	// Example 1: Three horses (child, father, mother)
	child := &Horse{
		Name: "Foal",
		DNA: map[string]string{
			"VHL20": "AA", "HTG4": "AA",
			"AHT4": "AA", "HMS7": "AA",
			"HTG6": "AA", "AHT5": "AA",
			"HMS6": "AA", "ASB23": "AA",
			"ASB2": "AA", "HTG10": "AA",
			"HTG7": "AA", "HMS3": "AA",
			"HMS2": "AA", "ASB17": "AA",
			"LEX3": "AA", "HMS1": "AA",
			"CA425": "AA",
		},
	}

	father := &Horse{
		Name: "Stallion",
		DNA: map[string]string{
			"VHL20": "AA", "HTG4": "AA",
			"AHT4": "AA", "HMS7": "AA",
			"HTG6": "AA", "AHT5": "AA",
			"HMS6": "AA", "ASB23": "AA",
			"ASB2": "AA", "HTG10": "AA",
			"HTG7": "AA", "HMS3": "AA",
			"HMS2": "AA", "ASB17": "AA",
			"LEX3": "AA", "HMS1": "AA",
			"CA425": "AA",
		},
	}

	mother := &Horse{
		Name: "Mare",
		DNA: map[string]string{
			"VHL20": "AA", "HTG4": "AA",
			"AHT4": "AA", "HMS7": "AA",
			"HTG6": "AA", "AHT5": "AA",
			"HMS6": "AA", "ASB23": "AA",
			"ASB2": "AA", "HTG10": "AA",
			"HTG7": "AA", "HMS3": "AA",
			"HMS2": "AA", "ASB17": "AA",
			"LEX3": "AA", "HMS1": "AA",
			"CA425": "BB",
		},
	}

	analyzer := NewDNAAnalyzer()

	fmt.Println("=== Three Horse Analysis (Child, Father, Mother) ===")
	results := analyzer.AnalyzeThreeHorses(child, father, mother, markers)
	PrintResults(results, true)

	/*	fmt.Println("\n=== Two Horse Analysis (Child, Parent) ===")
		results2 := analyzer.AnalyzeTwoHorses(child, father, markers)
		PrintResults(results2, false)*/
}
