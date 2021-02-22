# [Is Success Luck or Hard Work?](https://youtu.be/3LopI4YeC4I?t=218)

> In a competitive world, tiny advantages can make all the difference.

<p align="center">
    <a href="https://youtu.be/3LopI4YeC4I?t=218" alt="Veritasium's video link">
            <img src="https://img.youtube.com/vi/3LopI4YeC4I/0.jpg" />
    </a>
</p>

Simulation inspired by [_Is Success Luck or Hard Work?_](https://youtu.be/3LopI4YeC4I?t=218),
by [Veritasium](https://www.youtube.com/c/veritasium).

> When competition is fierce, being talented and hard-working is important, but it's not enough to guarantee success.
>
> You also need to catch a break.
>
> Largely, I think we're unaware of our good luck because, by definition, it's not something we did.\
> Like the housework done by your significant other, it goes unappreciated.
>
> And here's the crazy thing:
> downplaying the importance of chance events may actually improve your probability of success because
> if you perceive an outcome to be uncertain, you're less likely to invest effort in it,
> which further decreases your chances of success.
>
> So, it's a useful delusion to believe you are in full control of your destiny.

## Simulation description

<details>
<summary>Click to expand</summary>

> The importance of luck increases the greater the number of applicants applying for just a few spaces.
>
> Consider the most recent class of NASA astronauts.\
> From over 18,300 applicants in 2017,
> only 11 were selected and went on to graduate from the astronaut training program.
>
> Now we can make a toy model of the selection process:
> let's assume that astronauts are selected mostly based on skill,
> experience, and hard work, but also say five percent as a result of luck — fortunate circumstances.
> 
> For each applicant, I randomly generated a skill score out of a hundred,
> and I also randomly generated a luck score out of a hundred.\
> Then I added those numbers together, weighted in the 95-to-5 ratio to get an overall score.\
> This score represents the selector's judgments, meaning the top 11 by this metric would become astronauts.
>
> And I repeated this simulation a thousand times representing a thousand different astronaut selections.\
> And what I found was the astronauts who were picked were very lucky; they had an average luck score of 94.7.
>
> So how many of the selected astronauts would have been in the top 11 based on skill alone?\
> The answer was, on average, only 1.6.\
> That means, even with luck accounting for just 5% of the outcome, 9 or maybe 10 of the 11 applicants selected
> would have been different if luck played no role at all.
</details>

## Run the simulation

```shell script
go run sim.go
```

### Parameters

```
$ go run sim.go --help                                                                                                                                                                                                   10.38.55 PM, 21.02.2021
Usage:
  -applicants int
        Number of applicants (default 18300)
  -seed int
        Simulation seed (default 0)
  -simulations int
        Number of simulations to run (default 1000)
  -spaces int
        Number of applicants that will be selected (default 11)
```
