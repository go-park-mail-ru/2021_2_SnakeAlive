CREATE
OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at
= NOW();
RETURN NEW;
END;
$$
LANGUAGE plpgsql;

CREATE TABLE Users
(
    id          SERIAL      NOT NULL PRIMARY KEY,
    name        TEXT        NOT NULL,
    surname     TEXT        NOT NULL,
    email       TEXT        NOT NULL UNIQUE,
    password    TEXT        NOT NULL,
    avatar      TEXT                 DEFAULT 'https://snakehastrip.hb.bizmrg.com/default.jpg',
    description TEXT                 DEFAULT '',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON Users
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE Cookies
(
    id         SERIAL      NOT NULL PRIMARY KEY,
    hash       TEXT        NOT NULL,
    user_id    INT         NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE Places
(
    id          SERIAL NOT NULL PRIMARY KEY,
    name        TEXT   NOT NULL,
    country     TEXT   NOT NULL,
    rating      REAL   NOT NULL,
    lat         REAL   NOT NULL,
    lng         REAL   NOT NULL,
    description TEXT,
    tags        INTEGER[],
    photos      TEXT[],
    tsq         TSVECTOR
);

CREATE TABLE Tags
(
    id   SERIAL NOT NULL PRIMARY KEY,
    name TEXT
);

CREATE TABLE Trips
(
    id          SERIAL      NOT NULL PRIMARY KEY,
    title       TEXT        NOT NULL,
    origin      INT,
    description TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (origin) REFERENCES places (id) ON DELETE CASCADE
);

CREATE TABLE TripsUsers
(
    id      SERIAL NOT NULL PRIMARY KEY,
    user_id INT    NOT NULL,
    trip_id INT    NOT NULL,
    CONSTRAINT fk_trip FOREIGN KEY (trip_id) REFERENCES trips (id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON Trips
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE TripsPlaces
(
    id       SERIAL NOT NULL PRIMARY KEY,
    place_id INT    NOT NULL,
    trip_id  INT    NOT NULL,
    day      INT    NOT NULL,
    "order"  INT    NOT NULL,
    CONSTRAINT fk_trip FOREIGN KEY (trip_id) REFERENCES trips (id) ON DELETE CASCADE,
    CONSTRAINT fk_place FOREIGN KEY (place_id) REFERENCES places (id) ON DELETE CASCADE
);

CREATE TABLE Reviews
(
    id         SERIAL      NOT NULL PRIMARY KEY,
    title      TEXT,
    text       TEXT,
    rating     INT         NOT NULL,
    user_id    INT         NOT NULL,
    place_id   INT         NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (place_id) REFERENCES places (id) ON DELETE CASCADE
);

CREATE TABLE Countries
(
    id          SERIAL NOT NULL PRIMARY KEY,
    name        TEXT   NOT NULL,
    description TEXT   NOT NULL,
    translated  TEXT   NOT NULL,
    photo       TEXT
);

CREATE TABLE Albums
(
    id          SERIAL NOT NULL PRIMARY KEY,
    trip_id     INT    NOT NULL,
    author      INT    NOT NULL,
    title       TEXT,
    description TEXT,
    photos      TEXT[],
    CONSTRAINT fk_trip FOREIGN KEY (trip_id) REFERENCES trips (id) ON DELETE CASCADE,
    CONSTRAINT fk_author FOREIGN KEY (author) REFERENCES users (id) ON DELETE CASCADE
);

CREATE
OR REPLACE FUNCTION tsquery_trigger_function()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
AS
$$
BEGIN
UPDATE places
SET tsq = to_tsvector('russian', LOWER(new.name))
WHERE id = NEW.id;
RETURN NEW;
END;
$$;

CREATE TRIGGER tsquery_place
    AFTER INSERT
    ON places
    FOR EACH ROW
    EXECUTE PROCEDURE tsquery_trigger_function();


INSERT INTO Tags ("id", "name")
VALUES (1, 'Современные здания'),
       (2, 'Виды'),
       (3, 'Природа'),
       (4, 'Историческое место'),
       (5, 'Дворец'),
       (6, 'Архитектура'),
       (7, 'Церковь'),
       (8, 'Святое место'),
       (9, 'Резиденция'),
       (10, 'Заброшенное'),
       (11, 'Город'),
       (12, 'Мистическое место'),
       (13, 'Мост'),
       (14, 'Зоопарк'),
       (15, 'Крепость'), (16, 'Обзорное место'), (17, 'Музей');


INSERT INTO Countries ("name", "description", "photo", "translated")
VALUES ('Russia',
        'Россия – крупнейшая страна мира, расположенная в Восточной Европе и Северной Азии и омываемая водами Тихого и Северного Ледовитого океанов.',
        'russia.jpeg', 'Россия');
INSERT INTO Countries ("name", "description", "photo", "translated")
VALUES ('Germany',
        'Германия – государство в Западной Европе с лесами, реками, горными хребтами и пляжными курортами Северного моря.',
        'germany.jpeg', 'Германия');
INSERT INTO Countries ("name", "description", "photo", "translated")
VALUES ('USA',
        'Соединенные Штаты Америки – государство, состоящее из 50 штатов, занимает значительную часть Северной Америки. ',
        'usa.jpeg', 'Соединенные Штаты Америки');
INSERT INTO Countries ("name", "description", "photo", "translated")
VALUES ('UK',
        'Великобритания (официальное название – Соединенное Королевство Великобритании и Северной Ирландии) – островное государство на северо-западе Европы, состоящее из Англии, Шотландии, Уэльса и Северной Ирландии. ',
        'uk.jpeg', 'Великобритания');

INSERT INTO Users ("name", "surname", "password", "email", "description", "avatar")
VALUES ('Алексадра', 'Волынец', 'AAAcGFzc3dvcmQ=', 'alex@mail.ru', '', 'https://snakehastrip.hb.bizmrg.com/test.jpeg');
INSERT INTO Users ("name", "surname", "password", "email", "description", "avatar")
VALUES ('Никита', 'Черных', 'AAAZnJvbnRlbmQxMjM=', 'nikita@mail.ru', '',
        'https://snakehastrip.hb.bizmrg.com/test.jpeg');
INSERT INTO Users ("name", "surname", "password", "email", "description", "avatar")
VALUES ('Ксения', 'Самойлова', 'AAAMTIzNDU2Nzg=', 'ksenia@mail.ru', '', 'https://snakehastrip.hb.bizmrg.com/test.jpeg');
INSERT INTO Users ("name", "surname", "password", "email", "description", "avatar")
VALUES ('Андрей', 'Кравцов', 'AAAMDAwMTExMDAw', 'andrew@mail.ru', '', 'https://snakehastrip.hb.bizmrg.com/test.jpeg');

INSERT INTO Places ("name", "country", "lat", "lng", "rating", "description", "tags", "photos")
VALUES ('Москва-Cити',
        'Russia',
        55.749793,
        37.537393,
        5,
        '«Москва-Сити» — это Москва будущего, строящийся международный деловой квартал из ультрасовременных небоскрёбов. Уникальная для России и Восточной Европы зона деловой активности объединяет в себе апартаменты для жилья, офисные здания, многочисленные площадки для торговли и отдыха. Москвичей и гостей города «Москва-Сити» привлекает необычной конфигурацией сооружений, развитой социально-культурной инфраструктурой, — здесь можно посетить бутики, спа-салоны, рестораны, клубы, выставочные галереи, развлекательные центры.
        Создатели комплекса стремились не просто выстроить небоскрёбы, а сделать так, чтобы они органично вписались в ансамбль исторических памятников.',
        ARRAY[1, 2],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/moscow_city_0.jpeg',
        'https://snakehastrip.hb.bizmrg.com/moscow_city_1.jpeg',
        'https://snakehastrip.hb.bizmrg.com/moscow_city_2.jpeg',
        'https://snakehastrip.hb.bizmrg.com/moscow_city_3.jpeg',
        'https://snakehastrip.hb.bizmrg.com/moscow_city_4.jpeg']);

INSERT INTO Places ("name", "country", "lat", "lng", "rating", "description", "tags", "photos")
VALUES ('Воробьевы горы',
        'Russia',
        55.7077713,
        37.5394096,
        4,
        'Воробьевы горы — самый высокий из семи холмов, на которых стоит город. Отсюда открывается прекрасный панорамный вид на Москву, здесь снято множество лучших кинофильмов, в этом месте всегда огромное количество свадебных кортежей, байкеров и туристов. Карамзин рассказывает историю о том, как в начале XIX века известная французская художница Элизабет Виже-Лебрен приехала в Москву, чтобы написать знаменитый вид, открывающийся с Воробьевых гор, для императора Павла I. Она долго стояла на высоком берегу Москвы-реки, глядя на перспективу, а затем отбросила палитру, произнеся лишь: «Не смею...»
        ',
        ARRAY[3, 2],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/vorobievi_gory_0.jpeg',
        'https://snakehastrip.hb.bizmrg.com/vorobievi_gory_1.jpeg',
        'https://snakehastrip.hb.bizmrg.com/vorobievi_gory_2.jpeg']);
INSERT INTO Places ("name", "country", "lat", "lng", "rating", "description", "tags", "photos")
VALUES ('Дворцово-парковый ансамбль Петергоф',
        'Russia',
        59.8833300,
        29.9000000,
        5,
        'Дворцово-парковый ансамбль Петергоф — царство фонтанов, феерия играющей воды, дворцы, в которых оживает эпоха Петра Великого, блистательные интерьеры времен императрицы Елизаветы и царя Николая I.
        Петергоф был основан в самом начале XVIII в. императором Петром I как величественный памятник, прославляющий победу России в борьбе за выход к Балтийскому морю. Это самая роскошная летняя царская резиденция. Феерическое зрелище множества играющих водометов сделало его всемирно известным. В летнее время в парке просто не протолкнуться, особенно по выходным.
        Главная достопримечательность парка - уникальная фонтанная система, созданная в петровские времена под руководством первого русского инженера-гидравлика Туволкова. Ее часто сравнивают с Версальской, но в некоторых отношениях она даже превосходит французский аналог.
        Фонтаны Петергофа действуют по принципу сообщающихся сосудов за счет перепада высот рельефа и не требуют специального накачивания воды. Фонтаны и каскады питаются пресной водой, поступающей из источников Ропшинских высот по 22-километровому самотечному водоводу.
        Нижний парк растянулся по прибрежной полосе на 2 км и занимает площадь в 102 га. Особое своеобразие ему придает близость моря, с которым он так органично связан. Финский залив соединен специально прорытым Морским каналом с Большим каскадом - крупнейшим фонтанным сооружением мира, включающим 75 фонтанов и около 250 скульптур и декоративных украшений.
        Водопады, водометы, позолоченные статуи, барельефы, вазы, балюстрады, неумолкающий шум воды - все это поражает своим великолепием и создает торжественное и праздничное настроение.
        Особенно красив Большой каскад в праздничные дни, когда он становится площадкой для великолепных костюмированных представлений в обрамлении множества хрустальных струй воды, сопровождаемых светомузыкальными эффектами и фейерверками.
        ',
        ARRAY[4, 5],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/petergof_0.jpg']);
INSERT INTO Places ("name", "country", "lat", "lng", "rating", "description", "tags", "photos")
VALUES ('Аничков мост',
        'Russia',
        59.9332352,
        30.3433533,
        4,
        'Аничков мост — один из самых знаменитых мостов Санкт-Петербурга, история которого тесно связана с основанием Северной столицы. Сам мост не является шедевром архитектурной мысли; визитной карточкой и украшением Санкт-Петербурга он стал благодаря великолепным изваяниям скульптора Петра Клодта. Еще он просто находится на Невском проспекте.
        Петербургские жители с восторгом приняли творения Клодта. Пресса наперебой расхваливала талантливого скульптора. Ваятель удостоился похвалы и внимания самого царя — в 1841 году, вскоре после церемонии в честь открытия моста, Николай I пожаловал Клодту орден Святой Анны третьей степени.
        Тогда же родилось известное фривольное прозвище переправы — «Мост восемнадцати яиц». При подсчёте элементов мужского детородного органа учитывался и городовой, пост которого располагался на мосту вплоть до 1917 года.',
        ARRAY[6],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/anichkov_most_0.jpeg',
        'https://snakehastrip.hb.bizmrg.com/anichkov_most_1.jpeg',
        'https://snakehastrip.hb.bizmrg.com/anichkov_most_2.jpeg',
        'https://snakehastrip.hb.bizmrg.com/anichkov_most_3.jpeg',
        'https://snakehastrip.hb.bizmrg.com/anichkov_most_4.jpeg',
        'https://snakehastrip.hb.bizmrg.com/anichkov_most_5.jpeg']);
INSERT INTO Places ("name", "country", "lat", "lng", "rating", "description", "tags", "photos")
VALUES ('Эльбрус',
        'Russia',
        43.2577100,
        42.6443500,
        3,
        'Эльбрус — высочайшая гора России, расположенная на границе республик Кабардино-Балкария и Карачаево-Черкесия. Характерные двуглавые вершины горы, покрытые снегами – визитная карточка Северного Кавказа. Ослепительный Эльбрус также является самым высоким пиком Европы.
        Эльбрус сформировался более миллиона лет назад, раньше он был действующим вулканом, и до сих пор не утихают споры, потух он или просто спит. В пользу версии о спящем вулкане говорит тот факт, что горячие массы сохраняются в его глубинах и подогревают термальные источники до +60 °C. В недрах Эльбруса рождаются и насыщаются знаменитые минеральные воды курортов Северного Кавказа — Кисловодска, Пятигорска, Ессентуков, Железноводска. Гора состоит из чередующихся слоёв пепла, лавы и туфа. Последний раз исполин извергался в 50 году н. э.
        Высота восточной вершины горы – 5621 метр, западной – 5642 метра, между ними лежит седловина, уступающая вершинам по высоте 300 метров. Белый покров Эльбруса состоит из более 80 ледников, крупнейшие из них – Терскол, Большой Азау и Ирик. Ледники начинаются с высоты 3500 метров, их площадь – 145 км². Огромные ледяные массы дают начало рекам Кубани, Малке, Баксану и притокам Терека.Климат Приэльбрусья мягкий, влажность невысокая, благодаря чему морозы переносятся легко. А вот климат самого вулкана суровый, схожий с арктическим. Средняя зимняя температура — от 10 градусов мороза у подножия горы, до –25 °C на уровне 2000-3000 метров, и до –40 °C на вершине. Осадки на Эльбрусе частые и обильные, в основном это снег.
        Летом воздух прогревается до +10 °C — до высоты 2500 метров, а на высоте в 4200 метров даже в июле не бывает теплее –14 °C.
        Погода очень неустойчива: ясный безветренный день может мгновенно превратиться в снежное ненастье с сильным ветром.',
        ARRAY[3, 2],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/elbrus_0.jpeg',
        'https://snakehastrip.hb.bizmrg.com/elbrus_1.jpeg',
        'https://snakehastrip.hb.bizmrg.com/elbrus_2.jpeg',
        'https://snakehastrip.hb.bizmrg.com/elbrus_3.jpeg']);
INSERT INTO Places ("name", "country", "lat", "lng", "rating", "description", "tags", "photos")
VALUES ('Озеро Байкал',
        'Russia',
        53,
        108,
        5,
        'Озеро Байкал, расположенное на юге Восточной Сибири, на границе Иркутской области и Республики Бурятия, относится к числу самых древних водоемов нашей планеты. Но больше всего оно известно тем, что является самым глубоким озером на Земле и одновременно крупнейшим естественным резервуаром пресной воды – 19% всех мировых запасов.
        И сам Байкал, и прибрежные территории отличает неповторимая в своем разнообразии флора и фауна, что делает эти места поистине уникальными, неизменно привлекающими к себе внимание научных умов и многочисленных любителей путешествий и настоящих искателей приключений.
        По очертаниям Байкал похож на узкий полумесяц, настолько легко запоминающийся, что его без труда находят на карте России даже те, кто не особенно силен в географии. Простершийся с юго-запада на северо-восток на целых 636 километров, Байкал словно протискивается между горными массивами, а его водная гладь находится на высоте более 450 метров над уровнем моря, что дает все основания считать его горным озером. С запада к нему примыкают Байкальский и Приморский хребты, с востока и юго-востока – массивы Улан-Бургасы, Хамар-Дабан и Баргузинский. И весь этот природный ландшафт настолько гармоничен, что одно без другого трудно представить.
        Длина береговой линии сибирского «полумесяца» составляет 2100 км, на нем расположено 27 островов, самый большой из которых – Ольхон. Озеро находится в своеобразной котловине, которую, как было сказано выше, со всех сторон окружают горные хребты и сопки. Это дает основание предполагать, что береговая линия водоема на всем протяжении одинаковая. На самом же деле скалистым и обрывистым является только западное побережье Байкала. Рельеф же восточного более пологий: в некоторых местах горные вершины находятся на отдалении от берега на 10 и больше километров.',
        ARRAY[3, 2],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/ozero_baikal_0.jpeg',
        'https://snakehastrip.hb.bizmrg.com/ozero_baikal_1.jpeg',
        'https://snakehastrip.hb.bizmrg.com/ozero_baikal_2.jpeg',
        'https://snakehastrip.hb.bizmrg.com/ozero_baikal_3.jpeg',
        'https://snakehastrip.hb.bizmrg.com/ozero_baikal_4.jpeg']);
