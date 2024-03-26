---
title: Simple Web Site Approach using Golang, markdown, tailwind, Alpine js and Templ
description: Welcome to our GitHub repository, where we're bringing simplicity and efficiency back to web development.
---


# Simple Web Site Approach using Golang, markdown, tailwind, Alpine js and Templ

Welcome to our GitHub repository, where we're bringing simplicity and efficiency back to web development. Our journey
began with a realization: creating simple web pages shouldn't require complex solutions. This repository is our
exploration into making web development straightforward, focusing on static sites with minimal updates.

## Why Not WordPress?

For many, WordPress is the go-to solution for creating websites with a handful of pages and posts. However, it's like
having a swimming pool guard at the Olympics: unnecessary for the skilled and an overkill for the simple needs. Our
project seeks to avoid the overhead of databases for content that rarely changes, optimizing both development and page
load times.

## Introduction to Our Approach

Inspired by the elegance of Nuxt and the simplicity of Vue, we found a way to streamline web development without
sacrificing performance. Our goal was to develop a web application that is not only quick to develop but also
exceptionally fast for the end-user.

### The Turning Point

While experimenting with a Nuxt-driven site, the resource consumption for a low-traffic scenario was notably high. This
observation led us to rethink our approach and explore alternatives that maintain simplicity without the resource
overhead.

### Our Solution: A Golang Web App

We developed a lightweight web app in Golang, focusing on the following features:

- **Content-driven:** Dynamically serves content based on the URI, checking for corresponding .md or .html files or
  folders.
- **Integrated Templating and Markdown Rendering:** Utilizes Go's `http.ServeMux`, a powerful template engine, and
  Goldmark for markdown rendering.
- **Performance:** Achieves a Google Lighthouse score of 100, ensuring top-tier user experience.

### Why Golang?

Choosing Golang for this project was strategic. Its performance in terms of resource management is unparalleled,
especially when compared to more traditional choices for web development in similar scenarios. Our app demonstrates that
even for simple sites, efficiency doesn't have to be compromised.

## Features

- **Automatic Routing:** Based on file and folder names within the content directory.
- **Markdown and HTML Support:** Flexibility in content creation.
  Using [goldmark markdown](https://github.com/yuin/goldmark).
- **High Performance:** Optimized for speed, achieving perfect performance scores.
- **Templ engine:** for a more natural way of writing markup. [Templ](https://github.com/a-h/templ).
- **Tailwind CSS and Alpine.js:** For a modern, responsive design without the bloat.

## Getting Started

To set up this project locally and start contributing, follow these steps:

1. **Clone the repository** to your local machine.

    ```bash
    git clone <repository-url>
    ```

2. **Install Golang** if it's not already installed on your system.

3. **Environment Setup:**

    - Copy `example.env` to `.env` in the project root. This file contains necessary environment variables for running
      the application.

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

    - For dynamic template generation, use the `templ generate --watch` command. This will monitor your template files
      for changes and automatically regenerate them as needed.

        ```bash
        templ generate --watch
        ```

## Contributing

We welcome contributions from the community! Whether it's adding new features, improving documentation, or reporting
issues, your input is valuable. Please refer to our CONTRIBUTING.md for guidelines on how to contribute effectively.

## Conclusion

Our journey has shown that simplicity in web development doesn't have to compromise on performance or scalability. By
focusing on what truly matters, we've created a project that exemplifies efficiency and ease of use. We hope you join us
in refining and expanding this approach.
