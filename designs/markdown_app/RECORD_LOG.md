# RECORD_LOG

## Requirements

- Markdown based record log
  - Parse table of contents
  - Allow addition, editing and removal of entries
  - Structure extracted text in JSON & CSV
  - Allow tag based search and full text search
  - Allow exporting/importing logs to/from JSON and CSV files

## Design

### High Level Design

- Parse a markdown written in predefined format
- Save to MongoDB
- Enable CRUD

### Low Level Design

- In a common database in MongoDB, create a collection for each document, e.g. books_subitems
- Create a master collection that contains a list of all documents' headings and their table of contents
- subitems contains subitem lists as array of objects
- Write CRUD functions that edit the original markdown file as well the data in DB
- Expose CRUD through APIs
- Create an AST of the markdown file, save to MongoDB: maybe
  
## Use Cases

- Read log
  - Should allow entry of books currently being read, wishlist, books read
  - Configurable alarm, and whatsapp message to record page number of books currently being read, throw reminders for spaced repetitions of read books
  - Should remember day wise data entered
  - Should provide API to fetch all data

- Link recorder
  - TBD

- Todo list
  - TBD

- Course recorder
  - TBD