INSERT INTO Places ("name", "country", "lat", "lng", "rating", "description", "tags", "photos")
VALUES ('Ростовский кремль',
        'Russia',
        57.1838024,
        39.4146069,
        3,
        'Ростовский кремль – величественное каменное укрепление в старинном городе Ростове Великом. Территория кремля расположена в историческом центре, на небольшой возвышенности, недалеко от северо-западной оконечности озера Неро. Она очень красива и давно стала визитной карточкой древнего русского города. Белокаменные башни, мощные стены и купола церквей отлично вписаны в окружающий ландшафт. Многим силуэты храмов и башен Ростовского кремля знакомы по популярной комедии «Иван Васильевич меняет профессию».
        В Ростов приезжают не только, чтобы полюбоваться на архитектурные памятники. На территории кремля разместилось около десяти интересных музеев, рассказывающих об истории города, церковных реликвиях и знаменитой ростовской финифти. В музейных залах можно увидеть редкие произведения древнерусского искусства, старинные иконы и церковную утварь. А со стен кремля открываются прекрасные виды на городские кварталы и озеро Неро.
        ',
        ARRAY[4, 6],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/rostovskiy_kreml_0.jpeg',
        'https://snakehastrip.hb.bizmrg.com/rostovskiy_kreml_1.jpeg',
        'https://snakehastrip.hb.bizmrg.com/rostovskiy_kreml_2.jpeg',
        'https://snakehastrip.hb.bizmrg.com/rostovskiy_kreml_3.jpeg']);
