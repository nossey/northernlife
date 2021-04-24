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
-- Name: tags_posts_attachments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tags_posts_attachments (
    created_at timestamp with time zone,
    id uuid NOT NULL,
    post_id uuid NOT NULL,
    tag_id uuid NOT NULL
);


ALTER TABLE public.tags_posts_attachments OWNER TO postgres;

--
-- Name: COLUMN tags_posts_attachments.created_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.tags_posts_attachments.created_at IS '作成日時';


--
-- Name: COLUMN tags_posts_attachments.id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.tags_posts_attachments.id IS 'ID';


--
-- Name: COLUMN tags_posts_attachments.post_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.tags_posts_attachments.post_id IS 'ポストのID';


--
-- Name: COLUMN tags_posts_attachments.tag_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.tags_posts_attachments.tag_id IS 'タグのID';


--
-- Name: tags_posts_attachments tags_posts_attachment_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags_posts_attachments
    ADD CONSTRAINT tags_posts_attachment_pkey PRIMARY KEY (id);


--
-- Name: tags_posts_attachments fk_post_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags_posts_attachments
    ADD CONSTRAINT fk_post_id FOREIGN KEY (post_id) REFERENCES public.posts(id);


--
-- Name: tags_posts_attachments fk_tag_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags_posts_attachments
    ADD CONSTRAINT fk_tag_id FOREIGN KEY (tag_id) REFERENCES public.tags(id);


--
-- PostgreSQL database dump complete
--

