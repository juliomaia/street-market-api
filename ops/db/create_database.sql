create database if not exists street_market_db;
use street_market_db;

create table if not exists street_market(   
    id varchar(50),
    `long` decimal(11,7) NOT NULL, 
    lat	decimal(11,7) NOT NULL, 
    setcens bigint NOT NULL, 
    areap bigint NOT NULL, 
    coddist int NOT NULL, 
    distrito varchar(255) NOT NULL,
    codsubpref int NOT NULL,
    subprefe varchar(255) NOT NULL,
    regiao5 varchar(255) NOT NULL,
    regiao8 varchar(255) NOT NULL,
    nome_feira varchar(255) NOT NULL,
    registro varchar(255) NOT NULL,
    logradouro varchar(255) NOT NULL,
    numero varchar(255),
    bairro varchar(255) NOT NULL,
    referencia varchar(255), 
    created_at timestamp,
    updated_at timestamp,
PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;


INSERT INTO street_market_db.street_market
  (id,`long`, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, regiao5, regiao8, nome_feira, registro, logradouro, numero, bairro, referencia, created_at, updated_at)
VALUES ('a398b63e-f8b3-11eb-9a03-0242ac130003',-46.550164,-23.558733,355030885000091,3550308005040,87,'VILA FORMOSA',26,'ARICANDUVA-FORMOSA-CARRAO','Leste','Leste 1','VILA FORMOSA','4041-0','RUA MARAGOJIPE','S/N','VL FORMOSA','TV RUA PRETORIA', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO street_market_db.street_market
  (id,`long`, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, regiao5, regiao8, nome_feira, registro, logradouro, numero, bairro, referencia, created_at, updated_at)
VALUES ('d398b63e-f8b3-11eb-9a03-0242ac130003',-46.574716,-23.584852,355030893000035,3550308005042,95,'VILA PRUDENTE',29,'VILA PRUDENTE','Leste','Leste 1','PRACA SANTA HELENA','4045-2','RUA JOSE DOS REIS','909','VL ZELINA','RUA OLIVEIRA GOUVEIA',CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
