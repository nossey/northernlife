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
-- Name: posts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.posts (
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    id uuid NOT NULL,
    user_id character varying NOT NULL,
    title text NOT NULL,
    body text,
    plain_body text,
    published boolean,
    thumbnail character varying DEFAULT 'https://northernlife-content.net/lunch.jpg'::character varying NOT NULL
);


ALTER TABLE public.posts OWNER TO postgres;

--
-- Name: TABLE posts; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.posts IS 'ポスト';


--
-- Name: COLUMN posts.created_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.posts.created_at IS '作成日時';


--
-- Name: COLUMN posts.updated_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.posts.updated_at IS '更新日時';


--
-- Name: COLUMN posts.id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.posts.id IS 'ポストのID';


--
-- Name: COLUMN posts.title; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.posts.title IS 'ポストのタイトル';


--
-- Name: COLUMN posts.body; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.posts.body IS 'ポストの本文(markdown形式)';


--
-- Name: COLUMN posts.plain_body; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.posts.plain_body IS 'ポストの本文(レンダリングされたmarkdownを平文化)';


--
-- Name: COLUMN posts.published; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.posts.published IS 'ポストが公開されているか(trueなら公開, falseなら下書き状態)';


--
-- Name: COLUMN posts.thumbnail; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.posts.thumbnail IS 'サムネ画像のリンク';


--
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);


--
-- Name: posts fk_user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

