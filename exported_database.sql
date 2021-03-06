PGDMP                         x            goBlog    12.3    12.3                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                        0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            !           1262    16449    goBlog    DATABASE     �   CREATE DATABASE "goBlog" WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'English_United States.1251' LC_CTYPE = 'English_United States.1251';
    DROP DATABASE "goBlog";
                postgres    false                        3079    16450 	   uuid-ossp 	   EXTENSION     ?   CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;
    DROP EXTENSION "uuid-ossp";
                   false            "           0    0    EXTENSION "uuid-ossp"    COMMENT     W   COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';
                        false    2            �            1259    16461 
   base_table    TABLE     �   CREATE TABLE public.base_table (
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);
    DROP TABLE public.base_table;
       public         heap    postgres    false            �            1259    16475    post    TABLE     �   CREATE TABLE public.post (
    id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    title character varying(255) NOT NULL,
    body text,
    author_id uuid
)
INHERITS (public.base_table);
    DROP TABLE public.post;
       public         heap    postgres    false    2    203            �            1259    16464    user_account    TABLE     �   CREATE TABLE public.user_account (
    id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    email character varying(255) NOT NULL,
    password_hash character varying(255) NOT NULL
)
INHERITS (public.base_table);
     DROP TABLE public.user_account;
       public         heap    postgres    false    2    203                      0    16461 
   base_table 
   TABLE DATA           <   COPY public.base_table (created_at, updated_at) FROM stdin;
    public          postgres    false    203   �                 0    16475    post 
   TABLE DATA           R   COPY public.post (created_at, updated_at, id, title, body, author_id) FROM stdin;
    public          postgres    false    205                    0    16464    user_account 
   TABLE DATA           X   COPY public.user_account (created_at, updated_at, id, email, password_hash) FROM stdin;
    public          postgres    false    204   5       �
           2606    16483    post post_pkey 
   CONSTRAINT     L   ALTER TABLE ONLY public.post
    ADD CONSTRAINT post_pkey PRIMARY KEY (id);
 8   ALTER TABLE ONLY public.post DROP CONSTRAINT post_pkey;
       public            postgres    false    205            �
           2606    16474 #   user_account user_account_email_key 
   CONSTRAINT     _   ALTER TABLE ONLY public.user_account
    ADD CONSTRAINT user_account_email_key UNIQUE (email);
 M   ALTER TABLE ONLY public.user_account DROP CONSTRAINT user_account_email_key;
       public            postgres    false    204            �
           2606    16472    user_account user_account_pkey 
   CONSTRAINT     \   ALTER TABLE ONLY public.user_account
    ADD CONSTRAINT user_account_pkey PRIMARY KEY (id);
 H   ALTER TABLE ONLY public.user_account DROP CONSTRAINT user_account_pkey;
       public            postgres    false    204            �
           2606    16484    post post_author_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.post
    ADD CONSTRAINT post_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.user_account(id) ON DELETE CASCADE;
 B   ALTER TABLE ONLY public.post DROP CONSTRAINT post_author_id_fkey;
       public          postgres    false    204    205    2711                  x������ � �           x���Kj1����I�^>E]v3��U��ɤ	4.�PޘI�	f�uR�K�����@�#�j�����/�Ѿ�j=U��0�|�w^�?���|]�� ��z)��8�N7��q!�
�$�A���E������Vۆ�F��t�~�Z��*��,"#'
"��8��Һ2ӆ*�����5:)���7�|��MFU�P�Q��Q1��&j�~�� `��~~5-u�0�h�g~�:(
%B���%y�uA&a�RWoyT���G�O��i:��ա         �   x�}λR�@���v�s�������`ƌ�.��	8>�c����_��ǀEP��H�(�	��{쿅�p����XD"��,���}��?�u��f����q��^����j����~���T�1��]\l�k���S9k��aD
1E�I��KX�H#݁�H��p�_�Ek���.��ە���N����3��2RW��gr�o�2U�����:�?Q����NY\     