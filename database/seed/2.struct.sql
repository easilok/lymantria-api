-- Create users table

CREATE TABLE public.users (
	id bigserial NOT NULL,
	email text NOT NULL,
	"password" text NOT NULL,
	"name" text NOT NULL,
	permissions _text NOT NULL DEFAULT ARRAY[]::text[],
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE UNIQUE INDEX idx_users_email ON public.users USING btree (email);

-- Create languages table

CREATE TABLE public.languages (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	user_id int8 NOT NULL,
	"name" text NOT NULL,
	slug text NOT NULL,
	CONSTRAINT languages_pkey PRIMARY KEY (id),
	CONSTRAINT fk_languages_user FOREIGN KEY (user_id) REFERENCES public.users(id)
);
CREATE INDEX idx_languages_deleted_at ON public.languages USING btree (deleted_at);
CREATE UNIQUE INDEX idx_name_slug ON public.languages USING btree (name, slug);

-- Create butterflies table

CREATE TABLE public.butterflies (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	user_id int8 NOT NULL,
	described text NOT NULL,
	rarity public."enum_butterfly_rarity" NOT NULL,
	daytime public."enum_butterfly_daytime" NOT NULL,
	"group" text NULL,
	appearances int8 NOT NULL DEFAULT 0,
	"size" text NULL,
	image text NULL,
	scientific text NULL,
	"family" text NULL,
	CONSTRAINT butterflies_pkey PRIMARY KEY (id),
	CONSTRAINT butterflies_scientific_unique UNIQUE (scientific),
	CONSTRAINT fk_butterflies_user FOREIGN KEY (user_id) REFERENCES public.users(id)
);
CREATE INDEX idx_butterflies_deleted_at ON public.butterflies USING btree (deleted_at);

-- Create butterfly_descriptions table

CREATE TABLE public.butterfly_descriptions (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	language_id int8 NOT NULL,
	butterfly_id int8 NOT NULL,
	user_id int8 NOT NULL,
	description text NOT NULL,
	common_name text NULL,
	observation text NULL,
	CONSTRAINT butterfly_descriptions_pkey PRIMARY KEY (id),
	CONSTRAINT fk_butterflies_details FOREIGN KEY (butterfly_id) REFERENCES public.butterflies(id),
	CONSTRAINT fk_butterfly_descriptions_butterfly FOREIGN KEY (butterfly_id) REFERENCES public.butterflies(id),
	CONSTRAINT fk_butterfly_descriptions_language FOREIGN KEY (language_id) REFERENCES public.languages(id),
	CONSTRAINT fk_butterfly_descriptions_user FOREIGN KEY (user_id) REFERENCES public.users(id)
);
CREATE UNIQUE INDEX idx_butterfly_description_language ON public.butterfly_descriptions USING btree (language_id, butterfly_id);
CREATE INDEX idx_butterfly_descriptions_deleted_at ON public.butterfly_descriptions USING btree (deleted_at);

-- Create monitorings table

CREATE TABLE public.monitorings (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	registered_at timestamptz NOT NULL,
	"local" text NOT NULL,
	"name" text NOT NULL,
	longitude numeric NULL,
	latitude numeric NULL,
	observation text NULL,
	hosted_by int8 NOT NULL,
	timestamp_end timestamptz NULL,
	temperature_start numeric NULL,
	humidity_start numeric NULL,
	wind_start text NULL,
	precipitation_start numeric NULL,
	sky_start text NULL,
	temperature_end numeric NULL,
	humidity_end numeric NULL,
	wind_end text NULL,
	precipitation_end numeric NULL,
	sky_end text NULL,
	CONSTRAINT monitorings_pkey PRIMARY KEY (id),
	CONSTRAINT fk_monitorings_host FOREIGN KEY (hosted_by) REFERENCES public.users(id),
	CONSTRAINT fk_monitorings_user FOREIGN KEY (hosted_by) REFERENCES public.users(id)
);
CREATE INDEX idx_monitorings_deleted_at ON public.monitorings USING btree (deleted_at);

-- Create monitoring_appearances table

CREATE TABLE public.butterfly_appearances (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	monitoring_id int8 NOT NULL,
	butterfly_id int8 NOT NULL,
	quantity int8 NOT NULL DEFAULT 1,
	image text NULL,
	observation text NULL,
	registered_by int8 NOT NULL,
	CONSTRAINT butterfly_appearances_pkey PRIMARY KEY (id),
	CONSTRAINT fk_butterfly_appearances_butterfly FOREIGN KEY (butterfly_id) REFERENCES public.butterflies(id),
	CONSTRAINT fk_butterfly_appearances_monitoring FOREIGN KEY (monitoring_id) REFERENCES public.monitorings(id),
	CONSTRAINT fk_butterfly_appearances_register FOREIGN KEY (registered_by) REFERENCES public.users(id),
	CONSTRAINT fk_butterfly_appearances_user FOREIGN KEY (registered_by) REFERENCES public.users(id),
	CONSTRAINT fk_monitorings_appearances FOREIGN KEY (monitoring_id) REFERENCES public.monitorings(id),
	CONSTRAINT fk_monitorings_butterflies FOREIGN KEY (monitoring_id) REFERENCES public.monitorings(id)
);
CREATE INDEX idx_butterfly_appearances_deleted_at ON public.butterfly_appearances USING btree (deleted_at);
CREATE UNIQUE INDEX idx_butterfly_monitoring ON public.butterfly_appearances USING btree (monitoring_id, butterfly_id);
