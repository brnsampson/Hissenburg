INSERT INTO genders (gender) values ("Undefined");
INSERT INTO genders (gender) values ("Intersex");
INSERT INTO genders (gender) values ("Indeterminate");
INSERT INTO genders (gender) values ("Fluid");
INSERT INTO genders (gender) values ("NonBinary");
INSERT INTO genders (gender) values ("Female");
INSERT INTO genders (gender) values ("Male");

INSERT INTO names (name) values ("Wenlan");
INSERT INTO names (name) values ("Agune");
INSERT INTO names (name) values ("Beatrice");
INSERT INTO names (name) values ("Breagan");
INSERT INTO names (name) values ("Bronwyn");
INSERT INTO names (name) values ("Cannora");
INSERT INTO names (name) values ("Drelil");
INSERT INTO names (name) values ("Elgile");
INSERT INTO names (name) values ("Esme");
INSERT INTO names (name) values ("Groua");
INSERT INTO names (name) values ("Henaine");
INSERT INTO names (name) values ("Liranne");
INSERT INTO names (name) values ("Lirathil");
INSERT INTO names (name) values ("Lisabeth");
INSERT INTO names (name) values ("Moralil");
INSERT INTO names (name) values ("Morgwin");
INSERT INTO names (name) values ("Sybil");
INSERT INTO names (name) values ("Theune");
INSERT INTO names (name) values ("Ygwal");
INSERT INTO names (name) values ("Yslen");
INSERT INTO names (name) values ("Arwel");
INSERT INTO names (name) values ("Bevan");
INSERT INTO names (name) values ("Boroth");
INSERT INTO names (name) values ("Borrid");
INSERT INTO names (name) values ("Breagle");
INSERT INTO names (name) values ("Breglor");
INSERT INTO names (name) values ("Canhoreal");
INSERT INTO names (name) values ("Emrys");
INSERT INTO names (name) values ("Ethex");
INSERT INTO names (name) values ("Gringle");
INSERT INTO names (name) values ("Grinwit");
INSERT INTO names (name) values ("Gruwid");
INSERT INTO names (name) values ("Gruwth");
INSERT INTO names (name) values ("Gwestin");
INSERT INTO names (name) values ("Mannog");
INSERT INTO names (name) values ("Melnax");
INSERT INTO names (name) values ("Orthax");
INSERT INTO names (name) values ("Triunein");
INSERT INTO names (name) values ("Yirmeor");

INSERT INTO surnames (surname) values ("Abernathy");
INSERT INTO surnames (surname) values ("Addercap");
INSERT INTO surnames (surname) values ("Burl");
INSERT INTO surnames (surname) values ("Candlewick");
INSERT INTO surnames (surname) values ("Cormick");
INSERT INTO surnames (surname) values ("Crumwaller");
INSERT INTO surnames (surname) values ("Dunswallow");
INSERT INTO surnames (surname) values ("Getri");
INSERT INTO surnames (surname) values ("Glass");
INSERT INTO surnames (surname) values ("Harkness");
INSERT INTO surnames (surname) values ("Harper");
INSERT INTO surnames (surname) values ("Loomer");
INSERT INTO surnames (surname) values ("Malksmilk");
INSERT INTO surnames (surname) values ("Smythe");
INSERT INTO surnames (surname) values ("Sunderman");
INSERT INTO surnames (surname) values ("Swinney");
INSERT INTO surnames (surname) values ("Thatcher");
INSERT INTO surnames (surname) values ("Tolmen");
INSERT INTO surnames (surname) values ("Weaver");
INSERT INTO surnames (surname) values ("Wolder");

INSERT INTO backgrounds (title, description, image) values ("Alchemist", "", "");
INSERT INTO backgrounds (title, description, image) values ("Blacksmith", "", "");
INSERT INTO backgrounds (title, description, image) values ("Butcher", "", "");
INSERT INTO backgrounds (title, description, image) values ("Burglar", "", "");
INSERT INTO backgrounds (title, description, image) values ("Carpenter", "", "");
INSERT INTO backgrounds (title, description, image) values ("Cleric", "", "");
INSERT INTO backgrounds (title, description, image) values ("Gambler", "", "");
INSERT INTO backgrounds (title, description, image) values ("Gravedigger", "", "");
INSERT INTO backgrounds (title, description, image) values ("Herbalist", "", "");
INSERT INTO backgrounds (title, description, image) values ("Hunter", "", "");
INSERT INTO backgrounds (title, description, image) values ("Magician", "", "");
INSERT INTO backgrounds (title, description, image) values ("Mercenary", "", "");
INSERT INTO backgrounds (title, description, image) values ("Merchant", "", "");
INSERT INTO backgrounds (title, description, image) values ("Miner", "", "");
INSERT INTO backgrounds (title, description, image) values ("Outlaw", "", "");
INSERT INTO backgrounds (title, description, image) values ("Performer", "", "");
INSERT INTO backgrounds (title, description, image) values ("Pickpocket", "", "");
INSERT INTO backgrounds (title, description, image) values ("Smuggler", "", "");
INSERT INTO backgrounds (title, description, image) values ("Servant", "", "");
INSERT INTO backgrounds (title, description, image) values ("Ranger", "", "");

INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Agune" AND genders.gender="Female";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Beatrice" AND genders.gender="Female";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Breagan" AND genders.gender="Female";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Bronwyn" AND genders.gender="Female";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Cannora" AND genders.gender="Female";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Drelil" AND genders.gender="Female";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Elgile" AND genders.gender="Female";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Esme" AND genders.gender="Female";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Groua" AND genders.gender="Female";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Henaine" AND genders.gender="Female";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Liranne" AND genders.gender="Female";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Lirathil" AND genders.gender="Female";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Lisabeth" AND genders.gender="Female";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Moralil" AND genders.gender="Female";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Morgwin" AND genders.gender="Female";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Sybil" AND genders.gender="Female";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Theune" AND genders.gender="Female";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Wenlan " AND genders.gender="Female";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Ygwal" AND genders.gender="Female";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Yslen" AND genders.gender="Female";

INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Arwel" AND genders.gender="Male";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Bevan" AND genders.gender="Male";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Boroth" AND genders.gender="Male";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Borrid" AND genders.gender="Male";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Breagle" AND genders.gender="Male";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Breglor" AND genders.gender="Male";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Canhoreal" AND genders.gender="Male";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Emrys" AND genders.gender="Male";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Ethex" AND genders.gender="Male";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Gringle" AND genders.gender="Male";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Grinwit" AND genders.gender="Male";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Gruwid" AND genders.gender="Male";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Gruwth" AND genders.gender="Male";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Gwestin" AND genders.gender="Male";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Mannog" AND genders.gender="Male";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Melnax" AND genders.gender="Male";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Orthax" AND genders.gender="Male";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Triunein" AND genders.gender="Male";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Wenlan " AND genders.gender="Male";
INSERT INTO name_genders (name, gender) SELECT names.id, genders.id FROM names CROSS JOIN genders WHERE names.name="Yirmeor" AND genders.gender="Male";
