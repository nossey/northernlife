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
-- Name: tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tags (
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    id uuid NOT NULL,
    tag_name character varying,
    user_id character varying NOT NULL
);


ALTER TABLE public.tags OWNER TO postgres;

--
-- Name: TABLE tags; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.tags IS 'タグ';


--
-- Name: COLUMN tags.created_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.tags.created_at IS '作成日時';


--
-- Name: COLUMN tags.updated_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.tags.updated_at IS '更新日時';


--
-- Name: COLUMN tags.id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.tags.id IS 'ID';


--
-- Name: COLUMN tags.tag_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.tags.tag_name IS 'タグ名';


--
-- Name: COLUMN tags.user_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.tags.user_id IS 'タグの作成者';


--
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);


--
-- Name: tags tags_tag_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_tag_name_key UNIQUE (tag_name);


--
-- Name: tags fk_user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

