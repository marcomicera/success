package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sort"
	"sync"
)

type Applicant struct {
	skills int
	luck   int
	total  float32
}

// Applicant constructor
func NewApplicant(r *rand.Rand) Applicant {
	applicant := Applicant{
		skills: r.Intn(100),
		luck:   r.Intn(100),
	}
	applicant.total = float32(applicant.skills)*0.95 + float32(applicant.luck)*0.05
	return applicant
}

// Simulates one selection process
func selectionProcess(seed int64, applicantsNum int, spaces int) float32 {

	r := rand.New(rand.NewSource(seed))

	var applicants []Applicant

	// Score generation
	for i := 0; i < applicantsNum; i++ {
		applicants = append(applicants, NewApplicant(r))
	}

	// Selection process
	sort.Slice(applicants, func(i, j int) bool {
		return applicants[i].total > applicants[j].total
	})

	// Computing result
	var avgLuck float32 = 0.0
	for _, topApplicant := range applicants[:spaces] {
		avgLuck += float32(topApplicant.luck)
	}
	avgLuck /= float32(spaces)
	return avgLuck
}

func main() {

	// Params
	seed := flag.Int64("seed", 0, "Simulation seed")
	applicantsNum := flag.Int("applicants", 18300, "Number of applicants")
	spaces := flag.Int("spaces", 11, "Number of applicants that will be selected")
	numSimulations := flag.Int("simulations", 1000, "Number of simulations to run")
	flag.Parse()

	// Number of simulations check
	if *numSimulations < 1 {
		log.Fatalf("Invalid number of simulations to be run: %d. Terminating...", *spaces)
	}

	// Applicants number check
	if *applicantsNum <= 1 {
		log.Fatalf("Invalid number of applicants: %d. Terminating...", *applicantsNum)
	}

	// Number of spaces check
	if *spaces >= *applicantsNum {
		log.Fatalf("%d spaces for %d applicants. Terminating...", *spaces, *applicantsNum)
	}
	if *spaces <= 1 {
		log.Fatalf("Invalid spaces number: %d. Terminating...", *spaces)
	}

	// Final result: luck average across simulations
	avgLuckAcrossSimulations := float32(0)

	// Parallel simulations setup
	c := make(chan int64)
	var wg sync.WaitGroup
	wg.Add(*numSimulations)
	for ii := 0; ii < *numSimulations; ii++ {
		go func(c chan int64) {
			for {
				// Retrieving current seed
				currentSeed, more := <-c
				if more == false {
					wg.Done()
					return
				}

				// Starting a selection process
				avgLuck := selectionProcess(currentSeed, *applicantsNum, *spaces)
				//fmt.Printf("Average luck score of the top %d applicants: %f\n", spaces, avgLuck) // one sim. result
				avgLuckAcrossSimulations += avgLuck
			}
		}(c)
	}

	// Launching simulations
	for simSeed := *seed; simSeed < *seed+int64(*numSimulations); simSeed++ {
		c <- simSeed
	}
	close(c)
	wg.Wait()
	avgLuckAcrossSimulations /= float32(*numSimulations)
	fmt.Printf("The top %d applicants (out of %d) were %.2f%% lucky on average across %d simulations.\n",
		*spaces, *applicantsNum, avgLuckAcrossSimulations, *numSimulations)
}
