package static

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func ExampleHeadingAnchors() {
	r := ioutil.NopCloser(strings.NewReader(`
		<h1>Article 1</h1>
			<h2>Section 1</h2>
			<h2>Section 2</h2>
			<h2>Section 3</h2>
			  <h3>Sub Section 1</h3>
			    <h4>Sub Sub Section 1</h4>
		<h1>Article 2</h1>
		<h1>Article 3</h1>
		  <h2>Section 1</h2>
	`))

	r = HeadingAnchors(r)
	io.Copy(os.Stdout, r)
	// Output:
	// <html><head></head><body><h1 class="Heading"><a class="Anchor" aria-hidden="true" id="article_1" href="#article_1">
	// 	<svg xmlns="http://www.w3.org/2000/svg" aria-hidden="true" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-link"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"></path><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"></path></svg>
	// </a>Article 1</h1>
	// 			<h2 class="Heading"><a class="Anchor" aria-hidden="true" id="article_1.section_1" href="#article_1.section_1">
	// 	<svg xmlns="http://www.w3.org/2000/svg" aria-hidden="true" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-link"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"></path><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"></path></svg>
	// </a>Section 1</h2>
	// 			<h2 class="Heading"><a class="Anchor" aria-hidden="true" id="article_1.section_2" href="#article_1.section_2">
	// 	<svg xmlns="http://www.w3.org/2000/svg" aria-hidden="true" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-link"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"></path><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"></path></svg>
	// </a>Section 2</h2>
	// 			<h2 class="Heading"><a class="Anchor" aria-hidden="true" id="article_1.section_3" href="#article_1.section_3">
	// 	<svg xmlns="http://www.w3.org/2000/svg" aria-hidden="true" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-link"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"></path><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"></path></svg>
	// </a>Section 3</h2>
	// 			  <h3 class="Heading"><a class="Anchor" aria-hidden="true" id="article_1.section_3.sub_section_1" href="#article_1.section_3.sub_section_1">
	// 	<svg xmlns="http://www.w3.org/2000/svg" aria-hidden="true" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-link"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"></path><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"></path></svg>
	// </a>Sub Section 1</h3>
	// 			    <h4 class="Heading"><a class="Anchor" aria-hidden="true" id="article_1.section_3.sub_section_1.sub_sub_section_1" href="#article_1.section_3.sub_section_1.sub_sub_section_1">
	// 	<svg xmlns="http://www.w3.org/2000/svg" aria-hidden="true" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-link"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"></path><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"></path></svg>
	// </a>Sub Sub Section 1</h4>
	// 		<h1 class="Heading"><a class="Anchor" aria-hidden="true" id="article_2" href="#article_2">
	// 	<svg xmlns="http://www.w3.org/2000/svg" aria-hidden="true" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-link"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"></path><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"></path></svg>
	// </a>Article 2</h1>
	// 		<h1 class="Heading"><a class="Anchor" aria-hidden="true" id="article_3" href="#article_3">
	// 	<svg xmlns="http://www.w3.org/2000/svg" aria-hidden="true" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-link"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"></path><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"></path></svg>
	// </a>Article 3</h1>
	// 		  <h2 class="Heading"><a class="Anchor" aria-hidden="true" id="article_3.section_1" href="#article_3.section_1">
	// 	<svg xmlns="http://www.w3.org/2000/svg" aria-hidden="true" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-link"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"></path><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"></path></svg>
	// </a>Section 1</h2>
	// 	</body></html>
}

func ExampleNotes() {
	r := ioutil.NopCloser(strings.NewReader(`
Something here.

Note: Some note here about whatever nothing exciting.

More content.
	`))

	r = Notes(Markdown(r))
	io.Copy(os.Stdout, r)
	// Output:
	// <p>Something here.</p>
	//
	// <div class="Message"><p>Some note here about whatever nothing exciting.</p></div>
	//
	// <p>More content.</p>
}
