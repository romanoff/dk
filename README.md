DataKeeper
==========

DataKeeper allows you to easily create backups of data and restore them. It supports backups for databases as well as file backups. Here is an example of configuration file:
```toml
[mysql]
name=root
password=password
host=localhost
database=blog_development
```

This is an example of `.dk` configuration file. It specifies data source (in this case mysql database). Here are commands you can use to work with local data:

Commands:
---------
`dk init <dbname>` - Creates datakeeper config after couple questions
`dk config` - Show datakeeper configuration
`dk create <name>` - Creates database dump and stores it with specified name
`dk list` -> List all local dumps
`dk remove <name>` - Removes local database dump
`dk apply <name>` - Applies database dump locally
