
//user表
CREATE TABLE USER (
                      id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                      `name` VARCHAR(10) NOT NULL ,
                      UID VARCHAR(200),
                      BIRTHDAY BIGINT ,
                      `AREA` VARCHAR(5) ,
                      image VARCHAR(50),
                      email VARCHAR(20) NOT NULL,
                      slogan VARCHAR(20),
                      telephone VARCHAR(11),
                      `password` VARCHAR(20),
                      create_time BIGINT,
                      update_time BIGINT,
                      delete_time BIGINT);