INSERT INTO Places ("name", "country", "lat", "lng", "rating", "description", "tags", "photos")
VALUES ('Тауэрский мост',
        'UK',
        51.5054898,
        -0.0755067,
        5,
        'Тауэрский мост – разводная переправа через реку Темзу в центре Лондона, неподалеку от Тауэрской башни. Это одна из наиболее популярных достопримечательностей Лондона, которую легко узнают даже те, кто никогда не бывал в столице Соединенного Королевства. Ежегодно сюда стекаются тысячи туристов, открывающие для себя великолепие этого готического сооружения.
        Мост имеет общую длину 244 метра, посередине находятся две башни, каждая высотой 65 метров, между ними имеется пролет в 61 метр, который является разводным элементом. Это позволяет пропускать суда к городским причалам в любое время дня или ночи. Мощная гидравлическая система первоначально была водяной, в движение ее приводили большие паровые машины. Сегодня система полностью заменена на масляную и управляется с помощью компьютера.',
        ARRAY[6],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/tauerskiy-most_0.jpeg',
        'https://snakehastrip.hb.bizmrg.com/tauerskiy-most_1.jpeg',
        'https://snakehastrip.hb.bizmrg.com/tauerskiy-most_2.jpeg']);
INSERT INTO Places ("name", "country", "lat", "lng", "rating", "description", "tags", "photos")
VALUES ('Вестминстерское Аббатство',
        'UK',
        51.4993138,
        -0.1272882,
        4,
        'Вестминстерское Аббатство — не только самая большая церковь в Лондоне, но и средоточие государственной жизни страны, Здесь были коронованы 38 монархов, начиная с Вильгельма Завоевателя, ставшего английским королем в день Рождества 1666 г., т.е. все монархи, кроме Эдуарда V, убитого в 1483 г., и Эдуарда VIII, отрекшегося от престола в 1936 г.
        Все, что вам нужно знать об этом месте:  “Большую часть посетителей сюда привлекают надгробия.”',
        ARRAY[7, 8],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/vestminsterskoe-abbatstvo_0.jpeg',
        'https://snakehastrip.hb.bizmrg.com/vestminsterskoe-abbatstvo_1.jpeg',
        'https://snakehastrip.hb.bizmrg.com/vestminsterskoe-abbatstvo_2.jpeg']);
