package main

func main() {
	// - Create a master collection that contains a list of all documents' headings and their table of contents
	// - In a common database in MongoDB, create a collection for each document: e.g. books_subitems: connect to MongoDB Atlas
	// - headings and subitems contain table of contents and subitem lists as array of objects: parse markdown
	// - Write CRUD functions that edit the original markdown file as well the data in DB: allow editing of markdown file saved on GitHub through git logs
	// - Expose CRUD through APIs: write a server that exposes CRUDs
}
