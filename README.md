# Api-bible

![Build Status](https://travis-ci.org/alexfernandessd/api-bible.svg?branch=master)

This respository contains a simple bible api with `gets` only as the escriptures cannot be change.

## Build & Start

- To build the application you just need run the command: 
```sh
make build
```
- After build you can run the following command to start the application:
```sh
make run
```

## URLs for services

### Get /testament/{testament}

You can pass `old` or `new` on url param.

- Response 200:
```js
{
  books: [
    "book1",
    "book2",
  ],
}
```

### Get /book/{book}

You only need pass the book name on url param.

- Response 200:
```js
{}
```

### Get /book/{book}/chapter/{chapter}/verses

You need pass the book name and the chapter on url param.

- Response 200:
```js
{
  verses: [
    {
      "ID": 1,
      "Version": "aa",
      "Testament": 1,
      "Book": 0,
      "Chapter": 1,
      "Verse": 1,
      "Text": "No princípio criou Deus os céus e a terra."
    },
    {
      "ID": 2,
      "Version": "aa",
      "Testament": 1,
      "Book": 0,
      "Chapter": 1,
      "Verse": 2,
      "Text": "A terra era sem forma e vazia; e havia trevas sobre a face do abismo, mas o Espírito de Deus pairava sobre a face das águas."
    },
    ...,
  ],
}
```

### Get /book/{book}/chapter/{chapter}/verse/{verse}

You need pass the book name, chapter and the verse on url param.

- Response 200:
```js
{
  "ID": 1,
  "Version": "aa",
  "Testament": 1,
  "Book": 0,
  "Chapter": 1,
  "Verse": 1,
  "Text": "No princípio criou Deus os céus e a terra."
}
```