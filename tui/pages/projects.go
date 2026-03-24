package pages

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"portfolio/tui/router"
	"portfolio/tui/theme"
)

// ── Edit your projects here (ordered most recent first) ───────────────────────

var projects = []ProjectData{
	{
		Name: "Dropbox — Data Science Intern",
		Desc: "Customer success analytics for Dash product activation, 1B+ events",
		Links: []string{},
		Tech: "Databricks, MySQL, Python, Pandas, NumPy, SciPy, Logistic Regression, SQL",
		Body: `Data Science Intern - Customer Success Team
May 2025 - August 2025

Led a data analytics project to optimise onboarding for Dropbox's Dash product, defining scope and success metrics with stakeholders.

Built SQL pipelines in Databricks to classify and aggregate 1B+ user events across features, surfacing activation signals for the Customer Success team.

Created a dashboard for CSMs and delivered actionable insights to drive targeted customer outreach and improve onboarding outcomes.

Built a propensity model using logistic regression, validated with t-tests and permutation testing to ensure statistical robustness.`,
		Quote: `I had the pleasure of managing Adam on a project where he created product activation insights that support customer-facing teams. He took initiative in stakeholder management, reaching beyond the initial working group to include additional perspectives, and shared insights in a way that made them actionable for both customer-facing teams and customers.

Adam kept the project on track by consistently checking progress against milestones and timelines, while also building his technical skills quickly. He learned new tools such as Databricks and made effective use of visualization and automation features to validate solutions more efficiently.

What stood out most was his customer focus, grounding his work in real pain points uncovered through documentation and conversations with frontline teams, which made his contributions both practical and impactful. Adam brings a thoughtful balance of initiative, technical skill, and collaboration, and would be a valuable addition to any team.`,
		QuoteBy: "Daniil Voloshin, AI Data Analyst, Dropbox  (managed Adam directly, Sept 2025)",
	},
	{
		Name: "MetaTune — HackEurope",
		Desc: "AI fine-tuning platform cutting training jobs by up to 90%",
		Links: []string{
			"github.com/TitanSmash/hackeurope-MetaTune",
			"youtube.com/watch?v=LZW4yHROhzA",
			"devpost.com/software/metatune",
			"arxiv.org/abs/2205.13320",
			"proceedings.mlr.press/v188/ram22a/ram22a.pdf",
		},
		Tech: "Python, PyTorch, FastAPI, Bayesian Optimisation, Linear Programming, OpenShift, Kubernetes, Docker, TypeScript",
		Body: `Built for HackEurope in the sustainability track. MetaTune is an AI fine-tuning platform that aims to reduce the compute needed in trial-and-error hyperparameter search. It uses Bayesian Optimisation with Gaussian Processes to predict optimal configurations from historical runs, cutting required training jobs by up to 90% -- drastically reducing the carbon footprint of fine-tuning a model. Based on recent papers from Google and IBM.`,
	},
	{
		Name: "Big Data — Rented Bikes",
		Desc: "ETL pipeline and MapReduce benchmark on 14.8M bike trip records",
		Links: []string{
			"github.com/adam505x/Cyclistic_Data",
		},
		Tech: "Python, PySpark, Pandas, SQLite, MapReduce, SQL",
		Body: `Analysed 14.8 million bike trip records from Chicago's Divvy system, joined with historical weather data.

Built a full ETL pipeline handling schema normalisation across 34 CSV files, then benchmarked traditional single-threaded approaches (Pandas, SQLite) against Apache Spark MapReduce implementations across two high-cardinality aggregation queries -- to show when MapReduce is most optimally utilised.`,
	},
	{
		Name: "Crease Defenders — E-Commerce",
		Desc: "E-commerce brand -- 2,000+ customers across 37 countries (2020-2025)",
		Links: []string{},
		Tech: "",
		Body: `Crease Defenders was my longest running e-commerce store, which I ran from 2020-2025. I oversaw product development and imports from China, Poland, and the UK.

· 20,000,000+ impressions on organic content
· ~20 influencer campaigns
· 2,000+ customers across 37 countries
· 3,000+ email subscribers

## Strategy

## UI/UX Design
I spent a lot of time studying successful stores and replicating their winning designs. On my stores with ~1-3 products I focused on reducing the amount of clicks from first visit to order. My conversion rate was usually ~5%, which is considered very high for one-product stores.

## Marketing
I utilised micro-influencers (1k-15k followers) to generate UGC. I would send my product to them along with pointers on how to record a video. By letting them have higher creative control I found the results to be much more organic -- people can sense when ads seem fake. I then used this influencer content to create my own UGC and ads which generated more sales.

I also used Klaviyo for email marketing flows to target lost checkouts and repeat customers.

## Product Design
I shipped Crease Defenders in flexible plastic pouches so shipping costs per package were treated as envelopes. I also included a custom thank you card with a discount for their next purchase. This kept my gross margins at 80% and recurring customer rate at 15%.

## Python Scripts
· Exported CSVs of sales and customers to analyse trends, recurring customers, and projections
· Exported Facebook ad campaign data to calculate rank metrics between campaigns, focusing on improving CPA and CPC

## Other Ventures
During 2020-2025 I also ran ~7 other short-term e-commerce stores where I shipped products from 3PLs and hired agents in China to oversee product shipments and development. Some stores made money, some stores lost money.

There's a lot more fun scrappy tips and tricks I used during my e-commerce ventures -- maybe I'll write a few blog posts on this.`,
	},
	{
		Name: "Wordle Data Analysis",
		Desc: "Wordle data analysis and a bot with a better win rate than the average human",
		Links: []string{
			"github.com/adam505x/wordle_data_analysis",
			"adam505x.github.io/wordle_data_analysis/design.html",
			"youtube.com/watch?v=2SbWt6uUMHc",
		},
		Tech: "Python, SQL, Firebase",
		Body: `My first data science project. I learnt many essential skills in data analysis, designing a project to meet a brief, and uploading and retrieving data from cloud databases. Using my analysis I was able to create a bot to play Wordle with a much better win rate than humans.

For more information check out the project website, where I go over my design choices, planning, and testing. Or watch the video.`,
	},
	{
		Name: "London Accommodation",
		Desc: "Housing map for my Google internship, shared with other London interns",
		Links: []string{
			"github.com/adam505x/Google_London_Accomodation",
			"google-lon-accom-production.up.railway.app",
		},
		Tech: "HTML, CSS, JavaScript, Node.js, Express, CORS",
		Body: `I built this to keep track of all my options on where to stay in London for my Google internship. I used OpenStreetMap for the map, scraped Google Maps for listings, and Railway for hosting. Mainly vibe coded this to get it done quickly. I shared it in the WhatsApp group chat with the other London interns who found it useful.`,
	},
	{
		Name: "FIDE Chess Dashboard",
		Desc: "Interactive dashboard analysing FIDE chess ranking distributions",
		Links: []string{
			"github.com/adam505x/FIDE_interactive_dashboard",
		},
		Tech: "Python, VegaLite",
		Body: `Used VegaLite to make an interactive dashboard analysing FIDE Chess rankings distributions. Completed for a module at university.`,
	},
}

