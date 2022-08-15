CREATE TABLE "product" (
                         "id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
                         "name" VARCHAR NOT NULL,
                         "qty" bigint,
                         "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         "updated_at" TIMESTAMP(3) NOT NULL,

                         PRIMARY KEY ("id")
);



CREATE UNIQUE INDEX "product_key" ON "product"("name");