INSERT INTO Places ("name", "country", "lat", "lng", "rating", "description", "tags", "photos")
VALUES ('Букингемский дворец',
        'UK',
        51.5012171,
        -0.1420831,
        5,
        'Букингемский дворец – резиденция британских монархов в Лондоне. Сегодня там живет и работает Елизавета II. Во дворце кипит жизнь: проходят приемы и мероприятия государственного значения. Покой королевской семьи охраняют гвардейцы – их ярко-красные наряды видны издалека.',
        ARRAY[5, 9],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/bukingemskiy-dvorets_0.jpeg',
        'https://snakehastrip.hb.bizmrg.com/bukingemskiy-dvorets_1.jpeg',
        'https://snakehastrip.hb.bizmrg.com/bukingemskiy-dvorets_2.jpeg',
        'https://snakehastrip.hb.bizmrg.com/bukingemskiy-dvorets_3.jpeg',
        'https://snakehastrip.hb.bizmrg.com/bukingemskiy-dvorets_4.jpeg']);
INSERT INTO Places ("name", "country", "lat", "lng", "rating", "description", "tags", "photos")
VALUES ('Остров Мэн',
        'UK',
        54.163865,
        -4.487597,
        3,
        'Остров Мэн — горбатый остров длиной 50 км и шириной 16, расположен в Ирландском море между Англией и Ирландией. Остров имеет свой собственный парламент, обычаи и особую атмосферу. Лучший способ увидеть остров — это пройти 40 км по маршруту «Тысячелетний путь», от Рамси до Каслтауна вдоль всего горного хребта, проходящего по острову. Наилучший вид открывается с вершины Снэфелл (610 м), куда можно подняться на фуникулере.
        В XIX веке туризм стал самой важной отраслью экономики острова. Массовый туризм начался в 1830-х годах, в связи с организацией регулярного пароходного сообщения между островом (прежде всего — Дугласом) и Ливерпулем. Количество туристов, посещавших остров, росло на протяжении всего XIX века и начала XX века. Например, если в 1870-х годах каждый год остров посещало сто тысяч туристов, то в 1913 году остров посетило 553 000 туристов. После этого в связи с началом Первой мировой войны количество туристов снизилось, и пик 1913 года был побит только в 1948 году, но после этого года количество туристов начало снижаться. Связано это с постепенным ростом благосостояния населения и развитием авиации, следствием чего стал рост популярности курортов Южной Европы и более экзотичных мест.',
        ARRAY[3, 2],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/ostrov-men_0.jpeg',
        'https://snakehastrip.hb.bizmrg.com/ostrov-men_1.jpeg',
        'https://snakehastrip.hb.bizmrg.com/ostrov-men_2.jpeg']);

