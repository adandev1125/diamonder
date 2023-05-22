# diamonder

**diamonder** is a framework that allows for easy and efficient web development with Go. This framework is designed to be light-weight and modular, making it easy to customize and integrate into any project.

## Structure

The structure of **diamonder** is organized into the following directories:

- assets : Contains all static assets such as images, stylesheets and javascript files.
- config : Contains configuration files such as database configurations, application settings and more.
- controllers : Contains all controller files that are responsible for handling requests and interacting with models.
- models : Contains all model files that are responsible for interacting with the database.
- routes : Contains all route files that define the endpoints and their corresponding functions.
- system : Contains all system files that are responsible for managing the application.
- views : Contains all view files that are responsible for displaying data to the users.

## Third-party dependencies

**diamonder** currently utilizes the following third-party dependencies:

- github.com/go-sql-driver/mysql v1.7.1 : A MySQL driver for Go.
- github.com/gorilla/mux v1.8.0 : A powerful URL router and dispatcher.

## Getting Started

To get started with **diamonder**, you first clone the repository.

```
git clone https://github.com/adandev1125/diamonder
```

After clone is completed, you can start server with following commands.

```
go run main.go
```

## Quick Guide

### Adding Routes

All routes are defined in [routes/routes.go](routes/routes.go). You can also use subdirectories to organize yourself.

To add a route, you must define its path and controller.

For example,

``` go
router.HandleFunc("/user/login", controllers.UserLoginHandler).Methods("POST")
```

### Adding Controllers

You can add controllers in controllers folder.

Here is a code snippet for a simple controller.

``` go
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/welcome.html"))
	tmpl.Execute(w, WelcomePageData{
		Title:         "Diamond",
		FirstHtmlFile: "views/welcome.html",
	})
}
```

After finishing your controller, you must add it to the [routes](routes/routes.go) to see if it's working correctly.

### Configuring Database

To use database for you app, you have to configure for the database.

You can see the config in [config/config.yml](config/config.yml).

``` yaml
# Configuration for databases used by the application
databases:
  - name: 'test_db'
    driver: 'mysql'
    host: 'db'
    port: 3306
    database: 'test'
    username: 'test_user'
    password: 'test_password'
    parse_time: true # Setting to parse time is true to automatically convert time values into Go's time.Time format
```

### Using database

After configuring database in config/database.go, you can use it in any controller or any model.

For example:

``` go
... ... ...

var db = system.GetDatabase("test_db") // database name specified in config.yml

rows, error := db.Query("SELECT * FROM test_users")

if error != nil {
    log.Fatal(error)
}

defer rows.Close()

... ... ...
```

### Using static files
Static files are stored in [assets](assets/) folder.
You can use them in views like this:
``` html
... ... ...

<link rel="stylesheet" type="text/css" href="assets/welcome.css" />

... ... ...

<img style="width: 200px; margin-top: 80px;" src="assets/logo.png" />

... ... ...
```

### Templating Views
Templating views is supported by the html/template library.

To see more information about this please visit [https://pkg.go.dev/html/template](https://pkg.go.dev/html/template).

### Configuring Environments
To use another port, change the port in the [config](config/config.yml) file.

For example:
``` yaml
# Configuration for application "diamonder"
...
port: 8080
```

## Conclusion

diamonder is a great choice for web development with Go. With its easy-to-use structure and lightweight design, it is a perfect fit for any new or existing project. I hope you find it useful and I look forward to seeing what you create with it!

## Release Notes
### Version 0.0.3
Changed the way of configuration to YAML.

Enabled multiple databases.


### Version 0.0.2
Added docker files.

Added github workflows for CI/CD.


### Version 0.0.1
Initial Release.


