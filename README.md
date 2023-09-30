# Go Echo HTMX Kickstart project

## Overview

This is my kickstart project for building web applications. The stack is minimalist and includes the following:

- **Embedding**: Embed static folder and templates into the binary.
- **HTMX**: Enables dynamic content updates without requiring a full page reload.
- **Hyperscript**: Simplifies frontend interactivity with syntactic sugar.
- **Go Templates Rendering**: HTML rendering with Go's standard library.
- **Tailwind CSS**: Integrated for utility-first CSS styling (because i'm too lazy to write CSS).
- **Air for Hot Reloading**: Enables a smooth development experience by automatically reloading the application on code changes.
- **Sanitize HTML**: Sanitize HTML inputs with the [bluemonday](https://github.com/microcosm-cc/bluemonday) package.
  
The project also aims to maintain a structure that aligns with Golang's standard project layout.

## Getting Started

### Prerequisites

- **Install Air for Hot Reloading.**
    Follow the [Air installation guide](https://github.com/cosmtrek/air).
- **Install Bun for Tailwind CSS.**
    Follow the [Bun installation guide](https://bun.sh/docs/installation).

### Installation

1. **Clone the repository.**
    ```
    git clone https://github.com/bnema/kickstart-echo.git .
    ```
3. **Remove the .git folder and initialize a new one.**
    ```
    rm -rf .git && git init
    ```
4. **Remove go.mod, go.sum and create a new project.**
    ```
    rm go.mod go.sum && go mod init github.com/username/project
    ```
2. **Install dependencies.**
    ```
    go mod tidy
    ```
7. **Run Bun Install.**
    ```
    bun install
    ```
8. **Execute Tailwind with Bun.**
    ```
    bun run dev:css
    ```
9. **Create a .env at the root of the project (see .env.example).**
10. **Run the Application (Air will rebuild and restart the application on code changes).**
    ```
    air
    ```
11. **If the 2 previous commands are successful, you can run dev.sh who will run them in parallel.**
    ```
    chmod +x dev.sh
    ./dev.sh
    ```
12. **Bonus: The superlazy, one-liner command.**
    ```
    export MODULE_NAME=github.com/CHANGE_ME/I_MEAN_SERIOUSLY; git clone http://github.com/bnema/kickstart-echo.git . && go mod tidy && rm -rf .git && git init && rm go.mod go.sum && go mod init $MODULE_NAME && bun install && cp .env.example .env && chmod +x dev.sh && ./dev.sh
    ```

## License
Under the GPL-3.0 license. Please see the LICENSE file for more details.
