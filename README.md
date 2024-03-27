# Simple Web Site Approach using Golang, markdown, tailwind, Alpine js and Templ

Welcome to our GitHub repository, where we're bringing simplicity and efficiency back to web development. Our journey began with a realization: creating simple web pages shouldn't require complex solutions. This repository is our exploration into making web development straightforward, focusing on static sites with minimal updates.

Check out the demo [here].(https://magicms-medium.hakonleinan.com/)
## Why Not WordPress?

For many, WordPress is the go-to solution for creating websites with a handful of pages and posts. However, it's like having a swimming pool guard at the Olympics: unnecessary for the skilled and an overkill for the simple needs. Our project seeks to avoid the overhead of databases for content that rarely changes, optimizing both development and page load times.

## Introduction to Our Approach

Inspired by the elegance of Nuxt and the simplicity of Vue, we found a way to streamline web development without sacrificing performance. Our goal was to develop a web application that is not only quick to develop but also exceptionally fast for the end-user.

### The Turning Point

While experimenting with a Nuxt-driven site, the resource consumption for a low-traffic scenario was notably high. This observation led us to rethink our approach and explore alternatives that maintain simplicity without the resource overhead.

### Our Solution: A Golang Web App

We developed a lightweight web app in Golang, focusing on the following features:

- **Content-driven:** Dynamically serves content based on the URI, checking for corresponding .md or .html files or folders.
- **Integrated Templating and Markdown Rendering:** Utilizes Go's `http.ServeMux`, a powerful template engine, and Goldmark for markdown rendering.
- **Performance:** Achieves a Google Lighthouse score of 100, ensuring top-tier user experience.

### Why Golang?

Choosing Golang for this project was strategic. Its performance in terms of resource management is unparalleled, especially when compared to more traditional choices for web development in similar scenarios. Our app demonstrates that even for simple sites, efficiency doesn't have to be compromised.

## Features

- **Automatic Routing:** Based on file and folder names within the content directory.
- **Markdown and HTML Support:** Flexibility in content creation. Using [goldmark markdown](https://github.com/yuin/goldmark)
- **High Performance:** Optimized for speed, achieving perfect performance scores.
- - **Templ engine:** for a more natural way of writing markup. [Templ](https://github.com/a-h/templ)
- **Tailwind CSS and Alpine.js:** For a modern, responsive design without the bloat.

## Getting Started

To set up this project locally and start contributing, follow these steps:

1. **Clone the repository** to your local machine.

    ```bash
    git clone https://github.com/mediamagi/magicms-medium
    ```

2. **Install Golang** if it's not already installed on your system.

3. **Environment Setup:**

    - Copy `example.env` to `.env` in the project root. This file contains necessary environment variables for running the application.

        ```bash
        cp example.env .env
        ```

4. **Compile and run the Application:**

    - Navigate to the project directory and run `go run main.no` to compile and run the application.

5. **View the site:**

    - Navigate to `localhost:3000` in your web browser to view the site.

6. **Working with CSS:**

    - To automatically compile your Tailwind CSS, use the following command:

        ```bash
        npx tailwindcss -i ./src/input.css -o ./static/css/styles.css --watch
        ```

   This command watches for changes in your CSS and automatically recompiles the Tailwind output file.

7. **Generating Templates:**

    - For dynamic template generation, use the `templ generate --watch` command. This will monitor your template files for changes and automatically regenerate them as needed.

        ```bash
        templ generate --watch
        ```

# Adding Meta Title and Description in the Markdown File

To add metadata such as a title or description to your Markdown files, you can use the YAML front matter block at the beginning of the file. Here's an example:

```yaml
---
title: This is the title
description: This is the description
---
```

You can then access these meta fields in your `main.go` file using the following syntax:

```go
metaData["{field_name}"].(string)
```

## Handling Related Content in Markdown Files

Our platform supports associating related content with Markdown documents through the use of a `relation` field within the document's front matter. This allows for the seamless integration of additional resources, enhancing the richness and utility of your content.

### Specifying Related Content

To specify related content, include the `relation` field in the YAML front matter of your Markdown document. The related content must be located within the same directory as the referring document. Here’s how you can specify a related HTML file:

```yaml
---
relation: _additionalInfo.html
---
```

### Naming Conventions and Access

- Files named with a leading underscore (`_`) are treated specially. They are meant to be included as part of another document and are not directly accessible via a URL. This convention helps in organizing your content and keeping auxiliary files hidden.
- If you wish for the related file to be accessible as standalone content, do not prefix the filename with `_`. This makes the file directly accessible through its URL, allowing it to serve as both related content and a standalone resource.

### Examples

#### Including an HTML File

For Markdown documents that require supplementary HTML content to be displayed alongside the Markdown-rendered content, specify the HTML file using the `relation` field. The content of this file will be rendered in place when the Markdown document is accessed. Remember, if the filename starts with `_`, the content is intended for inclusion and not direct access:

```yaml
---
relation: _contactForm.html
---
```

#### Making Related Content Accessible

If you have a related document that should also be accessible on its own, simply ensure it does not begin with `_`. This is useful for supplementary guides, forms, or any content that serves a dual purpose—both as standalone and as related content.

```yaml
---
relation: termsOfService.html
---
```
