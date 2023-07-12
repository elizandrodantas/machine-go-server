<h1 align="center">
    🖨 Machine Go Server 🎛
</h1>

The main goal of machine-go-server is to provide a reliable system to manage machine access in freelance projects. By avoiding improper sharing, proper use of resources is ensured, increasing security and control over the machines involved.

In addition, the project aims to serve as a learning exercise and practical application of the knowledge acquired in Go. The responsible developer is studying the language and saw in machine-go-server an opportunity to create a real project, putting the learned concepts into practice. With this, he seeks to improve his skills in Go, as well as gain experience in the development of secure and scalable applications.

# 📰 Additional Information

- Machine visualization: machine-go-server offers all the necessary routes for administrators to be able to visualize requested or already authorized machines, allowing a complete view of the state of access.
- Machine validation: Machine validation is based on a unique identifier. If encryption is broken and access is attempted with the same ID on the same day, machine-go-server automatically disables the machine, allowing further analysis or permanent blocking.

# ✅ Requirements

- go1.19.3+
- postgres:13.5+

# 🛠 How to install?

To download Machine Go Server, you can clone the repository using the git clone command:

```bash
git clone https://github.com/elizandrodantas/machine-go-server.git
```

Navigate to the project directory:

```bash
cd machine-go-server/cmd/machine-go-server
```

Then complete the project using the go build command:

```go
go build
```

# ⚙ Tool

Before running Machine Go Server, you need to create the configuration file. Create a file named `config.toml` in the root directory of the executable and populate it with the following settings:

```toml
[api]
port = 3000

[database]
user = "postgres"
pass = "postgres"
host = "localhost"
port = 5432
name = "machine_api"
```

`this file must be in the root of the executable`

once configured, you can use the tools:

- To add tables to the database, run:

```bash
machine-go-server -c
```

`a user will be created, username: 'admin' and password: 'admin123'`

- To drop the tables in the database, run:

```bash
machine-go-server -d
```

- To start the HTTP server, run the following command:

```bash
machine-go-server
```

- For help, run:

```bash
machine-go-server --help
```

# Doc

### 🧩 Swagger

the swagger are in the [/swagger](/swagger) folder

### 🧩 Insominia Exports

the insominia exports are in the [/insominia](/insomnia) folder

# License

[MIT](https://github.com/elizandrodantas/machine-go-server/blob/main/LICENSE)