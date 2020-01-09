# Sparrow
a web framework,on module.
* log
* middleware
* controller
* orm
* config
## Install
#### step 1:
```
git clone
```
#### step 2:
```
go mod init main
```
why `main`?
The framework is not a module,exactly,that's a project.In order to be able to use module and better migrate, the module name should not be changed.
#### step 3:
```
go mod tidy
```
If download package fails due to network, you can set `proxy=https://goproxy.io`.
### Directory
```
--app               *usiness logic      controll or model
--bin               *bin                window exe,linux bin
--config            *config file        db.yaml,must yaml
--route             *route
--share             *global and share   all user share
--unit              *user exclusive     user exclusive
main.go             *main
```
### Run
main.go
```
//init mux
mux := http.NewServeMux()
route.InitRoute(mux)

//run server
share.Run(mux)
```
/route/route.go
```
mux.Handle("/index/getIndex", &unit.Context{Action: index.GetIndex})
```
/app/index/action.go
```
func GetIndex(b *unit.Context) string {
	return "hello world"
}
```
### Middleware
At /unit/context,then do anything yourself.