// ── ProjectData is what you fill in per project ───────────────────────────────

type ProjectData struct {
	Name    string   // short name, shown in the list
	Desc    string   // one-liner shown in the list
	Links   []string // github, live site, video, papers, etc.
	Tech    string   // tech stack
	Body    string   // full detail; prefix a line with "## " for a section heading
	Quote   string   // optional reference/testimonial text
	QuoteBy string   // attribution for the quote
}

// ─────────────────────────────────────────────────────────────────────────────

type ProjectsPage struct {
	cursor int
}

func (p ProjectsPage) Update(key string) router.Action {
	switch key {
	case "esc":
		return router.Pop{}
	case "enter":
		return router.Push{Page: NewProjectPage(projects[p.cursor])}
	case "down":
		if p.cursor < len(projects)-1 {
			p.cursor++
		}
		return router.Stay{Page: p}
	case "up":
		if p.cursor > 0 {
			p.cursor--
		}
		return router.Stay{Page: p}
	}
	return router.Stay{Page: p}
}

func (p ProjectsPage) View() string {
	// ── Left column: project list ──────────────────────────────────────────────
	var leftLines []string
	push := func(s string) { leftLines = append(leftLines, s) }

	push(theme.Heading.Render("PROJECTS"))
	push(theme.Dim.Render(strings.Repeat("-", 44)))
	push("")

	for i, proj := range projects {
		cursor := "  "
		if i == p.cursor {
			cursor = "> "
		}
		push(theme.Accent.Render(cursor + proj.Name))
		push(theme.Base.Render("    " + proj.Desc))
		push("")
	}

	push("")
	push(theme.Dim.Render("up/down  move   enter  open   esc  back"))

	// ── Right column: stars art ────────────────────────────────────────────────
	var rightLines []string
	if StarsArt != "" {
		rightLines = strings.Split(strings.TrimRight(StarsArt, "\n"), "\n")
	}

	maxLeft := 0
	for _, l := range leftLines {
		if w := lipgloss.Width(l); w > maxLeft {
			maxLeft = w
		}
	}

	const gap = "     "
	n := len(leftLines)
	if len(rightLines) > n {
		n = len(rightLines)
	}

	var b strings.Builder
	for i := 0; i < n; i++ {
		var left string
		if i < len(leftLines) {
			left = leftLines[i]
		}
		pad := strings.Repeat(" ", maxLeft-lipgloss.Width(left))

		var right string
		if i < len(rightLines) {
			right = rightLines[i]
		}
		b.WriteString(left + pad + gap + right + "\n")
	}

	return theme.Page.Render(b.String())
}
