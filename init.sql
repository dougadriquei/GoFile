-- Table: public.purchase

-- DROP TABLE public.purchase;

CREATE ROLE postgres 
    LOGIN SUPERUSER PASSWORD 'floripa@123';


CREATE TABLE public.purchase
(
    cpf_cnpj character varying(14) COLLATE pg_catalog."default" NOT NULL,
    private integer,
    incompleted integer,
    last_purchase_date character varying(10) COLLATE pg_catalog."default",
    average_ticket numeric,
    last_purchase_ticket numeric,
    most_frequent_store character varying(14) COLLATE pg_catalog."default",
    last_purchase_store character varying(14) COLLATE pg_catalog."default",
    idpurchase integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    CONSTRAINT purchase_pkey PRIMARY KEY (idpurchase)
);

ALTER TABLE public.purchase
    OWNER to postgres;

-- TODO - Dúvida ao usar integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
-- Para esse negócio acredito estamos tratando de muitas informações (2.147.483.647 de registros, pode ser esgotado, como o postgress trata com isso, cria o cache 2?).

-- TABLESPACE pg_default;

-- ALTER TABLE public.purchase
--     OWNER to postgres;