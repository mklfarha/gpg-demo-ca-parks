package config

var configDev = `
db:
  - name: ca_parks 
    host: 127.0.0.1
    port: 3306
    user: root
    pswd: root
    params: parseTime=true&loc=UTC&charset=utf8mb4&collation=utf8mb4_unicode_ci
    driver: "mysql"

aws:
  - region: 
    key_id: 
    secret: 
    bucket: 
`