INSERT INTO Places ("name", "country", "lat", "lng", "rating", "description", "tags", "photos")
VALUES ('Старый Собор',
        'Никарагуа',
        12.156542,
        -86.271268,
        4,
        'Старый Собор, имеющий второе название — Собора Святого Джеймса, собирался и затем был доставлен в Никарагуа с помощью кораблей из Бельгии в 1920 году. Храм был воздвигнут по проекту бельгийского архитектора Пабло Домбаче, проживавшего в столице Манагуа. Особенность этого собора состоит в том, что на всем западном полушарии до этого не было подобных сооружений, возведенных только из бетона и на металлическом каркасе.
        Местом установки была выбрана западная сторона площади Республики. Храм выполнен в неоклассическом стиле. В качестве фундамента послужила основа снесенной ранее церкви Сантьяго. Своим величественным видом, красотой и размерами собор просто притягивает внимание не только местного населения, но и посещающих страну туристов. ',
        ARRAY[6],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/old_sobor_0.jpeg']);

INSERT INTO Places ("name", "country", "lat", "lng", "rating", "description", "tags", "photos")
VALUES ('Вулкан Момбачо',
        'Никарагуа',
        11.8261075,
        -85.9675,
        4,
        'Вулкан Момбачо — стратовулкан в Никарагуа в 10 километрах от города Гранады. Вулкан и прилегающие к нему территория относится к заповеднику. Благодаря удивительной флоре, фауне и поразительным открывающимся видам, вулкан пользуется большой популярностью. На вершине находится туристический центр. Вулкан невысокий - 1344 метра над уровнем моря, но несмотря на это его хорошо видно с окружающих городов.
        Несмотря на то, что Момбачо относится к действующим вулканам, последний раз его активность наблюдалась в 1570 году. Почти круглый год вершина покрыта плотными облаками, что дает 100% влажность. Вулкан Момбачо похож на вечнозеленую гору посреди сухих тропиков. У его подножья раскинулся буйный лес с огромными цветами. Если Вы решили посетить парк на своей машине, то помните, на его территорию пускают только полноприводные автомобили 4х4, а все из-за того, что дорога на вулкан очень крутая и обычная машина туда просто не доедет. ',
        ARRAY[3, 2],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/vulkan-mombacho_0.jpeg',
        'https://snakehastrip.hb.bizmrg.com/vulkan-mombacho_1.jpeg',
        'https://snakehastrip.hb.bizmrg.com/vulkan-mombacho_2.jpeg',
        'https://snakehastrip.hb.bizmrg.com/vulkan-mombacho_3.jpeg',
        'https://snakehastrip.hb.bizmrg.com/vulkan-mombacho_4.jpeg',
        'https://snakehastrip.hb.bizmrg.com/vulkan-mombacho_5.jpeg']);

