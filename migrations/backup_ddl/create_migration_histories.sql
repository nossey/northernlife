--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.8
-- Dumped by pg_dump version 12.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

--
-- Name: migration_histories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.migration_histories (
    created_at timestamp with time zone,
    migration_name character varying
);


ALTER TABLE public.migration_histories OWNER TO postgres;

--
-- Name: TABLE migration_histories; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.migration_histories IS 'DBのMigration履歴用テーブル';


--
-- Name: COLUMN migration_histories.created_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.migration_histories.created_at IS '作成日時';


--
-- Name: migration_histories migration_histories_migration_name_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.migration_histories
    ADD CONSTRAINT migration_histories_migration_name_unique UNIQUE (migration_name);


--
-- PostgreSQL database dump complete
--

