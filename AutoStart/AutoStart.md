# Document of AutoStart

## Design Mentality

**Linux system
* using linux command "systemctl" to make all of service auto startup with linux operation system
* the service file of start must be copy to the path "/etc/systemd/system"
* using command "systemctl enable *service_name*" to make the service start with system
* if you change the file of service,execute "systemctl daemon-reload" to make the changed file working


## Design Implement

**Linux system

* create template for every service,parse the template and change  variable to insure that working for every time;

