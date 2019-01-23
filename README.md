# **TuTicket-Distribuidos**

## Proyecto de Sistemas Distribuidos Espol

Integrantes:
Allan Alarcón
Luigi Basantes Zambrano
Juan Crow

### Dependencias a instalar
### Instalar GoLan

wget https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz

### Calculamos el CheckSum
sha256sum go1.10.3.linux-amd64.tar.gz 

#### Verificar que el output del checksum quede igual a este
fa1b0e45d3b647c252f51f5e1204aba049cde4af177ef9f2181f43004f901035  go1.10.3.linux-amd64.tar.gz

#### Usar tar para extraer los archivos en /usr/local
sudo tar -C /usr/local -xzf go1.10.3.linux-amd64.tar.gz

#### Adjustar la variable de ambiente agregando la siguiente linea al archivo ~/.bash_profile

export PATH=$PATH:/usr/local/go/bin

#### Volvemos a cargar la variable de ambiente para poder empezar a usar GO lan

source ~/.bash_profile



### Instalar net

go get -u golang.org/x/net

### Instalar go-sql-driver

go get -u github.com/go-sql-driver/mysql

### Instalar go-redis

go get -u github.com/go-redis/redis

### Instalar gin

go get -u github.com/gin-gonic/gin

### Otras Dependencias.

### Redis

#### Para la instalacion de Redis seguir:

sudo yum install epel-release

sudo yum update

sudo yum install redis

#### Para habilitarlo y que se inicie al booteo

sudo systemctl start redis

sudo systemctl enable redis

### Mysql

#### Para instalar mysql primero visitar https://dev.mysql.com/downloads/repo/yum/ 

#### Una vez ahi descargar la version para centos7 (linux7)

wget https://dev.mysql.com/get/mysql80-community-release-el7-1.noarch.rpm

#### Comprobar la firma y comparar con el sitio web (curiosamente usan md5 un hash que ya ha demostrado vulnerabilidades)

md5sum https://dev.mysql.com/get/mysql80-community-release-el7-1.noarch.rpm


#### Output
739dc44566d739c5d7b893de96ee6848  mysql80-community-release-el7-1.noarch.rpm

#### Comprobar con el sitio web Una vez hecho esto instalar paquete y de ahi podremos acceder a yum para instalar mysql

sudo rpm -ivh mysql80-community-release-el7-1.noarch.rpm

sudo yum install mysql-server

#### Iniciar Mysql

sudo systemctl start mysqld

#### Nota: no es necesario hacer que se inicie con el boot, ya mysql lo hace por defecto

### Nginx

#### Para instalar Nginx

sudo yum install epel-release //en el caso de no haberlo instalado con redis

sudo yum install nginx

#### Para iniciarlo

sudo systemctl start nginx

#### Para permitir en el firewall

sudo firewall-cmd --permanent --zone=public --add-service=http 

sudo firewall-cmd --permanent --zone=public --add-service=https

sudo firewall-cmd --reload

#### Para que se inicie con el booteo

sudo systemctl enable nginx

