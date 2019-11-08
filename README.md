# email.go
send email using go 

### Basic Requirements
Operating System: **Ubuntu 18.04**. Later releases of Ubuntu may also be compatible with these instructions.

Install/configure the following applications:
* git ([install instructions](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git))
  * Don't forget to correctly set your git credentials:
```
git config --global user.name "FirstName LastName"
git config --global user.email "my@example.com"
```
* docker ([install instructions](https://docs.docker.com/install/linux/docker-ce/ubuntu/))
  * Make sure you can run docker as a non-root user ([instructions](https://docs.docker.com/install/linux/linux-postinstall/)).
    Note you will need to reboot your machine for the group changes (`groupadd`) to work.
* docker-compose ([install instructions](https://docs.docker.com/compose/install/))


### Setup

Get the code and get set up:
```bash
git https://github.com/Vikas14Web/c51_sym.git
cd c51_sym

# Run 
docker-compose up -d
```

### To bash into the main docker
```bash
docker-compose exec app bash
```

### Running Tests

You may now run those scripts by using `composer csfix`, `composer cscheck` and `composer phpstan`!
```
csfix: this script will try to fix as many errors as possible
cscheck: this script will checks potential errors in your code
composer phpstan: PHPStan focuses on finding errors in your code without actually running it
```
