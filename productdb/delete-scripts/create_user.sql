#DROP USER "AdminUser"@'localhost';
#DROP USER "AdminUser"@'%';

CREATE USER "AdminUser"@"localhost" IDENTIFIED BY "AdminPassword";

GRANT ALL PRIVILEGES ON *.* TO "AdminUser"@"localhost" WITH GRANT OPTION;

CREATE USER 'AdminUser'@'%' IDENTIFIED BY 'AdminPassword';

GRANT ALL PRIVILEGES ON *.* TO 'AdminUser'@'%' WITH GRANT OPTION;

FLUSH PRIVILEGES;