INSERT INTO Places ("name", "country", "lat", "lng", "rating", "description", "tags", "photos")
VALUES ('Архипелаг Солентинаме',
        'Никарагуа',
        11.133,
        -84.9964,
        5,
        'Архипелаг Солентинаме расположен в южной части озера Никарагуа. Он интересен как место обитания множества птиц, обезьян и других экзотических животных.
        Происхождение островов является вулканическим. Солентинаме состоит из четырех крупных островов, каждый в несколько километров в поперечнике, а также включает в себя примерно 32 мелких острова.
        На островах архипелага Солентинаме обнаружены древние петроглифы — рисунки на скалах, изображающие попугаев, обезьян и людей.
        Власти Никарагуа присвоили островам Солентинаме статус национального природного памятника Никарагуа. ',
        ARRAY[3, 2],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/arkhipelag-solentiname_0.jpeg',
        'https://snakehastrip.hb.bizmrg.com/arkhipelag-solentiname_1.jpeg',
        'https://snakehastrip.hb.bizmrg.com/arkhipelag-solentiname_2.jpeg',
        'https://snakehastrip.hb.bizmrg.com/varkhipelag-solentiname_3.jpeg']);

INSERT INTO Places ("name", "country", "lat", "lng", "rating", "description", "tags", "photos")
VALUES ('Хамберстоун',
        'Чили',
        -20.20675,
        -69.79319,
        5,
        'Хамберстоун — заброшенный шахтерский город на севере Чили в пустыне Атакама, находится в часе езды от города Икике. ЮНЕСКО внесла этот город-призрак в список объектов Всемирного Наследия, присвоив жутковатому месту статус музея под открытым небом. Если вы собираетесь в Чили, не упустите возможность взглянуть на то, к чему приводят мировые экономические бумы.
        Жители города переезжали туда для добычи селитры, необходимой для нитратных удобрений. Во второй половина 19 века город бурно развивался.
        В Хамберстоуне начала формироваться местная культура — «пампинос» — с особенными ценностями, фольклором и уникальным языком. Тут были собственные таможни и законы, дружелюбная атмосфера, всеобщая солидарность по отношению друг к другу, борьба за социальную справедливость и уважение к людям. Жители города существовали на условно отдельной территории, хотя формально Хамберстоун принадлежал к Чили и неустанно сколачивал стране огромный капитал. Настоящий мираж в пустыне, Хамберстоун обрастал новыми зданиями и улочками, заборами и фонарными столбами как бы назло природе. В городе была церковь и собственный театр, бары и рестораны.
        Однако в 1958 году компания, занимавшаяся разработками месторождения селитры, была закрыта, а ещё через два года месторождение было исчерпано.
        Чилийский город стал призраком, который пережил невероятный взлет, а теперь должен был навсегда быть похороненным в песках. Однако в конце 60-х годов, когда правительство страны искало любые способы борьбы с экономическим спадом, Хамберстоун был объявлен национальной достопримечательностью, а наполовину засыпанные песком дома и улицы — музеем под открытым небом. Целый отряд рабочих был отправлен для приведения города в порядок: вновь налажено освещение, проведена дорога, а на открытках появляются виды пустынного чуда. ',
        ARRAY[11, 10],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/khamberstoun_0.jpeg',
        'https://snakehastrip.hb.bizmrg.com/khamberstoun_1.jpeg',
        'https://snakehastrip.hb.bizmrg.com/khamberstoun_2.jpeg',
        'https://snakehastrip.hb.bizmrg.com/khamberstoun_3.jpeg']);

