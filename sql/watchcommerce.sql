CREATE DATABASE watchcommerce;

CREATE SCHEMA commerce;

-- table brand
CREATE TABLE commerce."brand"
(
    id         bigserial    NOT NULL,
    name       varchar(255) NULL,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now(),
    CONSTRAINT brand_pkey PRIMARY KEY (id)
);
CREATE INDEX index_brand_id ON commerce."brand" (id);

-- table product
CREATE TABLE commerce."product"
(
    id         bigserial    NOT NULL,
    brand_id   bigint       NOT NULL,
    name       varchar(255) NOT NULL,
    price      bigint       NOT NULL,
    quantity   bigint       NOT NULL,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now(),
    CONSTRAINT product_pkey PRIMARY KEY (id),
    CONSTRAINT brand_id_fkey FOREIGN KEY (brand_id) REFERENCES commerce."brand" (id) ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE INDEX index_product_id ON commerce."product" (id);
CREATE INDEX index_product_brand_id ON commerce."product" (brand_id);

-- table order
CREATE TABLE commerce."order"
(
    id         bigserial NOT NULL,
    total      bigint    NOT NULL,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now(),
    CONSTRAINT order_pkey PRIMARY KEY (id)
);
CREATE INDEX index_order_id ON commerce."order" (id);

-- table order_details
CREATE TABLE commerce."order_details"
(
    id         bigserial NOT NULL,
    order_id   bigint    NOT NULL,
    quantity   bigint    NOT NULL,
    price      bigint    NOT NULL,
    product_id bigint    NOT NULL,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now(),
    CONSTRAINT order_details_pkey PRIMARY KEY (id),
    CONSTRAINT order_details_fkey FOREIGN KEY (order_id) REFERENCES commerce."order" (id) ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE INDEX index_order_details_order_id ON commerce."order_details" (order_id);

CREATE
USER commerce WITH PASSWORD 'commerce';