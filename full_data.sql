--
-- PostgreSQL database dump
--

-- Dumped from database version 13.2
-- Dumped by pg_dump version 13.2

-- Started on 2022-05-06 00:17:03

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

SET default_table_access_method = heap;

--
-- TOC entry 205 (class 1259 OID 41288)
-- Name: group_subjects; Type: TABLE; Schema: public; Owner: setdatauser
--

CREATE TABLE public.group_subjects (
    id text NOT NULL,
    group_id text,
    teacher_subject_id text,
    created_date date
);


ALTER TABLE public.group_subjects OWNER TO setdatauser;

--
-- TOC entry 201 (class 1259 OID 41256)
-- Name: groups; Type: TABLE; Schema: public; Owner: setdatauser
--

CREATE TABLE public.groups (
    id text NOT NULL,
    name text,
    created_date date
);


ALTER TABLE public.groups OWNER TO setdatauser;

--
-- TOC entry 206 (class 1259 OID 41296)
-- Name: lessons; Type: TABLE; Schema: public; Owner: setdatauser
--

CREATE TABLE public.lessons (
    id text NOT NULL,
    name text,
    description text,
    video_file_url text,
    document_file_url text,
    group_subject_id text,
    created_date date
);


ALTER TABLE public.lessons OWNER TO setdatauser;

--
-- TOC entry 203 (class 1259 OID 41272)
-- Name: subjects; Type: TABLE; Schema: public; Owner: setdatauser
--

CREATE TABLE public.subjects (
    id text NOT NULL,
    name text,
    description text,
    created_date date
);


ALTER TABLE public.subjects OWNER TO setdatauser;

--
-- TOC entry 204 (class 1259 OID 41280)
-- Name: teacher_subjects; Type: TABLE; Schema: public; Owner: setdatauser
--

CREATE TABLE public.teacher_subjects (
    id text NOT NULL,
    teacher_id text,
    subject_id text,
    created_date date
);


ALTER TABLE public.teacher_subjects OWNER TO setdatauser;

--
-- TOC entry 202 (class 1259 OID 41264)
-- Name: user_groups; Type: TABLE; Schema: public; Owner: setdatauser
--

CREATE TABLE public.user_groups (
    id text NOT NULL,
    user_id text,
    group_id text,
    created_date date
);


ALTER TABLE public.user_groups OWNER TO setdatauser;

--
-- TOC entry 200 (class 1259 OID 41248)
-- Name: users; Type: TABLE; Schema: public; Owner: setdatauser
--

CREATE TABLE public.users (
    id text NOT NULL,
    username text,
    password text,
    email text,
    first_name text,
    last_name text,
    type text,
    created_date date
);


ALTER TABLE public.users OWNER TO setdatauser;

--
-- TOC entry 3028 (class 0 OID 41288)
-- Dependencies: 205
-- Data for Name: group_subjects; Type: TABLE DATA; Schema: public; Owner: setdatauser
--

COPY public.group_subjects (id, group_id, teacher_subject_id, created_date) FROM stdin;
55fe7685-743a-4733-8ed1-88d72add8f9b	e45ddb08-79d4-41ea-a9de-719a5dffe8d3	86e70d67-c034-46b5-85ce-a5bafd361a0b	2022-05-01
b2541136-ab19-4a54-9010-56e88c732e56	fed0aff4-00fb-4006-b4b5-4bc3ec9814e0	86e70d67-c034-46b5-85ce-a5bafd361a0b	2022-05-04
\.


--
-- TOC entry 3024 (class 0 OID 41256)
-- Dependencies: 201
-- Data for Name: groups; Type: TABLE DATA; Schema: public; Owner: setdatauser
--

COPY public.groups (id, name, created_date) FROM stdin;
e45ddb08-79d4-41ea-a9de-719a5dffe8d3	group2	2022-05-01
fed0aff4-00fb-4006-b4b5-4bc3ec9814e0	group1	2022-05-04
\.


--
-- TOC entry 3029 (class 0 OID 41296)
-- Dependencies: 206
-- Data for Name: lessons; Type: TABLE DATA; Schema: public; Owner: setdatauser
--

COPY public.lessons (id, name, description, video_file_url, document_file_url, group_subject_id, created_date) FROM stdin;
e17d611e-b540-483f-bf52-a90aa4a751fe	lesson2	lesson2 desc			55fe7685-743a-4733-8ed1-88d72add8f9b	2022-05-02
283f80fd-091e-4696-aa1e-2979f40aa58f	lesson1	lesson1 desc	123		55fe7685-743a-4733-8ed1-88d72add8f9b	2022-05-02
\.


--
-- TOC entry 3026 (class 0 OID 41272)
-- Dependencies: 203
-- Data for Name: subjects; Type: TABLE DATA; Schema: public; Owner: setdatauser
--