INSERT INTO Places ("name", "country", "lat", "lng", "rating", "description", "tags", "photos")
VALUES ('Остров Пасхи',
        'Чили',
        -27.1265,
        -109.2951,
        5,
        'В день праздника пасхи в 1722 г. голландский капитан Яков Роггевен наткнулся на остров в центральной части тихого океана. Он стал первым европейцем, ступившим на этот уединенный клочок суши. в судовом журнале роггевен отметил его как «Остров Пасхи».
        Остров Пасхи, или Рапа Нуи — остров в Тихом океане на территории Чили, известный благодаря гигантским каменным статуям ',
        ARRAY[12, 10],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/ostrov-paskhi_0.jpeg',
        'https://snakehastrip.hb.bizmrg.com/ostrov-paskhi_1.jpeg',
        'https://snakehastrip.hb.bizmrg.com/kostrov-paskhi_2.jpeg']);

INSERT INTO Reviews (id, title, text, rating, user_id, place_id, created_at)
VALUES (DEFAULT, 'title', 'text', 10, 1, 1, DEFAULT);
INSERT INTO Reviews (id, title, text, rating, user_id, place_id, created_at)
VALUES (DEFAULT, 'title2', 'text2', 11, 1, 1, DEFAULT);
INSERT INTO Reviews (id, title, text, rating, user_id, place_id, created_at)
VALUES (DEFAULT, 'title3', 'text3', 12, 1, 2, DEFAULT);

INSERT INTO Countries ("name", "description", "photo", "translated")
VALUES ('Czech',
        'Чехия, или Чешская Республика — центральноевропейское государство, граничащее на севере с Польшей, на западе — с Германией, на юге — с Австрией, и на востоке — со Словакией.',
        'https://snakehastrip.hb.bizmrg.com/czech.jpeg', 'Чехия');

INSERT INTO Places ("name", "country", "lat", "lng", "rating", "description", "tags", "photos")
VALUES ('Крепость Пражский Град',
        'Czech',
        50.0911,
        14.4016,
        4.4,
        'Пражский Град является символом чешского государства вот уже более тысячи лет. Он был заложен в 9-м столетии и стал резиденцией чешских королей, а потом – и президентов. Один из самых больших крепостных комплексов в мире, он объединяет дворцовые, служебные, фортификационные, религиозные сооружения, сады и живописные уголки. Панорама Града – незабываемое зрелище, которое останется в Вашей памяти навсегда.',
        ARRAY[4, 2, 6, 15],
        ARRAY[
        'https://snakehastrip.hb.bizmrg.com/krep_prazh_grad1.webp',
        'https://snakehastrip.hb.bizmrg.com/krep_prazh_grad2.webp',
        'https://snakehastrip.hb.bizmrg.com/krep_prazh_grad3.webp'
        'https://snakehastrip.hb.bizmrg.com/krep_prazh_grad4.webp',
           ]),
        ('Карлов Мост',
        'Czech',
        50.08635,
        14.4119534,
        4.8,
        'Карлов мост – символ Праги, самая популярная достопримечательность города и второй в истории каменный мост через Влтаву. Он до сих пор привлекает людей не только своей красотой и тридцатью барочными статуями XVII века, но и множеством легенд и загадок. Его история связана с магией, астрологией и неразгаданными тайнами. Удивительно, что мост преподносит сюрпризы исследователям до сих пор.',
        ARRAY[4, 2, 11],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/czech_carlov_most1.webp',
        'https://snakehastrip.hb.bizmrg.com/czech_carlov_most2.webp',
        'https://snakehastrip.hb.bizmrg.com/czech_carlov_most3.webp'
           ]),
        ('Собор Святого Вита',
        'Czech',
        50.0908918,
        14.4005114,
        4.8,
        'Собор Святого Вита – храм, украшающий исторический центр чешской столицы; доминанта Пражского града. Это один из самых знаменитых долгостроев Европы: подлинная средневековая готика органично сочетается в нем с элементами барокко и искусными неоготическими стилизациями.',
        ARRAY[4, 2, 6, 12],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/czech_st_vit1.webp',
        'https://snakehastrip.hb.bizmrg.com/czech_st_vit12.webp',
        'https://snakehastrip.hb.bizmrg.com/czech_st_vit3.webp'
            ]),
       ('Петршин Холм',
        'Czech',
        50.08138,
        14.39388,
        4.2,
        'Холм Петршин, как известно, раньше имел несколько других названий. Сначала он назывался горой, потом носил им Копец или Холм Св. Лаврентия, тогда там был установлен крест',
        ARRAY[2, 16],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/czech_petr_1.webp',
        'https://snakehastrip.hb.bizmrg.com/czech_petr_2.webp',
        'https://snakehastrip.hb.bizmrg.com/czech_petr_3.webp'
            ]),
       ('Йозефов',
        'Czech',
        50.089,
        14.419,
        4.6,
        'Йозефов – самый таинственный район Праги, овеянный множеством легенд, преданий и мифов. До 1850 года представлял собой сердце еврейской общины Праги, поэтому его также называют еврейским кварталом. Район находится между правым берегом реки Влтавы и Староместской площадью.',
        ARRAY[4, 12],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/czech_josefov1.webp',
        'https://snakehastrip.hb.bizmrg.com/czech_josefov2.webp',
        'https://snakehastrip.hb.bizmrg.com/czech_josefov3.webp'
            ]),
       ('Староместская площадь',
        'Czech',
        50.0875692,
        14.4211874,
        4.7,
        'Холм Петршин, как известно, раньше имел несколько других названий. Сначала он назывался горой, потом носил им Копец или Холм Св. Лаврентия, тогда там был установлен крест',
        ARRAY[4, 6, 8],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/czech_pl1.webp',
        'https://snakehastrip.hb.bizmrg.com/czech_pl2.webp',
        'https://snakehastrip.hb.bizmrg.com/czech_pl3.webp'
            ]);



