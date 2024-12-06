# Keygen Command

CLI tool to generate new JSON Web Tokens (JWT) to securely communicate with your services.

## Installation

### 1. Binary releases

To install `keygen` from binary releases, you can download the latest release from the [releases page](https://github.com/andrespd99/keygen/releases).

Then, place the binary in your PATH and you're good to go.

### 2. Using go install

To install `keygen` using go, run the following command:

```bash
go install github.com/andrespd99/keygen/cmd/keygen
```

## Usage

To generate a new JWT, run the following command:

```bash
keygen "{{SIGNING_SECRET_KEY}}"
```

This command will generate a new JWT and copy it to your clipboard.

### 3. Running directly from source

If you want to run the keygen command directly from source, you can do so by running the following command:

```bash
go run ./cmd/keygen/keygen.go "{{SIGNING_SECRET_KEY}}"
```

> Note: You will need Go 1.18 or higher to run this command.


## Setting the subject

You can set the subject of the generated JWT with the `-s` flag:

```bash
keygen -s "My custom subject" "{{SIGNING_SECRET_KEY}}"
```

The above command will generate a new JWT with the custom subject.

> NOTE: The -s flag must be passed before the signing secret key, else the command will fail.