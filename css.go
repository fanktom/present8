package main

// Base style
const css = `
/* minimal style */
body, h1, h2, h3, h4, h5, p {
	margin: 0;
	padding: 0;
}

body {
	background-color: #f6f8fa;
}

section.slide {
	position: relative;
	min-height: 100vh;
	display: flex;
	flex-direction: column;
	justify-content: center;
	background-color: white;
	margin-bottom: 20px;
}

@media print {
	section.slide {
		page-break-before: always;
	}
}

section.slide > div.padding {
	margin: 2em;
}

section.slide > div.pager {
	position: absolute;
	top: 0;
	right: 0;
	margin: 1em;
	color: rgba(0,0,0,0.4);
	font-size: 85%;
}

/* boxes can be used to put elements, such as logos into slides */
section.slide > div.box1,
section.slide > div.box2,
section.slide > div.box3,
section.slide > div.box4,
section.slide > div.box5,
section.slide > div.box6,
section.slide > div.box7,
section.slide > div.box8 {
	position: absolute;
	top: 0;
	right: 0;
	margin: 1em;
}

section.slide div.row {
	display: flex;
	justify-content: space-between;
}

section.slide div.column {
	width: 100%;
	padding-right: 1em;
}

section.slide div.column:last-child {
	padding-right: 0;
}

/* base style */
body {
	font-family: sans-serif;
	font-size: 1.5em;
	line-height: 1.5;
}

h1, h2, table, div.row, pre {
	margin-bottom: 1em;
}

h3, h4, h5, p {
	margin-bottom: 0.5em;
}

table {
	width: 100%;
	border-collapse: collapse;
}

table th {
	text-align: left;
	font-size: 85%;
}

table td, table th {
	padding: 0.4em 0.8em;
}

table tr:nth-child(even) {
	background-color: #f6f8fa;
}

img.center {
	display: block;
	margin: 0 auto;
}

code {
	padding: 0.4em;
	padding-top: 0.2em;
	padding-bottom: 0.2em;
	margin: 0;
	font-size: 85%;
	background-color: rgba(27,31,35,0.05);
	border-radius: 3px;	
}

pre {
	padding: 16px;
	overflow: auto;
	font-size: 85%;
	line-height: 1.45;
	background-color: #f6f8fa;
	border-radius: 3px;
}

pre > code {
	background: none;
}
`
