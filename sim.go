package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sort"
)

type Applicant struct {
	skills int
	luck   int
	total  float32
}

// Applicant constructor
func NewApplicant() Applicant {
	applicant := Applicant{
		skills: rand.Intn(100),
		luck:   rand.Intn(100),
	}
	applicant.total = float32(applicant.skills)*0.95 + float32(applicant.luck)*0.05
	return applicant
}

func main() {

	// Params
	seed := flag.Int64("seed", 0, "Simulation seed")
	applicantsNum := flag.Int("applicants", 18300, "Number of applicants")
	spaces := flag.Int("spaces", 11, "Number of applicants that will be selected")
	flag.Parse()

	// There cannot be more spaces than applicants
	if *spaces >= *applicantsNum {
		log.Fatalf("%d spaces for %d applicants. Terminating...", *spaces, *applicantsNum)
	}
	if *spaces <= 1 {
		log.Fatalf("Invalid spaces number: %d. Terminating...", *spaces)
	}

	rand.Seed(*seed)

	var applicants []Applicant

	// Score generation
	for i := 0; i < *applicantsNum; i++ {
		applicants = append(applicants, NewApplicant())
	}

	// Selection process
	sort.Slice(applicants, func(i, j int) bool {
		return applicants[i].total > applicants[j].total
	})

	// Computing results
	var avgLuck float32 = 0.0
	for _, topApplicant := range applicants[:(*spaces)] {
		avgLuck += float32(topApplicant.luck)
	}
	avgLuck /= float32(*spaces)
	fmt.Printf("Average luck score of the top %d applicants: %f\n", *spaces, avgLuck)
}
