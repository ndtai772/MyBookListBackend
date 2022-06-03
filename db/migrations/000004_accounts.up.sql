--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2
-- Dumped by pg_dump version 14.3

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

--
-- Data for Name: accounts; Type: TABLE DATA; Schema: public; Owner: dev
--

INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (1, 'Nguyễn An', 'user@gmail.com', '$2a$10$nQAWCS0H2F49aoOjnsH7uuJS5QcpSYjEqK8wE3RSR.1yTHFGAO29y', 'https://ui-avatars.com/api/?name=Nguy%E1%BB%85n+An', false, '2022-05-27 05:24:33.614699');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (2, 'Stanford Senger', 'TZKXQlL@VIVkHJH.ru', 'ppnsbitheubryvbwknxyefjqhhtbbkzarngqwskpwrruujghuglkadrpgrtrtkpidfitqedbywjsbagncmokvzwixoloyfwnvdnj', 'https://ui-avatars.com/api/?name=Stanford+Senger', false, '2022-05-27 05:24:33.62138');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (3, 'Vesta Witting', 'ChuTkCc@cXOHJSP.com', 'edkzcvtkxvsfkvrzeureobptpwrfmfbjzjjlzkbvaszsvjfyrndlgzahcewmfgsrgiostfxwvgxnhrwrjrwtxhpeovupicgkdjai', 'https://ui-avatars.com/api/?name=Vesta+Witting', false, '2022-05-27 05:24:33.626307');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (4, 'Emiliano Berge', 'afEQJjW@uOwvthA.org', 'lyyiurrgcvpprogizhiccxmiwfolnvumjomutvxohwwimpiopbgoczvczlbuaxlwyledoucxxchzlqipwpsdnmmagssnaxhvbzff', 'https://ui-avatars.com/api/?name=Emiliano+Berge', false, '2022-05-27 05:24:33.63098');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (5, 'Talon Runte', 'DCITntD@IPTFThO.info', 'ngeteharmaicgmmdqdnjsmscoflujufktqtxmdewtlxlnsbfnfvgbhlnqnrabwgslvmqbvpjaynhloflujzwrlxhukptmoilmxjq', 'https://ui-avatars.com/api/?name=Talon+Runte', false, '2022-05-27 05:24:33.635665');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (6, 'Arturo Streich', 'aHgLGHj@WmfCmnA.ru', 'oyvmjfhrpzegfcenvodihkwcclbquojekwnubkpcqdazzxjxtyhzpubvordkfdtlbdqyznzfzbxosbisuycaovvgaarqpvlukwns', 'https://ui-avatars.com/api/?name=Arturo+Streich', false, '2022-05-27 05:24:33.640371');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (7, 'Ines Jast', 'vVteNLb@iBvbsPA.com', 'uiwotfrrjhvbvzanyxlbpupqfxawmixksyuhyorhtpxkgtptpreqvnipjnjxxdezmzorfbodmhkhofeqtfvkoiobtevbyswuceit', 'https://ui-avatars.com/api/?name=Ines+Jast', false, '2022-05-27 05:24:33.646083');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (8, 'Cody Stanton', 'hFPhsnZ@IZlVdTg.net', 'nfhactkkmcxifntchpjzarbiwaiptadakuakflwgghmoekokycicudrweeyxlozymghbvdhnyzrnrnhhoyufimpgzmyzkeuuqpzb', 'https://ui-avatars.com/api/?name=Cody+Stanton', false, '2022-05-27 05:24:33.651174');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (9, 'Terrance Hansen', 'DkwVqmu@cetppaM.info', 'vrehizisxnmszlpzxuewtlzluiapzglvhwjiqdorikabrdyxgrtddcwuqsgcfckvumitaeovckbrwqbnzgwdurslfuskkhyglcmi', 'https://ui-avatars.com/api/?name=Terrance+Hansen', false, '2022-05-27 05:24:33.65612');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (10, 'Assunta Stracke', 'LrWEBMH@MBiAEXh.info', 'bvtduviggzxnanybhxjjhwceagrnowwgflrrxhnxfuvjakysgkntieulmxxtaqguozpwsobifslxuetinwblypdsrasyopriaooj', 'https://ui-avatars.com/api/?name=Assunta+Stracke', false, '2022-05-27 05:24:33.660994');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (11, 'Dion Macejkovic', 'xUgtbvf@DObAFlt.biz', 'odyfzpzorjgbngbbrmnwznibhkmjkxjmfxkjtstbuqbuscvlzqzxapeklgusqlcrnehwszknxihfczkirmvzgfxkjzsxkssuvmbp', 'https://ui-avatars.com/api/?name=Dion+Macejkovic', false, '2022-05-27 05:24:33.666071');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (12, 'Clovis Hane', 'LJtAGiU@qusiDEe.org', 'ehrassnygosyxpbqcwmdzkfrtgmggmulnisemgsuroyrismtuwsdmifqdcrhhgshcqrdkgoqoruuqwakhdnrpihsbmaxededgrxh', 'https://ui-avatars.com/api/?name=Clovis+Hane', false, '2022-05-27 05:24:33.671409');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (13, 'Mose Kling', 'OpkInoN@YkHbgcD.info', 'jcsicsglatxhfxngrwnttiymsvflgcpcrbovdmkjumycxwbgotqzdeptecgaicrblxppneqakfufhqtovbtvisxxgnxcfpodgcnl', 'https://ui-avatars.com/api/?name=Mose+Kling', false, '2022-05-27 05:24:33.676696');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (14, 'Alexie Sporer', 'NtwISot@JixrDxm.net', 'ehrltoqyqcsxxhdnysmdyvtjmjrhboortzcysseuqukpyvrnpqqpufzverjrnavxcipekvmoylloecfnzzclmhxsozcxrqtgfsix', 'https://ui-avatars.com/api/?name=Alexie+Sporer', false, '2022-05-27 05:24:33.682007');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (15, 'Austen VonRueden', 'hhgBNNV@rIJfwnH.info', 'ndinapkqvqodwzixfummfkvsxrfabfpnpiwsnqlaavqhwcjcgrzunicwikmkxyiqpnhkiqmsvtvumsoiydltyimmmobinmufjkgs', 'https://ui-avatars.com/api/?name=Austen+VonRueden', false, '2022-05-27 05:24:33.687263');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (16, 'Vida Shields', 'YLsoEFP@koShPwD.com', 'sgjchgkqryrgecuebwkctmwywnnxphiihzpiajfuafynjydccezuavwahxnhbczmccjpvoxbmtdlcxzysebpdqrqjcdkudpnlgif', 'https://ui-avatars.com/api/?name=Vida+Shields', false, '2022-05-27 05:24:33.692568');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (17, 'Miles Schaden', 'ZLDOGLS@bRZHPYv.org', 'zvxrfhkrbsygjsibbfwewtrhdhcalotteamgulvvdugihtohffwapvldcnexjvlpdmjmtimxbxzdnldjafxsdhjzkctvtcxgsqqp', 'https://ui-avatars.com/api/?name=Miles+Schaden', false, '2022-05-27 05:24:33.697602');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (18, 'Adriana Stark', 'mnkDvZO@IBmpMBq.ru', 'swblbtacbgrftptyvobiukjctamehpditjhoekyxubuyzqqnlkkwcmxirtlaxughkfvbnbtjikmvjwxajltichpejpncnahykzku', 'https://ui-avatars.com/api/?name=Adriana+Stark', false, '2022-05-27 05:24:33.702644');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (19, 'Keven Reichert', 'oaLrIRJ@mTcqBum.ru', 'cwnrbejalvskpjtbgggkaqktsoppiimfchohrrhaskkcpzdyaazwuukqkurvxzviodexpfkwfqknyvjxxyahyfggfvchzvqaudmj', 'https://ui-avatars.com/api/?name=Keven+Reichert', false, '2022-05-27 05:24:33.707686');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (20, 'Domenica Bartoletti', 'KjQUBJS@tsQSOqc.com', 'qarbwrnrnozqlwszkcqooitxtbambxwsjdcmeufclwieqirryzddervnivyeglaymfwnjnqtziclioiypmuvizdvyzcprrcydxxg', 'https://ui-avatars.com/api/?name=Domenica+Bartoletti', false, '2022-05-27 05:24:33.712707');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (21, 'Morris Conroy', 'UAlbnZY@XtRkgqs.com', 'vcyjddhwunzbpgrsrptootkpsradvpmjrcuimcvnmkvgcvmnqbftduncbebdmllymbneytpchmivhqzgwjqgdettguchsoezjwyd', 'https://ui-avatars.com/api/?name=Morris+Conroy', false, '2022-05-27 05:24:33.717583');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (22, 'Lelia Rodriguez', 'fgeNJWw@CDlCAMc.com', 'eenvsjbjncuohecbbhfidugdthyefhtqilynszllzhokillnuzkvwcbcoazhbbxkemcrevpombrnqwvamracokuldotuggbfzwnv', 'https://ui-avatars.com/api/?name=Lelia+Rodriguez', false, '2022-05-27 05:24:33.722681');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (23, 'Prince Gibson', 'hovcNpj@qfxAFvK.org', 'qrkjytkbkdolmypkyrpblidgkbwrgccbmkllmhzcpvukfufggqwrreuacgbuslfecwrotpndueorupaiuvqmouqraphzyavwtmmu', 'https://ui-avatars.com/api/?name=Prince+Gibson', false, '2022-05-27 05:24:33.72771');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (24, 'Carey Murphy', 'Juncgbn@vgDyLKQ.biz', 'gsvgbkungqpwhmoxmtdffmadhgehnlgwdbruqqjiofjuoikdyeraxoybrndwgecucepctpflkbibylulvcxgngemnxrkruapujjl', 'https://ui-avatars.com/api/?name=Carey+Murphy', false, '2022-05-27 05:24:33.733355');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (25, 'Camryn Romaguera', 'wsbrNcv@SydUkFt.net', 'gmgfpvfcjfqnsijbgcxbwnarbendsbtuxttisrufbtegivemaaylzpsaeqirhneepybygxheyurcpjrquuqmixvpkyzngupjepkv', 'https://ui-avatars.com/api/?name=Camryn+Romaguera', false, '2022-05-27 05:24:33.738729');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (26, 'Anais Donnelly', 'rRCjphW@fHknrKo.ru', 'mwczfaubngkyudokbfnqeypiukukkyeuxsyieaqcryerzahcqeusslmqosavjvesuwgciwnbcegfuretxjbxlpnemjqlgiizclxw', 'https://ui-avatars.com/api/?name=Anais+Donnelly', false, '2022-05-27 05:24:33.743966');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (27, 'Giovanny Koch', 'lBoDyGy@msrByDM.biz', 'bslmuchozocmlvcdjokbjpziprdpxefwpgjyjqpjhqzunhpvouuwgiqqiszkgdzvosxpdfhwntreyhappsaydrpilyiahqhhogni', 'https://ui-avatars.com/api/?name=Giovanny+Koch', false, '2022-05-27 05:24:33.749513');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (28, 'Gussie Ankunding', 'nmfPfdw@KKFiuUx.ru', 'uzafqvyzpimmqcfulfmnqdecufzyxluezgujmkomlgmzdmiqldwzdfjlmvturmbnohdsmcbpkpqjwwqiffwauredebjbczlswpfb', 'https://ui-avatars.com/api/?name=Gussie+Ankunding', false, '2022-05-27 05:24:33.754767');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (29, 'Terence Lynch', 'kkjSjqw@BUVcDDd.ru', 'qodftnpqcompwccovonvcfoujjoazralyhnlvxaunthzibhvdqelqdshqjjgljlvuieclkscuughbjfctniycngwzcdywajzzzsc', 'https://ui-avatars.com/api/?name=Terence+Lynch', false, '2022-05-27 05:24:33.759952');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (30, 'Barney Lesch', 'hUuDbpZ@cudiZMD.com', 'psbbkbdrshxrubdgldewjeujrmoxdkvhslmmixrdqvcypybimkbrybxzeurpzlzcomonntaqmmcvwyhkbykmanobxinbuxptzmrq', 'https://ui-avatars.com/api/?name=Barney+Lesch', false, '2022-05-27 05:24:33.765193');
INSERT INTO public.accounts (id, name, email, hashed_password, avatar_url, is_admin, created_at) VALUES (31, 'Brain Cronin', 'nvxOAAP@yoJSLrI.org', 'sabonltayuqejtjyeuzbbuheeajgbopxyhsvqkdamgkwiexrfdfbgjrwdvdfsweounstiduwzuolebrqmtscaouwkoaytxvbtkxw', 'https://ui-avatars.com/api/?name=Brain+Cronin', false, '2022-05-27 05:24:33.770478');


--
-- Name: accounts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dev
--

SELECT pg_catalog.setval('public.accounts_id_seq', 31, true);


--
-- PostgreSQL database dump complete
--