INSERT INTO Countries ("name", "description", "photo", "translated")
VALUES ('Netherlands',
        'Нидерланды или официально Королевство Нидерландов – европейское государство, включающее в себя материковую часть, расположенную в Западной Европе, и три острова в Карибском море – Бонэйр, Саба и Синт-Эстатиус. Помимо того, на правах членов королевства в состав страны входят три самоуправляемых государственных образования, занимающих острова Кюрасау, Аруба и Синт-Мартин.',
        'https://snakehastrip.hb.bizmrg.com/netherlands.jpeg', 'Нидерланды');
VALUES ('Музей Рембрандта',
        'Netherlands',
        52.369369,
        4.9012351
        4.89,
        'Музей Рембрандта — художественный музей на улице Йоденбреестраат в Еврейском квартале Амстердама. Музей открылся в 1911 году в доме, который Рембрандт приобрёл на пике славы в 1639 году и в котором проживал до своего банкротства в 1656 году.',
        ARRAY[17],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/neth_rembrandthuis1.webp',
        'https://snakehastrip.hb.bizmrg.com/neth_rembrandthuis2.webp',
        'https://snakehastrip.hb.bizmrg.com/neth_rembrandthuis3.webp'
           ]),
       ('Улица Красных Фонарей',
        'Netherlands',
        52.3724309,
        4.8975156,
        4.5,
        'Улица Красных Фонарей — район Амстердама, являющийся одной из самых известных и пикантных достопримечательностей города. Он состоит из сети переулков с невысокими зданиями, содержащими множество однокомнатных номеров с витринами.',
        ARRAY[1, 6],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/neth_red_light1.webp',
        'https://snakehastrip.hb.bizmrg.com/neth_red_light2.webp',
        'https://snakehastrip.hb.bizmrg.com/neth_red_light3.webp'
           ]),
       ('Музей Ван Гога',
        'Netherlands',
        52.3584159,
        4.8810756,
        4.1,
        'Музей Ван Гога в Амстердаме — самая большая коллекция творений Ван Гога в мире, здесь представлено 200 картин, 437 рисунков и 31 гравюра, включая знаменитые Подсолнухи, Белый сад и Желтый дом, а также множество автопортретов художника.',
        ARRAY[6, 17],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/neth_van1.webp',
        'https://snakehastrip.hb.bizmrg.com/neth_van2.webp',
        'https://snakehastrip.hb.bizmrg.com/neth_van3.webp'
            'https://snakehastrip.hb.bizmrg.com/neth_van4.webp',
           ]),
       ('Площадь Дам',
        'Netherlands',
        52.3730144,
        4.8933578,
        4.13,
        'Площадь Дам – центральное место Амстердама, где проходят ярмарки, концерты и прочие торжественные события. Она находится в 750 м от основного транспортного узла столицы Нидерландов – станции Сентрал. Данное место по форме является прямоугольником площадью 20 000 м². К площади примыкает бульвар Дамрак и улицы Рокин, Ньивендейк, Кальверстрат, Дамстрат.',
        ARRAY[2, 6, 11, 15],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/neth_sq_dam1.webp',
        'https://snakehastrip.hb.bizmrg.com/neth_sq_dam2.webp',
        'https://snakehastrip.hb.bizmrg.com/neth_sq_dam3.webp'
           ]),
       ('Рейксмузеум',
        'Netherlands',
        52.3599648,
        4.8853368,
        4.35,
        'Рейксмузеум — Амстердамский Государственный музей, крупнейший музей искусства в Нидерландах. В нем хранится более миллиона экспонатов. Рейксмузеум был открыт в 1885 году и стал достопримечательностью города, сочетая черты готического стиля и Возрождения.',
        ARRAY[4, 17],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/neth_wtf1.webp',
        'https://snakehastrip.hb.bizmrg.com/neth_wtf2.webp',
        'https://snakehastrip.hb.bizmrg.com/neth_wtf3.webp'
           ]),
       ('Парк цветов Кекенхоф',
        'Netherlands',
        52.2700188,
        4.5463893,
        4.93,
        'Кекенхоф — всемирно известный королевский парк цветов в Нидерландах. Находится на побережье между Амстердамом и Гаагой в небольшом городке Лиссе, при приближении к которому можно видеть указатели маршрута на Кекенхоф.',
        ARRAY[2, 3],
        ARRAY[
            'https://snakehastrip.hb.bizmrg.com/neth_park1.webp',
        'https://snakehastrip.hb.bizmrg.com/neth_park2.webp',
        'https://snakehastrip.hb.bizmrg.com/neth_park3.webp'
           ]);