# LitKeepService

LitKeep(a bill record app) backend service source code.

Using golang, fiber and gorm to make a good service.

## How to

### prepare your dababase.

Suppose you have a mysql or mariadb database.
Run sql files in `./sql` to initial the database.
The order of execution does not matter, there are no foreign keys here

### config your service like config.example.yaml

    the log util and redis is not used at all now!

`dev.yaml` for your development and `prod.yaml` for deployment.

### develop

1. If you like, you could use `air` to watch files' changement through development,
2. orjust rerun after `./script/restartcmd.sh`.
3. Or if you are using `pm2` from npm, run `./script/restart.sh` which will rebuild and rerun the service through pm2.
4. Or run `./script/build.sh` first and then run `pm2 start ./script/pm2.config.js`.

### deploy

Just like during development, 
1. if you are using `pm2` from npm, run `./script/restart.sh` which will rebuild and rerun the service through pm2.
2. Or run `./script/build.sh` first and then run `pm2 start ./script/pm2.config.js`.
3. Any other way, it is just a executable file after building.