COPY public.subjects (id, name, description, created_date) FROM stdin;
55125f30-3caa-4f7f-bd25-361c71576418	subject1	very good subject1	2022-05-01
9e5356ee-7671-4c54-95cc-a3c3fa81868b	subject2	very good subject2	2022-05-01
\.


--
-- TOC entry 3027 (class 0 OID 41280)
-- Dependencies: 204
-- Data for Name: teacher_subjects; Type: TABLE DATA; Schema: public; Owner: setdatauser
--

COPY public.teacher_subjects (id, teacher_id, subject_id, created_date) FROM stdin;
86e70d67-c034-46b5-85ce-a5bafd361a0b	1780570e-5277-4d49-9e47-e9eb0a97ebd3	55125f30-3caa-4f7f-bd25-361c71576418	2022-05-01
3fb3ec9e-0f56-4281-bc48-813551c39a41	9e145b40-9703-4d03-beda-b353534df2d0	55125f30-3caa-4f7f-bd25-361c71576418	2022-05-01
645c4842-be39-403b-baae-ffdc199095ce	9e145b40-9703-4d03-beda-b353534df2d0	9e5356ee-7671-4c54-95cc-a3c3fa81868b	2022-05-01
94e14304-b20f-4663-8b76-f4551b48d2bb	9e145b40-9703-4d03-beda-b353534df2d0	9e5356ee-7671-4c54-95cc-a3c3fa81868b	2022-05-04
1698aabf-1b5a-46cc-9bbe-850daf82773d	1780570e-5277-4d49-9e47-e9eb0a97ebd3	9e5356ee-7671-4c54-95cc-a3c3fa81868b	2022-05-04
\.


--
-- TOC entry 3025 (class 0 OID 41264)
-- Dependencies: 202
-- Data for Name: user_groups; Type: TABLE DATA; Schema: public; Owner: setdatauser
--

COPY public.user_groups (id, user_id, group_id, created_date) FROM stdin;
\.


--
-- TOC entry 3023 (class 0 OID 41248)
-- Dependencies: 200
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: setdatauser
--

COPY public.users (id, username, password, email, first_name, last_name, type, created_date) FROM stdin;
ef540487-5a33-4276-bab3-0c225b5e3e8a	admin	$2a$04$TC7.SVLfQ9LoNeOAQOkGJOURxjFfFzWD/Id89zLxfVHPv25QR8IVG	tleugazy98@gmail.com			admin	2022-04-30
1780570e-5277-4d49-9e47-e9eb0a97ebd3	teacher1	$2a$04$0vaeDegGXKVv.PE3gP9TRejH75aCcUAfsNdBASpKpud8okK..qgJC	teacher@gmail.com			teacher	2022-04-30
9e145b40-9703-4d03-beda-b353534df2d0	teacher2	$2a$04$Q9eS02xaQrOeCUdgveWHtehYVdLDTtqrNOsBUDTUSMY2X01lq7qx2	teacher2@gmail.com			teacher	2022-05-01
\.


--
-- TOC entry 2890 (class 2606 OID 41295)
-- Name: group_subjects group_subjects_pkey; Type: CONSTRAINT; Schema: public; Owner: setdatauser
--

ALTER TABLE ONLY public.group_subjects
    ADD CONSTRAINT group_subjects_pkey PRIMARY KEY (id);


--
-- TOC entry 2882 (class 2606 OID 41263)
-- Name: groups groups_pkey; Type: CONSTRAINT; Schema: public; Owner: setdatauser
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_pkey PRIMARY KEY (id);


--
-- TOC entry 2892 (class 2606 OID 41303)
-- Name: lessons lessons_pkey; Type: CONSTRAINT; Schema: public; Owner: setdatauser
--

ALTER TABLE ONLY public.lessons
    ADD CONSTRAINT lessons_pkey PRIMARY KEY (id);


--
-- TOC entry 2886 (class 2606 OID 41279)
-- Name: subjects subjects_pkey; Type: CONSTRAINT; Schema: public; Owner: setdatauser
--

ALTER TABLE ONLY public.subjects
    ADD CONSTRAINT subjects_pkey PRIMARY KEY (id);


--
-- TOC entry 2888 (class 2606 OID 41287)
-- Name: teacher_subjects teacher_subjects_pkey; Type: CONSTRAINT; Schema: public; Owner: setdatauser
--

ALTER TABLE ONLY public.teacher_subjects
    ADD CONSTRAINT teacher_subjects_pkey PRIMARY KEY (id);


--
-- TOC entry 2884 (class 2606 OID 41271)
-- Name: user_groups user_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: setdatauser
--

ALTER TABLE ONLY public.user_groups
    ADD CONSTRAINT user_groups_pkey PRIMARY KEY (id);


--
-- TOC entry 2880 (class 2606 OID 41255)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: setdatauser
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


-- Completed on 2022-05-06 00:17:04

--
-- PostgreSQL database dump complete
--

